-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email, password, avatar_name, bio)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;


-- name: CountRoles :many
SELECT count(*) FROM roles;


-- name: CreateRole :one
INSERT INTO roles (name) VALUES ($1)
RETURNING *;


-- name: AssignUserRole :one
INSERT INTO user_roles (user_id, role_id)
VALUES ($1, (SELECT id FROM roles WHERE name = $2))
RETURNING *;