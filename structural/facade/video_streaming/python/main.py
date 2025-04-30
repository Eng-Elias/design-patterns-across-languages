from video_streaming import VideoStreamingService


def main():
    # Create the video streaming service
    service = VideoStreamingService()

    print("--- Uploading and Processing Video ---")
    try:
        # Upload and process a video
        result = service.upload_and_process_video(
            username="test_user",
            password="password",
            file_path="/path/to/video.webm",
            title="My Awesome Video"
        )

        print("Video uploaded successfully!")
        print(f"Video ID: {result['video_id']}")
        print(f"Storage URL: {result['storage_url']}")
        print("CDN URLs:")
        for url in result['cdn_urls']:
            print(f"  - {url}")

        print("\n--- Streaming Video ---")
        # Stream the video
        stream_url = service.stream_video(
            username="test_user",
            password="password",
            video_id=result['video_id']
        )
        print(f"Stream URL: {stream_url}")

    except ValueError as e:
        print(f"Error: {e}")

    print("\n--- Testing Authentication Failure ---")
    try:
        # Try with wrong credentials
        service.upload_and_process_video(
            username="wrong_user",
            password="wrong_password",
            file_path="/path/to/video.webm",
            title="My Awesome Video"
        )
    except ValueError as e:
        print(f"Expected error: {e}")


if __name__ == "__main__":
    main() 