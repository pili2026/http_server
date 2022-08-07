package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var cachedConfig map[string]interface{}

func GetConfig() (m map[string]interface{}, err error) {
	if cachedConfig != nil {
		return cachedConfig, nil
	}

	var bytes []byte

	if bytes, err = ioutil.ReadFile(GetEnv().ConfigPath); err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(bytes, &m); err != nil {
		return nil, err
	}

	fmt.Printf("Loaded config \n\"%v\"\n", string(bytes))
	cachedConfig = m
	return m, nil
}
