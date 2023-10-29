package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// SetupRoutes sets up all the API routes
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Public routes
	r.POST("/register", func(c *gin.Context) { Register(c, db) })
	r.POST("/login", func(c *gin.Context) { Login(c, db) })

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(ValidateToken)
	{
		// Event Types
		authorized.POST("/event-types", func(c *gin.Context) { CreateEventType(c, db) })
		authorized.GET("/event-types", func(c *gin.Context) { ListEventTypes(c, db) })
		authorized.PUT("/event-types/:id", func(c *gin.Context) { UpdateEventType(c, db) })
		authorized.DELETE("/event-types/:id", func(c *gin.Context) { DeleteEventType(c, db) })

		// Events
		authorized.POST("/events", func(c *gin.Context) { CreateEvent(c, db) })
		authorized.GET("/events", func(c *gin.Context) { ListEvents(c, db) })
		authorized.PUT("/events/:id", func(c *gin.Context) { UpdateEvent(c, db) })
		authorized.DELETE("/events/:id", func(c *gin.Context) { DeleteEvent(c, db) })

		// Additional Properties
		authorized.POST("/additional-props", func(c *gin.Context) { CreateAdditionalProp(c, db) })
		authorized.GET("/additional-props", func(c *gin.Context) { ListAdditionalProps(c, db) })

		// Event Type Properties
		authorized.POST("/event-type-props", func(c *gin.Context) { CreateEventTypeProp(c, db) })
		authorized.GET("/event-type-props", func(c *gin.Context) { ListEventTypeProps(c, db) })
	}
}
