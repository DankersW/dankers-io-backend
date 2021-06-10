package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func trim_first_char(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}

func get_docker_info() []DockerInfo {
	ctx := context.Background()
	docker_cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	data_containers, err := docker_cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	docker_info := []DockerInfo{}
	for _, container := range data_containers {
		var info_item DockerInfo
		info_item.Name = trim_first_char(container.Names[0])
		image := strings.Split(container.Image, ":")

		info_item.Repo = image[0]
		info_item.Version = image[1]
		info_item.Status = container.State
		info_item.Uptime = container.Status
		info_item.Port = parse_ports(container.Ports)

		docker_info = append(docker_info, info_item)
	}
	return docker_info
}

func parse_ports(ports_struct []types.Port) string {
	var ports string = ""
	fmt.Println(ports_struct)
	for _, port := range ports_struct {
		ports += strconv.Itoa(int(port.PublicPort)) + " "
	}
	return "14, 18"
}
