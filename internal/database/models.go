// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Createdat time.Time
	Updatedat time.Time
	Name      string
}
