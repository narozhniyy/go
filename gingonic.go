package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("./request.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	router.Use(gin.Logger())

	v1 := router.Group("/v1")
	{
		v1.GET("/books", handleGetBooks)
/*		v1.GET("/book", handleGetBook)
		v1.POST("/addBook", posting)
		v1.PUT("/updateBook", putting)
		v1.DELETE("/deleteBook", deleting)
		v1.PATCH("/patchBook", patching)
		v1.HEAD("/head", head)
		v1.OPTIONS("/options", options)*/
	}

	router.Run(":8081") // listen and serve on 0.0.0.0:8081
}

func handleGetBooks(c *gin.Context) {
	c.String(http.StatusOK, "{}")
}