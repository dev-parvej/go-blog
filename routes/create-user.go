package route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dev-parvej/go-blog/config"
	"github.com/dev-parvej/go-blog/dto"
	"github.com/dev-parvej/go-blog/helper"
	model "github.com/dev-parvej/go-blog/models"
)

func CreateUserForm(w http.ResponseWriter, r *http.Request) {
	var session = config.Session(r)
	data := map[string]interface{}{
		"Title":         "Sign up",
		"Authenticated": session.Values["authenticated"],
	}
	config.Views().ExecuteTemplate(w, "sign-up", "main", data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var userRegister dto.UserRegister
	helper.ParseForm(r, &userRegister)

	errors := helper.ValidateForm(userRegister)

	if errors != nil {
		json.NewEncoder(w).Encode(errors.Error())
		return
	}

	hashed, _ := helper.HashPassword(userRegister.Password)
	fmt.Println(hashed)
	user := model.User{
		FullName: userRegister.Name,
		Email:    userRegister.Email,
		Password: hashed,
	}

	model.DB().Create(&user)

	http.Redirect(w, r, "/sign-in", http.StatusMovedPermanently)
}
