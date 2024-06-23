-- +goose Up
CREATE TABLE IF NOT EXISTS ticket_messages (
  id BIGSERIAL PRIMARY KEY,
  ticket_id BIGINT FOREIGN KEY REFERENCES tickets(id), 
  user_id BIGINT FOREIGN KEY REFERENCES users(id),
  content TEXT NOT NULL,  
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


-- +goose Down
DROP TABLE IF EXISTS ticket_messages;