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
			Id:           2,
			CategoryName: "อ้วน",
			Min:          25.0,
			Max:          29.9,
			BmiDesc:      "อ้วนในระดับหนึ่ง ถึงแม้จะไม่ถึงเกณฑ์ที่ถือว่าอ้วนมาก ๆ แต่ก็ยังมีความเสี่ยงต่อการเกิดโรค",
			BmiAdvice:    "ควรปรับพฤติกรรมการทานอาหาร ออกกำลังกาย และตรวจสุขภาพ",
		},
	}

	valueFilter := 20.0

	rows := sqlmock.NewRows([]string{"id", "category_name", "min", "max", "bmi_desc", "bmi_advice"}).
		AddRow(
			mockBmiConditions[0].Id, mockBmiConditions[0].CategoryName, mockBmiConditions[0].Min, mockBmiConditions[0].Max,
			mockBmiConditions[0].BmiDesc, mockBmiConditions[0].BmiAdvice,
		)

	query := "select id, category_name, min, max, bmi_desc, bmi_advice from bmi_condition where \\(min is not null and max is not null and \\? between min and max\\) or \\(min is null and \\? < max\\) or \\(max is null and \\? > min\\)"

	mock.ExpectQuery(query).WithArgs(valueFilter, valueFilter, valueFilter).WillReturnRows(rows)

	bmiRepository := repository.NewBmiRepository(db)
	result, err := bmiRepository.FindBmiCondition(valueFilter)
	if err != nil {
		t.Fatalf("FindBmiCondition unexpected error: %s", err)
	}

	expected := mockBmiConditions[0]
	if result.Id != expected.Id || result.BmiDesc != expected.BmiDesc {
		t.Errorf("unexpected result: got %+v, want %+v", result, expected)
	}

	assert.NoError(t, mock.ExpectationsWereMet())
}
