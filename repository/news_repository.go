package repository

import (
	"belajar-mongo/model/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type NewsRepository interface {
	Save(news domain.News)
	Update(filter bson.M, news bson.M)
	FindAll(filter bson.M) []domain.News
	FindBySlug(slug string) (domain.News, error)
}
