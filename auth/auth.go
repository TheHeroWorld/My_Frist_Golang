package auth

import (
	"My_Frist_Golang/db"
	"My_Frist_Golang/logging"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log = logging.GetLogger()

var (
	key []byte
	t   *jwt.Token
	s   string
)

func Auth(email string, password string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	JWTkey := os.Getenv("KEY")
	key := []byte(JWTkey)
	exp := jwt.NewNumericDate(time.Now().Add(6 * time.Hour))
	id, err := db.FindUser(email, password) // Ищем юзера в базе данных
	if err != nil || id == 0 {
		log.WithFields(logrus.Fields{
			"login": email,
		}).Info("Not found")
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
