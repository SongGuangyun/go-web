package config

type Jwt struct {
	SigningKey string `yaml:"signingKey"`
	ExpireTime int64  `yaml:"expireTime"`
}
