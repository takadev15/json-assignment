package updater

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func AutoReload() {
	for {
		min := 1
		max := 100
		wind := rand.Intn(max-min) + min
		water := rand.Intn(max-min) + min

		data := Status{}
		data.Wind = wind
		data.Water = water

		jsonData, err := json.Marshal(data)

		if err != nil {
			log.Fatal("error occured while marshalling status data:", err.Error())
		}
		err = ioutil.WriteFile("data.json", jsonData, 0644)

		if err != nil {
			log.Fatal("error occured while writing data to data.json file", err.Error())
		}
		time.Sleep(15 * time.Second)
	}
}

