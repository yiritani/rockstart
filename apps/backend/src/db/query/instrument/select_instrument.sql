-- name: GetInstrument :many
SELECT id, name, type, artist_id, description, tags
FROM instrument
WHERE type = ?
ORDER BY id ASC;
