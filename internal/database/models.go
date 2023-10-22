// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID
	Createdat time.Time
	Updatedat time.Time
	Name      string
	Url       string
	Userid    uuid.UUID
}

type FeedFollow struct {
	ID        uuid.UUID
	Createdat time.Time
	Updatedat time.Time
	Userid    uuid.UUID
	Feedid    uuid.UUID
}

type User struct {
	ID        uuid.UUID
	Createdat time.Time
	Updatedat time.Time
	Name      string
	ApiKey    string
}
