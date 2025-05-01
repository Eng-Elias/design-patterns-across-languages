import {
  LoggingMiddleware,
  AuthenticationMiddleware,
  AuthorizationMiddleware,
  FinalHandler,
  makeHttpRequest,
  Middleware,
  HttpResponse,
} from "./http_request_middleware";

// Mock console.log to prevent test output clutter and allow assertions
let consoleOutput: string[] = [];
const originalConsoleLog = console.log;
const mockedConsoleLog = (output: string) => {
  consoleOutput.push(output);
};

describe("HttpRequest Middleware Chain", () => {
  let chain: Middleware;
  let logger: LoggingMiddleware;
  let authenticator: AuthenticationMiddleware;
  let authorizer: AuthorizationMiddleware;
  let finalHandler: FinalHandler;

  beforeEach(() => {
    // Reset console log mock before each test
    console.log = mockedConsoleLog;
    consoleOutput = [];

    // Setup the chain
    logger = new LoggingMiddleware();
    authenticator = new AuthenticationMiddleware();
    authorizer = new AuthorizationMiddleware();
    finalHandler = new FinalHandler();

    // Properly set up the chain
    logger.setNext(authenticator);
    authenticator.setNext(authorizer);
    authorizer.setNext(finalHandler);

    // Start with the logger
    chain = logger;
  });

  afterEach(() => {
    // Restore original console.log after each test
    console.log = originalConsoleLog;
  });

  test("should process a successful request for admin role", async () => {
    const request = makeHttpRequest("/admin", "valid_token", "admin");
    const response = await chain.handle(request);

    expect(response.statusCode).toBe(200);
    expect(response.body).toContain("Successfully processed request");
    expect(response.body).toContain("user 456"); // Check user info propagated

    // Check logs to ensure all middleware ran
    const logString = consoleOutput.join("\n");
    expect(logString).toContain("[Log] Received request");
    expect(logString).toContain("[Auth] Authentication successful");
    expect(logString).toContain("[ACL] Authorization successful");
    expect(logString).toContain("[Final] Processing request");
  });

  test("should reject request with invalid authentication", async () => {
    const request = makeHttpRequest("/secure", "invalid_token");
    const response = await chain.handle(request);

    expect(response.statusCode).toBe(401);
    expect(response.body).toBe("Unauthorized");

    // Check logs to ensure chain stopped at Auth
    const logString = consoleOutput.join("\n");
    expect(logString).toContain("[Log] Received request");
    expect(logString).toContain(
      "[Auth] Authentication failed. Aborting request."
    );
    expect(logString).not.toContain("[ACL] Checking authorization");
    expect(logString).not.toContain("[Final] Processing request");
  });

  test("should reject request with insufficient authorization", async () => {
    const request = makeHttpRequest(
      "/admin/super",
      "valid_token",
      "superadmin"
    );
    const response = await chain.handle(request);

    expect(response.statusCode).toBe(403);
    expect(response.body).toBe("Forbidden");

    // Check logs to ensure chain stopped at ACL
    const logString = consoleOutput.join("\n");
    expect(logString).toContain("[Log] Received request");
    expect(logString).toContain("[Auth] Authentication successful");
    expect(logString).toContain("[ACL] Authorization failed.");
    expect(logString).not.toContain("[Final] Processing request");
  });

  test("should process successfully when no specific role is required", async () => {
    const request = makeHttpRequest("/public", "valid_token");
    const response = await chain.handle(request);

    expect(response.statusCode).toBe(200);
    expect(response.body).toContain("Successfully processed request");

    // Check logs to ensure ACL passed correctly
    const logString = consoleOutput.join("\n");
    expect(logString).toContain("[Log] Received request");
    expect(logString).toContain("[Auth] Authentication successful");
    expect(logString).toContain("[ACL] No specific role required. Proceeding.");
    expect(logString).toContain("[Final] Processing request");
  });

  test("FinalHandler should not allow setting next", () => {
    const final = new FinalHandler();
    const another = new LoggingMiddleware();
    expect(() => final.setNext(another)).toThrow(
      "FinalHandler cannot have a next handler."
    );
  });
});
