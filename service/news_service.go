package service

import "belajar-mongo/model/payload"

type NewsService interface {
	Create(request payload.CreateNewsRequest)
	Update(slug string, request payload.UpdateNewsRequest)
	List(request payload.GetNewsRequest) []payload.NewsResponse
	GetBySlug(slug string) payload.NewsResponse
}
