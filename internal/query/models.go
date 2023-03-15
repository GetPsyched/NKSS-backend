// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package query

import (
	"database/sql"
	"encoding/json"
	"time"
)

type AcademicAnnouncement struct {
	DateOfCreation time.Time      `json:"date_of_creation"`
	Title          string         `json:"title"`
	TitleLink      string         `json:"title_link"`
	Kind           string         `json:"kind"`
	Subtitle       sql.NullString `json:"subtitle"`
	SubtitleLink   sql.NullString `json:"subtitle_link"`
}

type AffiliatedGuild struct {
	GuildID        int64 `json:"guild_id"`
	Batch          int16 `json:"batch"`
	InfoChannel    int64 `json:"info_channel"`
	CommandChannel int64 `json:"command_channel"`
	GuestRole      int64 `json:"guest_role"`
}

type Book struct {
	BookName    string         `json:"book_name"`
	AuthorNames []string       `json:"author_names"`
	Publisher   string         `json:"publisher"`
	Edition     int16          `json:"edition"`
	Url         sql.NullString `json:"url"`
}

type BotPrefix struct {
	GuildID int64  `json:"guild_id"`
	Prefix  string `json:"prefix"`
}

type BranchSpecific struct {
	Code     string  `json:"code"`
	Branch   string  `json:"branch"`
	Semester int16   `json:"semester"`
	Credits  []int16 `json:"credits"`
}

type Club struct {
	Name             string         `json:"name"`
	Alias            sql.NullString `json:"alias"`
	Category         string         `json:"category"`
	Email            string         `json:"email"`
	ShortDescription string         `json:"short_description"`
	IsOfficial       bool           `json:"is_official"`
}

type ClubDetail struct {
	ClubName     string   `json:"club_name"`
	AboutUs      string   `json:"about_us"`
	WhyUs        string   `json:"why_us"`
	RoleOfSoph   string   `json:"role_of_soph"`
	RoleOfJunior string   `json:"role_of_junior"`
	RoleOfSenior string   `json:"role_of_senior"`
	Admins       []string `json:"admins"`
	Branch       []string `json:"branch"`
}

type ClubDiscord struct {
	ClubName      string `json:"club_name"`
	GuildID       int64  `json:"guild_id"`
	FreshmanRole  int64  `json:"freshman_role"`
	SophomoreRole int64  `json:"sophomore_role"`
	JuniorRole    int64  `json:"junior_role"`
	SeniorRole    int64  `json:"senior_role"`
	GuestRole     int64  `json:"guest_role"`
}

type ClubDiscordUser struct {
	Batch         int16          `json:"batch"`
	DiscordID     sql.NullInt64  `json:"discord_id"`
	Name          string         `json:"name"`
	Alias         sql.NullString `json:"alias"`
	GuildID       int64          `json:"guild_id"`
	GuildInvite   string         `json:"guild_invite"`
	FreshmanRole  int64          `json:"freshman_role"`
	SophomoreRole int64          `json:"sophomore_role"`
	JuniorRole    int64          `json:"junior_role"`
	SeniorRole    int64          `json:"senior_role"`
	GuestRole     int64          `json:"guest_role"`
}

type ClubFaculty struct {
	ClubName string `json:"club_name"`
	EmpID    int32  `json:"emp_id"`
}

type ClubMember struct {
	ClubName   string `json:"club_name"`
	RollNumber string `json:"roll_number"`
}

type ClubSocial struct {
	ClubName     string `json:"club_name"`
	PlatformType string `json:"platform_type"`
	Link         string `json:"link"`
}

type Course struct {
	Code       string   `json:"code"`
	Title      string   `json:"title"`
	Prereq     []string `json:"prereq"`
	Kind       string   `json:"kind"`
	Objectives []string `json:"objectives"`
	Content    []string `json:"content"`
	BookNames  []string `json:"book_names"`
	Outcomes   []string `json:"outcomes"`
}

type Event struct {
	GuildID    int64          `json:"guild_id"`
	EventTypes []string       `json:"event_types"`
	ChannelID  int64          `json:"channel_id"`
	Message    sql.NullString `json:"message"`
}

type Faculty struct {
	EmpID          int32          `json:"emp_id"`
	Name           string         `json:"name"`
	DpUrl          string         `json:"dp_url"`
	Designation    string         `json:"designation"`
	Qualification  string         `json:"qualification"`
	AreaOfInterest []string       `json:"area_of_interest"`
	Experience     []string       `json:"experience"`
	Mobile         string         `json:"mobile"`
	Mobile2        sql.NullString `json:"mobile_2"`
	Email          string         `json:"email"`
	Department     []string       `json:"department"`
	IsRegular      bool           `json:"is_regular"`
}

type Guild struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Locale    string        `json:"locale"`
	BotRole   int64         `json:"bot_role"`
	MuteRole  sql.NullInt64 `json:"mute_role"`
	EditLog   sql.NullInt64 `json:"edit_log"`
	DeleteLog sql.NullInt64 `json:"delete_log"`
}

type GuildRole struct {
	GuildID int64   `json:"guild_id"`
	Field   string  `json:"field"`
	Value   string  `json:"value"`
	RoleIds []int64 `json:"role_ids"`
}

type Hod struct {
	Department string `json:"department"`
	EmpID      int32  `json:"emp_id"`
}

type Hostel struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type JoinRole struct {
	GuildID int64 `json:"guild_id"`
	RoleID  int64 `json:"role_id"`
}

type ModRole struct {
	GuildID int64 `json:"guild_id"`
	RoleID  int64 `json:"role_id"`
}

type Student struct {
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
	IsVerified bool            `json:"is_verified"`
	Clubs      json.RawMessage `json:"clubs"`
}

type Warden struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	HostelID string `json:"hostel_id"`
}
