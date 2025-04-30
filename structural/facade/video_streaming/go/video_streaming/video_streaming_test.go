package video_streaming

import (
	"testing"
)

func TestVideoStreamingService(t *testing.T) {
	service := NewVideoStreamingService()

	t.Run("UploadAndProcessVideo", func(t *testing.T) {
		t.Run("successful upload", func(t *testing.T) {
			result, err := service.UploadAndProcessVideo(
				"test_user",
				"password",
				"/path/to/video.mp4",
				"Test Video",
			)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if result["video_id"] != "video_123" {
				t.Errorf("Expected video_id to be 'video_123', got %v", result["video_id"])
			}

			storageUrl := result["storage_url"].(string)
			if storageUrl != "https://storage.example.com/videos/video_123" {
				t.Errorf("Expected storage_url to be 'https://storage.example.com/videos/video_123', got %v", storageUrl)
			}

			cdnUrls := result["cdn_urls"].([]string)
			if len(cdnUrls) != 2 {
				t.Errorf("Expected 2 CDN URLs, got %d", len(cdnUrls))
			}
		})

		t.Run("authentication failure", func(t *testing.T) {
			_, err := service.UploadAndProcessVideo(
				"wrong_user",
				"wrong_password",
				"/path/to/video.mp4",
				"Test Video",
			)

			if err == nil {
				t.Error("Expected error, got nil")
			}

			if err.Error() != "authentication failed" {
				t.Errorf("Expected error message 'authentication failed', got %v", err)
			}
		})
	})

	t.Run("StreamVideo", func(t *testing.T) {
		t.Run("successful streaming", func(t *testing.T) {
			url, err := service.StreamVideo(
				"test_user",
				"password",
				"video_123",
			)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if url != "https://storage.example.com/videos/video_123" {
				t.Errorf("Expected URL to be 'https://storage.example.com/videos/video_123', got %v", url)
			}
		})

		t.Run("authentication failure", func(t *testing.T) {
			_, err := service.StreamVideo(
				"wrong_user",
				"wrong_password",
				"video_123",
			)

			if err == nil {
				t.Error("Expected error, got nil")
			}

			if err.Error() != "authentication failed" {
				t.Errorf("Expected error message 'authentication failed', got %v", err)
			}
		})
	})
} 