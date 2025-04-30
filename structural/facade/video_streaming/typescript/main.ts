import { VideoStreamingService } from "./video_streaming";

function main(): void {
  // Create the video streaming service
  const service = new VideoStreamingService();

  console.log("--- Uploading and Processing Video ---");
  try {
    // Upload and process a video
    const result = service.uploadAndProcessVideo(
      "test_user",
      "password",
      "/path/to/video.webm",
      "My Awesome Video"
    );

    console.log("Video uploaded successfully!");
    console.log(`Video ID: ${result.videoId}`);
    console.log(`Storage URL: ${result.storageUrl}`);
    console.log("CDN URLs:");
    result.cdnUrls.forEach((url) => console.log(`  - ${url}`));

    console.log("\n--- Streaming Video ---");
    // Stream the video
    const streamUrl = service.streamVideo(
      "test_user",
      "password",
      result.videoId
    );
    console.log(`Stream URL: ${streamUrl}`);
  } catch (error: any) {
    console.error(`Error: ${error.message}`);
  }

  console.log("\n--- Testing Authentication Failure ---");
  try {
    // Try with wrong credentials
    service.uploadAndProcessVideo(
      "wrong_user",
      "wrong_password",
      "/path/to/video.webm",
      "My Awesome Video"
    );
  } catch (error: any) {
    console.log(`Expected error: ${error.message}`);
  }
}

main();
