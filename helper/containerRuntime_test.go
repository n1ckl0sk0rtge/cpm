package helper

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestAvailable(t *testing.T) {

	// podman
	podman := exec.Command("sh", "-c", "podman ps")
	_, err := podman.Output()

	if err != nil {
		assert.Error(t, podmanAvailable())
	} else {
		assert.NoError(t, podmanAvailable())
	}

	// docker
	docker := exec.Command("sh", "-c", "docker ps")
	_, err = docker.Output()

	if err != nil {
		assert.Error(t, dockerAvailable())
	} else {
		assert.NoError(t, dockerAvailable())
	}

}
