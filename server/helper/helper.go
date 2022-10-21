package helper

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = "secret"
var timeout = 300 * time.Minute

func GenerateToken(id uint, email string) (string, error) {
	claim := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(timeout),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := parseToken.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(tokenStr string) (map[string]interface{}, error) {
	errResp := fmt.Errorf("need signin")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResp
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResp
	}

	var payload = map[string]interface{}{}
	claims := token.Claims.(jwt.MapClaims)
	payload["email"] = claims["email"]
	payload["id"] = claims["id"]

	exp := fmt.Sprintf("%v", claims["exp"])

	now := time.Now()
	expTime, _ := time.Parse(time.RFC3339, exp)

	if !now.Before(expTime) {
		return nil, fmt.Errorf("expired")
	}

	return payload, nil
}

func GeneratePasswordBrypt(password string) string {
	pwd := []byte(password)
	salt := 8
	hashedPassword, err := bcrypt.GenerateFromPassword(pwd, salt)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func ComparePwd(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
