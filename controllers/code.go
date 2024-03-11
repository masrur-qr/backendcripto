package controllers

import (
	"context"
	"fmt"
	"mongo/repo/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var SignUpTemp structs.SignUpStruct
	c.ShouldBindJSON(&SignUpTemp)

	if SignUpTemp.Login == "" || SignUpTemp.Surname == "" || SignUpTemp.Email == "" || SignUpTemp.Name == "" || SignUpTemp.Name == "" || SignUpTemp.Password == "" {
		c.JSON(404, "Error")
	} else {
		client, ctx := DBConnection()

		DBConnect := client.Database("MasrurCoins").Collection("Users")

		id := primitive.NewObjectID().Hex()
		Hashed, _ := HashPassword(SignUpTemp.Password)

		DBConnect.InsertOne(ctx, bson.M{
			"_id":      id,
			"name":    SignUpTemp.Login,
			"surname":    SignUpTemp.Login,
			"balance":    SignUpTemp.Login,
			"email":    SignUpTemp.Login,
			"login":    SignUpTemp.Login,
			"password": Hashed,
		})
	}
}
func Login(c *gin.Context) {
	var LoginTemp structs.SignUpStruct
	c.ShouldBindJSON(&LoginTemp)

	if LoginTemp.Login == "" || LoginTemp.Password == "" {
		c.JSON(404, "Error")
	} else {
		client, ctx := DBConnection()

		DBConnect := client.Database("MasDB").Collection("Soos")

		result := DBConnect.FindOne(ctx, bson.M{
			"login": LoginTemp.Login,
		})

		var userdata structs.SignUpStruct
		result.Decode(&userdata)
		isValidPass := CompareHashPasswords(userdata.Password,LoginTemp.Password)

		if isValidPass {
			http.SetCookie(c.Writer, &http.Cookie{
				Name: "CoinCookie",
				Value: userdata.Login,
				Expires: time.Now().Add(60*time.Second),
			})
			c.JSON(200,"success")
		}else{
			c.JSON(404,"Error")
		}
	}
}

// ! =============================    Helpers   ==================================
// ? ============ Connect To DB
func DBConnection() (*mongo.Client, context.Context) {
	url := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	NewCtx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	Client, err := mongo.Connect(NewCtx, url)
	if err != nil {
		fmt.Printf("errors: %v\n", err)
	}
	return Client, NewCtx
}

// ? ============ Hash The Password
func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

// ? ============ Compare The Password
func CompareHashPasswords(HashedPasswordFromDB, PasswordToCampare string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPasswordFromDB), []byte(PasswordToCampare))
	return err == nil
}

// ? ============ Handle Cors
func Cors(c *gin.Context)  {
	c.Writer.Header().Set("Access-Control-Allow-Origin","http://   :5500")
	c.Writer.Header().Set("Access-Control-Allow-Headers","Content-Type")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
	}

	c.Next()
}