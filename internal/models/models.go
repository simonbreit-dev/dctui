package models

type Project struct {
	Name            string
	ConfigFile      string
	WorkingDir      string
	ContainersCount int
}
