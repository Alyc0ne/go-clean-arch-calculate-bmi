package mocks

import (
	domain "github.com/bxcodec/go-clean-arch/domain"
	mock "github.com/stretchr/testify/mock"
)

type BmiRepository struct {
	mock.Mock
}

func (r *BmiRepository) FindBmiCondition(bmiRate float64) (*domain.BmiCondition, error) {
	ret := r.Called(bmiRate)
	if len(ret) == 0 {
		panic("no return value specified for FindBmiCondition")
	}

	// กำหนดค่าผลลัพธ์จาก mock
	result, _ := ret.Get(0).(*domain.BmiCondition)
	err, _ := ret.Get(1).(error)

	return result, err
}
