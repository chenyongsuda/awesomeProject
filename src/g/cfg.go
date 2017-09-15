package g

import (
	"encoding/json"
	"fmt"
	"github.com/toolkits/file"
	"log"
)

type HttpConf struct {
	Enabled bool  `json:"enabled"`
	Port    string `json:"port"`
}

type Config struct {
	Debug    bool     `json:"debug"`
	Server   string   `json:"server"`
	Interval int32    `json:"interval"`
	Http     HttpConf `json:"http"`
}

var confFile string
var pconf *Config

func Conf() *Config {
	return pconf
}

func ConfFile() string {
	return confFile
}

func ReadConf(fileName string) error {
	if fileName == "" {
		return fmt.Errorf("The File is empty")
	}

	if !file.IsExist(fileName) {
		return fmt.Errorf("The File is not exist")
	}

	confFile = fileName

	content, err := file.ToTrimString(fileName)
	if err != nil {
		return fmt.Errorf("Read Conf File %s err %s", fileName, err)
	}
	var c Config
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		return fmt.Errorf("Decode File %s err %s", fileName, err)
	}
	pconf = &c
	log.Println("Read File %s Success!", fileName)
	return nil
}
