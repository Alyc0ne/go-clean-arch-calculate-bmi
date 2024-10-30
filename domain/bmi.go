package domain

type (
	BmiReq struct {
		Weight float64 `json:"weight" validate:"required"`
		Height float64 `json:"height" validate:"required"`
	}

	BmiCondition struct {
		BmiId     string  `json:"bmi_id" gorm:"bmi_id"`
		Min       float64 `json:"min" gorm:"min"`
		Max       float64 `json:"max" gorm:"max"`
		BmiDesc   string  `json:"bmi_desc" gorm:"bmi_desc"`
		BmiAdvice string  `json:"bmi_advice" gorm:"bmi_advice"`
	}
)
