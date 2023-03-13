package login

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var store *session.Store = session.New()
var dbUser *gorm.DB
var client *mongo.Client

func ConnectMongoDb(clientMongo *mongo.Client) {
	client = clientMongo
}

type User struct {
	ClientId   string `bson:"clientId"`
	Credential string `bson:"credential"`
}

func SignIn(c *fiber.Ctx) error {
	newUser := new(User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	coll := client.Database("portalDeNovedades").Collection("login")
	clientId, _ := strconv.Atoi(c.Params("id"))
	cursor, err := coll.Find(context.TODO(), bson.M{"clientId": clientId})

	var user []User
	if err = cursor.All(context.Background(), &user); err != nil {
		fmt.Print(err)
	}

	if newUser.ClientId == "" {
		return c.SendString("usuario no habilitado")
	}

	return c.JSON(user)
}

//func Login(c *fiber.Ctx) error {
//	user, err := gothic.CompleteUserAuth(c.Context(), c)
//	if err != nil {
//		return c.SendStatus(fiber.StatusInternalServerError)
//	}
//	return c.JSON(user)
//}
