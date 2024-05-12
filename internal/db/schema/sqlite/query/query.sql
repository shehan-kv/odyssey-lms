-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email, password, avatar_name, bio)
VALUES (?,?,?,?,?,?)
RETURNING *;


-- name: CountRoles :many
SELECT count(*) FROM roles;


-- name: CreateRole :one
INSERT INTO roles (name) VALUES (?)
RETURNING *;


-- name: AssignUserRole :one
INSERT INTO user_roles (user_id, role_id)
VALUES (?, (SELECT id FROM roles WHERE name = ?))
RETURNING *;