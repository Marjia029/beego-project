package routers

import (
	"cat-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	// Serve static files
	beego.SetStaticPath("/static", "static")

	beego.Router("/", &controllers.MainController{})

	// New route for the API page
	beego.Router("/api", &controllers.ApiController{})

	// Route for cat breeds
	beego.Router("/cats", &controllers.CatController{})

	// Route for breeds with images
	beego.Router("/breeds", &controllers.CatController{}, "get:GetBreedsWithImages")

	// Route for cat images
	beego.Router("/cat-images", &controllers.CatController{}, "get:GetImages")

	beego.Router("/breed-images", &controllers.BreedImagesController{})

	beego.Router("/breed-search", &controllers.BreedSearchController{})

	beego.Router("/voting", &controllers.VotingController{})

}
