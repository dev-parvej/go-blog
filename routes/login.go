package route

import (
	"encoding/json"
	"net/http"

	"github.com/dev-parvej/go-blog/config"
	"github.com/dev-parvej/go-blog/dto"
	"github.com/dev-parvej/go-blog/helper"
	model "github.com/dev-parvej/go-blog/models"
)

func LoginForm(w http.ResponseWriter, r *http.Request) {
	var session = config.Session(r)
	data := map[string]interface{}{
		"Title":         "Login",
		"Authenticated": session.Values["authenticated"],
	}
	config.Views().ExecuteTemplate(w, "sign-in", "main", data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var session = config.Session(r)
	defer r.Body.Close()

	loginDto := dto.LoginDto{}
	helper.ParseForm(r, &loginDto)

	errors := helper.ValidateForm(loginDto)
	if errors != nil {
		json.NewEncoder(w).Encode(errors.Error())
		return
	}

	var user model.User
	model.DB().Model(model.User{Email: loginDto.Email}).First(&user)

	if !helper.ComparePassword(user.Password, loginDto.Password) {
		json.NewEncoder(w).Encode("Incorrect user name or password")
		return
	}
	session.Values["authenticated"] = true
	session.Values["user_id"] = user.Model.ID
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
