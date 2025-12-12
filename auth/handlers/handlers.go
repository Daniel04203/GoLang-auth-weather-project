package handlers

import (
	"log"
	"strconv"
	"strings"
	"time"

	"example.com/labwork_8/auth/cfg"
	"example.com/labwork_8/auth/db"
	"example.com/labwork_8/auth/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func checkPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByLogin(login string) (*model.User, error) {
	db := db.DB

	var user model.User
	if err := db.Where(&model.User{Login: login}).First(&user).Error; err != nil {
		// if errors.Is(err, gorm.ErrRecordNotFound) {
		// 	return nil, nil
		// }
		return nil, err
	}

	return &user, nil
}

func getToken(login string) (string, error) {
	tokenExp, _ := strconv.Atoi(cfg.GetProperty("JWT_EXP_HRS"))

	claims := jwt.MapClaims{
		"login": login,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * time.Duration(tokenExp)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(cfg.GetProperty("JWT_SECRET")))
	return t, err
}

func AuthHandler(c *fiber.Ctx) error {
	var user model.User

	c.BodyParser(&user)

	userDB, err := getUserByLogin(user.Login)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Login or password are incorrect!",
		})
	}

	if !checkPasswordHash(userDB.Password, user.Password) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Login or password are incorrect!",
		})
	}

	t, err := getToken(userDB.Login)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Success login", "data": t})
}

func VerifyHandler(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(cfg.GetProperty("JWT_SECRET")), nil
	})

	if err != nil {
		log.Println("Token parse error:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		log.Println("Invalid token claims")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token claims",
		})
	}

	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			log.Println("Token expired")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Token expired",
			})
		}
	}

	return c.SendStatus(fiber.StatusOK)
}
