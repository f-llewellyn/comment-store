-- create comment table
CREATE TABLE IF NOT EXISTS COMMENT (
    id SERIAL PRIMARY KEY,
    username VARCHAR,
    timestamp TIMESTAMPTZ DEFAULT NOW(),
    content VARCHAR 
);

-- insert default comments
INSERT INTO COMMENT (username, content)
VALUES 
    ('f-llewellyn', 'Hello World!'),
    ('otherUzr', 'This is another comment');