package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	log "github.com/sirupsen/logrus"
)

//LoadConfig reads in a configuration file in JSON format and unmarshalls it into a pre-defined data type
func  LoadConfig(filename string, config interface{}) (err error) {
	if len(filename) == 0 {
		log.Debug("filename has not been set")
	}

	file, err := os.Open(filename)
	if(err!=nil) {
		log.Fatal("Couldn't load file ", filename)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if(err!=nil) {
		log.Fatal("Couldn't read data")
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Unable to unmarshall data into data type")
	}

	return
}