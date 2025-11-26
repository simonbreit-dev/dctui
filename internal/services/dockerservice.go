package service

import (
	"context"
	"dctui/internal/models"
	"dctui/internal/utils"
	"slices"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerService struct {
	docker *client.Client
}

func NewDockerService() *DockerService {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic("shit")
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

func (s *DockerService) FetchContainersForProject(projectName string) []container.Summary {
	ctx := context.Background()
	containers, err := s.docker.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic("shit")
	}
	containers = utils.Filter[container.Summary](containers, func(c container.Summary) bool {
		projectLabel, exists := c.Labels["com.docker.compose.project"]
		if exists && projectLabel != "" {
			return projectLabel == projectName
		} else {
			return false
		}
	})
	return containers
}

func (s *DockerService) FetchProjects() []models.Project {
	ctx := context.Background()
	containers, err := s.docker.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic("Error fetching containers.")
	}
	projects := make([]models.Project, 0)

	for _, summary := range containers {
		projectLabel, exists := summary.Labels["com.docker.compose.project"]
		if exists && projectLabel != "" {
			projectIdx := slices.IndexFunc(projects, func(p models.Project) bool {
				return p.Name == projectLabel
			})
			if projectIdx >= 0 {
				projects[projectIdx].ContainersCount++
			} else {
				configFile, _ := summary.Labels["com.docker.compose.project.config_files"]
				workingDir, _ := summary.Labels["com.docker.compose.project.working_dir"]
				newProject := models.Project{
					Name:            projectLabel,
					ConfigFile:      configFile,
					WorkingDir:      workingDir,
					ContainersCount: 1,
				}
				projects = append(projects, newProject)
			}
		}

	}
	return projects
}
