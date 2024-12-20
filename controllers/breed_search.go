package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type BreedImage struct {
	URL string `json:"url"`
}

type CatBreed struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Origin       string `json:"origin"`
	WikipediaURL string `json:"wikipedia_url"`
}

type BreedSearchController struct {
	web.Controller
}

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

	var breeds []CatBreed
	if err := json.Unmarshal(body, &breeds); err != nil {
		c.Data["Error"] = "Failed to parse breed list."
		c.TplName = "breed_search.tpl"
		return
	}

	// Set default selected breed (first breed)
	if len(breeds) > 0 {
		c.fetchBreedDetailsConcurrently(breeds[0].ID)
	}

	c.Data["ActiveTab"] = "breeds"
	c.Data["Breeds"] = breeds
	c.TplName = "breed_search.tpl"
}

func (c *BreedSearchController) Post() {
	breedID := c.GetString("breed_id")
	if breedID == "" {
		c.Redirect("/breed-search", 302)
		return
	}

	c.fetchBreedDetailsConcurrently(breedID)
	c.TplName = "breed_search.tpl"
	c.Data["ActiveTab"] = "breeds"
}

// Fetch breed details and images concurrently using channels
func (c *BreedSearchController) fetchBreedDetailsConcurrently(breedID string) {
	apiKey := "live_GWXcPdnWze27MNMJSjinKshtfsnVsi4EdrXfKUNhOmXsLakl5N7MwJCShLvC5Rxo"

	// Channel to receive breed details
	breedDetailsChan := make(chan CatBreed)
	// Channel to receive breed images
	imagesChan := make(chan []BreedImage)
	// Channel to handle errors
	errChan := make(chan error)

	// Fetch breed details concurrently
	go func() {
		breedDetailsURL := "https://api.thecatapi.com/v1/breeds"
		req, _ := http.NewRequest("GET", breedDetailsURL, nil)
		req.Header.Set("x-api-key", apiKey)

		client := &http.Client{}
		resp, _ := client.Do(req)
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

		breedDetailsChan <- selectedBreed
	}()

	// Fetch breed images concurrently
	go func() {
		imagesURL := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=8", breedID)
		req, _ := http.NewRequest("GET", imagesURL, nil)
		req.Header.Set("x-api-key", apiKey)

		client := &http.Client{}
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		var images []BreedImage
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &images)

		imagesChan <- images
	}()

	// Wait for all goroutines to complete and handle errors
	select {
	case selectedBreed := <-breedDetailsChan:
		c.Data["SelectedBreed"] = selectedBreed
	case err := <-errChan:
		c.Data["Error"] = err.Error()
		c.TplName = "breed_search.tpl"
		return
	}

	select {
	case breedImages := <-imagesChan:
		c.Data["BreedImages"] = breedImages
	case err := <-errChan:
		c.Data["Error"] = err.Error()
		c.TplName = "breed_search.tpl"
		return
	}

	// Fetch breeds again for the sidebar (if needed)
	req, _ := http.NewRequest("GET", "https://api.thecatapi.com/v1/breeds", nil)
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	var breeds []CatBreed
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &breeds)

	c.Data["Breeds"] = breeds
}
