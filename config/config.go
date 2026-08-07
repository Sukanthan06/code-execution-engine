package config

import "time"

type Config struct {
	Timeout     time.Duration
	DockerImage string
	MemoryLimit string
	CPULimit    string
	NetworkMode string
}

var AppConfig = Config{
	Timeout:     5 * time.Second,
	DockerImage: "code-runner",
	MemoryLimit: "128m",
	CPULimit:    "1",
	NetworkMode: "none",
}
