package main

import (
	"github.com/birwin93/hippo"
	"github.com/birwin93/hippo/demo/handlers/sessions"
	"github.com/birwin93/hippo/demo/handlers/users"
)

func main() {
	app := hippo.NewApp()

	app.Add("/login", &sessions.LoginHandler{})
	app.Add("/logout", &sessions.LogoutHandler{})

	app.Add("/users", &users.FindByIdHandler{})

	app.Start("127.0.0.1", "3000")
}
