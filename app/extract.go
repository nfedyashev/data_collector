package app

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Error interface {
	Error() string
}

func ExtractVideoDevices() ([]string, Error) {
	cmd := exec.Command("imagesnap", "-l")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return []string{}, err
	}

	raw := strings.Split(out.String(), "\n")

	result := []string{}
	for v := 0; v < len(raw); v++ {
		fmt.Println(raw[v])

		if raw[v] != "Video Devices:" && raw[v] != "" {
			result = append(result, raw[v])
		}
	}

	return result, nil
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func ExtractPreferableDeviceName(input []string) string {
	if contains(input, "HD Pro Webcam C920") {
		return "HD Pro Webcam C920"
	}

	if contains(input, "FaceTime HD Camera") {
		return "FaceTime HD Camera"
	}
	return input[0]
}
