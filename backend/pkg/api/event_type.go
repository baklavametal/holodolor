package api

import (
	"github.com/baklavametal/holodolor/backend/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// CreateEventType - Create a new event type
func CreateEventType(c *gin.Context, db *gorm.DB) {
	var eventType model.EventType
	if err := c.ShouldBindJSON(&eventType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	db.Create(&eventType)
	c.JSON(http.StatusOK, gin.H{"message": "Event type created successfully", "data": eventType})
}

// ListEventTypes - List all event types
func ListEventTypes(c *gin.Context, db *gorm.DB) {
	var eventTypes []model.EventType
	db.Find(&eventTypes)
	c.JSON(http.StatusOK, gin.H{"data": eventTypes})
}

// UpdateEventType - Update an existing event type by ID
func UpdateEventType(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var eventType model.EventType

	if db.First(&eventType, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event type not found"})
		return
	}

	if err := c.ShouldBindJSON(&eventType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	db.Save(&eventType)
	c.JSON(http.StatusOK, gin.H{"message": "Event type updated successfully", "data": eventType})
}

// DeleteEventType - Delete an event type by ID
func DeleteEventType(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var eventType model.EventType

	if db.First(&eventType, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event type not found"})
		return
	}

	db.Delete(&eventType)
	c.JSON(http.StatusOK, gin.H{"message": "Event type deleted successfully"})
}
