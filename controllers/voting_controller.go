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

	// Create channels for concurrent fetching
	imageChan := make(chan VotingCatImage)
	errChan := make(chan error)

	// Fetch random cat image concurrently
	go func() {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			errChan <- err
			return
		}
		req.Header.Set("x-api-key", apiKey)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}

		var images []VotingCatImage
		if err := json.Unmarshal(body, &images); err != nil {
			errChan <- err
			return
		}

		if len(images) > 0 {
			imageChan <- images[0] // Send the first image to the channel
		}
	}()

	// Wait for either the image or an error
	select {
	case img := <-imageChan:
		// Set the image URL to be displayed
		c.Data["ImageURL"] = img.URL
	case err := <-errChan:
		// Handle error
		c.Data["Error"] = "Failed to fetch cat image: " + err.Error()
		c.TplName = "voting.tpl"
		return
	}

	// Pass the favorites list to the template
	c.Data["Favorites"] = favorites

	c.TplName = "voting.tpl"
}

// Post method to handle like, dislike, and saving to favorites
func (c *VotingController) Post() {
	action := c.GetString("action")
	imageURL := c.GetString("image_url")

	// Handle favorite action
	if action == "favorite" {
		// Add the image to favorites
		favorites = append(favorites, imageURL)
	}

	// Fetch a new random cat image if like/dislike action
	if action == "like" || action == "dislike" || action == "favorite" {
		c.Get() // This will fetch a new random image and update `c.Data["ImageURL"]`
	}

	// Send the updated image URL and favorites list as a JSON response
	c.Data["json"] = map[string]interface{}{"image_url": c.Data["ImageURL"], "favorites": favorites}
	c.ServeJSON()
}
