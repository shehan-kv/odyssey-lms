-- +goose Up
CREATE TABLE IF NOT EXISTS course_sections (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  course_id BIGINT FOREIGN KEY REFERENCES courses(id)
);

-- +goose Down
DROP TABLE IF EXISTS course_sections;