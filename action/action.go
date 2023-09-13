package action

import (
	"fmt"

	"github.com/raspiantoro/go-type-driven/auth/privileges"
	"github.com/raspiantoro/go-type-driven/auth/roles"
	"github.com/raspiantoro/go-type-driven/auth/users"
	"github.com/raspiantoro/go-type-driven/resources"
)

func Write[
	WP privileges.Write[WR],
	WR roles.Roles,
	RP privileges.Read[RR],
	RR roles.Roles,
	DP privileges.Delete[DR],
	DR roles.Roles,
	T string,
](user users.User[WR], file *resources.File[WP, WR, RP, RR, DP, DR, T]) {
	fmt.Printf("%s write on file %s\n", user.Name, file.Name)
}

func Read[
	WP privileges.Write[WR],
	WR roles.Roles,
	RP privileges.Read[RR],
	RR roles.Roles,
	DP privileges.Delete[DR],
	DR roles.Roles,
	T string,
](user users.User[RR], file *resources.File[WP, WR, RP, RR, DP, DR, T]) {
	fmt.Printf("%s read file %s\n", user.Name, file.Name)
}

func Delete[
	WP privileges.Write[WR],
	WR roles.Roles,
	RP privileges.Read[RR],
	RR roles.Roles,
	DP privileges.Delete[DR],
	DR roles.Roles,
	T string,
](user users.User[DR], file *resources.File[WP, WR, RP, RR, DP, DR, T]) {
	fmt.Printf("%s deleting file %s\n", user.Name, file.Name)
}
