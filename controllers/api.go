package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

// Post represents a post structure from the JSONPlaceholder API
type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	// Fetch data from the JSONPlaceholder API
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error fetching data: %s", err))
		return
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error decoding response: %s", err))
		return
	}

	// Pass the posts data to the template
	c.Data["Posts"] = posts
	c.TplName = "posts.tpl"
}
