package login

import (
	"context"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Token    int    `json:"token"`
	User     string `json:"user"`
	Email    string `json:"email"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

func SignIn(c *fiber.Ctx, googleUser) error {
	coll := googleUser.getBasicProfile("https://apis.google.com/js/platform.js").Content("226207107235-qiemrp4uajoakgvu2gi4obckqmgvdg9c.apps.googleusercontent.com")
	profile := coll.Find(context.TODO())

	email := (("email") + profile.getEmail())
	coll.Where("email = ?", profile.getEmail).First(&user)
	if email.User == "" {
		return c.SendString("usuario no habilitado")
	}

	nombre := (("nombre") + profile.getNombre())
	apellido := (("apellido") + profile.getApellido())

	user := (("user") + profile.getUser())

	token = googleUser.getAuthResponse()
	console.log(token)

	return (profile)
}
