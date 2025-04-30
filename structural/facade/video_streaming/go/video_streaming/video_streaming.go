package video_streaming

import (
	"errors"
	"fmt"
	"strings"
)

// VideoFormat represents the format of a video
type VideoFormat string

const (
	MP4  VideoFormat = "mp4"
	WEBM VideoFormat = "webm"
	MOV  VideoFormat = "mov"
)

// VideoQuality represents the quality of a video
type VideoQuality string

const (
	LOW    VideoQuality = "low"
	MEDIUM VideoQuality = "medium"
	HIGH   VideoQuality = "high"
	ULTRA  VideoQuality = "ultra"
)

// Video represents a video file
type Video struct {
	ID       string
	Title    string
	Format   VideoFormat
	Quality  VideoQuality
	Size     int64
	Duration int
}

// User represents a user of the service
type User struct {
	ID       string
	Username string
	Email    string
}

// VideoUploader handles video file uploads
type VideoUploader struct{}

func (u *VideoUploader) UploadVideo(filePath, title string) Video {
	fmt.Printf("Uploading video: %s from %s\n", title, filePath)
	// Simulate video upload
	return Video{
		ID:       "video_123",
		Title:    title,
		Format:   WEBM,
		Quality:  HIGH,
		Size:     1024 * 1024 * 100, // 100MB
		Duration: 300,               // 5 minutes
	}
}

// VideoEncoder handles video encoding
type VideoEncoder struct{}

func (e *VideoEncoder) EncodeVideo(video Video, targetFormat VideoFormat) Video {
	fmt.Printf("Encoding video %s to %s\n", video.ID, targetFormat)
	// Simulate video encoding
	return Video{
		ID:       video.ID,
		Title:    video.Title,
		Format:   targetFormat,
		Quality:  video.Quality,
		Size:     video.Size,
		Duration: video.Duration,
	}
}

// StorageService handles video storage
type StorageService struct{}

func (s *StorageService) StoreVideo(video Video) string {
	fmt.Printf("Storing video %s in cloud storage\n", video.ID)
	// Simulate storage
	return fmt.Sprintf("https://storage.example.com/videos/%s", video.ID)
}

func (s *StorageService) GetVideoUrl(videoId string) string {
	return fmt.Sprintf("https://storage.example.com/videos/%s", videoId)
}

// CDNService handles content delivery
type CDNService struct{}

func (c *CDNService) DistributeVideo(videoUrl string) []string {
	fmt.Printf("Distributing video %s through CDN\n", videoUrl)
	// Simulate CDN distribution
	videoId := videoUrl[strings.LastIndex(videoUrl, "/")+1:]
	return []string{
		fmt.Sprintf("https://cdn-1.example.com/videos/%s", videoId),
		fmt.Sprintf("https://cdn-2.example.com/videos/%s", videoId),
	}
}

// AuthenticationService handles user authentication
type AuthenticationService struct{}

func (a *AuthenticationService) AuthenticateUser(username, password string) (*User, error) {
	fmt.Printf("Authenticating user: %s\n", username)
	// Simulate authentication
	if username == "test_user" && password == "password" {
		return &User{
			ID:       "user_123",
			Username: username,
			Email:    fmt.Sprintf("%s@example.com", username),
		}, nil
	}
	return nil, errors.New("authentication failed")
}

func (a *AuthenticationService) AuthorizeVideoAccess(user *User, videoId string) bool {
	fmt.Printf("Authorizing user %s for video %s\n", user.Username, videoId)
	// Simulate authorization
	return true
}

// NotificationService handles user notifications
type NotificationService struct{}

func (n *NotificationService) SendUploadNotification(user *User, video Video) {
	fmt.Printf("Sending upload notification to %s\n", user.Email)
	// Simulate notification
}

func (n *NotificationService) SendProcessingNotification(user *User, video Video) {
	fmt.Printf("Sending processing notification to %s\n", user.Email)
	// Simulate notification
}

func (n *NotificationService) SendReadyNotification(user *User, video Video) {
	fmt.Printf("Sending ready notification to %s\n", user.Email)
	// Simulate notification
}

// VideoStreamingService is the facade that simplifies the interaction with the subsystems
type VideoStreamingService struct {
	uploader    *VideoUploader
	encoder     *VideoEncoder
	storage     *StorageService
	cdn         *CDNService
	auth        *AuthenticationService
	notifier    *NotificationService
}

// NewVideoStreamingService creates a new VideoStreamingService instance
func NewVideoStreamingService() *VideoStreamingService {
	return &VideoStreamingService{
		uploader:    &VideoUploader{},
		encoder:     &VideoEncoder{},
		storage:     &StorageService{},
		cdn:         &CDNService{},
		auth:        &AuthenticationService{},
		notifier:    &NotificationService{},
	}
}

// UploadAndProcessVideo handles the entire video upload and processing workflow
func (s *VideoStreamingService) UploadAndProcessVideo(
	username, password, filePath, title string,
) (map[string]interface{}, error) {
	// Authenticate user
	user, err := s.auth.AuthenticateUser(username, password)
	if err != nil {
		return nil, err
	}

	// Upload video
	video := s.uploader.UploadVideo(filePath, title)
	s.notifier.SendUploadNotification(user, video)

	// Encode video
	encodedVideo := s.encoder.EncodeVideo(video, MP4)
	s.notifier.SendProcessingNotification(user, encodedVideo)

	// Store video
	storageUrl := s.storage.StoreVideo(encodedVideo)

	// Distribute through CDN
	cdnUrls := s.cdn.DistributeVideo(storageUrl)
	s.notifier.SendReadyNotification(user, encodedVideo)

	return map[string]interface{}{
		"video_id":    video.ID,
		"storage_url": storageUrl,
		"cdn_urls":    cdnUrls,
	}, nil
}

// StreamVideo handles video streaming
func (s *VideoStreamingService) StreamVideo(
	username, password, videoId string,
) (string, error) {
	// Authenticate user
	user, err := s.auth.AuthenticateUser(username, password)
	if err != nil {
		return "", err
	}

	// Authorize access
	if !s.auth.AuthorizeVideoAccess(user, videoId) {
		return "", errors.New("unauthorized access")
	}

	// Get video URL
	return s.storage.GetVideoUrl(videoId), nil
} 