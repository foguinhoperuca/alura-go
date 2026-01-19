package service

import (
	"errors"
	"pizzeria/internal/models"
)

func ValidateReviewRating(review *models.Review) error {
	if review.Rating < 0 || review.Rating > 5 {
		return errors.New("Rating must be between 1 and 5!")
	}

	return nil
}
