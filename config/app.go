package config

type App struct {
	AppName string `yaml:"appName"`
	AppEnv  string `yaml:"appEnv"`
	AppPort int    `yaml:"appPort"`
}
