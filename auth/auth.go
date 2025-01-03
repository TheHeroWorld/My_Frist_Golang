package auth

import (
	"My_Frist_Golang/db"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	key []byte
	t   *jwt.Token
	s   string
)

func Auth(email string, password string) (string, error) {
	key := []byte(`K0IxiQZOBwHGejUGCTwEz7J9EKi6l1evwEdET/Zy6mg=`)
	exp := jwt.NewNumericDate(time.Now().Add(6 * time.Hour))
	id, err := db.FindUser(email, password) // Ищем юзера в базе данных
	if err != nil || id == 0 {
		return "", fmt.Errorf("invalid login or password")
	}
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ // Создаем JWT
		"id":    id,
		"email": email,
		"exp":   exp.Unix(),
	})
	s, _ = t.SignedString(key) // Возвращаем JWT
	return s, nil
}
