// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	UserIp       string    `json:"user_ip"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID                int64     `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	HashedPassword    []byte    `json:"hashed_password"`
	Activated         bool      `json:"activated"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	Version           int32     `json:"version"`
}

type UsersPermission struct {
	UserID       int64 `json:"user_id"`
	PermissionID int64 `json:"permission_id"`
}
