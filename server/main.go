package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/consul/api"
)

var (
	service_id   = "petservice_2"
	service_name = "petservice"
)

func main() {

	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	registeration := &api.AgentServiceRegistration{
		Kind:              "a",
		ID:                service_id,
		Name:              service_name,
		Tags:              []string{"pet", "dog"},
		Port:              9002,
		Address:           "localhost",
		SocketPath:        "",
		TaggedAddresses:   map[string]api.ServiceAddress{},
		EnableTagOverride: false,
		Meta:              map[string]string{},
		Weights:           &api.AgentWeights{Passing: 1, Warning: 1},
		Check:             &api.AgentServiceCheck{},
		Checks:            []*api.AgentServiceCheck{},
		Proxy:             &api.AgentServiceConnectProxyConfig{},
		Connect:           &api.AgentServiceConnect{},
		Namespace:         "",
		Partition:         "",
	}

	registeration.Check = &api.AgentServiceCheck{
		CheckID:                        "petservice_check_1",
		Name:                           "petservice_check",
		Interval:                       "5s",
		Timeout:                        "3s",
		HTTP:                           "http://localhost:8200/check",
		DeregisterCriticalServiceAfter: "30s",
	}

	err = client.Agent().ServiceRegister(registeration)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte(r.RemoteAddr + "hello"))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", 8200), nil)
}

func Deregister(serviceId string) error {

	// Get a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Agent().ServiceDeregister("petservice_1")
	if err != nil {
		return err
	}
	return nil
}
