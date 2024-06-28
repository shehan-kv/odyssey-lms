-- +goose Up
CREATE TABLE IF NOT EXISTS course_sections_complete (
  user_id INTEGER NOT NULL,
  section_id INTEGER NOT NULL,
  PRIMARY KEY (user_id, section_id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (section_id) REFERENCES course_sections(id)
);

-- +goose Down
DROP TABLE IF EXISTS course_enroll;