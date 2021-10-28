package database

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"

	"github.com/waleedElgazar/resturant/configration"
	"github.com/waleedElgazar/resturant/models"
)

var DB *sql.DB

func AddUser(user models.User) {
	DB = configration.OpenConnection()
	defer DB.Close()
	query := "INSERT INTO newResturant.users set userName=?, userPassword=?, userEmail=?, userPhone=?"
	insert, err := DB.Prepare(query)
	if err != nil {
		fmt.Println("error while executing insert query")
		return
	}
	_, err = insert.Exec(user.UserName, user.UserPassword, user.UserEmail, user.UserPhone)
	if err != nil {
		fmt.Println("error while parsing inserting data",err)
		return
	}

}

func GetUser(search string) models.User {
	var user models.User
	DB = configration.OpenConnection()
	defer DB.Close()
	query := "SELECT userId, userName, userEmail,userPassword,userPhone from users WHERE userEmail=?"
	result, err := DB.Query(query, search)
	if err != nil {
		fmt.Println("error while finding the user", err)
		return user
	}
	var name, password, email, phone string
	var id uint
	for result.Next(){
		err = result.Scan(&id, &name, &email, &password, &phone)
		user = models.User{
			UserName:     name,
			UserPassword: []byte(password),
			IdUser:       id,
			UserEmail:    email,
			UserPhone:    phone,
		}
	}
	
	if err != nil {
		fmt.Println("error while feteching data", err)
		fmt.Println(user.UserPassword,"------id ",user.IdUser)
	}
	return user

}
func GetUserWithId(search uint) models.User {
	var user models.User
	DB = configration.OpenConnection()
	defer DB.Close()
	query := "SELECT userId, userName, userEmail,userPassword,userPhone from users WHERE userId=?"
	result, err := DB.Query(query, search)
	if err != nil {
		fmt.Println("error while finding the user", err)
		return user
	}
	var name, password, email, phone string
	var id uint
	for result.Next(){
		err = result.Scan(&id, &name, &email, &password, &phone)
		user = models.User{
			UserName:     name,
			UserPassword: []byte(password),
			IdUser:       id,
			UserEmail:    email,
			UserPhone:    phone,
		}
	}
	
	if err != nil {
		fmt.Println("error while feteching data", err)
		fmt.Println(user.UserPassword,"------id ",user.IdUser)
	}
	return user

}

func AddVerification(verify models.Verification){
	DB = configration.OpenConnection()
	defer DB.Close()
	query := "INSERT INTO newResturant.verify set userEmail=?, verificationcode=?"
	insert, err := DB.Prepare(query)
	if err != nil {
		fmt.Println("error while executing insert query")
		return
	}
	_, err = insert.Exec(verify.UserAccount,verify.VerificationCode)
	if err != nil {
		fmt.Println("error while parsing inserting data",err)
		return
	}
}

func CreateOTP() string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	max := 6
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}