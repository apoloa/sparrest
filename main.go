package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/apoloa/sparrest/model"
	"net/http"
)

func main() {
	path := flag.String("path", "./server.yaml", "path of the file to load")

	flag.Parse()
	data, err := ioutil.ReadFile(*path)

	if err != nil {
		fmt.Println("Invalid path")
		panic(err)
	}

	server := model.Server{}

	err = yaml.Unmarshal(data, &server)

	if err != nil {
		fmt.Println("Invalid yaml")
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		defer fmt.Printf("[%v] - %v - %v\n", r.Method, r.RemoteAddr, r.URL.Path[1:])
		w.Header().Add("Content-Type", "application/json")
		for _, route := range server.Routes {
			if route.Method == r.Method && r.URL.Path[1:] == route.Route {
				w.WriteHeader(route.StatusCode)
				w.Write([]byte(route.Response))
				return
			}
		}
		w.WriteHeader(404)

	})
	err = http.ListenAndServe(fmt.Sprintf(":%v", server.Port), nil)
	if err != nil {
		fmt.Println("Error creating the server")
		fmt.Println(err)
	}
	fmt.Println("Server running")
}
