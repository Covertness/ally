package stat

import (
	"net/http"

	"github.com/Covertness/ally/pkg/account"

	"github.com/Covertness/ally/pkg/address"
	"github.com/Covertness/ally/pkg/storage"
	"github.com/cockroachdb/apd"
	"github.com/gin-gonic/gin"
)

// Register routes
func Register(router *gin.RouterGroup) {
	router.GET("", current)
}

func current(c *gin.Context) {
	type balanceResult struct {
		Amount *apd.Decimal
	}

	balance, err := address.GetTotalDebts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var accountBalance balanceResult
	if err := storage.GetDB().Table(storage.GetDB().NewScope(account.Account{}).TableName()).
		Select("sum(balance) as amount").Scan(&accountBalance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"realBalance":    balance.Text('f'),
		"accountBalance": accountBalance.Amount.Text('f'),
	})
}
