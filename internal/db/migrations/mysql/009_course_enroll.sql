-- +goose Up
CREATE TABLE IF NOT EXISTS course_enroll (
  user_id BIGINT FOREIGN KEY REFERENCES users(id),
  course_id BIGINT FOREIGN KEY REFERENCES courses(id),
  PRIMARY KEY (user_id, course_id)
);

-- +goose Down
DROP TABLE IF EXISTS course_enroll;