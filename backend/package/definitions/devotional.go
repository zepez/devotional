package definitions

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Devotional struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Source      string             `json:"source" bson:"source"`
	Name        string             `json:"name" bson:"name"`
	Html        string             `json:"html" bson:"html"`
	Plain       string             `json:"plain_text" bson:"plain_text"`
	Image       string             `json:"image" bson:"image"`
	Target_Date string             `json:"target_date" bson:"target_date"`
	Created_at  time.Time          `json:"created_at" bson:"created_at"`
	Updated_at  time.Time          `json:"updated_at" bson:"updated_at"`
}
