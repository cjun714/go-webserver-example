package main

import (
	"net/http"
	"os"
	"util/log"

	"github.com/julienschmidt/httprouter"
)

/**
* main() handle below URL:
* /
* /hello
* /user/get/{id}
* /user/del/{id}
**/
func main() {
	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("static")) //'/static' as file root path
	router.GET("/", index)
	router.GET("/hello", hello)
	router.GET("/user/get/:id", getUser)
	router.GET("/user/del/:id", delUser)

	port := getPort()
	log.H("Server is started on: " + port)
	log.E(http.ListenAndServe(":"+port, router)) // start server and listen on 8080
}

func getPort() string {
	if len(os.Args) == 2 {
		return os.Args[1]
	}

	return "80"
}
