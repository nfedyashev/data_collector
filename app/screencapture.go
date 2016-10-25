package app

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func Screencapture(pathToScreenshot string) Error {
	t := time.Now()
	fmt.Println(t.Format("3:04PM"), "starting to screencapture...")

	cmd := exec.Command("screencapture", "-x", pathToScreenshot)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	return nil
}

func UploadToS3(localPath string, destinationBucketPath string) Error {
	cmd := exec.Command("s3cmd", "put", localPath, destinationBucketPath)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	return nil
}

func Cleanup(pathToLocalFile string) Error {
	cmd := exec.Command("rm", pathToLocalFile)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	return nil
}

func Snapshot(pathToSnapshot string) Error {
	cmd := exec.Command("imagesnap", "-q", "-w", "-d", "HD Pro Webcam C920", "-w", "1", pathToSnapshot)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	return nil
}
