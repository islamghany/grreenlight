// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddPermissionForUser(ctx context.Context, arg AddPermissionForUserParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	DeleteAllSessionForUser(ctx context.Context, userID int64) error
	DeleteSession(ctx context.Context, id uuid.UUID) error
	GetAllPermissionsForUser(ctx context.Context, id int64) ([]string, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	InsertSession(ctx context.Context, arg InsertSessionParams) (Session, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
