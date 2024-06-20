-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  last_login DATETIME,
  bio TEXT
);

CREATE UNIQUE INDEX IF NOT EXISTS users_email_idx ON users(email);

-- +goose Down
DROP TABLE IF EXISTS users;