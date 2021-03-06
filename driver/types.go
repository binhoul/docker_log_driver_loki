package driver

import (
	"time"
)

type jsonLogLine struct {
	ContainerId   string            `json:"container_id"`
	ContainerName string            `json:"container_name"`
	StackName     string            `json:"stack_name"`
	ServiceName   string            `json:"service_name"`
	ImageId       string            `json:"image_id"`
	ImageName     string            `json:"image_name"`
	Command       string            `json:"command"`
	Tag           string            `json:"tag"`
	Extra         map[string]string `json:"extra"`
	Host          string            `json:"host"`
	Timestamp     time.Time         `json:"timestamp"`
}
