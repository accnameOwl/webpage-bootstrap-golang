package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "./handler"
	"github.com/julienschmidt/httprouter"
)

//PageTemplate is the page
var router *httprouter.Router

func init() {
	router = httprouter.New()

	router.GET("/", handlers.Homepage)
	router.GET("/path" /*path*/, nil /*handlers.someHandlerFunon*/)
	router.POST("/path" /*path*/, nil /*handlers.someHandlerFunction*/)
}

func main() {
	fmt.Println("Starting app.go...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func doSomeStuff() {

}
