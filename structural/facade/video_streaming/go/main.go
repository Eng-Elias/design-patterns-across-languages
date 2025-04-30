package main

import (
	"facade_pattern_video_streaming_go/video_streaming"
	"fmt"
)

func main() {
	// Create the video streaming service
	service := video_streaming.NewVideoStreamingService()

	fmt.Println("--- Uploading and Processing Video ---")
	// Upload and process a video
	result, err := service.UploadAndProcessVideo(
		"test_user",
		"password",
		"/path/to/video.webm",
		"My Awesome Video",
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Video uploaded successfully!")
	fmt.Printf("Video ID: %s\n", result["video_id"])
	fmt.Printf("Storage URL: %s\n", result["storage_url"])
	fmt.Println("CDN URLs:")
	for _, url := range result["cdn_urls"].([]string) {
		fmt.Printf("  - %s\n", url)
	}

	fmt.Println("\n--- Streaming Video ---")
	// Stream the video
	streamUrl, err := service.StreamVideo(
		"test_user",
		"password",
		result["video_id"].(string),
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Stream URL: %s\n", streamUrl)

	fmt.Println("\n--- Testing Authentication Failure ---")
	// Try with wrong credentials
	_, err = service.UploadAndProcessVideo(
		"wrong_user",
		"wrong_password",
		"/path/to/video.webm",
		"My Awesome Video",
	)

	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
} 