-- +goose Up
CREATE TABLE IF NOT EXISTS courses (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL, 
  code VARCHAR(50) NOT NULL, 
  description TEXT NOT NULL, 
  image TEXT NOT NULL, 
  category_id BIGINT FOREIGN KEY REFERENCES course_categories(id), 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS courses;