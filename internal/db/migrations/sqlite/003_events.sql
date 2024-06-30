-- +goose Up
CREATE TABLE IF NOT EXISTS events (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  type TEXT, 
  description TEXT,
  severity TEXT
);


-- +goose Down
DROP TABLE IF EXISTS events;