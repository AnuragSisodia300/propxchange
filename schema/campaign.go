package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Campaign struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	UserID     primitive.ObjectID `json:"user_id" bson:"user_id"`
	PropertyID primitive.ObjectID `json:"property_id" bson:"property_id"`
	StartDate  time.Time          `json:"start_date" bson:"start_date"`
	EndDate    time.Time          `json:"end_date" bson:"end_date"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
