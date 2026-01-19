package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"pizzeria/internal/data"
	"pizzeria/internal/handler"
	"pizzeria/internal/models"
)


var hardcoded_pizzas = [] models.Pizza {
	{ID: 1, Name: "[MODELS HARDCODED] Tuscany",          Price: 49.5},
	{ID: 2, Name: "[MODELS HARDCODED] Margherita",       Price: 79.5},
	{ID: 3, Name: "[MODELS HARDCODED] Tuna with cheese", Price: 69.5},
}

func main() {
	data.LoadPizzas()

	router := gin.Default()
	router.GET("fixed_pizzas", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"pizza01": "[ANON FUNC] TODO Get struct Tuscany",
			"pizza02": "[ANON FUNC] TODO Get struct Margherita",
			"pizza03": "[ANON FUNC] TODO Get struct Tuna with cheese",
		})
	})

	router.GET("pizzas",        handler.GetPizzas)
	router.POST("pizzas",       handler.PostPizza)
	router.GET("pizzas/:id",    handler.GetPizzaByID)
	router.DELETE("pizzas/:id", handler.DeletePizzaByID)
	router.PUT("pizzas/:id",    handler.UpdatePizzaByID)

	router.POST("pizzas/:id/reviews", handler.PostReview)

	router.POST("people", handler.PostPerson)
	
	router.Run(":8080")
}

func helloWorld() {
	var pizzashop_name string = "Go Pizza Shop"
	instagram, whatsapp := "@go_pizzashop", 15994785333
	msg01 := "Hello"
	msg02 :="World"
	fmt.Println(pizzashop_name, instagram, whatsapp)
	fmt.Printf("%s %s!! Go get the %s !!", msg01, pizzashop_name, msg02)

	fmt.Println(hardcoded_pizzas)
}
