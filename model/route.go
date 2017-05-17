package model

// Route defines the routes inside the server
type Route struct {
	Method     string `yaml:"method"`
	Route      string `yaml:"route"`
	Response   string `yaml:"response"`
	StatusCode int    `yaml:"statusCode"`
}

// Server model that defines a server with the port to listen and the routes
type Server struct {
	Port   int     `yaml:"port"`
	Routes []Route `yaml:"routes"`
}
