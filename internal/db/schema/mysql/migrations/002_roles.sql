-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL UNIQUE,
  is_default BOOLEAN NOT NULL
);

INSERT INTO roles (name, is_default) VALUES ('administrator', true);
INSERT INTO roles (name, is_default) VALUES ('instructor', true);
INSERT INTO roles (name, is_default) VALUES ('student', true);

-- +goose Down
DROP TABLE IF EXISTS roles;