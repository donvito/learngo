package main

import (
	"fmt"

	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	//List all volumes
	fmt.Print("\n\n\n")
	fmt.Println("VOLUMES")
	fmt.Printf("-----------------------\n")
	myFilters := filters.NewArgs()
	myFilters.Add("dangling", "false")

	volumes, err := cli.VolumeList(context.Background(), myFilters)
	if err != nil {
		panic(err)
	}

	fmt.Println(volumes)
	// for _, volume := range volumes {

	// }

}
