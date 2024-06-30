-- +goose Up
CREATE TABLE IF NOT EXISTS tickets (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  subject TEXT NOT NULL, 
  description TEXT NOT NULL, 
  user_id INTEGER NOT NULL, 
  type TEXT NOT NULL, 
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  closed_at DATETIME,
  status TEXT,
  FOREIGN KEY (user_id) REFERENCES users(id)
);


-- +goose Down
DROP TABLE IF EXISTS tickets;