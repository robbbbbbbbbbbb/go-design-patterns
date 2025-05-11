package models

import "time"

type DogBreed struct {
	ID               int        `json:"id"`
	Breed            string     `json:"breed"`
	WeightLowLbs     float64    `json:"weight_low_lbs"`
	WeightHighLbs    float64    `json:"weight_high_lbs"`
	AverageWeight    float64    `json:"average_weight"`
	LifeSpan         float64    `json:"life_span"`
	Details          string     `json:"details"`
	AlternateNames   string     `json:"alternate_names"`
	GeographicOrigin string     `json:"geographic_origin"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

type CatBreed struct {
	ID               int        `json:"id"`
	Breed            string     `json:"breed"`
	WeightLowLbs     float64    `json:"weight_low_lbs"`
	WeightHighLbs    float64    `json:"weight_high_lbs"`
	AverageWeight    float64    `json:"average_weight"`
	LifeSpan         float64    `json:"lifespan"`
	Details          string     `json:"details"`
	AlternateNames   string     `json:"alternate_names"`
	GeographicOrigin string     `json:"geographic_origin"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

type Dog struct {
	ID               int        `json:"id"`
	DogName          string     `json:"dog_name"`
	BreedID          int        `json:"breed_id"`
	BreederID        int        `json:"breeder_id"`
	Color            string     `json:"color"`
	DateOfBirth      *time.Time `json:"date_of_birth"`
	SpayedOrNeutered int        `json:"spayed_or_neutered"`
	Description      string     `json:"description"`
	Weight           float64    `json:"weight"`
	Breed            DogBreed   `json:"breed"`
	Breeder          Breeder    `json:"breeder"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

type Cat struct {
	ID               int        `json:"id"`
	CatName          string     `json:"cat_name"`
	BreedID          int        `json:"breed_id"`
	BreederID        int        `json:"breeder_id"`
	Color            string     `json:"color"`
	DateOfBirth      *time.Time `json:"date_of_birth"`
	SpayedOrNeutered int        `json:"spayed_or_neutered"`
	Description      string     `json:"description"`
	Weight           float64    `json:"weight"`
	Breed            CatBreed   `json:"breed"`
	Breeder          Breeder    `json:"breeder"`
	CreatedAt        *time.Time `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

type Breeder struct {
	ID          int         `json:"int"`
	BreederName string      `json:"breeder_name"`
	Address     string      `json:"address"`
	City        string      `json:"city"`
	ProvState   string      `json:"prov_state"`
	Country     string      `json:"country"`
	ZipCode     string      `json:"zip_code"`
	PhoneNumber string      `json:"phone_number"`
	Email       string      `json:"email"`
	Status      int         `json:"status"`
	DogBreeds   []*DogBreed `json:"dog_breeds"`
	CatBreeds   []*CatBreed `json:"cat_breeds"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at"`
}

type Pet struct {
	Species     string  `json:"species"`
	Breed       string  `json:"breed"`
	MinWeight   float64 `json:"min_weight"`
	MaxWeight   float64 `json:"max_weight"`
	Description string  `json:"description"`
	LifeSpan    float64 `json:"lifespan"`
}
