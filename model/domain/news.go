package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type News struct {
	Id       string `bson:"_id"`
	Title    string
	Slug     string
	ImageURL string `bson:"image_url"`
	Content  string
	Category string
	Tags     primitive.A
}
