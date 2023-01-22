package register

import (
	"testing"
)

func TestGetRegistry(t *testing.T) {
	registry := GetRegistry()
	registry.Image = "sonar"

	newRegistry := GetRegistry()

	if newRegistry.Image != registry.Image {
		t.Errorf("Expected %s, got %s", newRegistry.Image, registry.Image)
	}

}
