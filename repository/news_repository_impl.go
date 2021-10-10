package repository

import (
	"belajar-mongo/model/domain"
	"belajar-mongo/pkg/database"
	"belajar-mongo/pkg/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewsRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewNewsRepositoryImpl(database *mongo.Database) *NewsRepositoryImpl {
	return &NewsRepositoryImpl{Collection: database.Collection("news")}
}

func (repository *NewsRepositoryImpl) Save(news domain.News) {
	ctx, cancel := database.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"title":     news.Title,
		"slug":      news.Slug,
		"image_url": news.ImageURL,
		"content":   news.Content,
		"category":  news.Category,
		"tags":      news.Tags,
	})
	exception.PanicIfNeeded(err)
}

func (repository *NewsRepositoryImpl) Update(filter bson.M, news bson.M) {
	ctx, cancel := database.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.UpdateOne(ctx, filter, news)
	exception.PanicIfNeeded(err)
}

func (repository *NewsRepositoryImpl) FindAll(filter bson.M) []domain.News {
	ctx, cancel := database.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, filter)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	var news []domain.News
	for _, document := range documents {
		news = append(news, domain.News{
			Id:       document["_id"].(primitive.ObjectID).Hex(),
			Title:    document["title"].(string),
			Slug:     document["slug"].(string),
			ImageURL: document["image_url"].(string),
			Content:  document["content"].(string),
			Category: document["category"].(string),
			Tags:     document["tags"].(primitive.A),
		})
	}

	return news
}

func (repository *NewsRepositoryImpl) FindBySlug(slug string) (domain.News, error) {
	ctx, cancel := database.NewMongoContext()
	defer cancel()

	var news domain.News
	err := repository.Collection.FindOne(ctx, bson.M{"slug": slug}).Decode(&news)

	return news, err
}
