package bmi_test

import (
	"testing"

	"github.com/bxcodec/go-clean-arch/bmi"
	"github.com/bxcodec/go-clean-arch/bmi/mocks"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCalculateBMI(t *testing.T) {
	mockBmiRepository := new(mocks.BmiRepository)
	bmiService := bmi.NewBmiService(mockBmiRepository)

	bmiVeryOverWeight := &domain.BmiCondition{
		Id:           1,
		CategoryName: "อ้วนมาก",
		Min:          30.0,
		BmiDesc:      "ค่อนข้างอันตราย เสี่ยงต่อการเกิดโรคร้ายแรงที่แฝงมากับความอ้วน",
		BmiAdvice:    "ควรปรับพฤติกรรมการทานอาหาร และเริ่มออกกำลังกาย หาก BMI สูงกว่า 40.0 ควรไปตรวจสุขภาพและปรึกษาแพทย์",
	}

	bmiOverWeight := &domain.BmiCondition{
		Id:           2,
		CategoryName: "อ้วน",
		Min:          25.0,
		Max:          29.9,
		BmiDesc:      "อ้วนในระดับหนึ่ง ถึงแม้จะไม่ถึงเกณฑ์ที่ถือว่าอ้วนมาก ๆ แต่ก็ยังมีความเสี่ยงต่อการเกิดโรค",
		BmiAdvice:    "ควรปรับพฤติกรรมการทานอาหาร ออกกำลังกาย และตรวจสุขภาพ",
	}

	bmiNormalWeight := &domain.BmiCondition{
		Id:           3,
		CategoryName: "น้ำหนักปกติ เหมาะสม",
		Min:          18.5,
		Max:          24.9,
		BmiDesc:      "น้ำหนักที่เหมาะสมสำหรับคนไทยคือค่า BMI ระหว่าง 18.5-24",
		BmiAdvice:    "ควรรักษาระดับ BMI ให้อยู่ในช่วงนี้ให้นานที่สุด และตรวจสุขภาพทุกปี",
	}

	bmiUnderWeight := &domain.BmiCondition{
		Id:           4,
		CategoryName: "ผอมเกินไป",
		Max:          18.4,
		BmiDesc:      "น้ำหนักน้อยกว่าปกติ อาจเสี่ยงต่อการได้รับสารอาหารไม่เพียงพอ",
		BmiAdvice:    "ควรรับประทานอาหารเพียงพอ และออกกำลังกายเสริมสร้างกล้ามเนื้อเพื่อเพิ่มค่า BMI",
	}

	mockBmiRepository.On("FindBmiCondition", mock.MatchedBy(func(bmiRate float64) bool {
		return bmiRate >= 30.0
	})).Return(bmiVeryOverWeight, nil)

	mockBmiRepository.On("FindBmiCondition", mock.MatchedBy(func(bmiRate float64) bool {
		return bmiRate >= 25.0 && bmiRate <= 29.9
	})).Return(bmiOverWeight, nil)

	mockBmiRepository.On("FindBmiCondition", mock.MatchedBy(func(bmiRate float64) bool {
		return bmiRate >= 18.5 && bmiRate < 25.0
	})).Return(bmiNormalWeight, nil)

	mockBmiRepository.On("FindBmiCondition", mock.MatchedBy(func(bmiRate float64) bool {
		return bmiRate < 18.5
	})).Return(bmiUnderWeight, nil)

	tests := []struct {
		name      string
		req       *domain.BmiReq
		expected  *domain.BmiCondition
		expectErr bool
	}{
		{
			name:     bmiVeryOverWeight.CategoryName,
			req:      &domain.BmiReq{Weight: 95, Height: 170}, // BMI = 32.87
			expected: bmiVeryOverWeight,
		},
		{
			name:     bmiOverWeight.CategoryName,
			req:      &domain.BmiReq{Weight: 65, Height: 160}, // BMI = 25.39
			expected: bmiOverWeight,
		},
		{
			name:     bmiNormalWeight.CategoryName,
			req:      &domain.BmiReq{Weight: 56, Height: 170}, // BMI = 19.38
			expected: bmiNormalWeight,
		},
		{
			name:     bmiUnderWeight.CategoryName,
			req:      &domain.BmiReq{Weight: 32, Height: 150}, // BMI = 14.22
			expected: bmiUnderWeight,
		},
		{
			name:      "ทดสอบน้ำหนักติดลบ",
			req:       &domain.BmiReq{Weight: -70, Height: 170},
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "ทดสอบส่วนสูงติดลบ",
			req:       &domain.BmiReq{Weight: 70, Height: -170},
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := bmiService.CalculateResultBmi(tt.req)

			if tt.expectErr {
				assert.Error(t, err)
				assert.Nil(t, result, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
