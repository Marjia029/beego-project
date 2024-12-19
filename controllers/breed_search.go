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
		c.fetchBreedDetails(breeds[0].ID)
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

	c.fetchBreedDetails(breedID)
	c.TplName = "breed_search.tpl"
	c.Data["ActiveTab"] = "breeds"
}

func (c *BreedSearchController) fetchBreedDetails(breedID string) {
	apiKey := "live_GWXcPdnWze27MNMJSjinKshtfsnVsi4EdrXfKUNhOmXsLakl5N7MwJCShLvC5Rxo"

	// Fetch breed details
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

	// Fetch images
	imagesURL := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=8", breedID)
	req, _ = http.NewRequest("GET", imagesURL, nil)
	req.Header.Set("x-api-key", apiKey)

	resp, _ = client.Do(req)
	defer resp.Body.Close()

	var images []BreedImage
	body, _ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &images)

	c.Data["SelectedBreed"] = selectedBreed
	c.Data["BreedImages"] = images
	c.Data["Breeds"] = breeds

}
