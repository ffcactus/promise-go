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
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	nodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}
	for _, node := range nodes {
		fmt.Printf("%s %v %v\n", node.ID[:10], node.Status.State, node.ManagerStatus.Leader)
	}

}
