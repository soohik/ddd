// geektutu.com
// main.go
package main

import (
	"backend/cmd/backend/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg,_ := config.Config()
	fmt.Println(cfg)
	
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}