package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = map[string]album{
	"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumById(context *gin.Context) {
	id := context.Param("id")

	album, ok := albums[id]
	if ok {
		context.IndentedJSON(http.StatusOK, album)
	} else {
		context.IndentedJSON(http.StatusNotFound, fmt.Sprintf("Album not found for ID: %s", id))
	}
}

func CreateAlbum(context *gin.Context) {
	var newAlbum album

	if err := context.BindJSON(&newAlbum); err != nil {
		context.IndentedJSON(http.StatusBadRequest, "Invalid album payload provided")
		return
	}
	_, ok := albums[newAlbum.ID]
	if ok {
		context.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("Album already exists for ID: %s", newAlbum.ID))
		return
	}

	albums[newAlbum.ID] = newAlbum
	context.IndentedJSON(http.StatusCreated, newAlbum)

}
