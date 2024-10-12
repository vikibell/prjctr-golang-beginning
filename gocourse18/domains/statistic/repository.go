package statistic

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewStatisticsRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll(page int, limit int) (statistics []Statistic, total int64, err error) {
	offset := (page - 1) * limit

	err = r.db.
		Model(&Statistic{}).
		Count(&total).
		Limit(limit).
		Offset(offset).
		Find(&statistics).Error

	return
}

func (r *Repository) FindOneById(id int) (statistic Statistic, err error) {
	err = r.db.
		Model(&Statistic{}).
		First(&statistic, id).Error

	return
}
