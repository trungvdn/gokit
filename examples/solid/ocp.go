package main

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

type City struct {
}

func GetCities(sourceType string, source string) ([]City, error) {
	var data []byte
	var err error

	if sourceType == "file" {
		data, err = ioutil.ReadFile(source)
		if err != nil {
			return nil, err
		}
	} else if sourceType == "link" {
		resp, err := http.Get(source)
		if err != nil {
			return nil, err
		}

		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
	}

	var cities []City
	err = yaml.Unmarshal(data, &cities)
	if err != nil {
		return nil, err
	}

	return cities, nil
}

type DataReader func(source string) ([]byte, error)

func ReadFromFile(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ReadFromLink(link string) ([]byte, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return data, nil
}

func GetCitiesOCP(reader DataReader, source string) ([]City, error) {
	data, err := reader(source)
	if err != nil {
		return nil, err
	}

	var cities []City
	err = yaml.Unmarshal(data, &cities)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
