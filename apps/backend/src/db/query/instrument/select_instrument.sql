-- name: GetInstrument :one
SELECT id, name, type, artist_id, description, tags
FROM m_instrument
WHERE id = $1;

-- name: ListInstruments :many
SELECT id, name, type, artist_id, description, tags
FROM m_instrument
ORDER BY name;
