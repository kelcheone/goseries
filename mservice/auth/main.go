package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	handleRequest()

}

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["client"] = "kelche"
	claims["aud"] = " billing.jwt.go.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(MySigningKey)

	if err != nil {
		fmt.Printf("Something went wrong. %s", err.Error())
		return "", err
	}

	return tokenString, nil

}

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	fmt.Println(validToken)

	if err != nil {
		fmt.Println("Failed to generate the token ")
	}

	fmt.Fprintf(w, string(validToken))
}

func handleRequest() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
