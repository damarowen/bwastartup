package user

import (
	"errors"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
	UpdateUser(user User) (User, error)
	FindAllUser() ([]User, error)

}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}
func (r *UserRepository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Take(&user).Error
	if err != nil {
		return user, errors.New("Email not found")
	}
	return user, nil
}

func (r *UserRepository) FindById(id int) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).Take(&user).Error
	if err != nil {
		return user, errors.New("User not found")
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *UserRepository) FindAllUser() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}