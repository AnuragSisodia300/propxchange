package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Property struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Description  string             `bson:"description" json:"description"`
	Price        float64            `bson:"price" json:"price"`
	Location     string             `bson:"location" json:"location"`
	Bedrooms     int                `bson:"bedrooms" json:"bedrooms"`
	Bathrooms    int                `bson:"bathrooms" json:"bathrooms"`
	SquareFeet   float64            `bson:"square_feet" json:"square_feet"`
	OwnerID      primitive.ObjectID `bson:"owner_id" json:"owner_id"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	Types        []string           `bson:"types" json:"types"`
	MinIRR       float64            `bson:"min_irr" json:"min_irr"`
	MaxIRR       float64            `bson:"max_irr" json:"max_irr"`
	PerUnitPrice float64            `bson:"per_unit_price" json:"per_unit_price"`
	FundRaised   float64            `bson:"fund_raised" json:"fund_raised"`
	Units        int                `bson:"units" json:"units"`
}
