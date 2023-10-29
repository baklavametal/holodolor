package api

import (
	"github.com/baklavametal/holodolor/backend/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// CreateEventTypeProp - Create a new event type property
func CreateEventTypeProp(c *gin.Context, db *gorm.DB) {
	var eventTypeProp model.EventTypeProp
	if err := c.ShouldBindJSON(&eventTypeProp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	db.Create(&eventTypeProp)
	c.JSON(http.StatusOK, gin.H{"message": "Event type property created successfully", "data": eventTypeProp})
}

// ListEventTypeProps - List all event type properties
func ListEventTypeProps(c *gin.Context, db *gorm.DB) {
	var eventTypeProps []model.EventTypeProp
	db.Find(&eventTypeProps)
	c.JSON(http.StatusOK, gin.H{"data": eventTypeProps})
}
