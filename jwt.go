package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func Hmac() string {
	mySigningKey := []byte("TestSecretKey")

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtString, _ := token.SignedString(mySigningKey)

	fmt.Printf("%v \n", jwtString)
	fmt.Printf("%v \n", token.Header)

	return jwtString
}

func main() {
	ss := Hmac()
	fmt.Printf("%v \n", ss)
}
