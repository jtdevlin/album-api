package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jtdevlin/album-api/internal/controller"
)

func main() {
	router := gin.Default()
	router.GET("albums", controller.GetAlbums)
	router.GET("albums/:id", controller.GetAlbumById)
	router.POST("albums", controller.CreateAlbum)
	router.Run()
}
