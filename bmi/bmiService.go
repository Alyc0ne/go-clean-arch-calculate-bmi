package bmi

import (
	"errors"
	"fmt"
	"math"

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

func (s *bmiService) roundToDecimals(value float64) float64 {
	//* fix to 2 decimals
	decimals := 2
	factor := math.Pow(10, float64(decimals))
	return math.Round(value*factor) / factor
}

func (s *bmiService) CalculateResultBmi(req *domain.BmiReq) (*domain.BmiCondition, error) {
	if req.Weight < 0 || req.Height < 0 {
		return nil, errors.New("weight and height cannot less than 0")
	}

	//* Convert Height From Centimeter to Meter.
	heightMeters := req.Height / 100

	bmi := s.roundToDecimals(req.Weight / (heightMeters * heightMeters))
	bmiCondition, err := s.bmiRepo.FindBmiCondition(bmi)
	if err != nil {
		return nil, domain.ErrInternalServerError
	}

	if bmiCondition == nil {
		return nil, fmt.Errorf("bmi condition not found")
	}

	bmiCondition.Bmi = bmi

	return bmiCondition, nil
}
