-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  last_login TIMESTAMP,
  is_active BOOLEAN,
  bio TEXT,
  role BIGINT FOREIGN KEY REFERENCES roles(id)
);


SET @indexCount := (SELECT count(*) FROM information_schema.statistics WHERE table_name = 'users' AND index_name = 'users_email_idx' AND table_schema = DATABASE());
SET @createIndex := IF( @indexCount > 0, 'SELECT ''Index exists.''', 'ALTER TABLE users ADD INDEX users_email_idx (email);');
PREPARE stmt FROM @createIndex;
EXECUTE stmt;

-- +goose Down
DROP TABLE IF EXISTS users;