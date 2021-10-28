package models


type User struct{
	IdUser			uint		`json:"-"`
	UserName		string		`json:"name"`
	UserPassword	[]byte		`json:"-"`
	UserEmail		string		`json:"email"`
	UserPhone		string		`json:"phone"`
}