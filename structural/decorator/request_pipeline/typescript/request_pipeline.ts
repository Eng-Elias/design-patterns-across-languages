// --- Types ---
export interface Request {
  headers: Record<string, string>;
  body: Record<string, any>;
  path: string;
  method: string;
}

export interface Response {
  statusCode: number;
  body: Record<string, any>;
}

// --- Component Interface ---
export interface RequestHandler {
  handle(request: Request): Response;
}

// --- Concrete Component ---
export class BaseHandler implements RequestHandler {
  public handle(request: Request): Response {
    // In a real application, this would contain the core business logic
    return {
      statusCode: 200,
      body: { message: "Request processed successfully" },
    };
  }
}

// --- Decorator Base Class ---
export abstract class Middleware implements RequestHandler {
  protected handler: RequestHandler;

  constructor(handler: RequestHandler) {
    this.handler = handler;
  }

  public abstract handle(request: Request): Response;
}

// --- Concrete Decorators ---
export class AuthMiddleware extends Middleware {
  public handle(request: Request): Response {
    // Check if the request has a valid auth token
    if (!request.headers["Authorization"]) {
      return {
        statusCode: 401,
        body: { error: "Unauthorized" },
      };
    }

    // If authenticated, pass to the next handler
    return this.handler.handle(request);
  }
}

export class LoggingMiddleware extends Middleware {
  public handle(request: Request): Response {
    // Log request details
    console.log(`Request received: ${request.method} ${request.path}`);
    console.log(`Headers: ${JSON.stringify(request.headers)}`);
    console.log(`Body: ${JSON.stringify(request.body)}`);

    // Process the request
    const response = this.handler.handle(request);

    // Log response
    console.log(`Response: ${response.statusCode}`);
    console.log(`Body: ${JSON.stringify(response.body)}`);

    return response;
  }
}

export class RateLimitMiddleware extends Middleware {
  private maxRequests: number;
  private requestCount: number = 0;

  constructor(handler: RequestHandler, maxRequests: number = 100) {
    super(handler);
    this.maxRequests = maxRequests;
  }

  public handle(request: Request): Response {
    this.requestCount++;

    if (this.requestCount > this.maxRequests) {
      return {
        statusCode: 429,
        body: { error: "Too many requests" },
      };
    }

    return this.handler.handle(request);
  }
}
