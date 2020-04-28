package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string        `json:"title,omitempty" bson:"title,omitempty"`
	CreatedAt time.Time     `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
