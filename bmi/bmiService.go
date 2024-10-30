package bmi

import (
	"fmt"

	"github.com/bxcodec/go-clean-arch/domain"
)

type (
	BmiRepository interface {
		FindBmiCondition(bmiRate float64) (*domain.BmiCondition, error)
	}

	bmiService struct {
		bmiRepo BmiRepository
	}
)

func NewBmiService(bmiRepository BmiRepository) *bmiService {
	return &bmiService{bmiRepo: bmiRepository}
}

func (s *bmiService) CalculateResultBmi(req *domain.BmiReq) (*domain.BmiCondition, error) {
	//* Convert Height From Centimeter to Meter.
	heightMeters := req.Height / 100

	bmi := req.Weight / (heightMeters * heightMeters)

	bmiCondition, err := s.bmiRepo.FindBmiCondition(bmi)
	if err != nil {
		return nil, domain.ErrInternalServerError
	}

	if bmiCondition == nil {
		return nil, fmt.Errorf("bmi condition not found")
	}

	return bmiCondition, nil
}
