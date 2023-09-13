package resources

import (
	"github.com/raspiantoro/go-type-driven/auth/privileges"
	"github.com/raspiantoro/go-type-driven/auth/roles"
)

type File[
	WP privileges.Write[WR],
	WR roles.Roles,
	RP privileges.Read[RR],
	RR roles.Roles,
	DP privileges.Delete[DR],
	DR roles.Roles,
	T string,
] struct {
	write  WP
	read   RP
	delete DP
	Name   T
}

func New[
	WP privileges.Write[WR],
	WR roles.Roles,
	RP privileges.Read[RR],
	RR roles.Roles,
	DP privileges.Delete[DR],
	DR roles.Roles,
	T string,
](wp WP, rp RP, dp DP, name T) *File[WP, WR, RP, RR, DP, DR, T] {
	return &File[WP, WR, RP, RR, DP, DR, T]{
		write:  wp,
		read:   rp,
		delete: dp,
		Name:   name,
	}
}
