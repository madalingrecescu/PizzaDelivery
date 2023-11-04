// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package users_db

import (
	"database/sql"
)

type User struct {
	UserID         int32          `json:"userId"`
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	HashedPassword string         `json:"hashedPassword"`
	PhoneNumber    string         `json:"phoneNumber"`
	Role           sql.NullString `json:"role"`
}
