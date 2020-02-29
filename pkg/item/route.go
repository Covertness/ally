package item

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/cockroachdb/apd"
	"github.com/gin-gonic/gin"
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

	newItem, err := Create(req.ExternalID, req.Price, req.Meta)
	if err != nil {
		if !strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{"error": "externalID has exists"})
			return
		}

		log.Printf("create item err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         newItem.ID,
		"externalID": newItem.ExternalID,
		"price":      newItem.Price,
		"meta":       newItem.Meta,
		"createdAt":  newItem.CreatedAt,
		"updatedAt":  newItem.UpdatedAt,
	})
}

type createReq struct {
	ExternalID string       `binding:"required"`
	Price      *apd.Decimal `binding:"required"`
	Meta       map[string]interface{}
}

func show(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		id = 0
	}

	mItem, err := GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "object not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         mItem.ID,
		"externalID": mItem.ExternalID,
		"price":      mItem.Price,
		"meta":       mItem.Meta,
		"createdAt":  mItem.CreatedAt,
		"updatedAt":  mItem.UpdatedAt,
	})
}
