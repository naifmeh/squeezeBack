package common

import (
	"io/ioutil"
	"log"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
	"errors"
	"strings"
)
type AppClaims struct {
	DeviceName string `json:"deviceName"`
	DeviceMac string `json:"deviceMac"`
	Role string `json:"role"`
	jwt.StandardClaims
}

/* Private key sign JWT, public Key verify JWT in reauest */
const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath = "keys/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

func initKeys() {
	var err error
	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s \n", err)
	}
}

func GenerateJWT(deviceName, deviceMac, role string) (string, error) {
	claims := AppClaims{
		deviceName,
		deviceMac,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer: "admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256,claims)

	ss,err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

/*Middleware for validating the token*/
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func (token *jwt.Token)(interface{}, error) {
		return verifyKey, nil
	}, request.WithClaims(&AppClaims{}))

	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				DisplayAppError(
					w,
					err,
					"Access token is expired",
					401,
				)
				return
			default:
				DisplayAppError(w,
					err,
					"Error while parsing JWT",
					500)
				return
			}

		default:
			DisplayAppError(w,
				err,
				"Error while parsing JWT",
				500)
			return

		}
	}
	if token.Valid {
		context.Set(r,"device",token.Claims.(*AppClaims).DeviceName)
		next(w,r) // Next handkler
	}
}

func TokenFromAuthHeader(r *http.Request) (string, error) {
	if ah := r.Header.Get("Authorization"); ah != "" {
		if len(ah) >6 && strings.ToUpper(ah[0:6]) == "BEARER" {
			return ah[7:],nil
		}
	}
	return "", errors.New("No token in the HTTP request")
}