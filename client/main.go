package main

import (
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {

	conn, err := Dial("petservice", "", 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	_ = conn
}

// Dial dials a service through gRPC and returns a new connection.
func Dial(serviceName, tag string, timeout time.Duration) (grpc.ClientConnInterface, error) {

	// cfg := consul.Config()
	//target := fmt.Sprintf("consul://%s:%s@%s/%s?tag=%s", "", "", "localhost", serviceName, tag)

	target := fmt.Sprintf("consul://127.0.0.1:8500/petservice")

	opts := []grpc.DialOption{
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	return grpc.Dial(target, opts...)
}
