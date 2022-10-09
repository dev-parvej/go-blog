package dto

type LoginDto struct {
	Email    string `validate:"required" json:"email" mapstructure:"email"`
	Password string `validate:"required" json:"password" mapstructure:"password"`
}
