package main

import (
	"encoding/json"
	"path"
	"strings"

	"../storage"
)

type Config struct {
	LogToStderr bool
	Addr        string
	Http        struct {
		Mode            string
		KeepAlivePeriod int
		ReadTimeout     int
		WriteTimeout    int
		Certificate     string
		PrivateKey      string
	}
	GroupCache struct {
		Addr  string
		Peers []string
	}
	Filters struct {
		Request   []string
		RoundTrip []string
		Response  []string
	}
}

func NewConfig(uri, filename string) (*Config, error) {
	store, err := storage.OpenURI(uri)
	if err != nil {
		return nil, err
	}

	config := new(Config)

	fileext := path.Ext(filename)
	filename1 := strings.TrimSuffix(filename, fileext) + ".user" + fileext

	for i, name := range []string{filename, filename1} {
		object, err := store.GetObject(name, -1, -1)
		if err != nil {
			if i == 0 {
				return nil, err
			} else {
				continue
			}
		}

		rc := object.Body()
		defer rc.Close()

		data, err := storage.ReadJson(rc)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, config)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}
