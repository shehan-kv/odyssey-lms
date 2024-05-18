-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE,
  is_default BOOLEAN NOT NULL
);

INSERT INTO roles (name, is_default) VALUES ('administrator', true);
INSERT INTO roles (name, is_default) VALUES ('instructor', true);
INSERT INTO roles (name, is_default) VALUES ('student', true);

-- +goose Down
DROP TABLE IF EXISTS roles;