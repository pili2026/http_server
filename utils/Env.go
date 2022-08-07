package utils

import (
	"os"
)

type envType struct {
	ConfigPath  string
	ServicePort string
}

var defaultEnvConfig = envType{
	ConfigPath:  "res/config.yml",
	ServicePort: "8000",
}

var env *envType = nil

func GetEnv() envType {

	if env == nil {
		env = &envType{}
	}

	env.ServicePort = os.Getenv("PORT")
	if len(env.ServicePort) <= 0 {
		env.ServicePort = defaultEnvConfig.ServicePort
	}

	env.ConfigPath = os.Getenv("CONFIG_PATH")
	if len(env.ConfigPath) <= 0 {
		env.ConfigPath = defaultEnvConfig.ConfigPath
	}
	return *env
}
