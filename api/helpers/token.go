package helpers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Uid string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(userId string) (signedToken string, signedRefreshToken string, err error) {

	claims := &SignedDetails{
		Uid: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(170)).Unix(),
		},
	}

	token, err := createRefreshOrToken(*claims)

	refreshToken, err := createRefreshOrToken(*refreshClaims)

	if err != nil {
		log.Panic(err)
	}

	return token, refreshToken, err

}

func createRefreshOrToken(claims SignedDetails) (token string, err error) {

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
}

func DecodeJWT(tokenStr string) (SignedDetails, error) {

	secretString := []byte(SECRET_KEY)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secretString, nil

	})

	if err != nil {
		return SignedDetails{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return SignedDetails{
			Uid: fmt.Sprintf("%v", claims["Uid"]),
		}, nil
	} else {
		return SignedDetails{}, errors.New("Invalid token")
	}

}
