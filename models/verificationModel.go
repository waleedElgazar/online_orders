package models

type Verification struct {
	Id					uint	`json:"-"`
	UserAccount			string	`json:"user_account"`
	VerificationCode	string		`json:"verification_code"`
}