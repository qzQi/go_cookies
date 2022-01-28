package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 只有可导出的成员才可以转化为json字段
type album struct {
	ID     string  `json:"id`
	Title  string  `json:"title`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "blue train", Artist: "qzy", Price: 99},
	{ID: "2", Title: "blue train", Artist: "qzy", Price: 99},
	{ID: "3", Title: "blue train", Artist: "qzy", Price: 99},
	{ID: "4", Title: "blue train", Artist: "qzy", Price: 99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// 动态路由
func getAlumByID(c *gin.Context) {
	// queryString里面的id=12，应该就是可以解析为12
	// 讨论错了，这个是动态路由

	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	// 貌似是不管这里是什么都把消息给了这个动态路由
	router.GET("/albums/:id", getAlumByID)

	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}
