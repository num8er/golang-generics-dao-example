package routes

import (
	"github.com/gin-gonic/gin"

	"go-generics-dao-example/controllers/customers"
)

func AttachCustomersRoutes(r *gin.RouterGroup) {
	// r.GET("", customers.List)

	r.POST("", customers.Create)
	r.GET("/:id", customers.Get)
	// r.PUT("/:id", customers.Update)
	// r.DELETE("/:id", customers.Delete)
}
