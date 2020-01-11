package admin

import (
	"net/http"

	"github.com/Covertness/ally/pkg/address"
	"github.com/Covertness/ally/pkg/etherscan"
	"github.com/Covertness/ally/pkg/hdwallet"
	"github.com/Covertness/ally/pkg/transactiongroup"

	"github.com/cockroachdb/apd"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// Register routes
func Register(router *gin.RouterGroup) {
	router.GET("/rootAddress", rootAddress)
	router.POST("/withdraw", withdraw)
}

func rootAddress(c *gin.Context) {
	root, err := hdwallet.RootAddress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	balance, err := etherscan.GetInstance().GetBalance(root.Address.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"address": root.Address,
		"balance": balance,
	})
}

func withdraw(c *gin.Context) {
	var req withdrawReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	availableBalance, err := address.GetTotalDebts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if availableBalance.Cmp(req.Amount) < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not enough balance"})
		return
	}

	if !common.IsHexAddress(req.To) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid address"})
		return
	}

	group, err := transactiongroup.CreateAdminWithdrawGroup(req.Amount, req.To)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactionGroupID": group.ID,
	})
}

type withdrawReq struct {
	To     string       `binding:"required"`
	Amount *apd.Decimal `binding:"required"`
}
