package transactiongroup

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Register routes
func Register(router *gin.RouterGroup) {
	router.GET("/:id", show)
}

func show(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		id = 0
	}

	mGroup, err := GetByID(uint(id))

	transactions := []map[string]interface{}{}
	for _, t := range mGroup.Transactions {
		transactions = append(transactions, map[string]interface{}{
			"id":         t.ID,
			"address":    t.Address.Address,
			"status":     t.Status,
			"nonce":      t.Nonce,
			"hash":       t.Hash,
			"to":         t.To,
			"amount":     t.Amount,
			"meta":       t.Meta,
			"created_at": t.CreatedAt,
			"updated_at": t.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":           mGroup.ID,
		"type":         mGroup.Type,
		"transactions": transactions,
	})
}
