package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thiiluh/kombibeer/api/handlers"
)

func InitRoutes() error {
	r := gin.Default()
	handlers.BeersRoutes(r)
	return r.Run()
}
