package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"pizzeria/internal/data"
	"pizzeria/internal/models"
)

func GetPizzas(c *gin.Context) {
	c.JSON(200, gin.H {
		"pizzas": data.Pizzas,
	})
}

func PostPizza(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H {
			"error": err.Error(),
		})
		return 
	}
	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizza()
	c.JSON(201, newPizza)
}

func GetPizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	fmt.Println("idParam: ", idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		error_message := fmt.Sprintf("[ERROR] %s", err.Error())
		c.JSON(400, gin.H {
			"message": error_message,
		})
		return
	}

	for _, p := range data.Pizzas {
		if p.ID == id {
			c.JSON(200, gin.H {
				"found": id,
				"pizza": p,
			})
			return
		}
	}
	error_message := fmt.Sprintf("ID %d not found!", id)				// FIXME doing python in go?!
	c.JSON(404, gin.H {
		"message": error_message,
	})

}

func DeletePizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		error_message := fmt.Sprintf("[ERROR] %s", err.Error())
		c.JSON(400, gin.H {
			"message": error_message,
		})
		return
	}

	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizza()
			message := fmt.Sprintf("Pizza %d was deleted", id)
			c.JSON(200, gin.H {
				"message": message,
			})
			return
		}
	}
	c.JSON(404, gin.H {"message": "Pizza *NOT* found!!"})
}

func UpdatePizzaByID(c *gin.Context) {
	var updatedPizza models.Pizza
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		error_message := fmt.Sprintf("[ERROR] %s", err.Error())
		c.JSON(400, gin.H {
			"message": error_message,
		})
		return
	}

	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(400, gin.H{"erro": err.Error()})
		return
	}

	for i, p:= range data.Pizzas {
		if p.ID == id {
			data.Pizzas[i] = updatedPizza
			data.SavePizza()
			message := fmt.Sprintf("Pizza # %d was updated", id)
			c.JSON(200, gin.H {
				"message": message,
				"savedPizza": data.Pizzas[i],
			})
			return
		}
	}

	c.JSON(404, gin.H {
		"method": "pizza not found!",
	})
}
