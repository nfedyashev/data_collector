package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScreencapture(t *testing.T) {
	err := Screencapture("screenshot.png")

	assert.Equal(t, err, nil)
}

func TestSnapshot(t *testing.T) {
	err := Snapshot("snapshot.png", "HD Pro Webcam C920")

	assert.Equal(t, err, nil)
}
