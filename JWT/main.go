package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_Secret_Key")

var users = map[string]string{
	"Mohammad": "1234",
}

type loginInfo struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

func main() {
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/welcome", welcome)

	log.Fatal(http.ListenAndServe(":8301", nil))
}

func signin(w http.ResponseWriter, r *http.Request) {
	var login loginInfo
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pass, ok := users[login.Username]

	if !ok || pass != login.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expTime := time.Now().Add(5 * time.Minute)

	cliams := &Claims{
		Username: login.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	stringToken, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "JWTToken",
		Expires:  expTime,
		Value:    stringToken,
		HttpOnly: true,
	})

}

func welcome(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("JWTToken")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString := c.Value

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "خوش آمدید.")
}
