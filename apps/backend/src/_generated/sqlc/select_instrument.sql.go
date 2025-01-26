// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: select_instrument.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getInstrument = `-- name: GetInstrument :one
SELECT id, name, type, artist_id, description, tags
FROM m_instrument
WHERE id = $1
`

type GetInstrumentRow struct {
	ID          pgtype.UUID
	Name        string
	Type        pgtype.Text
	ArtistID    pgtype.UUID
	Description pgtype.Text
	Tags        []string
}

func (q *Queries) GetInstrument(ctx context.Context, dollar_1 pgtype.UUID) (GetInstrumentRow, error) {
	row := q.db.QueryRow(ctx, getInstrument, dollar_1)
	var i GetInstrumentRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.ArtistID,
		&i.Description,
		&i.Tags,
	)
	return i, err
}

const listInstruments = `-- name: ListInstruments :many
SELECT id, name, type, artist_id, description, tags
FROM m_instrument
ORDER BY name
`

type ListInstrumentsRow struct {
	ID          pgtype.UUID
	Name        string
	Type        pgtype.Text
	ArtistID    pgtype.UUID
	Description pgtype.Text
	Tags        []string
}

func (q *Queries) ListInstruments(ctx context.Context) ([]ListInstrumentsRow, error) {
	rows, err := q.db.Query(ctx, listInstruments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListInstrumentsRow{}
	for rows.Next() {
		var i ListInstrumentsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.ArtistID,
			&i.Description,
			&i.Tags,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
