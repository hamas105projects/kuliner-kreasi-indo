package user

import (
	"errors"
	"kuliner-kreasi-indo/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) CreateCashier(name, email, password string) (User, error) {
	hash, _ := utils.HashPassword(password)
	user := User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: hash,
		Role:     "cashier",
	}
	err := s.DB.Create(&user).Error
	return user, err
}

func (s *UserService) GetAllCashiers(offset, limit int) ([]User, error) {
	var users []User
	err := s.DB.Where("role = ?", "cashier").Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (s *UserService) GetCashierByID(id string) (User, error) {
	var u User
	err := s.DB.Where("id = ? AND role = ?", id, "cashier").First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, nil
	}
	return u, err
}

func (s *UserService) UpdateCashier(id, name, email, password string) (User, error) {
	var u User
	if err := s.DB.Where("id = ? AND role = ?", id, "cashier").First(&u).Error; err != nil {
		return u, err
	}

	if name != "" {
		u.Name = name
	}
	if email != "" {
		u.Email = email
	}
	if password != "" {
		hash, _ := utils.HashPassword(password)
		u.Password = hash
	}

	err := s.DB.Save(&u).Error
	return u, err
}

func (s *UserService) DeleteCashier(id string) error {
	return s.DB.Where("id = ? AND role = ?", id, "cashier").Delete(&User{}).Error
}
