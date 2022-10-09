package dto

type PostCreate struct {
	Title string `json:"title" validate:"required,min=1,max=255" mapstructure:"title"`
	Body  string `json:"body" validate:"required,min=10" mapstructure:"body"`
}
