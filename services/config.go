package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config Service type.
type configService struct {
	config *map[string]string
}

// Config Service constructor function.
func NewConfigService() *configService {
	config, err := setConfig("app/config.json")
	if err != nil {
		fmt.Println("Unable to read 'app/config.json'. Is the filepath correct?")
		panic(err)
	}

	return &configService{config: &config}
}

func setConfig(path string) (map[string]string, error) {

	// Read the config data.
	rawData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal the data into a JSON wrapper.
	var config map[string]string
	err = json.Unmarshal(rawData, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (cs *configService) GetConfig(key string) (string, error) {

	// If the requested key exists, return the value.
	config := *cs.config
	if val, ok := config[key]; ok {
		return val, nil
	}

	// Otherwise, return an error.
	return "", fmt.Errorf("Config Service: Key \"%v\" not found.", key)
}
