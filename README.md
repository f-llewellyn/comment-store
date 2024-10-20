# Comment Store

This is a very basic rest API written in go as a learning exercise.

## Getting Started

Before getting started, ensure you have golang and docker installed on your machine.

To start the server, run the following in the project root directory:

```
go mod download
go run main.go
```

If you have go installed on your machine, just run `air` in the root project directory.

## Endpoints

All endpoints are codumented in a bruno collection whcih can be found in the `api-docs` directory.

### GET /comment

_Fetches all comments_

```
// Example response
[
  {
    "Id": 2,
    "Username": "otherUzr",
    "Timestamp": "2024-10-11T16:28:52.85517+01:00",
    "Content": "This is another comment"
  },
  {
    "Id": 3,
    "Username": "Finn 2",
    "Timestamp": "2024-10-17T22:51:32+01:00",
    "Content": "Hello World!"
  },
  {
    "Id": 4,
    "Username": "Finn 4",
    "Timestamp": "2024-10-17T22:52:04+01:00",
    "Content": "Hello World!"
  },
]
```

### GET /comment/{id}

_Fetches comment with specified id_

```
// Example response
{
  "Id": 2,
  "Username": "otherUzr",
  "Timestamp": "2024-10-11T16:28:52.85517+01:00",
  "Content": "This is another comment"
}
```

### POST /comment

_Creates comment with specified data_

```
// Example request
{
  "Username": "User 1",
  "Content": "This is an example comment"
}

// Example response
{
  "Id": 32,
}
```

### PUT /comment/{id}

_Updates comment with specified id_

```
// Example request
{
  "Content": "This is an updated comment"
}

// Example response
{
  "Id": 32,
  "Username: "Finn",
  "Content": "This is an updated comment",
  "Timestamp": "2024-10-20T14:30:54+01:00"
}
```

### DELETE /comment/{id}

_Deletes comment with specified id_
