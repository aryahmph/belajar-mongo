package payload

type GetNewsRequest struct {
	Category string `json:"category"`
	Tag      string `json:"tag"`
}

type CreateNewsRequest struct {
	Title    string `json:"title" validate:"required"`
	Slug     string `json:"slug"`
	ImageURL string `json:"image_url" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Category string `json:"category" validate:"required"`
	Tags     string `json:"tags" validate:"required"`
}

type UpdateNewsRequest struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Tags     string `json:"tags"`
}

type NewsResponse struct {
	Id       string        `json:"id"`
	Title    string        `json:"title"`
	Slug     string        `json:"slug"`
	ImageURL string        `json:"image_url"`
	Content  string        `json:"content"`
	Category string        `json:"category"`
	Tags     []interface{} `json:"tags"`
}
