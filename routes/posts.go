package route

import (
	"net/http"

	"github.com/dev-parvej/go-blog/config"
	model "github.com/dev-parvej/go-blog/models"
	"github.com/gorilla/mux"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var session = config.Session(r)
	params := mux.Vars(r)

	post := model.Post{}

	model.DB().Preload("User").First(&post, params["id"])

	data := map[string]interface{}{
		"Title":         "Login",
		"Authenticated": session.Values["authenticated"],
		"post":          post,
	}

	config.Views().ExecuteTemplate(w, "post", "main", data)
}
