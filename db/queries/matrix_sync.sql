-- name: MatrixSaveFilterID :exec
INSERT INTO matrix_filters (user_id, filter_id) VALUES (?, ?) ON CONFLICT (user_id) DO UPDATE SET filter_id = excluded.filter_id;

-- name: MatrixLoadFilterID :one
SELECT filter_id FROM matrix_filters WHERE user_id = ?;

-- name: MatrixSaveNextBatch :exec
INSERT INTO matrix_next_batch (user_id, token) VALUES (?, ?) ON CONFLICT (user_id) DO UPDATE SET token = excluded.token;

-- name: MatrixLoadNextBatch :one
SELECT token FROM matrix_next_batch WHERE user_id = ?;
