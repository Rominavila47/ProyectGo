package login

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	TokenId int    `json:"tokenId" gorm:"primaryKey"`
	Email   string `json:"email"`
	User    string `json:"user"`
}

func onSignIn(c *fiber.Ctx) error {
	profile := googleUser.getBasicProfile()
	console := googleUser()
	console("ID: " + profile.getId())
	console("Name: " + profile.getName())
	console("Email: " + profile.getEmail())

	id_token := googleUser.getAuthResponse().id_token
	console(id_token)

	fmt.Println(onSignIn)
}
