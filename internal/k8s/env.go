package k8s

import (
	"fmt"
	"os/exec"
	"strings"
)

type Env string

const (
	EnvUnknown       Env = "unknown"
	EnvGKE               = "gke"
	EnvMinikube          = "minikube"
	EnvDockerDesktop     = "docker-for-desktop"
)

func DetectEnv() (Env, error) {
	cmd := exec.Command("kubectl", "config", "current-context")
	outputBytes, err := cmd.Output()
	if err != nil {
		exitErr, isExit := err.(*exec.ExitError)
		if isExit {
			return EnvUnknown, fmt.Errorf("DetectEnv failed. Output:\n%s", string(exitErr.Stderr))
		}
		return EnvUnknown, fmt.Errorf("DetectEnv: %v", err)
	}

	output := strings.TrimSpace(string(outputBytes))
	return EnvFromString(output), nil
}

func EnvFromString(s string) Env {
	if s == EnvMinikube {
		return EnvMinikube
	} else if s == EnvDockerDesktop {
		return EnvDockerDesktop
	} else if strings.HasPrefix(s, EnvGKE) {
		// GKE context strings look like:
		// gke_blorg-dev_us-central1-b_blorg
		return EnvGKE
	}
	return EnvUnknown
}
