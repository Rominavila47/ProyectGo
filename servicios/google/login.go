package login

import "github.com/gofiber/fiber/v2"

//	"context"

//	"github.com/gofiber/fiber"
//	"github.com/gofiber/fiber/v2"

type User struct {
	User       string `bson:"user"`
	ClientId   string `bson:"clientId"`
	Credential string `bson:"credential"`
}

func SignIn(c *fiber.Ctx) error {
	User := new(User)
	if err := c.BodyParser(User); err != nil {
		return c.Status(503).SendString(err.Error())
	}

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
