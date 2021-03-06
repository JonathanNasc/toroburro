package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

const fileName string = "settings.json"

var settings Settings

type Source struct {
	Name string `json:"name"`
	ApiKey string `json:"api-key"`
}

type Asset struct {
	Source string `json:"source"`
	Codes []string `json:"codes"`
}

type Settings struct {
	Sources []Source `json:"sources"`
	Assets []Asset `json:"assets"`
}

func GetSettings() Settings {
	var onlyOnce sync.Once
	onlyOnce.Do(func() {
		settings = LoadSettings()
	})

	return settings
}

func LoadSettings() Settings {
	file, _ := ioutil.ReadFile(fileName)
	err := json.Unmarshal(file, &settings)
	if err != nil {
		panic("Something went wrong during file read: " + err.Error())
	}

	return settings
}
