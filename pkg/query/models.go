// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package query

import (
	"database/sql"
)

type Course struct {
	Code       string
	Title      string
	Branch     string
	Semester   int16
	Credits    []int16
	Prereq     []string
	Type       string
	Objectives string
	Content    string
	Books      string
	Outcomes   string
}

type Group struct {
	Name        string
	Alias       sql.NullString
	Branch      sql.NullString
	Kind        sql.NullString
	Description sql.NullString
}

type GroupAdmin struct {
	GroupName  sql.NullString
	Position   sql.NullString
	RollNumber sql.NullInt32
}

type GroupDiscord struct {
	Name          sql.NullString
	ID            sql.NullInt64
	Invite        sql.NullString
	FresherRole   sql.NullInt64
	SophomoreRole sql.NullInt64
	JuniorRole    sql.NullInt64
	SeniorRole    sql.NullInt64
	GuestRole     sql.NullInt64
}

type GroupDiscordUser struct {
	Batch         int16
	DiscordUid    sql.NullInt64
	Name          string
	Alias         sql.NullString
	ID            sql.NullInt64
	Invite        sql.NullString
	FresherRole   sql.NullInt64
	SophomoreRole sql.NullInt64
	JuniorRole    sql.NullInt64
	SeniorRole    sql.NullInt64
	GuestRole     sql.NullInt64
}

type GroupFaculty struct {
	GroupName string
	Name      string
	Mobile    sql.NullInt64
}

type GroupMember struct {
	RollNumber int32
	GroupName  string
}

type GroupSocial struct {
	Name sql.NullString
	Type sql.NullString
	Link sql.NullString
}

type Hostel struct {
	Number     string
	Name       sql.NullString
	WardenName sql.NullString
}

type Student struct {
	RollNumber   int32
	Section      string
	SubSection   string
	Name         string
	Gender       sql.NullString
	Mobile       sql.NullString
	Birthday     sql.NullTime
	Email        string
	Batch        int16
	HostelNumber sql.NullString
	RoomNumber   sql.NullString
	DiscordUid   sql.NullInt64
	Verified     bool
}
