package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// To Do: Clean up and comment...

var config = map[string]string{
	"FlickrAppId": "testvalue01",
}
var path string

func TestMain(m *testing.M) {
	jsonConf, err := json.Marshal(config)
	if err != nil {
		log.Fatalf("Error initializing TestMain for Config Service, %s", err)
	}
	path = setup(jsonConf)

	exitCode := m.Run()
	teardown(path)
	os.Exit(exitCode)
}

func setup(jsonConf []byte) string {
	config := []byte(jsonConf)
	file, err := ioutil.TempFile(os.TempDir(), "config.json")
	err = ioutil.WriteFile(file.Name(), config, 0644)
	if err != nil {
		log.Fatalf("Error setting up TestMain for Config Service: %s", err)
	}

	return file.Name()
}

func teardown(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatalf("Error tearing down TestMain for Config Service: %s", err)
	}
}

func TestSetConfig(t *testing.T) {
	_, err := setConfig(path)
	if err != nil {
		t.Errorf("Error encountered when attempting to setConfig: %s", err)
	}
}

func TestGetConfig(t *testing.T) {
	cs := &configService{config: &config}

	got, err := cs.GetConfig("FlickrAppId")
	if err != nil {
		t.Errorf("Error encountered when attempting to getConfig: %s", err)
	}

	if config["FlickrAppId"] != got {
		t.Errorf("GetConfig did not fetch the correct value for: %s, got %s instead", config["FlickrAppId"], got)
	}
}
