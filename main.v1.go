package main

import (
	"log"
	"net/http"

	"github.com/andrepinto/keit/v1"
	"github.com/emicklei/go-restful"
)

func main() {
	u := v1.APIResource{}
	restful.DefaultContainer.Add(u.WebService())

	log.Printf("start listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
