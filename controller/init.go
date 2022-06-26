package controller

import (
	"assignment-3/updater"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WeatherStatus(c *gin.Context) {
	fileData, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal("Error when reading json file", err.Error())
	}

	var statusData updater.Status

	err = json.Unmarshal(fileData, &statusData)
	if err != nil {
		log.Fatal("Error when reading json file", err.Error())
	}

	waterData := statusData.Water
	windData := statusData.Wind

	var (
		waterStatus string
		windStatus  string
		result      gin.H
	)

	switch {
	case waterData <= 5:
		waterStatus = "Aman"
	case waterData >= 6 && waterData <= 8:
		waterStatus = "Siaga"
	case waterData > 8:
		waterStatus = "Bahaya"
	default:
		waterStatus = "Water Value not defined"
	}

	switch {
	case windData <= 6:
		windStatus = "Aman"
	case windData >= 7 && waterData <= 15:
		windStatus = "Siaga"
	case windData > 15:
		windStatus = "Bahaya"
	default:
		windStatus = "Water Value not defined"
	}

	result = gin.H{
		"waterStatus": waterStatus,
		"windStatus":  windStatus,
		"waterVal":    waterData,
		"windVal":     windData,
	}

  c.HTML(http.StatusOK, "index.html", result)
}
