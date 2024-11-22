package http

import (
	"net/http"
	"sync"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (c *CustomValidator) Validate(i interface{}) error {
	c.lazyInit()
	if err := c.validate.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (c *CustomValidator) lazyInit() {
	c.once.Do(func() {
		c.validate = validator.New()
	})
}
