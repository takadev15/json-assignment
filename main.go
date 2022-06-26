package main

import (
	"assignment-3/controller"
	"assignment-3/updater"

	"github.com/gin-gonic/gin"
)

func main() {
  go updater.AutoReload()
  r := gin.Default()
  r.GET("/", controller.WeatherStatus)
  // r.LoadHTMLGlob("view/*.html")
  r.LoadHTMLFiles("./view/index.html")
  r.Run(":3030")
}
