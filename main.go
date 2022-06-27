package main

import (
	"assignment-3/controller"
	"assignment-3/updater"

	"github.com/gin-gonic/gin"
  // "path/filepath"
)

func main() {
  go updater.AutoReload()
  r := gin.Default()
  r.GET("/", controller.WeatherStatus)
  r.LoadHTMLGlob("./view/*")
  // filePrefix, _ := filepath.Abs("./view/index.html")
  // r.LoadHTMLFiles(filePrefix)
  r.Run(":3030")
}
