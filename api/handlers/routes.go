package handlers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() error {
	r := gin.Default()
	BeersRoutes(r)
	return r.Run()
}
