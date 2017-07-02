package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/docker/swarmkit/cmd/swarmd/defaults"
	"github.com/docker/swarmkit/node"
)

func main() {
	n, err := node.New(&node.Config{
		Hostname:         "test",
		ForceNewCluster:  false,
		ListenControlAPI: defaults.ControlAPISocket,
		//ListenRemoteAPI:    addr,
		//AdvertiseRemoteAPI: advertiseAddr,
		//JoinAddr:           managerAddr,
		StateDir: "./state",
		//JoinToken:          joinToken,
		//ExternalCAs:        externalCAOpt.Value(),
		//Executor:           executor,
		//HeartbeatTick:      hb,
		//ElectionTick:       election,
		//AutoLockManagers:   autolockManagers,
		//UnlockKey:          unlockKey,
	})
	fmt.Println("Node Created")
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	if err := n.Start(ctx); err != nil {
		fmt.Println(n)
	}
	fmt.Println("Node Started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		select {
		case <-n.Ready():
		case <-ctx.Done():
		}
		if ctx.Err() == nil {
			fmt.Println("node is ready")
		} else {
			fmt.Println(ctx.Err())
		}
	}()

	<-c
	n.Stop(ctx)
}
