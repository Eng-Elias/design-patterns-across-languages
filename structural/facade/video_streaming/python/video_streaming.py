from dataclasses import dataclass
from typing import Dict, List, Optional
from enum import Enum


class VideoFormat(Enum):
    MP4 = "mp4"
    WEBM = "webm"
    MOV = "mov"


class VideoQuality(Enum):
    LOW = "low"
    MEDIUM = "medium"
    HIGH = "high"
    ULTRA = "ultra"


@dataclass
class Video:
    id: str
    title: str
    format: VideoFormat
    quality: VideoQuality
    size: int
    duration: int


@dataclass
class User:
    id: str
    username: str
    email: str


# --- Subsystems ---
class VideoUploader:
    def upload_video(self, file_path: str, title: str) -> Video:
        print(f"Uploading video: {title} from {file_path}")
        # Simulate video upload
        return Video(
            id="video_123",
            title=title,
            format=VideoFormat.WEBM,
            quality=VideoQuality.HIGH,
            size=1024 * 1024 * 100,  # 100MB
            duration=300  # 5 minutes
        )


class VideoEncoder:
    def encode_video(self, video: Video, target_format: VideoFormat) -> Video:
        print(f"Encoding video {video.id} to {target_format.value}")
        # Simulate video encoding
        return Video(
            id=video.id,
            title=video.title,
            format=target_format,
            quality=video.quality,
            size=video.size,
            duration=video.duration
        )


class StorageService:
    def store_video(self, video: Video) -> str:
        print(f"Storing video {video.id} in cloud storage")
        # Simulate storage
        return f"https://storage.example.com/videos/{video.id}"

    def get_video_url(self, video_id: str) -> str:
        return f"https://storage.example.com/videos/{video_id}"


class CDNService:
    def distribute_video(self, video_url: str) -> List[str]:
        print(f"Distributing video {video_url} through CDN")
        # Simulate CDN distribution
        return [
            f"https://cdn-1.example.com/videos/{video_url.split('/')[-1]}",
            f"https://cdn-2.example.com/videos/{video_url.split('/')[-1]}"
        ]


class AuthenticationService:
    def authenticate_user(self, username: str, password: str) -> Optional[User]:
        print(f"Authenticating user: {username}")
        # Simulate authentication
        if username == "test_user" and password == "password":
            return User(
                id="user_123",
                username=username,
                email=f"{username}@example.com"
            )
        return None

    def authorize_video_access(self, user: User, video_id: str) -> bool:
        print(f"Authorizing user {user.username} for video {video_id}")
        # Simulate authorization
        return True


class NotificationService:
    def send_upload_notification(self, user: User, video: Video) -> None:
        print(f"Sending upload notification to {user.email}")
        # Simulate notification
        pass

    def send_processing_notification(self, user: User, video: Video) -> None:
        print(f"Sending processing notification to {user.email}")
        # Simulate notification
        pass

    def send_ready_notification(self, user: User, video: Video) -> None:
        print(f"Sending ready notification to {user.email}")
        # Simulate notification
        pass


# --- Facade ---
class VideoStreamingService:
    def __init__(self):
        self.uploader = VideoUploader()
        self.encoder = VideoEncoder()
        self.storage = StorageService()
        self.cdn = CDNService()
        self.auth = AuthenticationService()
        self.notifier = NotificationService()

    def upload_and_process_video(
        self,
        username: str,
        password: str,
        file_path: str,
        title: str
    ) -> Dict[str, str]:
        # Authenticate user
        user = self.auth.authenticate_user(username, password)
        if not user:
            raise ValueError("Authentication failed")

        # Upload video
        video = self.uploader.upload_video(file_path, title)
        self.notifier.send_upload_notification(user, video)

        # Encode video
        encoded_video = self.encoder.encode_video(video, VideoFormat.MP4)
        self.notifier.send_processing_notification(user, encoded_video)

        # Store video
        storage_url = self.storage.store_video(encoded_video)

        # Distribute through CDN
        cdn_urls = self.cdn.distribute_video(storage_url)
        self.notifier.send_ready_notification(user, encoded_video)

        return {
            "video_id": video.id,
            "storage_url": storage_url,
            "cdn_urls": cdn_urls
        }

    def stream_video(
        self,
        username: str,
        password: str,
        video_id: str
    ) -> str:
        # Authenticate user
        user = self.auth.authenticate_user(username, password)
        if not user:
            raise ValueError("Authentication failed")

        # Authorize access
        if not self.auth.authorize_video_access(user, video_id):
            raise ValueError("Unauthorized access")

        # Get video URL
        return self.storage.get_video_url(video_id) 