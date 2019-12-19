package controllers

import (
	"github.com/iris/repositories"
	"github.com/iris/services"
	"github.com/kataras/iris/mvc"
)

type MovieController struct {

}

func (c *MovieController) Get() mvc.View {
	movieRepo := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepo)
	res := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: res,
	}
}
