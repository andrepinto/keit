package main

import (
	"log"
	"net/http"

	api "github.com/andrepinto/keit/v2"
	"github.com/emicklei/go-restful"
)

func main() {
	u := api.APIResource{}
	restful.DefaultContainer.Add(u.WebService())

	log.Printf("start listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
