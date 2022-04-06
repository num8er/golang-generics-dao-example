package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	db "go-generics-dao-example/db/common/mysql"
	"go-generics-dao-example/routes"
)

func init() {
	_ = godotenv.Load()

	db.Connect()
}

func main() {
	r := gin.Default()
	routes.Attach(r)
	r.Run("127.0.0.1:8080")

	defer db.Close()
}
