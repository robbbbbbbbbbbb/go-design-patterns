package pets

import "go-breeders/models"

func New(species string) *models.Pet {
	return &models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0.0,
		MaxWeight:   0.0,
		Description: "no description yet",
		LifeSpan:    0.0,
	}
}
