package repository_test

import (
	"log"
	"testing"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/internal/repository"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

func TestFindBmiCondition(t *testing.T) {
	db, mock := NewMockDB()

	mockBmiConditions := []domain.BmiCondition{
		{
			BmiId:     "1",
			Min:       0,
			Max:       18.4,
			BmiDesc:   "ผอมเกินไป",
			BmiAdvice: "น้ำหนักน้อยกว่าปกติ อาจเสี่ยงต่อการได้รับสารอาหารไม่เพียงพอ ควรรับประทานอาหารเพียงพอ และออกกำลังกายเสริมสร้างกล้ามเนื้อเพื่อเพิ่มค่า BMI",
		},
		{
			BmiId:     "2",
			Min:       18.5,
			Max:       24.9,
			BmiDesc:   "น้ำหนักปกติ เหมาะสม",
			BmiAdvice: "น้ำหนักที่เหมาะสมสำหรับคนไทยคือค่า BMI ระหว่าง 18.5-24 ควรรักษาระดับ BMI ให้อยู่ในช่วงนี้ให้นานที่สุด และตรวจสุขภาพทุกปี",
		},
	}

	rows := sqlmock.NewRows([]string{"bmi_id", "min", "max", "bmi_desc", "bmi_advice"}).
		AddRow(
			mockBmiConditions[0].BmiId, mockBmiConditions[0].Min, mockBmiConditions[0].Max,
			mockBmiConditions[0].BmiDesc, mockBmiConditions[0].BmiAdvice,
		).
		AddRow(
			mockBmiConditions[0].BmiId, mockBmiConditions[0].Min, mockBmiConditions[0].Max,
			mockBmiConditions[0].BmiDesc, mockBmiConditions[0].BmiAdvice,
		)

	query := "SELECT bmi_id, min, max, bmi_desc, bmi_advice from bmi_condition WHERE \\? between min and max"

	mock.ExpectQuery(query).WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

	bmiRepository := repository.NewBmiRepository(db)
	result, err := bmiRepository.FindBmiCondition(20.0)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	expected := mockBmiConditions[1]
	if result.BmiId != expected.BmiId || result.BmiDesc != expected.BmiDesc {
		t.Errorf("unexpected result: got %+v, want %+v", result, expected)
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}
