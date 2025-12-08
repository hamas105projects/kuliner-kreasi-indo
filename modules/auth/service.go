package auth

import (
	"kuliner-kreasi-indo/modules/user"
	"kuliner-kreasi-indo/utils"

	"gorm.io/gorm"
)

type AuthService struct {
	DB        *gorm.DB
	JWTSecret string
	ExpireHrs int
}

func (s *AuthService) Login(email, password string) (string, error) {
	var u user.User
	if err := s.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return "", err
	}

	if !utils.CheckPassword(u.Password, password) {
		return "", gorm.ErrRecordNotFound
	}

	token, err := GenerateToken(u.ID.String(), u.Role, s.JWTSecret, s.ExpireHrs)
	if err != nil {
		return "", err
	}
	return token, nil
}
