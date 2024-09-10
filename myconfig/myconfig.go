package myconfig

import (
	"bytes"
	"encoding/json"
	"github.com/houyanzu/work-box/config"
	"io/ioutil"
	"os"
)

type MyConfig struct {
	config.Config
	Common commonConfig `json:"common"`
}

type commonConfig struct {
}

var myConfig *MyConfig

func ParseConfig(configName string) (err error) {
	dat, err := ioutil.ReadFile(configName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dat, &myConfig)
	if err != nil {
		return err
	}

	config.ParseConfig(&myConfig.Config)

	return nil
}

func GetConfig() *MyConfig {
	return myConfig
}

func CreateConfigFile() {
	var conf MyConfig
	js, err := json.Marshal(conf)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	_ = json.Indent(&out, js, "", "  ")
	f, _ := os.OpenFile("config.json", os.O_WRONLY|os.O_CREATE, 0777)
	out.WriteTo(f)
}
