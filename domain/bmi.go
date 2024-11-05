package domain

type (
	BmiReq struct {
		Weight float64 `json:"weight" validate:"required"`
		Height float64 `json:"height" validate:"required"`
	}

	BmiCondition struct {
		Id           int     `json:"-" gorm:"id"`
		Bmi          float64 `json:"bmi" gorm:"-"`
		CategoryName string  `json:"category_name" gorm:"category_name"`
		Min          float64 `json:"-" gorm:"min"`
		Max          float64 `json:"-" gorm:"max"`
		BmiDesc      string  `json:"bmi_desc" gorm:"bmi_desc"`
		BmiAdvice    string  `json:"bmi_advice" gorm:"bmi_advice"`
	}
)
