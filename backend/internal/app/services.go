package app

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt            = "sjkdfnshbghjslbgdfg1234"
	signingKey      = "4353KSFjH$8#9qrkjk#4#"
	tokenExpiration = time.Hour * 48
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (h *Handler) GenerateToken(email, password string) (string, error) {
	user, err := h.rep.GetUserAuth(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	if user.Id == 0 {
		return "", errors.New("bad credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpiration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	if err = h.rep.UpdateLastLogin(user.Id); err != nil {
		return "", err
	}

	return token.SignedString([]byte(signingKey))
}

func (h *Handler) ParseToken(accessToken string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return &TokenClaims{}, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return &TokenClaims{}, errors.New("token claims are not of the supposed type")
	}

	return claims, nil
}
