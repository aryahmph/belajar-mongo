package controller

import (
	"belajar-mongo/model/payload"
	"belajar-mongo/pkg/exception"
	"belajar-mongo/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type NewsController struct {
	NewsService service.NewsService
}

func NewNewsController(newsService service.NewsService) *NewsController {
	return &NewsController{NewsService: newsService}
}

func (controller *NewsController) Route(app *fiber.App) {
	app.Post("/api/news", controller.Create)
	app.Get("/api/news", controller.List)
	app.Get("/api/news/:slug", controller.GetBySlug)
	app.Patch("/api/news/:slug", controller.Update)
}

func (controller *NewsController) Create(ctx *fiber.Ctx) error {
	request := payload.CreateNewsRequest{
		Title:    strings.Trim(ctx.FormValue("title"), " "),
		Content:  strings.Trim(ctx.FormValue("content"), " "),
		Category: strings.Trim(ctx.FormValue("category"), " "),
		Tags:     strings.Trim(ctx.FormValue("tags"), " "),
	}

	file, err := ctx.FormFile("image")
	exception.PanicIfNeeded(err)
	request.ImageURL = file.Filename

	err = ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
	exception.PanicIfNeeded(err)

	controller.NewsService.Create(request)
	return ctx.JSON(payload.WebResponse{
		Code:   200,
		Status: "OK",
		Error:  "",
	})
}

func (controller *NewsController) Update(ctx *fiber.Ctx) error {
	request := payload.UpdateNewsRequest{
		Title:    strings.Trim(ctx.FormValue("title"), " "),
		Content:  strings.Trim(ctx.FormValue("content"), " "),
		Category: strings.Trim(ctx.FormValue("category"), " "),
		Tags:     strings.Trim(ctx.FormValue("tags"), " "),
	}

	file, err := ctx.FormFile("image")
	if err == nil {
		request.ImageURL = file.Filename
		err = ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))
		exception.PanicIfNeeded(err)
	}

	controller.NewsService.Update(ctx.Params("slug"), request)
	return ctx.JSON(payload.WebResponse{
		Code:   200,
		Status: "OK",
		Error:  "",
	})
}

func (controller *NewsController) List(ctx *fiber.Ctx) error {
	var request payload.GetNewsRequest
	err := ctx.QueryParser(&request)
	exception.PanicIfNeeded(err)

	responses := controller.NewsService.List(request)
	return ctx.JSON(payload.WebResponse{
		Code:   200,
		Status: "OK",
		Error:  "",
		Data:   responses,
	})
}

func (controller *NewsController) GetBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	response := controller.NewsService.GetBySlug(slug)
	return ctx.JSON(payload.WebResponse{
		Code:   200,
		Status: "OK",
		Error:  "",
		Data:   response,
	})
}
