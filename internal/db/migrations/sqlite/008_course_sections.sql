-- +goose Up
CREATE TABLE IF NOT EXISTS course_sections (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  course_id INTEGER NOT NULL,
  FOREIGN KEY (course_id) REFERENCES courses(id)
);

-- +goose Down
DROP TABLE IF EXISTS course_sections;