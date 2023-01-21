package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id        primitive.ObjectID `bson:"_id"`
	Message   string             `json:"message" validate:"required,min=2,max=200"`
	Channel   string             `json:"channel"`
	Status    string             `json:"status" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

const (
	PENDING   string = "PENDING"
	SENT             = "SENT"
	CANCELLED        = "CANCELLED"
)
