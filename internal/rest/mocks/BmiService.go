package mocks

import (
	domain "github.com/bxcodec/go-clean-arch/domain"
	mock "github.com/stretchr/testify/mock"
)

type BmiService struct {
	mock.Mock
}

func (s *BmiService) CalculateResultBmi(req *domain.BmiReq) (*domain.BmiCondition, error) {
	ret := s.Called(req)
	if len(ret) == 0 {
		panic("no return value specified for CalculateResultBmi")
	}

	// กำหนดค่าผลลัพธ์จาก mock
	result, _ := ret.Get(0).(*domain.BmiCondition)
	err, _ := ret.Get(1).(error)

	return result, err
}
