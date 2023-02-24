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

func Val(tokenString string) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return our secretkey here
		return []byte("TestSecretKey"), nil
	})

	// fmt.Printf("%v \n", err.Error())
	// fmt.Printf("%v \n", token.Header)

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		// do what you want here
		fmt.Printf("%v", claims.Foo)
	} else {
		fmt.Println(err)
		fmt.Println(token.Valid)
	}
}

func main() {
	ss := Hmac()
	fmt.Printf("%v \n", ss)

	Val(ss)
}
