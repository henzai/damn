package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var cli *client.Client

// init is
func init() {
	docker, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}
	cli = docker
}

// List is
func List() {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("id: %s, image: %s, name: %s\n", container.ID[:10], container.Image, container.Names[0][1:])
	}
}

func getContainers() []types.Container {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	return containers
}

func HasContainer(value string) (string, bool) {
	if id, hasContainer := HasContainerByID(value); hasContainer == true {
		return id, hasContainer
	}
	if id, hasContainer := HasContainerByName(value); hasContainer == true {
		return id, hasContainer
	}
	return HasContainerByName(value)

}

// HasContainerByID is
func HasContainerByID(id string) (string, bool) {
	containers := getContainers()
	for _, container := range containers {
		if id == container.ID[:12] {
			return id, true
		}
	}
	return "", false
}

// HasContainerByName is
func HasContainerByName(name string) (string, bool) {
	containers := getContainers()
	for _, container := range containers {
		for _, containerName := range container.Names {
			if name == containerName[1:] {
				return container.ID[:10], true
			}
		}
	}
	return "", false
}
