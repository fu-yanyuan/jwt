package main

import (
	"fmt"
	"time"

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
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
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
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return our secretkey here
		return []byte("TestSecretKey"), nil
	})

	// fmt.Printf("************%v \n", err.Error())
	fmt.Printf("%v \n", token.Header)

	// if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
	if token.Valid {
		// do what you want here
		// fmt.Printf("%v %v\n", claims.Foo, claims.ExpiresAt)

		// ipfs request
		ipfsReq()

	} else {
		fmt.Println(err)
		fmt.Println(token.Valid)
	}
}

func main() {
	ss := Hmac()
	fmt.Printf("%v \n", ss)

	Val(ss)
	time.Now().Add(24 * time.Hour).Unix()
	fmt.Println(time.Now())

	now := jwt.TimeFunc().Unix()
	fmt.Println(now)
}
