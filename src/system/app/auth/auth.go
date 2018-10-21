package auth

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

var (
	SignKey   *rsa.PrivateKey
	VerifyKey *rsa.PublicKey
)

func initKeys() {
	var err error
	SignKeyFile, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}
	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(SignKeyFile)

	if err != nil {
		log.Fatal("Error make RSA private key")
		return
	}
	VerifyKeyFile, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(VerifyKeyFile)

	if err != nil {
		log.Fatal("Error make RSA public key")
		return
	}
}

func Auth(id int) []byte {
	var str Token
	initKeys()
	signer := jwt.NewWithClaims(jwt.GetSigningMethod("RS512"), jwt.MapClaims{
		"id":   id,
		"role": "admin",
		"exp":  0,
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

func JsonResponseByVar(ok string, data interface{}) []byte {
	var res struct {
		Ok   string      `json:"ok"`
		Data interface{} `json:"data"`
	}
	res.Ok = ok
	res.Data = data
	json, err := json.Marshal(res)
	if err != nil {
		panic(err.Error())
	}

	return json
}

func IsTokenValid(val string) (int64, error) {
	initKeys()
	token, err := jwt.Parse(val, func(token *jwt.Token) (interface{}, error) {
		return VerifyKey, nil
	})

	switch err.(type) {
	case nil:
		if !token.Valid {
			return 0, errors.New("Token is invalid")
		}

		var userID int64

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return 0, errors.New("Token is invalid")
		}

		userID = int64(claims["id"].(float64))

		return userID, nil
	case *jwt.ValidationError:
		vErr := err.(*jwt.ValidationError)

		switch vErr.Errors {
		case jwt.ValidationErrorExpired:
			return 0, errors.New("Token Expired, get a new one.")
		default:
			fmt.Println(vErr)
			return 0, errors.New("Error while Parsing Token!")
		}
	default:
		return 0, errors.New("Unable to parse token")
	}
}
