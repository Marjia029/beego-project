package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BreedImage struct { // Renamed to avoid conflicts
	URL string `json:"url"`
}

type CatBreed struct { // Renamed Breed to CatBreed
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BreedSearchController struct {
	web.Controller
}

// Get handles the search page rendering
func (c *BreedSearchController) Get() {
	apiKey := "live_GWXcPdnWze27MNMJSjinKshtfsnVsi4EdrXfKUNhOmXsLakl5N7MwJCShLvC5Rxo"
	url := "https://api.thecatapi.com/v1/breeds"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Data["Error"] = "Failed to create API request."
		c.TplName = "breed_search.tpl"
		return
	}
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Data["Error"] = "Failed to fetch breed list."
		c.TplName = "breed_search.tpl"
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Data["Error"] = "Failed to read API response."
		c.TplName = "breed_search.tpl"
		return
	}

	var breeds []CatBreed // Directly using CatBreed struct here
	if err := json.Unmarshal(body, &breeds); err != nil {
		c.Data["Error"] = "Failed to parse breed list."
		c.TplName = "breed_search.tpl"
		return
	}

	c.Data["Breeds"] = breeds
	c.TplName = "breed_search.tpl"
}

// Post handles the breed search and displays the result
func (c *BreedSearchController) Post() {
	breedID := c.GetString("breed_id")
	if breedID == "" {
		c.Data["Error"] = "Please select a breed."
		c.TplName = "breed_search.tpl"
		return
	}

	apiKey := "live_GWXcPdnWze27MNMJSjinKshtfsnVsi4EdrXfKUNhOmXsLakl5N7MwJCShLvC5Rxo"

	// Fetch breed details
	breedDetailsURL := "https://api.thecatapi.com/v1/breeds"
	req, err := http.NewRequest("GET", breedDetailsURL, nil)
	if err != nil {
		c.Data["Error"] = "Failed to create breed details request."
		c.TplName = "breed_search.tpl"
		return
	}
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.Data["Error"] = "Failed to fetch breed details."
		c.TplName = "breed_search.tpl"
		return
	}
	defer resp.Body.Close()

	var breeds []CatBreed
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &breeds)

	var selectedBreed CatBreed
	for _, breed := range breeds {
		if breed.ID == breedID {
			selectedBreed = breed
			break
		}
	}

	// Fetch images for the breed
	imagesURL := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=10", breedID)
	req, _ = http.NewRequest("GET", imagesURL, nil)
	req.Header.Set("x-api-key", apiKey)

	resp, err = client.Do(req)
	if err != nil {
		c.Data["Error"] = "Failed to fetch breed images."
		c.TplName = "breed_search.tpl"
		return
	}
	defer resp.Body.Close()

	var images []BreedImage // Renamed to BreedImage to avoid conflict
	body, _ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &images)

	c.Data["SelectedBreed"] = selectedBreed
	c.Data["BreedImages"] = images
	c.TplName = "breed_search.tpl"
}
