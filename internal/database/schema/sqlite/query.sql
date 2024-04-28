-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email, password, avatar_name, bio)
VALUES (?,?,?,?,?,?)
RETURNING *;