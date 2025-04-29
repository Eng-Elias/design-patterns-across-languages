import {
  MessageSender, // Import the interface for mocking
  InfoNotification,
  WarningNotification,
  UrgentNotification,
} from "./notification_system";

// Create a mock implementation for the MessageSender interface
const mockSender: MessageSender = {
  sendMessage: jest.fn(),
};

describe("Notification System Bridge Tests", () => {
  beforeEach(() => {
    // Reset the mock before each test
    (mockSender.sendMessage as jest.Mock).mockClear();
  });

  test("InfoNotification should send correctly formatted message", () => {
    // Arrange
    const infoNotification = new InfoNotification(mockSender);
    const message = "Test info message.";
    const expectedSubject = "Info";
    const expectedBody = `[INFO] ${message}`;

    // Act
    infoNotification.send(message);

    // Assert
    expect(mockSender.sendMessage).toHaveBeenCalledTimes(1);
    expect(mockSender.sendMessage).toHaveBeenCalledWith(
      expectedSubject,
      expectedBody
    );
  });

  test("WarningNotification should send correctly formatted message", () => {
    // Arrange
    const warningNotification = new WarningNotification(mockSender);
    const message = "Test warning message.";
    const expectedSubject = "Warning";
    const expectedBody = `[WARNING] ${message}`;

    // Act
    warningNotification.send(message);

    // Assert
    expect(mockSender.sendMessage).toHaveBeenCalledTimes(1);
    expect(mockSender.sendMessage).toHaveBeenCalledWith(
      expectedSubject,
      expectedBody
    );
  });

  test("UrgentNotification should send correctly formatted message", () => {
    // Arrange
    const urgentNotification = new UrgentNotification(mockSender);
    const message = "Test urgent message.";
    const expectedSubject = "** URGENT **";
    const expectedBody = `[URGENT ACTION REQUIRED] ${message}`;

    // Act
    urgentNotification.send(message);

    // Assert
    expect(mockSender.sendMessage).toHaveBeenCalledTimes(1);
    expect(mockSender.sendMessage).toHaveBeenCalledWith(
      expectedSubject,
      expectedBody
    );
  });

  test("Different notifications should use the same sender instance", () => {
    // Arrange - using the same mockSender defined above
    const infoNotification = new InfoNotification(mockSender);
    const urgentNotification = new UrgentNotification(mockSender);

    // Act
    infoNotification.send("Info 1");
    urgentNotification.send("Urgent 1");

    // Assert
    expect(mockSender.sendMessage).toHaveBeenCalledTimes(2);
    expect(mockSender.sendMessage).toHaveBeenCalledWith(
      "Info",
      "[INFO] Info 1"
    );
    expect(mockSender.sendMessage).toHaveBeenCalledWith(
      "** URGENT **",
      "[URGENT ACTION REQUIRED] Urgent 1"
    );
  });
});
