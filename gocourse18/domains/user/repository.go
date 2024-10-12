package user

import (
	"github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/user/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll(page int, limit int) (users []model.User, total int64, err error) {
	offset := (page - 1) * limit

	err = r.db.
		Model(&model.User{}).
		Count(&total).
		Limit(limit).
		Offset(offset).
		Find(&users).Error

	return
}

func (r *Repository) FindOneById(id int) (user model.User, err error) {
	err = r.db.
		Model(&model.User{}).
		First(&user, id).Error

	return
}
