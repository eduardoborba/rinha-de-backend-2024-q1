package actions

import (
	"net/http"
	"rinha_de_backend/models"

	"github.com/gobuffalo/buffalo"
)

// TransactionsCreateHandler is a the handler to create a transaction
func TransactionsCreateHandler(c buffalo.Context) error {
	client := models.Client{}
	err := models.DB.Find(&client, c.Param("id"))
	if err != nil {
		return c.Render(http.StatusNotFound, r.JSON(map[string]string{"message": "Client not found"}))
	}

	tx := &models.Transaction{}
	if err := c.Bind(tx); err != nil {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]string{"message": "Failed to parse request"}))
	}

	tx.ClientID = client.ID

	if err := models.DB.Save(tx); err != nil {
		return c.Render(http.StatusUnprocessableEntity, r.JSON(map[string]string{"message": "Failed to create transaction"}))
	}
	return c.Render(http.StatusOK, r.JSON(map[string]interface{}{"limit": client.Limit, "saldo": client.Balance}))
}
