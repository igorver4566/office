package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const (
	privKeyPath = "./keys/app.rsa"
	pubKeyPath  = "./keys/app.rsa.pub"
)

type Token struct {
	Ok    string `json:"ok"`
	Token string `json:"token"`
}

var VerifyKey, SignKey []byte

func initKeys() {
	var err error

	SignKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	VerifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

func Auth(id int) []byte {
	var str Token
	initKeys()
	signer := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), jwt.MapClaims{
		"id":   id,
		"role": "admin",
		"exp":  time.Now().Add(time.Minute * 20).Unix(),
	})
	tokenString, err := signer.SignedString(SignKey)
	if err != nil {
		str.Ok = "false"
		str.Token = "Ошибка получения ключа"
		log.Printf("Error signing token: %v\n", err)
	} else {
		str.Ok = "true"
		str.Token = tokenString
	}
	return JsonResponse(str)
}
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return VerifyKey, nil
		})

	if err == nil {

		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorised access to this resource")
	}
}
func JsonResponse(response interface{}) []byte {

	json, err := json.Marshal(response)
	if err != nil {
		panic(err.Error())
	}

	return json
}

func JsonResponseByVar(ok string, data string) []byte {
	var res struct {
		Ok   string `json:"ok"`
		Data string `json:"data"`
	}
	res.Ok = ok
	res.Data = data
	json, err := json.Marshal(res)
	if err != nil {
		panic(err.Error())
	}

	return json
}
