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
	query := `
		select
			*
		from
			bmi_condition
		where
			? between min and max
	`
	bmiCondition := new(domain.BmiCondition)
	tx := r.db.Raw(query, bmiRate).Scan(&bmiCondition)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return bmiCondition, nil
}
