package repository

import (
	"github.com/bxcodec/go-clean-arch/domain"
	"gorm.io/gorm"
)

type bmiRepository struct {
	db *gorm.DB
}

func NewBmiRepository(db *gorm.DB) *bmiRepository {
	return &bmiRepository{db}
}

func (r *bmiRepository) FindBmiCondition(bmiRate float64) (*domain.BmiCondition, error) {
	query := "select id, category_name, min, max, bmi_desc, bmi_advice from bmi_condition where (min is not null and max is not null and ? between min and max) or (min is null and ? < max) or (max is null and ? > min)"

	bmiCondition := new(domain.BmiCondition)
	tx := r.db.Raw(query, bmiRate, bmiRate, bmiRate).Scan(&bmiCondition)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return bmiCondition, nil
}
