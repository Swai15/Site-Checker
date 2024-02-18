package main

import (
	"encoding/json"
	"os"
)

func readTrackedWebsitesFromFile() error {
	file, err := os.Open(dataFileName)
	if err != nil {
		if os.IsNotExist(err) {
			trackedWebsites = []string{}
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&trackedWebsites)
	if err != nil {
		return err
	}
	return nil

}

func writeTrackedWebsitesToFile() error {
	file, err := os.Create(dataFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(trackedWebsites)
	if err != nil {
		return err
	}

	return nil
}

