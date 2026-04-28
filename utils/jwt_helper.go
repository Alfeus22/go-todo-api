package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var secretKey = []byte("RahasiaDapurGalih22")

func GenerateToken(userID bson.ObjectID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.Hex(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
