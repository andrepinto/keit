package v2

import (
	"fmt"
	"time"

	restful "github.com/emicklei/go-restful"
)

//APIResource ...
type APIResource struct {
}

// WebService creates a new service that can handle REST requests for User resources.
func (a APIResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/api").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.GET("/").To(a.Info).
		// docs
		Doc("get api info").
		Writes(ApiData{}).
		Returns(200, "OK", ApiData{}))

	ws.Route(ws.GET("/5xx").To(a.Err5xx).
		// docs
		Doc("error 500").
		Writes(ApiData{}).
		Returns(500, "OK", ApiData{}))

	ws.Route(ws.GET("/delay/{delay}").To(a.Delay).
		// docs
		Doc("get api info").
		Writes(ApiData{}).
		Returns(500, "OK", ApiData{}))

	return ws
}

//Info ...
func (a APIResource) Info(request *restful.Request, response *restful.Response) {

	api := ApiData{
		Name:        "keit api demo",
		Version:     "v2",
		Meta:        request.HeaderParameter("meta"),
		Description: "this is a testing api",
	}

	response.WriteEntity(api)
}

//Err5xx ...
func (a APIResource) Err5xx(request *restful.Request, response *restful.Response) {

	response.WriteErrorString(500, "error 5xx")
}

//Delay ...
func (a APIResource) Delay(request *restful.Request, response *restful.Response) {

	id := request.PathParameter("delay")

	d, _ := time.ParseDuration(id)

	time.Sleep(d)

	response.Write([]byte(fmt.Sprintf("delay:%s", id)))
}
