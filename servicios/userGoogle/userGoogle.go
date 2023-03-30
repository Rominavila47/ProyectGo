package userGoogle

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func ConnectMongoDb(clientMongo *mongo.Client) {
	client = clientMongo
}

type User struct {
	Email string `bson:"email"`
}

type GoogleClaims struct {
	Email         string `bson:"email"`
	EmailVerified bool   `bson:"email_verified"`
	FirstName     string `bson:"given_name"`
	LastName      string `bson:"family_name"`
	jwt.StandardClaims
}

func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}
	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}

func ValidateGoogleJWT(c *fiber.Ctx) error {
	tokenString := c.Params("tokenString")
	claimsStruct := GoogleClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,

		func(token *jwt.Token) (interface{}, error) {
			pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
			if err != nil {
				return nil, err
			}
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*GoogleClaims)
	if !ok {
		return errors.New("Invalid Google JWT")
	}

	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
		return errors.New("iss is invalid")
	}

	if claims.Audience != os.Getenv("GOOGLEKEY") {
		return errors.New("aud is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return errors.New("JWT is expired")
	}

	coll := client.Database("portalDeNovedades").Collection("login")
	var newUser User

	err = coll.FindOne(context.TODO(), bson.D{{"email", claims.Email}}).Decode(&newUser)
	if err != nil {
		fmt.Print(err)
	}

	return c.SendString("usuario invalido")
}

func InsertEmail(c *fiber.Ctx) error {
	newUser := new(User)
	if err := c.BodyParser(newUser); err != nil {
		return err
	}

	coll := client.Database("portalDeNovedades").Collection("login")

	result, err := coll.InsertOne(context.Background(), newUser)

	if err != nil {
		fmt.Print(result)
	}

	return c.SendString("usuario agregado")
}

func DeleteEmail(c *fiber.Ctx) error {
	coll := client.Database("portalDeNovedades").Collection("login")
	newUser := new(User)
	if err := c.BodyParser(newUser); err != nil {
		return err
	}

	result, err := coll.DeleteOne(context.Background(), bson.M{"email": newUser.Email})
	if err != nil {
		fmt.Print(result)
	}
	fmt.Print("usuario eliminado")
	return c.SendString("usuario eliminado")
}

func MakeJWT(email string, secret string) (string, error) {
	claims := GoogleClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/*
func loginHandler(c *fiber.Ctx) {
	defer r.Body.Close()

	// parse the GoogleJWT that was POSTed from the front-end
	type parameters struct {
		GoogleJWT *string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Couldn't decode parameters")
		return
	}

	// Validate the JWT is valid
	claims, err := auth.ValidateGoogleJWT(*params.GoogleJWT)
	if err != nil {
		respondWithError(w, 403, "Invalid google auth")
		return
	}
	if claims.Email != user.Email {
		respondWithError(w, 403, "Emails don't match")
		return
	}

	// create a JWT for OUR app and give it back to the client for future requests
	tokenString, err := auth.MakeJWT(claims.Email, cfg.JWTSecret)
	if err != nil {
		respondWithError(w, 500, "Couldn't make authentication token")
		return
	}

	respondWithJSON(w, 200, struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	})
}
*/
