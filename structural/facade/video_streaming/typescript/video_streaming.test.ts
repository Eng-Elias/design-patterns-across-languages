import { VideoStreamingService } from "./video_streaming";

describe("VideoStreamingService", () => {
  let service: VideoStreamingService;

  beforeEach(() => {
    service = new VideoStreamingService();
  });

  describe("uploadAndProcessVideo", () => {
    it("should successfully upload and process a video with valid credentials", () => {
      const result = service.uploadAndProcessVideo(
        "test_user",
        "password",
        "/path/to/video.webm",
        "Test Video"
      );

      expect(result.videoId).toBe("video_123");
      expect(result.storageUrl).toMatch(
        /^https:\/\/storage\.example\.com\/videos\//
      );
      expect(result.cdnUrls).toHaveLength(2);
      expect(result.cdnUrls[0]).toMatch(
        /^https:\/\/cdn-1\.example\.com\/videos\//
      );
      expect(result.cdnUrls[1]).toMatch(
        /^https:\/\/cdn-2\.example\.com\/videos\//
      );
    });

    it("should throw an error with invalid credentials", () => {
      expect(() => {
        service.uploadAndProcessVideo(
          "wrong_user",
          "wrong_password",
          "/path/to/video.webm",
          "Test Video"
        );
      }).toThrow("Authentication failed");
    });
  });

  describe("streamVideo", () => {
    it("should successfully stream a video with valid credentials", () => {
      const url = service.streamVideo("test_user", "password", "video_123");

      expect(url).toMatch(/^https:\/\/storage\.example\.com\/videos\//);
    });

    it("should throw an error with invalid credentials", () => {
      expect(() => {
        service.streamVideo("wrong_user", "wrong_password", "video_123");
      }).toThrow("Authentication failed");
    });
  });
});
