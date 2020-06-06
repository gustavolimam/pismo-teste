package transaction

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gustavolimam/pismo-teste/database"
	"github.com/gustavolimam/pismo-teste/types"
	"github.com/labstack/echo/v4"
)

func createTransaction(c echo.Context) error {
	log.Println("Iniciando processo de criação de uma nova transação")

	tr := new(types.Transaction)
	if err := c.Bind(tr); err != nil {
		return c.JSON(http.StatusBadRequest, "Erro nos dados de transação recebidos!")
	}

	switch tr.OperationTypeID {
	case 1:
		tr.Amount = -tr.Amount

	case 2:
		tr.Amount = -tr.Amount

	case 3:
		tr.Amount = -tr.Amount
	}

	if err := database.InsertTransaction(*tr); err != nil {
		return c.JSON(http.StatusBadRequest, &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintln("Erro ao tentar inserir os dados da transação", err),
		})
	}

	return c.JSON(http.StatusOK, "Sucesso")
}
