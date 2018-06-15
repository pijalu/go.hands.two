package main

import (
	"log"
	"os"
	"strconv"

	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/pijalu/go.hands.two/frinsultdata/handler"
	"github.com/pijalu/go.hands.two/frinsultproto"
)

func main() {
	var service micro.Service
	// Create service
	useK8S, _ := strconv.ParseBool(os.Getenv("USE_K8S"))
	if useK8S {
		service = k8s.NewService(
			micro.Name("frinsult.srv.micro"),
			micro.Version("latest"),
		)
	} else {
		service = micro.NewService(
			micro.Name("frinsult.srv.micro"),
			micro.Version("latest"),
		)
	}
	// Register
	frinsultproto.RegisterFrinsultServiceHandler(
		service.Server(),
		new(handler.FrinsultHandler))

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
