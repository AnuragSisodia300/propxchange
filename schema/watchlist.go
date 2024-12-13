package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Watchlist represents the watchlist that stores user preferences for properties.
type Watchlist struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	PropertyID primitive.ObjectID `json:"property_id" bson:"property_id"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}
