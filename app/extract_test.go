package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractVideoDevices(t *testing.T) {
	result, _ := ExtractVideoDevices()

	assert.Equal(t, result, []string{"USB Video device", "HD Pro Webcam C920"})
}

func TestExtractPreferableDeviceName(t *testing.T) {
	actual1 := ExtractPreferableDeviceName([]string{"USB Video device", "HD Pro Webcam C920"})
	assert.Equal(t, "HD Pro Webcam C920", actual1)

	actual2 := ExtractPreferableDeviceName([]string{"FaceTime HD Camera"})
	assert.Equal(t, "FaceTime HD Camera", actual2)

	actual3 := ExtractPreferableDeviceName([]string{"iSight", "DV"})
	assert.Equal(t, "iSight", actual3)
}
