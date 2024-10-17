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

func (r *Repository) FindAll(page int, limit int) ([]Statistic, int64, error) {
	offset := (page - 1) * limit
	var statistics []Statistic
	var total int64

	err := r.db.
		Model(&Statistic{}).
		Count(&total).
		Limit(limit).
		Offset(offset).
		Find(&statistics).Error

	return statistics, total, err
}

func (r *Repository) FindOneByID(id int) (Statistic, error) {
	var statistic Statistic

	err := r.db.
		Model(&Statistic{}).
		First(&statistic, id).Error

	return statistic, err
}
