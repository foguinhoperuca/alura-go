package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"pizzeria/internal/data"
	"pizzeria/internal/models"
	"pizzeria/internal/service"
)

func PostReview(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		error_message := fmt.Sprintf("[ERROR] %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H {
			"message": error_message,
		})
		return
	}

	var newReview models.Review
	if err := c.ShouldBindJSON(&newReview); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if err := service.ValidateReviewRating(&newReview); err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H {
			"error": err.Error(),
		})

		return
	}

	for i, p:= range data.Pizzas {
		if p.ID == id {
			p.Reviews = append(p.Reviews, newReview)
			data.Pizzas[i] = p
			data.SavePizza()
			message := fmt.Sprintf("Pizza # %d was updated", id)
			c.JSON(http.StatusCreated, gin.H {
				"message": message,
				"savedPizza": data.Pizzas[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H {
		"method": "pizza not found!",
	})
}
