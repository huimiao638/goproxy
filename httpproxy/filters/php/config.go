package php

import (
	"encoding/json"

	"../../../storage"
)

type Config struct {
	FetchServers []struct {
		URL       string
		Password  string
		SSLVerify bool
	}
	Sites     []string
	Transport string
}

func NewConfig(uri, path string) (*Config, error) {
	store, err := storage.OpenURI(uri)
	if err != nil {
		return nil, err
	}

	object, err := store.GetObject(path, -1, -1)
	if err != nil {
		return nil, err
	}

	rc := object.Body()
	defer rc.Close()

	data, err := storage.ReadJson(rc)
	if err != nil {
		return nil, err
	}

	config := new(Config)
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
