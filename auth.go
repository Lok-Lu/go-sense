package sense

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func EncodeJwtToken(ak string, sk string) (string, error) {
	payload := jwt.MapClaims{
		"iss": ak,
		"exp": time.Now().Add(10000 * time.Hour).Unix(),
		"nbf": time.Now().Add(-5 * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte(sk))
	if err != nil {
		return "", fmt.Errorf("error encoding JWT token: %s", err.Error())
	}
	return signedToken, nil
}

func CheckJwtToken(tokenString, sk string) (bool, error) {
	body, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(sk), nil
	})
	if err != nil {
		return false, err
	}
	if !body.Valid {
		return false, errors.New("invalid token")
	}

	token := body.Claims.(jwt.MapClaims)
	exp := token["exp"].(float64)
	return int64(exp) > time.Now().Unix(), nil
}
