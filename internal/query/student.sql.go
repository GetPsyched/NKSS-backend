// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: student.sql

package query

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/lib/pq"
)

const getDiscordLinkStatus = `-- name: GetDiscordLinkStatus :one
SELECT is_verified FROM student WHERE discord_id = $1
`

func (q *Queries) GetDiscordLinkStatus(ctx context.Context, discordID sql.NullInt64) (bool, error) {
	row := q.db.QueryRowContext(ctx, getDiscordLinkStatus, discordID)
	var is_verified bool
	err := row.Scan(&is_verified)
	return is_verified, err
}

const getHostels = `-- name: GetHostels :many
SELECT hostel.id, hostel.name, hostel.email, JSON_AGG(JSON_BUILD_OBJECT('name', warden.name, 'mobile', warden.mobile)) AS "wardens"
FROM hostel
LEFT JOIN warden ON warden.hostel_id = hostel.id
GROUP BY hostel.id
`

type GetHostelsRow struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Wardens json.RawMessage `json:"wardens"`
}

func (q *Queries) GetHostels(ctx context.Context) ([]GetHostelsRow, error) {
	rows, err := q.db.QueryContext(ctx, getHostels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetHostelsRow
	for rows.Next() {
		var i GetHostelsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Wardens,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStudent = `-- name: GetStudent :one
SELECT
    roll_number, section, name, gender, mobile, birth_date, email, batch, hostel_id, room_id, discord_id, clubs, is_verified, (
        SELECT
            JSON_AGG(JSON_BUILD_OBJECT(
                'name', cm.club_name, 'position',
                COALESCE((SELECT position FROM club_admin WHERE roll_number = cm.roll_number), 'Member')
            ))
        FROM
            club_member AS cm
        WHERE
            cm.roll_number = $1
    ) AS clubs
FROM
    student
WHERE roll_number = $1
`

type GetStudentRow struct {
	RollNumber string          `json:"roll_number"`
	Section    string          `json:"section"`
	Name       string          `json:"name"`
	Gender     sql.NullString  `json:"gender"`
	Mobile     sql.NullString  `json:"mobile"`
	BirthDate  sql.NullTime    `json:"birth_date"`
	Email      string          `json:"email"`
	Batch      int16           `json:"batch"`
	HostelID   string          `json:"hostel_id"`
	RoomID     sql.NullString  `json:"room_id"`
	DiscordID  sql.NullInt64   `json:"discord_id"`
	Clubs      json.RawMessage `json:"clubs"`
	IsVerified bool            `json:"is_verified"`
	Clubs_2    json.RawMessage `json:"clubs_2"`
}

func (q *Queries) GetStudent(ctx context.Context, rollNumber string) (GetStudentRow, error) {
	row := q.db.QueryRowContext(ctx, getStudent, rollNumber)
	var i GetStudentRow
	err := row.Scan(
		&i.RollNumber,
		&i.Section,
		&i.Name,
		&i.Gender,
		&i.Mobile,
		&i.BirthDate,
		&i.Email,
		&i.Batch,
		&i.HostelID,
		&i.RoomID,
		&i.DiscordID,
		&i.Clubs,
		&i.IsVerified,
		&i.Clubs_2,
	)
	return i, err
}

const getStudentByDiscordID = `-- name: GetStudentByDiscordID :one
SELECT
    roll_number, section, name, gender, mobile, birth_date, email, batch, hostel_id, room_id, discord_id, clubs, is_verified,
    CAST(ARRAY(SELECT club.alias FROM club JOIN club_member AS cm ON cm.club_name = club.name WHERE cm.roll_number = s.roll_number) AS VARCHAR[]) AS clubs
FROM
    student AS s
WHERE discord_id = $1
`

type GetStudentByDiscordIDRow struct {
	RollNumber string          `json:"roll_number"`
	Section    string          `json:"section"`
	Name       string          `json:"name"`
	Gender     sql.NullString  `json:"gender"`
	Mobile     sql.NullString  `json:"mobile"`
	BirthDate  sql.NullTime    `json:"birth_date"`
	Email      string          `json:"email"`
	Batch      int16           `json:"batch"`
	HostelID   string          `json:"hostel_id"`
	RoomID     sql.NullString  `json:"room_id"`
	DiscordID  sql.NullInt64   `json:"discord_id"`
	Clubs      json.RawMessage `json:"clubs"`
	IsVerified bool            `json:"is_verified"`
	Clubs_2    []string        `json:"clubs_2"`
}

func (q *Queries) GetStudentByDiscordID(ctx context.Context, discordID sql.NullInt64) (GetStudentByDiscordIDRow, error) {
	row := q.db.QueryRowContext(ctx, getStudentByDiscordID, discordID)
	var i GetStudentByDiscordIDRow
	err := row.Scan(
		&i.RollNumber,
		&i.Section,
		&i.Name,
		&i.Gender,
		&i.Mobile,
		&i.BirthDate,
		&i.Email,
		&i.Batch,
		&i.HostelID,
		&i.RoomID,
		&i.DiscordID,
		&i.Clubs,
		&i.IsVerified,
		pq.Array(&i.Clubs_2),
	)
	return i, err
}
