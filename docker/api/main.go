package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/images", Images)
	router.HandleFunc("/containers", Containers)
	router.HandleFunc("/networks", Networks)
	router.HandleFunc("/swarm-nodes", SwarmNodes)
	log.Fatal(http.ListenAndServe(":9090", router))

}

func Images(w http.ResponseWriter, r *http.Request) {

	cli, err := newDockerClient()
	if err != nil {
		panic(err)
	}

	//List all images available locally
	images, err := cli.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		panic(err)
	}

	htmlOutput := "<html>"
	for _, image := range images {
		htmlOutput += image.ID + " | " + strconv.Itoa(int(image.Size)) + "<br/>"
	}
	htmlOutput += "</html>"
	fmt.Fprint(w, htmlOutput)
}

func Containers(w http.ResponseWriter, r *http.Request) {

	cli, err := newDockerClient()
	if err != nil {
		panic(err)
	}

	//Retrieve a list of containers
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	//Iterate through all containers and display each container's properties
	//fmt.Println("Image ID | Repo Tags | Size")
	htmlOutput := "<html>"
	for _, container := range containers {
		htmlOutput += strings.Join(container.Names, ",") + " | " + container.Image + "<br/>"
	}
	htmlOutput += "</html>"
	fmt.Fprint(w, htmlOutput)

}

func Networks(w http.ResponseWriter, r *http.Request) {

	cli, err := newDockerClient()
	if err != nil {
		panic(err)
	}

	networks, err := cli.NetworkList(context.Background(), network.ListOptions{})
	if err != nil {
		panic(err)
	}

	//List all networks
	htmlOutput := "<html>"
	//fmt.Println("Network Name | ID")
	for _, network := range networks {
		htmlOutput += network.Name + " | " + network.ID + "<br/>"
	}
	htmlOutput += "</html>"
	fmt.Fprint(w, htmlOutput)

}

func SwarmNodes(w http.ResponseWriter, r *http.Request) {

	cli, err := newDockerClient()
	if err != nil {
		panic(err)
	}

	swarmNodes, err := cli.NodeList(context.Background(), swarm.NodeListOptions{})
	if err != nil {
		panic(err)
	}

	//List all nodes - works only in Swarm Mode
	htmlOutput := "<html>"
	//fmt.Println("Name | Role | Leader | Status")
	for _, swarmNode := range swarmNodes {
		htmlOutput += swarmNode.Description.Hostname
	}
	htmlOutput += "</html>"
	fmt.Fprint(w, htmlOutput)

}

func newDockerClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
}
