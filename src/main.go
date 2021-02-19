package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//HashPassword get a password and hash
func HashPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(passwordBytes), err
}

//CheckPasswordHash check the password hash match
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "password is secret"
	hash, err := HashPassword(password)
	if err != nil {
		fmt.Errorf("Could not hash password")
	}

	fmt.Println("Password:", password)
	fmt.Println("Hash:", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:", match)
}
