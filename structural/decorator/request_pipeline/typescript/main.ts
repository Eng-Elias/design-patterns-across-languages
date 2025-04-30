import {
  Request,
  BaseHandler,
  AuthMiddleware,
  LoggingMiddleware,
  RateLimitMiddleware,
} from "./request_pipeline";

function printResponse(response: {
  statusCode: number;
  body: Record<string, any>;
}): void {
  console.log(`Status Code: ${response.statusCode}`);
  console.log(`Body: ${JSON.stringify(response.body)}\n`);
}

function main(): void {
  // Create a sample request
  const request: Request = {
    headers: { "Content-Type": "application/json" },
    body: { data: "test" },
    path: "/api/test",
    method: "POST",
  };

  console.log("--- Simple Request Processing ---");
  // Process request with base handler only
  const baseHandler = new BaseHandler();
  let response = baseHandler.handle(request);
  printResponse(response);

  console.log("--- Request with Auth Middleware ---");
  // Process request with auth middleware
  const authHandler = new AuthMiddleware(baseHandler);
  response = authHandler.handle(request);
  printResponse(response);

  // Add authorization header and try again
  request.headers["Authorization"] = "Bearer valid-token";
  response = authHandler.handle(request);
  printResponse(response);

  console.log("--- Request with Rate Limiting ---");
  // Process request with rate limiting
  const rateLimitHandler = new RateLimitMiddleware(baseHandler, 2);
  for (let i = 0; i < 3; i++) {
    console.log(`Request ${i + 1}:`);
    response = rateLimitHandler.handle(request);
    printResponse(response);
  }

  console.log("--- Request with Multiple Middleware ---");
  // Create a chain of middleware: Auth -> RateLimit -> Logging -> BaseHandler
  const handler = new AuthMiddleware(
    new RateLimitMiddleware(new LoggingMiddleware(new BaseHandler()), 2)
  );

  // Process request through the middleware chain
  response = handler.handle(request);
  printResponse(response);
}

main();
