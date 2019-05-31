package sn

import (
	"os/exec"
	"strings"
)

func get() (string, error) {
	cmd := exec.Command("wmic", "csproduct", "get", "uuid")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(strings.Replace(strings.Replace(string(out), "-", "", -1), "UUID", "", -1)), nil
}
