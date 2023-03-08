package login

import (
)
//	"context"

//	"github.com/gofiber/fiber"
//	"github.com/gofiber/fiber/v2"

type User struct {
	ClientId   string `json:"clientId"`
	Credential string `json:"credential"`
}

func SignIn(googleUser) error {
	var profile := googleUser.getBasicProfile()

}

//func SignIn(c *fiber.Ctx) error {

//	coll := googleUser.getBasicProfile("https://apis.google.com/js/platform.js").Content("226207107235-qiemrp4uajoakgvu2gi4obckqmgvdg9c.apps.googleusercontent.com")
//	profile := coll.Find(context.TODO())

//	email := (("email") + profile.getEmail())
//	coll.Where("email = ?", profile.getEmail).First(&email)
//	if email.User == "" {
//		return c.SendString("usuario no habilitado")
//	}

//	nombre := (("nombre") + profile.getNombre())
//	apellido := (("apellido") + profile.getApellido())

//	user := (("user") + profile.getUser())

//	token = googleUser.getAuthResponse()
//	console.log(token)
//	console.log(token)

//	return (profile)
//}
