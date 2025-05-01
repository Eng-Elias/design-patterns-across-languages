// Define types (no changes needed)
export type HttpRequest = {
  path: string;
  headers: { [key: string]: string };
  user?: { id: number; role: string };
  requiredRole?: string;
};

export type HttpResponse = {
  statusCode: number;
  body: string;
};

// Interface for all middleware handlers
export interface Middleware {
  setNext(middleware: Middleware): Middleware;
  handle(request: HttpRequest): Promise<HttpResponse>;
}

// Abstract class manages the link, defines the handle signature
export abstract class AbstractMiddleware implements Middleware {
  protected nextMiddleware: Middleware | null = null; // Changed to protected for clarity

  public setNext(middleware: Middleware): Middleware {
    this.nextMiddleware = middleware;
    return middleware;
  }

  // Define the abstract method signature that subclasses must implement
  abstract handle(request: HttpRequest): Promise<HttpResponse>;
}

// Concrete Middleware Implementations

export class LoggingMiddleware extends AbstractMiddleware {
  public async handle(request: HttpRequest): Promise<HttpResponse> {
    console.log(`[Log] Received request: ${request.path}`);
    // Explicitly call the next handler's handle method if it exists
    if (this.nextMiddleware) {
      return await this.nextMiddleware.handle(request);
    }
    // End of chain if this were the last handler
    return {
      statusCode: 200,
      body: "Request logged but not processed further",
    };
  }
}

export class AuthenticationMiddleware extends AbstractMiddleware {
  public async handle(request: HttpRequest): Promise<HttpResponse> {
    console.log("[Auth] Checking authentication...");
    const token = request.headers["authorization"];

    if (token === "valid_token") {
      console.log("[Auth] Authentication successful.");
      request.user = { id: 456, role: "admin" }; // Mutate request
      // Explicitly call the next handler's handle method if it exists
      if (this.nextMiddleware) {
        return await this.nextMiddleware.handle(request);
      }
      // Auth passed, but end of chain
      return {
        statusCode: 200,
        body: "Authentication successful but not processed further",
      };
    } else {
      console.log("[Auth] Authentication failed. Aborting request.");
      // Stop the chain by returning the response directly
      return { statusCode: 401, body: "Unauthorized" };
    }
  }
}

export class AuthorizationMiddleware extends AbstractMiddleware {
  public async handle(request: HttpRequest): Promise<HttpResponse> {
    console.log("[ACL] Checking authorization...");
    const user = request.user;
    const requiredRole = request.requiredRole;

    if (!user) {
      console.log(
        "[ACL] No user found in request (likely due to failed auth). Aborting."
      );
      // Stop the chain
      return { statusCode: 401, body: "Authentication Required" };
    }

    let proceed = false;
    if (requiredRole && user.role === requiredRole) {
      console.log(`[ACL] Authorization successful (Role: ${user.role}).`);
      proceed = true;
    } else if (!requiredRole) {
      console.log("[ACL] No specific role required. Proceeding.");
      proceed = true;
    }

    if (proceed) {
      // Explicitly call the next handler's handle method if it exists
      if (this.nextMiddleware) {
        return await this.nextMiddleware.handle(request);
      }
      // Auth passed, but end of chain
      return {
        statusCode: 200,
        body: "Authorization successful but not processed further",
      };
    } else {
      // This case handles: required_role exists, but user role doesn't match
      console.log(
        `[ACL] Authorization failed. User role '${user.role}' does not match required role '${requiredRole}'. Aborting.`
      );
      // Stop the chain
      return { statusCode: 403, body: "Forbidden" };
    }
  }
}

export class FinalHandler extends AbstractMiddleware {
  // Implement the abstract handle method
  public async handle(request: HttpRequest): Promise<HttpResponse> {
    console.log("[Final] Processing request...");
    const userId = request.user?.id ?? "anonymous"; // Correctly access potential user ID
    const body = `Successfully processed request for path: ${request.path} for user ${userId}`;
    console.log("[Final] Request processed.");
    // This is the end of the chain, return the final response
    return { statusCode: 200, body: body };
  }

  // Override setNext to prevent chaining further (good practice)
  public setNext(middleware: Middleware): Middleware {
    throw new Error("FinalHandler cannot have a next handler.");
  }
}

// Helper function (no changes needed)
export function makeHttpRequest(
  path: string,
  token?: string,
  requiresRole?: string
): HttpRequest {
  const request: HttpRequest = {
    path: path,
    headers: {},
  };
  if (token) {
    request.headers["authorization"] = token;
  }
  if (requiresRole) {
    request.requiredRole = requiresRole;
  }
  return request;
}
