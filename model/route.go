package model

type Route struct {
	Method string `yaml:"method"`
	Route string `yaml:"route"`
	Response string `yaml:"response"`
	StatusCode int `yaml:"statusCode"`
}

type Server struct {
	Port int `yaml:"port"`
	Routes []Route `yaml:"routes"`
}
