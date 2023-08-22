package util

import (
	"api-go-gin/properties"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

const InvalidTokenMessage = "token has invalid claims: token is expired"

func GenerateToken(id int) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iss"] = fmt.Sprintf("http://%s:%d", properties.Properties.Server.Host, properties.Properties.Server.Port)

	tokenSigned, err := token.SignedString([]byte(properties.Properties.Jwt.Secret))

	if err != nil {
		log.Panic(err)
	}

	return tokenSigned
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(properties.Properties.Jwt.Secret), nil
	})

	if err != nil {
		log.Println("Error during token verification:", err)

		if err.Error() == InvalidTokenMessage {
			return nil, errors.New("EXPIRED_TOKEN")
		}

		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		log.Println("Invalid token")
		return nil, errors.New("INVALID_TOKEN")
	}

	return token, nil
}
