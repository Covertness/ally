package order

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Covertness/ally/pkg/item"

	"github.com/Covertness/ally/pkg/account"
	"github.com/Covertness/ally/pkg/storage"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Register routes
func Register(router *gin.RouterGroup) {
	router.POST("", create)
	router.GET("/:id", show)
}

func create(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var mItem item.Item
	if err := storage.GetDB().First(&mItem, req.ItemID).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var acc account.Account
	if err := storage.GetDB().FirstOrCreate(&acc, account.Account{
		Name: req.Account,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newOrder, err := Create(req.Sequence, &acc, &mItem)
	if err != nil {
		if !strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{"error": "sequence has exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":        newOrder.ID,
		"status":    newOrder.Status,
		"createdAt": newOrder.CreatedAt,
		"updatedAt": newOrder.UpdatedAt,
	})
}

type createReq struct {
	ItemID   uint   `binding:"required"`
	Sequence uint   `binding:"required"`
	Account  string `binding:"required"`
}

func show(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		id = 0
	}

	mOrder, err := GetByID(uint(id))

	c.JSON(http.StatusOK, gin.H{
		"id":        mOrder.ID,
		"status":    mOrder.Status,
		"address":   mOrder.Address.Address,
		"createdAt": mOrder.CreatedAt,
		"updatedAt": mOrder.UpdatedAt,
	})
}
