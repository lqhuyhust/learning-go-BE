package services

import (
	"errors"
	"httpServer/models"
	"httpServer/repositories"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type AuthService struct {
	UserRepository         *repositories.UserRepository
	RefreshTokenRepository *repositories.RefreshTokenRepository
}

func NewAuthService(userRepository *repositories.UserRepository, refreshTokenRepository *repositories.RefreshTokenRepository) *AuthService {
	return &AuthService{
		UserRepository:         userRepository,
		RefreshTokenRepository: refreshTokenRepository,
	}
}

// hash password
func (s *AuthService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// verify password
func (s *AuthService) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// generate jwt access token
func (s *AuthService) GenerateAccessJWT(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// generate jwt refresh token
func (s *AuthService) GenerateRefreshJWT(user *models.User) (string, error) {
	// delete old refresh token
	if err := s.RefreshTokenRepository.DeleteByUserID(user.ID); err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	// save new refresh token
	newRefreshTokenModel := models.RefreshToken{
		RefreshToken: tokenString,
		UserID:       user.ID,
		ExpiredAt:    time.Now().Add(time.Hour * 24),
	}
	if err := s.RefreshTokenRepository.Save(&newRefreshTokenModel); err != nil {
		return "", err
	}

	return tokenString, nil
}

// verify jwt access token
func (s *AuthService) VerifyAccessJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	return uint(userID), nil
}

// register new user
func (s *AuthService) Register(username, password, profile string) error {
	// hash password
	hashedPassword, err := s.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: username,
		Password: hashedPassword,
		Profile:  profile,
	}

	return s.UserRepository.Save(user)
}

// login
func (s *AuthService) Login(username, password string) (string, string, error) {
	user, err := s.UserRepository.FindByUsername(username)
	if err != nil {
		return "", "", err
	}

	if err := s.VerifyPassword(user.Password, password); err != nil {
		return "", "", err
	}

	// generate jwt access token
	accessToken, err := s.GenerateAccessJWT(user)
	if err != nil {
		return "", "", err
	}

	// generate jwt refresh token
	refreshToken, err := s.GenerateRefreshJWT(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Logout(userID uint) error {
	return s.RefreshTokenRepository.DeleteByUserID(userID)
}

func (s *AuthService) RefreshToken(refreshToken string) (string, string, error) {
	// find refresh token
	token, err := s.RefreshTokenRepository.FindByRefreshToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	// check if refresh token expired
	if token.ExpiredAt.Before(time.Now()) {
		return "", "", errors.New("refresh token expired")
	}

	// get user
	user, err := s.UserRepository.FindByID(token.UserID)
	if err != nil {
		return "", "", err
	}

	// generate new access token
	newAccessToken, err := s.GenerateAccessJWT(user)
	if err != nil {
		return "", "", err
	}

	// generate new refresh token
	newRefreshToken, err := s.GenerateRefreshJWT(user)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
