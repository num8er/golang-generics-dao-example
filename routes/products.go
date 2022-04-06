package routes

import (
	"github.com/gin-gonic/gin"

	"go-generics-dao-example/controllers/products"
)

func AttachProductsRoutes(r *gin.RouterGroup) {
	// r.GET("", products.List)

	r.POST("", products.Create)
	r.GET("/:id", products.Get)
	// r.PUT("/:id", products.Update)
	// r.DELETE("/:id", products.Delete)
}
