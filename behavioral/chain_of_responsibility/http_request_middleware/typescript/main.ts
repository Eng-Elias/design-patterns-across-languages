import {
  Middleware,
  LoggingMiddleware,
  AuthenticationMiddleware,
  AuthorizationMiddleware,
  FinalHandler,
  makeHttpRequest,
  HttpResponse,
} from "./http_request_middleware";

function setupMiddlewareChain(): Middleware {
  const logger = new LoggingMiddleware();
  const authenticator = new AuthenticationMiddleware();
  const authorizer = new AuthorizationMiddleware();
  const finalHandler = new FinalHandler();

  // Build the chain: Log -> Auth -> Authorize -> Final
  logger.setNext(authenticator).setNext(authorizer).setNext(finalHandler);

  return logger; // Return the entry point of the chain
}

async function runSimulations() {
  const middlewareChain = setupMiddlewareChain();

  console.log("--- Simulating a successful request ---");
  const request1 = makeHttpRequest("/admin/dashboard", "valid_token", "admin");
  const response1 = await middlewareChain.handle(request1);
  console.log(`Response 1: ${JSON.stringify(response1)}\n`);

  console.log("--- Simulating a request with invalid authentication ---");
  const request2 = makeHttpRequest("/user/profile", "invalid_token");
  const response2 = await middlewareChain.handle(request2);
  console.log(`Response 2: ${JSON.stringify(response2)}\n`);

  console.log("--- Simulating a request with insufficient authorization ---");
  const request3 = makeHttpRequest(
    "/admin/settings",
    "valid_token",
    "superadmin"
  );
  const response3 = await middlewareChain.handle(request3);
  console.log(`Response 3: ${JSON.stringify(response3)}\n`);

  console.log("--- Simulating a request requiring no specific role ---");
  const request4 = makeHttpRequest("/public/info", "valid_token");
  const response4 = await middlewareChain.handle(request4);
  console.log(`Response 4: ${JSON.stringify(response4)}\n`);
}

// Execute the simulations
runSimulations().catch((error) => {
  console.error("An error occurred during simulation:", error);
});
