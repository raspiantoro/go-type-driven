package privileges

import "github.com/raspiantoro/go-type-driven/auth/roles"

type Read[R roles.Roles] struct{}
type Write[R roles.Roles] struct{}
type Delete[R roles.Roles] struct{}

type Privileges[R roles.Roles] interface {
	Read[R] | Write[R] | Delete[R]
}
