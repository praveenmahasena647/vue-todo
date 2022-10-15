package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDo struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Act  string             `json:"act"`
	Done bool               `json:"done"`
}
