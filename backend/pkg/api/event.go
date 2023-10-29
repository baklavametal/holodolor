package api

import (
	"github.com/baklavametal/holodolor/backend/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// CreateEvent - Create a new event
func CreateEvent(c *gin.Context, db *gorm.DB) {
	var event model.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	db.Create(&event)
	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully", "data": event})
}

// ListEvents - List all events
func ListEvents(c *gin.Context, db *gorm.DB) {
	var events []model.Event
	db.Find(&events)
	c.JSON(http.StatusOK, gin.H{"data": events})
}

// UpdateEvent - Update an existing event by ID
func UpdateEvent(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var event model.Event

	if db.First(&event, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	db.Save(&event)
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "data": event})
}

// DeleteEvent - Delete an event by ID
func DeleteEvent(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")
	var event model.Event

	if db.First(&event, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	db.Delete(&event)
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
