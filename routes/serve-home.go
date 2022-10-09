package route

import (
	"fmt"
	"net/http"

	"github.com/dev-parvej/go-blog/config"
	model "github.com/dev-parvej/go-blog/models"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	var session = config.Session(r)

	var posts []model.Post
	model.DB().Model(&model.Post{}).Select("id", "title", "SUBSTRING(`body`, 1, 140) as body", "user_id").Preload("User").Order("id DESC").Find(&posts)

	for _, post := range posts {
		fmt.Println("SS", post.User)
	}
	data := map[string]interface{}{
		"Title":         "Blog with golang",
		"Authenticated": session.Values["authenticated"],
		"posts":         posts,
	}

	config.Views().ExecuteTemplate(w, "index", "main", data)
}
