package server

import (
	"encoding/json"
	"net/http"

	"github.com/libopenstorage/openstorage/cluster"
	"github.com/libopenstorage/openstorage/config"
)

type clusterApi struct {
	restBase
}

func newClusterAPI(name string) restServer {
	return &clusterApi{restBase{version: config.Version, name: name}}
}

func (c *clusterApi) Routes() []*Route {
	return []*Route{
		&Route{verb: "GET", path: clusterPath("/enumerate"), fn: c.enumerate},
		&Route{verb: "GET", path: clusterPath("/inspect/{id}"), fn: c.inspect},
		&Route{verb: "DELETE", path: clusterPath(""), fn: c.delete},
		&Route{verb: "DELETE", path: clusterPath("/{id}"), fn: c.delete},
		&Route{verb: "PUT", path: snapPath("shutdown"), fn: c.shutdown},
		&Route{verb: "PUT", path: snapPath("shutdown/{id}"), fn: c.shutdown},
	}
}

func (c *clusterApi) String() string {
	return c.name
}

func (c *clusterApi) enumerate(w http.ResponseWriter, r *http.Request) {
	method := "enumerate"
	inst, err := cluster.Inst()
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}
	cluster, err := inst.Enumerate()
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cluster)
}

func (c *clusterApi) inspect(w http.ResponseWriter, r *http.Request) {
	method := "inspect"
	c.sendNotImplemented(w, method)
}

func (c *clusterApi) delete(w http.ResponseWriter, r *http.Request) {
	method := "delete"
	c.sendNotImplemented(w, method)
}

func (c *clusterApi) shutdown(w http.ResponseWriter, r *http.Request) {
	method := "shutdown"
	c.sendNotImplemented(w, method)
}

func (c *clusterApi) sendNotImplemented(w http.ResponseWriter, method string) {
	c.sendError(c.name, method, w, "Not implemented.", http.StatusNotImplemented)
}

func clusterVersion(route string) string {
	return "/" + config.Version + "/" + route
}

func clusterPath(route string) string {
	return clusterVersion("cluster" + route)
}
