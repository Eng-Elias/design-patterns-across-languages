import {
  Request,
  Response,
  BaseHandler,
  AuthMiddleware,
  LoggingMiddleware,
  RateLimitMiddleware,
} from "./request_pipeline";

describe("Request Processing Pipeline", () => {
  let request: Request;

  beforeEach(() => {
    request = {
      headers: { "Content-Type": "application/json" },
      body: { data: "test" },
      path: "/api/test",
      method: "POST",
    };
  });

  test("BaseHandler processes request correctly", () => {
    const handler = new BaseHandler();
    const response = handler.handle(request);

    expect(response.statusCode).toBe(200);
    expect(response.body).toEqual({
      message: "Request processed successfully",
    });
  });

  test("AuthMiddleware blocks unauthorized requests", () => {
    const handler = new AuthMiddleware(new BaseHandler());
    const response = handler.handle(request);

    expect(response.statusCode).toBe(401);
    expect(response.body).toEqual({ error: "Unauthorized" });
  });

  test("AuthMiddleware allows authorized requests", () => {
    request.headers["Authorization"] = "Bearer valid-token";
    const handler = new AuthMiddleware(new BaseHandler());
    const response = handler.handle(request);

    expect(response.statusCode).toBe(200);
    expect(response.body).toEqual({
      message: "Request processed successfully",
    });
  });

  test("RateLimitMiddleware limits requests", () => {
    const handler = new RateLimitMiddleware(new BaseHandler(), 2);

    // First request should succeed
    let response = handler.handle(request);
    expect(response.statusCode).toBe(200);

    // Second request should succeed
    response = handler.handle(request);
    expect(response.statusCode).toBe(200);

    // Third request should fail
    response = handler.handle(request);
    expect(response.statusCode).toBe(429);
    expect(response.body).toEqual({ error: "Too many requests" });
  });

  test("Middleware chain works correctly", () => {
    // Create a chain: Auth -> RateLimit -> BaseHandler
    const handler = new AuthMiddleware(
      new RateLimitMiddleware(new BaseHandler(), 2)
    );

    // Test with valid auth
    request.headers["Authorization"] = "Bearer valid-token";
    let response = handler.handle(request);
    expect(response.statusCode).toBe(200);

    // Test with invalid auth
    delete request.headers["Authorization"];
    response = handler.handle(request);
    expect(response.statusCode).toBe(401);
  });
});
