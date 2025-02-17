package garden_categories_models

import "time"

type GardenCategories struct {
	ID             string `bson:"_id,omitempty" json:"_id"`
	Name string `bson:"name" json:"name"`
	BackgroundColor string `bson:"backgroundColor" json:"backgroundColor"`
	ImageURL string `bson:"imageURL" json:"imageURL"`
	CreatedDate time.Time `bson:"createdDate" json:"createdDate"`
	LastUpdate time.Time `bson:"lastUpdate" json:"lastUpdate"`
}