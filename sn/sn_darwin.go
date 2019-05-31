package sn

import (
	"os/exec"
	"strings"
)

func get() (string, error) {
	cmd := exec.Command("/bin/bash", "-l", "-c", "system_profiler SPHardwareDataType | awk '/Hardware UUID/ {print $NF}'")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(strings.Replace(string(out), "-", "", -1)), nil
}
