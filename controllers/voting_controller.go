package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type VotingController struct {
	web.Controller
}

type VotingCatImage struct {
	URL string `json:"url"`
}

var favorites []string

// Get method to fetch a random cat image
func (c *VotingController) Get() {
	c.Data["ActiveTab"] = "voting"
	apiKey := "live_GWXcPdnWze27MNMJSjinKshtfsnVsi4EdrXfKUNhOmXsLakl5N7MwJCShLvC5Rxo"
	url := "https://api.thecatapi.com/v1/images/search"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Data["Error"] = "Failed to create API request."
		c.TplName = "voting.tpl"
		return
	}
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Data["Error"] = "Failed to fetch random cat image."
		c.TplName = "voting.tpl"
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Data["Error"] = "Failed to read API response."
		c.TplName = "voting.tpl"
		return
	}

	var images []VotingCatImage
	if err := json.Unmarshal(body, &images); err != nil {
		c.Data["Error"] = "Failed to parse image response."
		c.TplName = "voting.tpl"
		return
	}

	// Set the image URL to be displayed
	if len(images) > 0 {
		c.Data["ImageURL"] = images[0].URL
	}

	// Pass the favorites list to the template
	c.Data["Favorites"] = favorites

	c.TplName = "voting.tpl"
}

// Post method to handle like, dislike, and saving to favorites
func (c *VotingController) Post() {
	action := c.GetString("action")
	imageURL := c.GetString("image_url")

	switch action {
	case "like", "dislike":
		// Simply get a new image after the like/dislike action
		c.Redirect("/voting", 302)
	case "favorite":
		// Add the image to favorites
		favorites = append(favorites, imageURL)
		c.Redirect("/voting", 302)
	}
}
