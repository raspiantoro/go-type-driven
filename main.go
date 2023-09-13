package main

import (
	"github.com/raspiantoro/go-type-driven/action"
	"github.com/raspiantoro/go-type-driven/auth/privileges"
	"github.com/raspiantoro/go-type-driven/auth/roles"
	"github.com/raspiantoro/go-type-driven/auth/users"
	"github.com/raspiantoro/go-type-driven/resources"
)

func main() {

	file := resources.New(
		privileges.Write[roles.Manager]{},
		privileges.Read[roles.Staff]{},
		privileges.Delete[roles.Admin]{},
		"rahasia.txt")

	admin := users.User[roles.Admin]{Name: "admin"}
	manager := users.User[roles.Manager]{Name: "manager"}
	staff := users.User[roles.Staff]{Name: "staff"}

	action.Write(manager, file)
	action.Read(staff, file)
	action.Delete(admin, file)
}
