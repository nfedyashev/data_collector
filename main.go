package main

import (
	. "github.com/nfedyashev/myapp/app"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if os.Getenv("BUCKET_NAME") == "" {
		log.Fatal("oops, BUCKET_NAME system variable has to be provided")
	}

	destinationBucketPath := strings.Join([]string{"s3://", os.Getenv("BUCKET_NAME"), "/"}, "")

	//hostName := os.Hostname()

	ticker := time.NewTicker(15 * time.Minute)
	for {
		select {
		case <-ticker.C:
			t := time.Now()
			fmt.Println(t.Format("3:04PM"), "starting...")

			pathToScreenshot := strings.Join([]string{fmt.Sprintf("%v", int32(time.Now().Unix())), "-screenshot.png"}, "")
			pathToSnapshot := strings.Join([]string{fmt.Sprintf("%v", int32(time.Now().Unix())), "-snapshot.jpg"}, "")

			if Screencapture(pathToScreenshot) == nil {
				fmt.Println("Successfully saved screenshot")

				UploadToS3(pathToScreenshot, destinationBucketPath)
				Cleanup(pathToScreenshot)
			}

			if Snapshot(pathToSnapshot) == nil {
				fmt.Println("Successfully made snapshot")

				UploadToS3(pathToSnapshot, destinationBucketPath)
				Cleanup(pathToSnapshot)
			}

			fmt.Println("")
		}
	}
}
