import unittest
from video_streaming import VideoStreamingService


class TestVideoStreamingService(unittest.TestCase):
    def setUp(self):
        self.service = VideoStreamingService()

    def test_upload_and_process_video_success(self):
        result = self.service.upload_and_process_video(
            username="test_user",
            password="password",
            file_path="/path/to/video.webm",
            title="Test Video"
        )

        self.assertEqual(result["video_id"], "video_123")
        self.assertTrue(result["storage_url"].startswith("https://storage.example.com/videos/"))
        self.assertEqual(len(result["cdn_urls"]), 2)
        self.assertTrue(all(url.startswith("https://cdn-") for url in result["cdn_urls"]))

    def test_upload_and_process_video_auth_failure(self):
        with self.assertRaises(ValueError):
            self.service.upload_and_process_video(
                username="wrong_user",
                password="wrong_password",
                file_path="/path/to/video.webm",
                title="Test Video"
            )

    def test_stream_video_success(self):
        url = self.service.stream_video(
            username="test_user",
            password="password",
            video_id="video_123"
        )

        self.assertTrue(url.startswith("https://storage.example.com/videos/"))

    def test_stream_video_auth_failure(self):
        with self.assertRaises(ValueError):
            self.service.stream_video(
                username="wrong_user",
                password="wrong_password",
                video_id="video_123"
            )


if __name__ == '__main__':
    unittest.main() 