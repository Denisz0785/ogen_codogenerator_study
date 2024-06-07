package handlers

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const salt = "hjb%&^jkl354h!"
const signingKey = "jhgjhgjhgjhgjhgjhg"
const tokenTTL = 12 * time.Hour

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func hashPassword(pass string) string {
	hash := sha256.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}

func (h *Handler) generateToken(name, pass string) (string, error) {
	user, err := h.repo.GetUser(name, hashPassword(pass))
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func parseToken(inputToken string) (int, error) {
	token, err := jwt.ParseWithClaims(inputToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("unexpected signing method")
			return nil, errors.New("unexpected signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		log.Println(err.Error())
		return -1, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return -1, errors.New("invalid token claims")
	}
	return claims.UserId, nil
}
