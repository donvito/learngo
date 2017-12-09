package main

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	listImages(cli)
	listCointainers(cli)
	listNetworks(cli)
	listSwarmNodes(cli)

	fmt.Printf("\n")

}

func listImages(cli *client.Client) {

	//List all images available locally
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("LIST IMAGES\n-----------------------")
	fmt.Println("Image ID | Repo Tags | Size")
	for _, image := range images {
		fmt.Printf("%s | %s | %d\n", image.ID, image.RepoTags, image.Size)
	}

}

func listCointainers(cli *client.Client) {
	//Retrieve a list of containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Print("\n\n\n")
	fmt.Println("LIST CONTAINERS\n-----------------------")
	fmt.Println("Container Names | Image | Mounts")
	//Iterate through all containers and display each container's properties
	for _, container := range containers {
		fmt.Printf("%s | %s | %s\n", container.Names, container.Image, container.Mounts)
	}

}

func listNetworks(cli *client.Client) {
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic(err)
	}

	//List all networks
	fmt.Print("\n\n\n")
	fmt.Println("LIST NETWORKS\n-----------------------")
	fmt.Println("Network Name | ID")
	for _, network := range networks {
		fmt.Printf("%s | %s\n", network.Name, network.ID)
	}

}

func listSwarmNodes(cli *client.Client) {
	swarmNodes, err := cli.NodeList(context.Background(), types.NodeListOptions{})
	if err != nil {
		panic(err)
	}

	//List all nodes - works only in Swarm Mode
	fmt.Print("\n\n\n")
	fmt.Println("LIST SWARM NODES\n-----------------------")
	fmt.Println("Name | Role | Leader | Status")
	for _, swarmNode := range swarmNodes {
		fmt.Printf("%s | %s | isLeader = %t | %s\n", swarmNode.Description.Hostname, swarmNode.Spec.Role, swarmNode.ManagerStatus.Leader, swarmNode.Status.State)
	}

}
