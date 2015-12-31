package db

import (
	"golang.org/x/crypto/scrypt"
)

type User struct {
	uID int
	username string
	salt string
	algo string
	parameters string
}

func Login (username, password string) (*User, error){
	
}
