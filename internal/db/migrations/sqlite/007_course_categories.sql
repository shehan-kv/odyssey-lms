-- +goose Up
CREATE TABLE IF NOT EXISTS course_categories (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS course_categories;