package api

import (
	"github.com/baklavametal/holodolor/backend/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// CreateAdditionalProp - Create a new additional property
func CreateAdditionalProp(c *gin.Context, db *gorm.DB) {
	var additionalProp model.AdditionalProp
	if err := c.ShouldBindJSON(&additionalProp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	db.Create(&additionalProp)
	c.JSON(http.StatusOK, gin.H{"message": "Additional property created successfully", "data": additionalProp})
}

// ListAdditionalProps - List all additional properties
func ListAdditionalProps(c *gin.Context, db *gorm.DB) {
	var additionalProps []model.AdditionalProp
	db.Find(&additionalProps)
	c.JSON(http.StatusOK, gin.H{"data": additionalProps})
}
