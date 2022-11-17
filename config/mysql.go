package config

type Mysql struct {
	DBHost     string `yaml:"dbHost"`
	DBPort     string `yaml:"dbPort"`
	DBDatabase string `yaml:"dbDatabase"`
	DBUsername string `yaml:"dbUsername"`
	DBPassword string `yaml:"dbPassword"`
}
