package docker

import (
	"os/exec"
)

// Restart will restart a docker or throw an error if something happens
func Restart(dockerID string) error {
	_, err := exec.Command("docker", "restart", dockerID).Output()
	if err != nil {
		return err
	}

	return nil
}
