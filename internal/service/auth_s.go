package service

import (
	"errors"
	"fmt"
	"os"
	"soulvent/internal/model"
	"soulvent/internal/repository"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUP(username, email, password, gender, city string) (*model.User, string, error) {
	username = strings.ToLower(username)
	email = strings.ToLower(email)
	existingUser, _ := repository.GetUser("", email, "")
	if existingUser != nil {
		return nil, "", errors.New("email already exists")
	}
	existingUser, _ = repository.GetUser("", "", username)
	if existingUser != nil {
		return nil, "", errors.New("username already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user := &model.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		Gender:       gender,
		City:         city,
	}
	err = repository.CreateUser(user)
	if err != nil {
		return nil, "", err
	}
	token, err := generateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func LogIn(usernameOrEmail, password string) (*model.User, string, error) {
	//make lower case
	usernameOrEmail = strings.ToLower(usernameOrEmail)

	user, err := repository.GetUser("", usernameOrEmail, "")
	if err != nil || user == nil {
		user, err = repository.GetUser("", "", usernameOrEmail)
		if err != nil || user == nil{
			return nil, "", errors.New("invalid username or email, no user found ")
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := generateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func generateToken(user *model.User) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET is not set")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET is not set")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", errors.New("invalid user ID in token")
	}

	return userID, nil
}
