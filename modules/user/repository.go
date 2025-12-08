package user

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(u *User) error {
	return r.DB.Create(u).Error
}

func (r *UserRepository) GetByID(id string) (User, error) {
	var u User
	err := r.DB.First(&u, "id = ?", id).Error
	return u, err
}

func (r *UserRepository) GetAll(offset, limit int) ([]User, error) {
	var users []User
	err := r.DB.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (r *UserRepository) Update(u *User) error {
	return r.DB.Save(u).Error
}

func (r *UserRepository) Delete(id string) error {
	return r.DB.Delete(&User{}, "id = ?", id).Error
}
