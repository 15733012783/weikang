package client

import (
	"fmt"
	"github.com/15733012783/mysql/consul"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
)

func Client(serverName, Address string) (*grpc.ClientConn, error) {
	conn, err := consul.GetClient(serverName, Address)
	if err != nil {
		return nil, err
	}
	fmt.Println(conn)
	return grpc.Dial("consul://10.2.171.70:8500/"+serverName+"?wait=14s", grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`))
}
