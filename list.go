package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/docker/swarmkit/api"
	"github.com/docker/swarmkit/cmd/swarmd/defaults"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"time"
)

func main() {
	opts := []grpc.DialOption{}
	addr := defaults.ControlAPISocket
	insecureCreds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	opts = append(opts, grpc.WithTransportCredentials(insecureCreds))
	opts = append(opts, grpc.WithDialer(
		func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", addr, timeout)
		}))
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	c := api.NewControlClient(conn)
	c.ListNodes(ctx, &api.ListNodesRequest{})
}
