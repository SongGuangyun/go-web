package config

type Server struct {
	Mysql Mysql
	App   App
	Log   Log
	Jwt   Jwt
}
