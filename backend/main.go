package main

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		indexFile, _ := filepath.Abs("./frontend/dist/index.html")
		c.File(indexFile)
	})

	staticDir, _ := filepath.Abs("./frontend/dist/static")
	r.Static("/static", staticDir)

	favicon, _ := filepath.Abs("./frontend/src/assets/logo.png")
	r.StaticFile("favicon.ico", favicon)

	r.Run() // listen and serve on 0.0.0.0:8080
}