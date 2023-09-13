package users

import (
	"github.com/raspiantoro/go-type-driven/auth/roles"
)

type User[T roles.Roles] struct {
	Name string
}
