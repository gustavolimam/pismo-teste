package account

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gustavolimam/pismo-teste/database"

	"github.com/labstack/echo/v4"
)

func createAccount(c echo.Context) error {
	log.Println("Iniciando processo de criação de uma nova conta")

	account := new(struct {
		DocNumber string `json:"document_number"`
	})
	if err := c.Bind(account); err != nil {
		return c.JSON(http.StatusBadRequest, "Erro nos dados de account recebidos!")
	}

	ID, err := database.InsertAccount(account.DocNumber)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintln("Erro ao tentar inserir os dados da conta", err),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"ID":      ID,
		"Message": "Sucesso",
	})
}

func readAccount(c echo.Context) error {
	log.Println("Iniciando processo para retornar dados da conta:", c.Param("account_id"))

	accountResponse, err := database.GetAccount(c.Param("account_id"))
	if err != nil {
		if err != sql.ErrNoRows {
			return c.JSON(http.StatusBadRequest, &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintln("Erro ao tentar retornar os dados da conta", err),
			})
		}
	}

	return c.JSON(http.StatusOK, accountResponse)
}
