package main

import (
	"log"
	"net/http"

	"github.com/possawang/go-service-lib-common/routerutils"
)

func helloWorld(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello world"))
}

func globalMdwHelloWorld(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Global MDW from %s\n", r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

func mdwHelloWorld(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request hello world from %s\n", r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

func main() {
	endpoints := make(map[string]routerutils.Endpoint)
	endpoints["/"] = routerutils.Endpoint{
		Execution: helloWorld,
		Method:    "GET",
		Mdw:       mdwHelloWorld,
	}
	routerutils.StartingService(endpoints, globalMdwHelloWorld)
}
