package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(
		client.WithHost("http://localhost:2376"),
		client.WithVersion("1.35"),
	)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("=== container ===\n")
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("=== node ===\n")
	for _, node := range nodes {
		fmt.Printf("%s %v %v %v\n", node.ID[:10], node.Status.State, node.ManagerStatus.Leader, node.ManagerStatus.Addr)
	}

	swarm, err := cli.SwarmInspect(context.Background())
	fmt.Printf("=== join token ===\n")
	fmt.Printf("Worker = %s\n", swarm.JoinTokens.Worker)
	fmt.Printf("Manager = %s\n", swarm.JoinTokens.Manager)

}
