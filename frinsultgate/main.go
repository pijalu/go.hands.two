package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/pijalu/go.hands.two/env"
	"github.com/pijalu/go.hands.two/frinsultproto"
)

var friService frinsultproto.FrinsultService

func main() {
	var service micro.Service
	// Create service
	useK8S, _ := strconv.ParseBool(os.Getenv("USE_K8S"))
	if useK8S {
		service = k8s.NewService(
			micro.Name("frinsult.client.micro"),
		)
	} else {
		service = micro.NewService(
			micro.Name("frinsult.client.micro"),
		)
	}

	service.Init()
	friService = frinsultproto.NewFrinsultService("frinsult.srv.micro", service.Client())

	// start gateway
	port := env.GetEnvWithDefault("PORT", "8080")
	r := mux.NewRouter()

	// We will use ingress on K8S for routing
	prefix := ""
	if !useK8S {
		prefix = "/api"
	}

	r.Methods("GET").Path(prefix + "/insults/{id:[0-9]+}").HandlerFunc(getInsultByID)
	r.Methods("DELETE").Path(prefix + "/insults/{id:[0-9]+}").HandlerFunc(deleteInsultByID)
	r.Methods("PATCH").Path(prefix + "/insults/{id:[0-9]+}").HandlerFunc(updateInsultByID)
	r.Methods("PUT").Path(prefix + "/insults").HandlerFunc(putInsult)

	r.Methods("POST").Path(prefix + "/insults/upvote/{id:[0-9]+}").HandlerFunc(upvoteInsultByID)
	r.Methods("POST").Path(prefix + "/insults/downvote/{id:[0-9]+}").HandlerFunc(downvoteInsultByID)

	r.Methods("GET").Path(prefix + "/insults").HandlerFunc(getInsults)

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
