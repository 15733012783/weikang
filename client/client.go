package client

import (
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
)

func Client(serverName, Address string) (*grpc.ClientConn, error) {
	return grpc.Dial("consul://10.2.171.70:8500/"+serverName+"?wait=14s", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`))
}
