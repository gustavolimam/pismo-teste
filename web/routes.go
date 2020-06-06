package web

import (
	"github.com/gustavolimam/pismo-teste/internal/account"
	"github.com/gustavolimam/pismo-teste/internal/transaction"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes function to start server
func RegisterRoutes(e *echo.Echo) {

	account.Load(e)
	transaction.Load(e)
}
