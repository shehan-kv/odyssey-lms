-- +goose Up
CREATE TABLE IF NOT EXISTS course_categories (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS course_categories;