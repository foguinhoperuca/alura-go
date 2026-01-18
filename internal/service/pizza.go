package service

import (
	"errors"
	"pizzeria/internal/models"
)

func ValidatePizzaPrice(pizza *models.Pizza) error {
	if pizza.Price < 0 {
		return errors.New("The pizza's price can't be negative")
	}

	return nil
}
