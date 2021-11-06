package controller

import (
	"fmt"
	"gopkg.in/mail.v2"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/waleedElgazar/resturant/database"
	"github.com/waleedElgazar/resturant/models"
	"golang.org/x/crypto/bcrypt"
)

var CommonUserId int
var verifyCode models.Verification

func Register(ctx *fiber.Ctx) error {
	var data map[string]string
	err := ctx.BodyParser(&data)
	if err != nil {
		fmt.Println("error while parsing data", err)
		return err
	}
	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		fmt.Println("error while encrypting password")
		return err
	}
	verifyCode=models.Verification{
		UserAccount: data["email"],
		VerificationCode: database.CreateOTP(),
	}
	user := models.User{
		UserName:     data["name"],
		UserEmail:    data["email"],
		UserPhone:    data["phone"],
		UserPassword: password,
	}
	database.AddVerification(verifyCode)
	SendVerifyCode(verifyCode)
	database.AddUser(user)
	return ctx.JSON(user)
}

func Login(ctx *fiber.Ctx) error {
	var data map[string]string
	err := ctx.BodyParser(&data)
	if err != nil {
		fmt.Println("error parsing data", err)
		return err
	}
	user := database.GetUser(data["email"])
	CommonUserId = int(user.IdUser)
	pass := data["password"]
	err = bcrypt.CompareHashAndPassword(user.UserPassword, []byte(pass))
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(
			fiber.Map{
				"message": "the password isn't correct",
			},
		)
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.IdUser)),
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	})
	token, err := claims.SignedString([]byte(os.Getenv("MySecretKey")))
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "couldn't login",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * 10),
		HTTPOnly: true,
	}
	ctx.Cookie(&cookie)
	return ctx.JSON(fiber.Map{
		"message": "login",
	})
}

func GetUser(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("MySecretKey")), nil
		})
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	id, _ := strconv.Atoi(claims.Issuer)
	user := database.GetUserWithId(uint(id))
	return ctx.JSON(user)
}

func IsAuthorized(ctx *fiber.Ctx) error {
	//must run login function to get the token first
	cookie := ctx.Cookies("jwt")
	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("MySecretKey")), nil
		})
	if err != nil {
		ctx.Status(fiber.StatusUnauthorized)
		return err
	}
	return nil
}

func LogOut(ctx *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	ctx.Cookie(&cookie)
	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func VerifyAccount(ctx *fiber.Ctx) error {
	var data models.Verification
	err:=ctx.BodyParser(&data)
	if err!=nil {
		fmt.Println("error while verify",err)
		return nil
	}

	if data.VerificationCode!= verifyCode.VerificationCode{
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(
			fiber.Map{
				"message":"verification failed ",
			},
		)
	}
	return ctx.JSON(
		fiber.Map{
			"message":"verification success ",
		},
	)
}

func SendVerifyCode(data models.Verification){
	adc := mail.NewMessage()
	adc.SetHeader("From", os.Getenv("EMAIL"))
	fmt.Println("email",os.Getenv("EMAIL"),os.Getenv("PASSWORD"))
	adc.SetHeader("To", data.UserAccount)
	adc.SetHeader("Subject", "hi from golang")
	adc.SetBody("text/plain", "your verification code is "+data.VerificationCode)
	a := mail.NewDialer("smtp.gmail.com", 587, "walidreda427@gmail.com", os.Getenv("PASSWORD"))
	if err := a.DialAndSend(adc); err != nil {
		fmt.Println("error ", err)
		panic(err)
	}
}