-- +goose Up
CREATE TABLE IF NOT EXISTS courses (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL, 
  code TEXT NOT NULL, 
  description TEXT NOT NULL, 
  image TEXT NOT NULL, 
  category_id INTEGER NOT NULL, 
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (category_id) REFERENCES course_categories(id)
);

-- +goose Down
DROP TABLE IF EXISTS courses;