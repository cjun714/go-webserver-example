package main

import (
	"dao"
	"fmt"
	"net/http"
	"strconv"
	"unsafe"
	"util/json"
	"util/log"

	"github.com/julienschmidt/httprouter"
)

// URL: /
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// URL: /hello
func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	writeObj(w, "Hello world!")
}

// URL: /user/get/:id
func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, e := strconv.Atoi(ps.ByName("id")) // string --> int
	if e != nil {
		log.E(e)
		return // TODO should return error msg, errorcode = 400
	}

	usr := dao.GetUserByID(id)
	writeObj(w, usr)
}

// URL: /user/del/:id
func delUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, e := strconv.Atoi(ps.ByName("id")) // string --> int
	if e != nil {
		log.E(e)
		return // TODO should return error msg, errorcode = 400
	}

	result := dao.DelUserByID(id)
	writeObj(w, result)
}

func writeObj(w http.ResponseWriter, obj interface{}) {
	bytes := json.ToJSON(obj)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, token")
	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, *(*string)(unsafe.Pointer(&bytes)))
}
