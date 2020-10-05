CREATE TABLE todos (
    id    SERIAL PRIMARY KEY,
    content varchar(255) NOT NULL CHECK (content <> ''),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
