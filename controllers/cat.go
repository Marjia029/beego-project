package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

// Breed represents the structure of a cat breed from The Cat API
type Breed struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Origin           string `json:"origin"`
	Description      string `json:"description"`
	Temperament      string `json:"temperament"`
	ReferenceImageID string `json:"reference_image_id"`
	ImageURL         string
}

type CatController struct {
	beego.Controller
}

func (c *CatController) Get() {
	// The Cat API URL for fetching breeds
	url := "https://api.thecatapi.com/v1/breeds"

	// Optional: Add your API key to the request headers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error creating request: %s", err))
		return
	}
	// Add your API key (if available)
	// req.Header.Set("x-api-key", "your-api-key")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error making request: %s", err))
		return
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var breeds []Breed
	if err := json.NewDecoder(resp.Body).Decode(&breeds); err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error decoding response: %s", err))
		return
	}

	// Pass the data to the template
	c.Data["Breeds"] = breeds
	c.TplName = "cat_breeds.tpl"
}

func (c *CatController) GetBreedsWithImages() {
	// The Cat API URL for fetching breeds
	breedsURL := "https://api.thecatapi.com/v1/breeds"

	// Fetch the list of breeds
	resp, err := http.Get(breedsURL)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error fetching breeds: %s", err))
		return
	}
	defer resp.Body.Close()

	var breeds []Breed
	if err := json.NewDecoder(resp.Body).Decode(&breeds); err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error decoding breeds response: %s", err))
		return
	}

	// Fetch image URLs using reference_image_id
	for i, breed := range breeds {
		if breed.ReferenceImageID != "" {
			imageURL := fmt.Sprintf("https://api.thecatapi.com/v1/images/%s", breed.ReferenceImageID)
			imageResp, err := http.Get(imageURL)
			if err == nil {
				var imageResponse struct {
					URL string `json:"url"`
				}
				if err := json.NewDecoder(imageResp.Body).Decode(&imageResponse); err == nil {
					breeds[i].ImageURL = imageResponse.URL
				}
				imageResp.Body.Close()
			}
		}
	}

	// Pass the data to the template
	c.Data["Breeds"] = breeds
	c.TplName = "breed_info.tpl"
}

func (c *CatController) GetImages() {
	// The Cat API URL for fetching random images
	url := "https://api.thecatapi.com/v1/images/search?limit=5"

	resp, err := http.Get(url)
	if err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error making request: %s", err))
		return
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var images []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&images); err != nil {
		c.Ctx.WriteString(fmt.Sprintf("Error decoding response: %s", err))
		return
	}

	// Pass the data to the template
	c.Data["Images"] = images
	c.TplName = "cat_images.tpl"
}
