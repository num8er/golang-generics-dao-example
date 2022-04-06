package routes

import (
	"github.com/gin-gonic/gin"
)

func Attach(r *gin.Engine) {
	AttachCustomersRoutes(r.Group("/customers"))
	AttachProductsRoutes(r.Group("/products"))
}
