package config

import (
	"testing"
	"time"
)

func TestAppConfigDefaults(t *testing.T) {
	if AppConfig.Timeout != 5*time.Second {
		t.Errorf("expected Timeout to be 5s, got %v", AppConfig.Timeout)
	}
	if AppConfig.DockerImage != "code-runner" {
		t.Errorf("expected DockerImage to be 'code-runner', got %v", AppConfig.DockerImage)
	}
	if AppConfig.MemoryLimit != "128m" {
		t.Errorf("expected MemoryLimit to be '128m', got %v", AppConfig.MemoryLimit)
	}
	if AppConfig.CPULimit != "1" {
		t.Errorf("expected CPULimit to be '1', got %v", AppConfig.CPULimit)
	}
	if AppConfig.NetworkMode != "none" {
		t.Errorf("expected NetworkMode to be 'none', got %v", AppConfig.NetworkMode)
	}
}
