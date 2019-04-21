package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"pnt/economy/fabricAPI/router"
)

func main() {
	app := gin.Default()
	router.Route(app)
	gin.SetMode(gin.DebugMode)
	addr := fmt.Sprintf(":%d", 4090)
	log.Printf("fabricApi Server start listen addr %s", addr)
	app.Run(addr)

}
