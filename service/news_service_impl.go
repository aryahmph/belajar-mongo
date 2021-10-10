package service

import (
	"belajar-mongo/model/domain"
	"belajar-mongo/model/payload"
	"belajar-mongo/pkg/exception"
	"belajar-mongo/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

type NewsServiceImpl struct {
	NewsRepository repository.NewsRepository
	validate       *validator.Validate
}

func NewNewsServiceImpl(newsRepository repository.NewsRepository, validate *validator.Validate) *NewsServiceImpl {
	return &NewsServiceImpl{NewsRepository: newsRepository, validate: validate}
}

func (service *NewsServiceImpl) Create(request payload.CreateNewsRequest) {
	err := service.validate.Struct(request)
	exception.PanicIfNeeded(err)

	// Check news already exist or not
	_, err = service.NewsRepository.FindBySlug(request.Slug)
	if err != nil || err != mongo.ErrNoDocuments {
		exception.PanicIfNeeded(exception.AlreadyExistError)
	}

	// Tags from []string to []interface{}
	tagSlice := strings.Split(request.Tags, ",")
	tags := make([]interface{}, len(tagSlice))
	for i, value := range tagSlice {
		tags[i] = value
	}

	news := domain.News{
		Title:    request.Title,
		Slug:     request.Slug,
		ImageURL: request.ImageURL,
		Content:  request.Content,
		Category: request.Category,
		Tags:     tags,
	}

	if news.Slug == "" {
		news.Slug = slug.Make(news.Title)
	}

	service.NewsRepository.Save(news)
}

func (service *NewsServiceImpl) Update(slug string, request payload.UpdateNewsRequest) {
	// Check news already exist or not
	news, err := service.NewsRepository.FindBySlug(slug)
	if err == mongo.ErrNoDocuments {
		exception.PanicIfNeeded(exception.NotFoundError)
	}
	exception.PanicIfNeeded(err)

	// Replace data
	if request.Title != "" {
		news.Title = request.Title
	}
	if request.ImageURL != "" {
		news.ImageURL = request.ImageURL
	}
	if request.Content != "" {
		news.Content = request.Content
	}
	if request.Category != "" {
		news.Category = request.Category
	}
	if request.Tags != "" {
		// Tags from []string to []interface{}
		tagSlice := strings.Split(request.Tags, ",")
		tags := make([]interface{}, len(tagSlice))
		for i, value := range tagSlice {
			tags[i] = value
		}
		news.Tags = tags
	}

	filter := bson.M{"slug": slug}

	updateStatement := bson.M{"$set" : bson.M{
		"title":     news.Title,
		"image_url": news.ImageURL,
		"content":   news.Content,
		"category":  news.Category,
		"tags":      news.Tags,
	}}

	service.NewsRepository.Update(filter, updateStatement)
}

func (service *NewsServiceImpl) List(request payload.GetNewsRequest) []payload.NewsResponse {
	filter := bson.M{}
	if request.Category != "" {
		filter = bson.M{"category": request.Category}
	} else if request.Tag != "" {
		filter = bson.M{"tags": request.Tag}
	}

	news := service.NewsRepository.FindAll(filter)
	var newsResponses []payload.NewsResponse
	for _, item := range news {
		newsResponses = append(newsResponses, payload.NewsResponse{
			Id:       item.Id,
			Title:    item.Title,
			Slug:     item.Slug,
			ImageURL: "http://localhost:3000/uploads/" + item.ImageURL,
			Content:  item.Content,
			Category: item.Category,
			Tags:     item.Tags,
		})
	}

	return newsResponses
}

func (service *NewsServiceImpl) GetBySlug(slug string) payload.NewsResponse {
	news, err := service.NewsRepository.FindBySlug(slug)
	if err == mongo.ErrNoDocuments {
		err = exception.NotFoundError
	}
	exception.PanicIfNeeded(err)

	return payload.NewsResponse{
		Id:       news.Id,
		Title:    news.Title,
		Slug:     news.Slug,
		ImageURL: "http://localhost:3000/uploads/" + news.ImageURL,
		Content:  news.Content,
		Category: news.Category,
		Tags:     news.Tags,
	}
}
