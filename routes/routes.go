package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"kirin.com/subscription-service/controller"
)

// RegisterRoutes sets up the routes for the API
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	controller.RegisterUserRoutes(router, db)
	controller.RegisterPlanRoutes(router, db)
	controller.RegisterSubscriptionRoutes(router, db)
	controller.RegisterBillingRoutes(router, db)
}
