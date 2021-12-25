package user

import "gorm.io/gorm"

type IUserRepository interface {
	Save(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &repository{db}
}
func (r *repository ) Save(user User) (User,error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user,err
	}
	return user, nil
}