package route

import (
	"encoding/json"
	"net/http"

	"github.com/dev-parvej/go-blog/config"
	"github.com/dev-parvej/go-blog/dto"
	"github.com/dev-parvej/go-blog/helper"
	model "github.com/dev-parvej/go-blog/models"
)

func CreatePostForm(w http.ResponseWriter, r *http.Request) {
	var session = config.Session(r)
	data := map[string]interface{}{
		"Title":         "Create post",
		"Authenticated": session.Values["authenticated"],
	}
	config.Views().ExecuteTemplate(w, "create-post", "main", data)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var session = config.Session(r)
	userId := session.Values["user_id"]

	var user = model.User{}
	var postDto = dto.PostCreate{}

	model.DB().Model(model.User{}).Where("id=?", userId).First(&user)

	helper.ParseForm(r, &postDto)
	err := helper.ValidateForm(postDto)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	post := model.Post{
		UserID: int(user.Model.ID),
		Title:  postDto.Title,
		Body:   postDto.Body,
	}

	model.DB().Create(&post)

	http.Redirect(w, r, "/posts/new", http.StatusMovedPermanently)
}
