-- +goose Up
CREATE TABLE IF NOT EXISTS user_roles (
  user_id INTEGER NOT NULL,
  role_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (role_id) REFERENCES roles(id)
);


-- +goose Down
DROP TABLE IF EXISTS user_roles;