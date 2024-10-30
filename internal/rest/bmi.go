package rest

import (
	"net/http"

	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

type (
	ResponseError struct {
		Message string `json:"message"`
	}

	BmiService interface {
		CalculateResultBmi(req *domain.BmiReq) (*domain.BmiCondition, error)
	}

	BmiHandler struct {
		Service BmiService
	}
)

func NewBmiHandler(e *echo.Echo, bmiService BmiService) {
	handler := &BmiHandler{
		Service: bmiService,
	}
	e.POST("/calculateBmi", handler.CalculateBmi)
}

func isRequestValidate(m *domain.BmiReq) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (h *BmiHandler) CalculateBmi(c echo.Context) error {
	var bmiReq domain.BmiReq
	err := c.Bind(&bmiReq)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValidate(&bmiReq); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	bmiCondition, err := h.Service.CalculateResultBmi(&bmiReq)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, bmiCondition)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
