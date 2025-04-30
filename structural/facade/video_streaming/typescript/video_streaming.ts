// Enums
export enum VideoFormat {
  MP4 = "mp4",
  WEBM = "webm",
  MOV = "mov",
}

export enum VideoQuality {
  LOW = "low",
  MEDIUM = "medium",
  HIGH = "high",
  ULTRA = "ultra",
}

// Types
export interface Video {
  id: string;
  title: string;
  format: VideoFormat;
  quality: VideoQuality;
  size: number;
  duration: number;
}

export interface User {
  id: string;
  username: string;
  email: string;
}

// Subsystems
class VideoUploader {
  uploadVideo(filePath: string, title: string): Video {
    console.log(`Uploading video: ${title} from ${filePath}`);
    // Simulate video upload
    return {
      id: "video_123",
      title,
      format: VideoFormat.WEBM,
      quality: VideoQuality.HIGH,
      size: 1024 * 1024 * 100, // 100MB
      duration: 300, // 5 minutes
    };
  }
}

class VideoEncoder {
  encodeVideo(video: Video, targetFormat: VideoFormat): Video {
    console.log(`Encoding video ${video.id} to ${targetFormat}`);
    // Simulate video encoding
    return {
      ...video,
      format: targetFormat,
    };
  }
}

class StorageService {
  storeVideo(video: Video): string {
    console.log(`Storing video ${video.id} in cloud storage`);
    // Simulate storage
    return `https://storage.example.com/videos/${video.id}`;
  }

  getVideoUrl(videoId: string): string {
    return `https://storage.example.com/videos/${videoId}`;
  }
}

class CDNService {
  distributeVideo(videoUrl: string): string[] {
    console.log(`Distributing video ${videoUrl} through CDN`);
    // Simulate CDN distribution
    return [
      `https://cdn-1.example.com/videos/${videoUrl.split("/").pop()}`,
      `https://cdn-2.example.com/videos/${videoUrl.split("/").pop()}`,
    ];
  }
}

class AuthenticationService {
  authenticateUser(username: string, password: string): User | null {
    console.log(`Authenticating user: ${username}`);
    // Simulate authentication
    if (username === "test_user" && password === "password") {
      return {
        id: "user_123",
        username,
        email: `${username}@example.com`,
      };
    }
    return null;
  }

  authorizeVideoAccess(user: User, videoId: string): boolean {
    console.log(`Authorizing user ${user.username} for video ${videoId}`);
    // Simulate authorization
    return true;
  }
}

class NotificationService {
  sendUploadNotification(user: User, video: Video): void {
    console.log(`Sending upload notification to ${user.email}`);
    // Simulate notification
  }

  sendProcessingNotification(user: User, video: Video): void {
    console.log(`Sending processing notification to ${user.email}`);
    // Simulate notification
  }

  sendReadyNotification(user: User, video: Video): void {
    console.log(`Sending ready notification to ${user.email}`);
    // Simulate notification
  }
}

// Facade
export class VideoStreamingService {
  private uploader: VideoUploader;
  private encoder: VideoEncoder;
  private storage: StorageService;
  private cdn: CDNService;
  private auth: AuthenticationService;
  private notifier: NotificationService;

  constructor() {
    this.uploader = new VideoUploader();
    this.encoder = new VideoEncoder();
    this.storage = new StorageService();
    this.cdn = new CDNService();
    this.auth = new AuthenticationService();
    this.notifier = new NotificationService();
  }

  uploadAndProcessVideo(
    username: string,
    password: string,
    filePath: string,
    title: string
  ): { videoId: string; storageUrl: string; cdnUrls: string[] } {
    // Authenticate user
    const user = this.auth.authenticateUser(username, password);
    if (!user) {
      throw new Error("Authentication failed");
    }

    // Upload video
    const video = this.uploader.uploadVideo(filePath, title);
    this.notifier.sendUploadNotification(user, video);

    // Encode video
    const encodedVideo = this.encoder.encodeVideo(video, VideoFormat.MP4);
    this.notifier.sendProcessingNotification(user, encodedVideo);

    // Store video
    const storageUrl = this.storage.storeVideo(encodedVideo);

    // Distribute through CDN
    const cdnUrls = this.cdn.distributeVideo(storageUrl);
    this.notifier.sendReadyNotification(user, encodedVideo);

    return {
      videoId: video.id,
      storageUrl,
      cdnUrls,
    };
  }

  streamVideo(username: string, password: string, videoId: string): string {
    // Authenticate user
    const user = this.auth.authenticateUser(username, password);
    if (!user) {
      throw new Error("Authentication failed");
    }

    // Authorize access
    if (!this.auth.authorizeVideoAccess(user, videoId)) {
      throw new Error("Unauthorized access");
    }

    // Get video URL
    return this.storage.getVideoUrl(videoId);
  }
}
