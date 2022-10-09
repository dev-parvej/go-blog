package dto

type UserRegister struct {
	Name     string `validate:"required" json:"name" mapstructure:"name"`
	Email    string `validate:"required,email" json:"email" mapstructure:"email"`
	Password string `validate:"required,min=6,max=12" json:"password" mapstructure:"password"`
}
