package main

import (
	"net/http"
	"os"
	"util/log"

	"github.com/julienschmidt/httprouter"
)

/**
* main() handle below URLs:
* /
* /hello
* /user/get/{id}
* /user/del/{id}
**/
func main() {
	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("static")) // set '/static' as file root path
	router.GET("/", index)                                // mapping '/' to index()
	router.GET("/hello", hello)                           // mapping '/hello' to hello()
	router.GET("/user/get/:id", getUser)                  // mapping '/user/get/:id' to getUser()
	router.GET("/user/del/:id", delUser)                  // mapping '/user/del/:id' to delUser()

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
