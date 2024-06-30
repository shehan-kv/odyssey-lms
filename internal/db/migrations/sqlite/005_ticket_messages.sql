-- +goose Up
CREATE TABLE IF NOT EXISTS ticket_messages (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  ticket_id INTEGER NOT NULL, 
  user_id INTEGER NOT NULL, 
  content TEXT NOT NULL,  
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (ticket_id) REFERENCES tickets(id)
);


-- +goose Down
DROP TABLE IF EXISTS ticket_messages;