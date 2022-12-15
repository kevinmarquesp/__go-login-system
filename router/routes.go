package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer(port string) {
	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.HandleFunc("/", Index)
	muxRouter.HandleFunc("/signin", SignIn)
	muxRouter.HandleFunc("/newuser", NewUser)

	err := http.ListenAndServe(port, muxRouter)
	if err != nil {
		log.Fatal(err)
	}
}
