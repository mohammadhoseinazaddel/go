package jwt

import (
	"main/datalayer"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_Secret_Key")

type LoginInfo struct {
	Password string
	Username string
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, login LoginInfo, dbhandler datalayer.SQLHandler) bool {

	user, err := dbhandler.GetUserByEmail(login.Username)
	if err != nil {
		return false
	}

	if user.Password != login.Password {
		return false
	}

	expTime := time.Now().Add(30 * (24 * time.Hour))

	cliams := &Claims{
		Username: login.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	stringToken, err := token.SignedString(jwtKey)

	if err != nil {
		return false
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "JWTToken",
		Expires:  expTime,
		Value:    stringToken,
		HttpOnly: true,
	})
	return true
}

func IsLogedin(r *http.Request) (LoginInfo, bool) {

	c, err := r.Cookie("JWTToken")
	if err != nil {
		return LoginInfo{}, false
	}

	tokenString := c.Value

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		return LoginInfo{}, false
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return LoginInfo{}, false
		}
		return LoginInfo{}, false
	}

	return LoginInfo{
		Username: claims.Username,
	}, true
}
