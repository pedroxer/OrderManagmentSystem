-- name: CreateGood :one
INSERT INTO goods (name, category_id, price, description, weight) VALUES ($1, $2, $3, $4,$5) RETURNING id;
