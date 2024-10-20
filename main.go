package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Comment struct {
	Id int
	Username string
	Timestamp string
	Content string
}

type CommentCreate struct {
	Username string
	Content string
}

type CommentUpdate struct {
	Content string
}

func main() {
	fmt.Println("Building REST APIs in Go 1.22!")

	var db *sqlx.DB
	var err error

	db, err = sqlx.Connect("pgx", "host=localhost user=root password=root dbname=commentstore port=5432 sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r*http.Request) {
		fmt.Fprint(w, "hello world!")
	})

	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r*http.Request) {
		comments := []Comment{}

		err := db.Select(&comments, "SELECT * FROM comment;")
		handleServerErrorHTTPError(w, err)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(comments)
	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r*http.Request) {
		id := r.PathValue("id")
		comment := Comment{}
		query := fmt.Sprintf("SELECT * FROM comment WHERE id = %s LIMIT 1;", id)
		err := db.Get(&comment, query)
		handleServerErrorHTTPError(w, err)
		

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(comment)
	})
	
	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r*http.Request) {
		var commentData CommentCreate
		var err error
		err = json.NewDecoder(r.Body).Decode(&commentData)
		if err != nil {
			handleBadRequestHTTPError(w, err)
		}
		query := fmt.Sprintf("INSERT INTO comment (username, timestamp, content) VALUES ('%s', '%s', '%s') RETURNING id;", commentData.Username, time.Now().Format(time.RFC3339), commentData.Content)
		id := 0
		err = db.QueryRow(query).Scan(&id)
		handleServerErrorHTTPError(w, err)
		responseBody := fmt.Sprintf(`{"Id": %d}`, id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(responseBody)
	})

	mux.HandleFunc("PUT /comment/{id}", func(w http.ResponseWriter, r*http.Request) {
		id := r.PathValue("id")
		var commentData CommentUpdate
		var err error
		err = json.NewDecoder(r.Body).Decode(&commentData)
		if err != nil {
			handleBadRequestHTTPError(w, err)
		}
		query := fmt.Sprintf("UPDATE comment SET content = '%s' WHERE id = %s RETURNING *", commentData.Content, id)
		updatedComment := Comment{}
		err = db.QueryRowx(query).StructScan(&updatedComment)
		handleServerErrorHTTPError(w, err)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedComment)
	})

	mux.HandleFunc("DELETE /comment/{id}", func(w http.ResponseWriter, r*http.Request) {
		id := r.PathValue("id")
		query := fmt.Sprintf("DELETE FROM comment WHERE id = %s", id)
		_, err = db.Query(query)
		handleServerErrorHTTPError(w, err)
		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	if err:= http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println((err.Error()))
	}
}

func handleServerErrorHTTPError(w http.ResponseWriter, err error) {
	if err != nil && err != pgx.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	
}

func handleBadRequestHTTPError(w http.ResponseWriter, err error) {
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	
}