package services

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerService struct {
	docker *client.Client
}

func NewDockerService() *DockerService {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		print("shit")
	}
	return &DockerService{
		docker: cli,
	}
}

func (s *DockerService) FetchContainers() []container.Summary {
	ctx := context.Background()
	containers, err := s.docker.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic("shit")
	}
	return containers
}
