-- +goose Up
CREATE TABLE IF NOT EXISTS course_enroll (
  user_id INTEGER NOT NULL,
  course_id INTEGER NOT NULL,
  PRIMARY KEY (user_id, course_id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (course_id) REFERENCES courses(id)
);

-- +goose Down
DROP TABLE IF EXISTS course_enroll;