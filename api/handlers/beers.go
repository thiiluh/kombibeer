package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/thiiluh/kombibeer/internal/repository"
	"github.com/thiiluh/kombibeer/pkg/beers"
	"github.com/thiiluh/kombibeer/pkg/beers/usecase"
)

var usecaseBeer = usecase.NewBeerUseCase(repository.NewBeerRepository())

func BeersRoutes(router *gin.Engine) {
	router.GET("/beers", findAll)
	router.GET("/beers/:id", findOne)
	router.POST("/beers", create)
	router.PATCH("/beers/:id", update)
	router.PUT("/beers/:id", updateAllFileds)
	router.DELETE("/beers/:id", remove)
}

func findAll(c *gin.Context) {
	beers, err := usecaseBeer.FindAll()
	if err != nil {
		fmt.Errorf(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if len(beers) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.IndentedJSON(http.StatusOK, beers)
}

func findOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	beer, err := usecaseBeer.FindOne(id)

	if err != nil {
		if err.Error() == "record not found" {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, beer)
}

func create(c *gin.Context) {
	var newBeer beers.Beer

	if err := c.BindJSON(&newBeer); err != nil {
		fmt.Errorf(err.Error())
		c.Status(http.StatusBadRequest)
		return
	}

	beer, err := usecaseBeer.Create(newBeer)

	if err != nil {
		fmt.Errorf(err.Error())

		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, beer)
}

func update(c *gin.Context) {
	beers := beers.Beer{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err = c.BindJSON(&beers); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	updateBeers, err := usecaseBeer.Update(id, beers)

	if err != nil {
		if err.Error() == "resource not found" {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, updateBeers)
}

func updateAllFileds(c *gin.Context) {
	beers := beers.Beer{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err = c.BindJSON(&beers); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	updateBeers, err := usecaseBeer.UpdateAllFields(id, beers)

	if err != nil {
		if err.Error() == "resource not found" {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusNotFound)
		return
	}

	c.IndentedJSON(http.StatusOK, updateBeers)
}

func remove(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := usecaseBeer.Remove(id); err != nil {
		if err.Error() == "resource not found" {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusNoContent)
}
