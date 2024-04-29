-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  avatar_name TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  bio TEXT
);


-- +goose Down
DROP TABLE IF EXISTS users;