package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiiluh/kombibeer/internal/config"

	"github.com/thiiluh/kombibeer/internal/repository"
	"github.com/thiiluh/kombibeer/pkg/beers"
	"github.com/thiiluh/kombibeer/pkg/beers/usecase"
)

var usercase = usecase.NewBeerUseCase(repository.NewBeerRepository())

func BeersRoutes(router *gin.Engine) {
	router.GET("/beers", findAll)
	router.GET("/beers/:id", findOne)
	router.POST("/beers", create)
	router.PUT("/beers", update)
	router.DELETE("/beers", remove)
}

func findAll(c *gin.Context) {
	beers, err := usercase.FindAll()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, beers)
}

func findOne(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "teste")
}

func create(c *gin.Context) {
	var newBeer beers.Beer

	if err := c.BindJSON(&newBeer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	config.DB.Create(&newBeer)

	c.IndentedJSON(http.StatusOK, newBeer)
}

func update(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "teste")
}

func remove(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "teste")
}
