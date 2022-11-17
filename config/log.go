package config

type Log struct {
	StorageLocation     string `yaml:"storageLocation"`
	RotationTime        int    `yaml:"rotationTime"`
	RemainRotationCount int    `yaml:"remainRotationCount"`
}
