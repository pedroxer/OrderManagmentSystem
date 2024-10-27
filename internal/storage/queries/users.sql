-- name: SaveUser :one
INSERT INTO users (username, email, password_hash, role) VALUES ($1, $2, $3, 0) RETURNING user_id;

-- name: GetUser :one
SELECT * FROM users WHERE email = $1;