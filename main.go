package main

import (
	. "github.com/nfedyashev/myapp/app"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"github.com/nfedyashev/data_collector/app"
)

func main() {
	if os.Getenv("BUCKET_NAME") == "" {
		log.Fatal("oops, BUCKET_NAME system variable has to be provided")
	}

	destinationBucketPath := strings.Join([]string{"s3://", os.Getenv("BUCKET_NAME"), "/"}, "")

	all_devices, err  := ExtractVideoDevices()
	if err != nil {
		log.Fatal("can not get list of devices")
	}

	device_name := ExtractPreferableDeviceName(all_devices)
	fmt.Println("imagesnap'ing using device: " + device_name)


	//hostName := os.Hostname()

	ticker := time.NewTicker(15 * time.Minute)
	for {
		select {
		case <-ticker.C:
			t := time.Now()
			fmt.Println(t.Format("3:04PM"), "starting...")

			pathToScreenshot := strings.Join([]string{"tmp/", fmt.Sprintf("%v", int32(time.Now().Unix())), "-screenshot.png"}, "")
			pathToSnapshot := strings.Join([]string{"tmp/", fmt.Sprintf("%v", int32(time.Now().Unix())), "-snapshot.jpg"}, "")

			if app.Screencapture(pathToScreenshot) == nil {
				fmt.Println("Successfully saved screenshot")

				UploadToS3(pathToScreenshot, destinationBucketPath)
				Cleanup(pathToScreenshot)
			}

			if app.Snapshot(pathToSnapshot, device_name) == nil {
				fmt.Println("Successfully made snapshot")

				UploadToS3(pathToSnapshot, destinationBucketPath)
				Cleanup(pathToSnapshot)
			}

			fmt.Println("")
		}
	}
}
