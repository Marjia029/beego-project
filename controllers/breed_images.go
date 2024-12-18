package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type BreedImagesController struct {
	beego.Controller
}

type CatImage struct {
	URL string `json:"url"`
}

// Get handles the request to fetch breed images
func (c *BreedImagesController) Get() {
	breedID := c.GetString("breed_id") // Get breed_id from the query parameter
	if breedID == "" {
		c.Data["Error"] = "Breed ID is required."
		c.TplName = "breed_images.tpl"
		return
	}

	// Fetch images for the given breed
	apiKey := "live_GWXcPdnWze27MNMJSjinKshtfsnVsi4EdrXfKUNhOmXsLakl5N7MwJCShLvC5Rxo"
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=10", breedID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Data["Error"] = "Failed to create API request."
		c.TplName = "breed_images.tpl"
		return
	}
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Data["Error"] = "Failed to fetch breed images."
		c.TplName = "breed_images.tpl"
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Data["Error"] = "Failed to read API response."
		c.TplName = "breed_images.tpl"
		return
	}

	var images []CatImage
	if err := json.Unmarshal(body, &images); err != nil {
		c.Data["Error"] = "Failed to parse API response."
		c.TplName = "breed_images.tpl"
		return
	}

	c.Data["BreedImages"] = images
	c.TplName = "breed_images.tpl"
}
