package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type FavoritesController struct {
	web.Controller
}

// Get method to display favorite cat images
func (c *FavoritesController) Get() {
	// Set the active tab to "favs" (this can be used in the navbar for styling)
	c.Data["ActiveTab"] = "favs"

	// Pass the favorites list (from the global variable) to the template
	c.Data["Favorites"] = favorites

	// Render the favorites.tpl template
	c.TplName = "favorites.tpl"
}
