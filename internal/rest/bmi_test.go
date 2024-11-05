package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/internal/rest"
	"github.com/bxcodec/go-clean-arch/internal/rest/mocks"
	faker "github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCalculateBmi(t *testing.T) {
	var mockBmiCondition domain.BmiCondition
	err := faker.FakeData(&mockBmiCondition)
	assert.NoError(t, err)

	mockBmiService := new(mocks.BmiService)
	mockBmiService.On("CalculateResultBmi", mock.Anything).Return(mockBmiCondition, nil)

	bmiReq := domain.BmiReq{
		Weight: 70,
		Height: 175,
	}
	body, err := json.Marshal(bmiReq)
	require.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/calculateBmi", bytes.NewBuffer(body))
	require.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	handler := rest.BmiHandler{
		Service: mockBmiService,
	}
	err = handler.CalculateBmi(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockBmiService.AssertExpectations(t)
}
