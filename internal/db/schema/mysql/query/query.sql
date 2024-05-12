-- name: CreateUser :execresult
INSERT INTO users (first_name, last_name, email, password, avatar_name, bio)
VALUES (?,?,?,?,?,?);


-- name: CountRoles :many
SELECT count(*) FROM roles;


-- name: CreateRole :execresult
INSERT INTO roles (name) VALUES (?);


-- name: AssignUserRole :execresult
INSERT INTO user_roles (user_id, role_id)
VALUES (?, (SELECT id FROM roles WHERE name = ?));