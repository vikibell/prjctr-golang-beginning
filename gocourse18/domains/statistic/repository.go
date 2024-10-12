package statistic

import (
	model2 "github.com/vikibell/prjctr-golang-beginning/gocourse18/domains/statistic/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewStatisticsRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll(page int, limit int) (statistics []model2.Statistics, total int64, err error) {
	offset := (page - 1) * limit

	err = r.db.
		Model(&model2.Statistics{}).
		Count(&total).
		Limit(limit).
		Offset(offset).
		Find(&statistics).Error

	return
}

func (r *Repository) FindOneById(id int) (statistic model2.Statistics, err error) {
	err = r.db.
		Model(&model2.Statistics{}).
		First(&statistic, id).Error

	return
}
