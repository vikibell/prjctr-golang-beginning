package user

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll(page int, limit int) ([]User, int64, error) {
	var users []User
	var total int64

	offset := (page - 1) * limit
	err := r.db.
		Model(&User{}).
		Count(&total).
		Limit(limit).
		Offset(offset).
		Find(&users).Error

	return users, total, err
}

func (r *Repository) FindOneByID(id int) (User, error) {
	var user User

	err := r.db.
		Model(&User{}).
		First(&user, id).Error

	return user, err
}
