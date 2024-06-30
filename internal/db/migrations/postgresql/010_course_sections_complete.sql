-- +goose Up
CREATE TABLE IF NOT EXISTS course_sections_complete (
  user_id BIGINT FOREIGN KEY REFERENCES users(id),
  section_id BIGINT FOREIGN KEY REFERENCES course_sections(id),
  PRIMARY KEY (user_id, section_id)
);

-- +goose Down
DROP TABLE IF EXISTS course_enroll;