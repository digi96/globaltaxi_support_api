-- name: CreatePayload :one
INSERT INTO payload(body, created_at)
VALUES ($1, $2)
RETURNING *;
-- name: GetPaylodById :one
SELECT *
FROM payload
WHERE payload_id = $1
LIMIT 1;
-- name: ListUndoPayloads :many
SELECT *
FROM payload
WHERE status = '0'
ORDER BY payload_id
LIMIT $1 OFFSET $2;
-- name: UpdatePayload :one
UPDATE payload
SET status = '1'
WHERE payload_id = sqlc.arg('payload_id')
RETURNING *;