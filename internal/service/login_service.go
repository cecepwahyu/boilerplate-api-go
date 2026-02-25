package service

import (
	"errors"
	"time"

	_ "github.com/cecepwahyu/rest-api-go/internal/model"
	"github.com/cecepwahyu/rest-api-go/internal/repository"
	"github.com/cecepwahyu/rest-api-go/internal/utilities"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	Login(email, password string) (map[string]interface{}, error)
}

type loginService struct {
	repo repository.LoginRepository
}

var jwtSecret = []byte("SECRET_KEY")

func NewLoginService(repo repository.LoginRepository) LoginService {
	return &loginService{repo}
}

func (s *loginService) Login(email, password string) (map[string]interface{}, error) {

	if !utilities.IsValidEmail(email) {
		return nil, errors.New("invalid email format")
	}

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"idPeserta": user.IDPeserta,
		"email":     user.Email,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(jwtSecret)

	user.Token = tokenString
	user.TokenUpdatedAt = time.Now()
	s.repo.Save(user)

	return map[string]interface{}{
		"idPeserta": user.IDPeserta,
		"token":     tokenString,
		"email":     user.Email,
		"isActive":  user.IsActive,
	}, nil
}
