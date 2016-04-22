package db

import (
	"crypto/rand"
	//"encoding/base64"
)

const (
	SaltSize = 16 // 128 bits
)

type User struct {
	uID        int
	username   string
	salt       string
	algo       string
	parameters string
}

func Login(username, password string) (*User, error) {
	return nil, nil
}

func genSalt() ([]byte, error) {
	b := make([]byte, SaltSize)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func CreateUser(username, password string) (*User, error) {
	return nil, nil
}

type kdf interface {
	//hash(password string, salt []byte,
}
