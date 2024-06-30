-- +goose Up
CREATE TABLE IF NOT EXISTS ticket_messages (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  ticket_id BIGINT FOREIGN KEY REFERENCES tickets(id), 
  user_id BIGINT FOREIGN KEY REFERENCES users(id),
  content TEXT NOT NULL,  
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- +goose Down
DROP TABLE IF EXISTS ticket_messages;