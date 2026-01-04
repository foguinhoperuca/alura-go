package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"pizzeria/models"
)

var pizzas [] models.Pizza
var hardcoded_pizzas = [] models.Pizza {
	{ID: 1, Name: "[MODELS HARDCODED] Tuscany",          Price: 49.5},
	{ID: 2, Name: "[MODELS HARDCODED] Margherita",       Price: 79.5},
	{ID: 3, Name: "[MODELS HARDCODED] Tuna with cheese", Price: 69.5},
}

func main() {
	loadPizzas()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})
	router.GET("fixed_pizzas", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"pizza01": "[ANON FUNC] TODO Get struct Tuscany",
			"pizza02": "[ANON FUNC] TODO Get struct Margherita",
			"pizza03": "[ANON FUNC] TODO Get struct Tuna with cheese",
		})
	})
	router.GET("pizzas",    getPizzas)
	router.POST("pizzas",   postPizzas)
	router.GET("pizza/:id", getPizzaByID)
	router.Run()
}

func getPizzas(c *gin.Context) {
	c.JSON(200, gin.H {
		"pizzas": pizzas,
	})
}

func postPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H {
			"error": err.Error(),
		})
		return 
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	savePizza()
	c.JSON(201, newPizza)
}

func getPizzaByID(c *gin.Context) {
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

	for _, p := range pizzas {
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

func loadPizzas() {
	file, err := os.Open("data/pizzas.json")
	if err != nil {
		fmt.Println("Error file: ", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error *decoding* JSON: ", err)
	}
}

func savePizza() {
	file, err := os.Create("data/pizzas.json")
	if err != nil {
		fmt.Println("Error file: ", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error *encoding* JSON", err)
	}
	
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
