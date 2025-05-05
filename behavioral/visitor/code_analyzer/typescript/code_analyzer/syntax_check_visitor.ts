import { Visitor } from "./visitor";
import {
  FunctionDefinitionNode,
  VariableDeclarationNode,
  AssignmentStatementNode,
  IfStatementNode,
  ExpressionStatementNode,
  // ExpressionNode // Needed if checking expression internals
} from "./ast_nodes";

// Performs simple syntax checks, like tracking declared variables using scopes.
export class SyntaxCheckVisitor implements Visitor {
  private errors: string[] = [];
  // Stack of Sets to manage variable scopes (using Set for efficient lookup)
  private scopeStack: Set<string>[] = [new Set<string>()]; // Start with global scope

  getErrors(): string[] {
    return this.errors;
  }

  private addError(message: string): void {
    this.errors.push(message);
  }

  private pushScope(): void {
    this.scopeStack.push(new Set<string>());
  }

  private popScope(): void {
    if (this.scopeStack.length > 1) {
      // Avoid popping the global scope
      this.scopeStack.pop();
    }
  }

  private declareVariable(name: string, nodeInfo: string = "variable"): void {
    const currentScope = this.scopeStack[this.scopeStack.length - 1];
    if (currentScope?.has(name)) {
      this.addError(
        `Syntax Error: Identifier '${name}' already declared as a ${nodeInfo} in this scope.`
      );
    } else {
      currentScope?.add(name);
    }
  }

  private isDeclared(name: string): boolean {
    // Check from current scope outwards
    for (let i = this.scopeStack.length - 1; i >= 0; i--) {
      if (this.scopeStack[i]?.has(name)) {
        return true;
      }
    }
    return false;
  }

  // --- Visitor Methods ---

  visitFunctionDefinition(element: FunctionDefinitionNode): void {
    // Declare function name in the current (outer) scope
    this.declareVariable(element.name, "function");
    // Push a new scope for the function body
    this.pushScope();
    // Declare parameters in the new scope
    element.parameters.forEach((param) =>
      this.declareVariable(param, "parameter")
    );

    // Explicitly traverse the function body within the new scope
    element.body.forEach((stmt) => stmt.accept(this));

    // Pop the scope after visiting the body
    this.popScope();
  }

  visitVariableDeclaration(element: VariableDeclarationNode): void {
    // Declare in the current scope
    this.declareVariable(element.name);
    // Check initializer if present
    element.initializer?.accept(this); // Visit initializer
  }

  visitAssignmentStatement(element: AssignmentStatementNode): void {
    // Check if the target variable exists
    if (!this.isDeclared(element.targetVariable)) {
      this.addError(
        `Syntax Error: Variable '${element.targetVariable}' used before declaration.`
      );
    }
    // Check RHS expression variables if needed
    element.value.accept(this); // Visit RHS value
  }

  visitIfStatement(element: IfStatementNode): void {
    // Check condition variables
    element.condition.accept(this); // Visit condition

    // Push scope for the 'then' branch
    this.pushScope();
    // Explicitly traverse the 'then' branch statements
    element.thenBranch.forEach((stmt) => stmt.accept(this));
    // Pop scope after 'then' branch traversal
    this.popScope();

    if (element.elseBranch) {
      // Push scope for the 'else' branch
      this.pushScope();
      // Explicitly traverse the 'else' branch statements
      element.elseBranch.forEach((stmt) => stmt.accept(this));
      // Pop scope after 'else' branch traversal
      this.popScope();
    }
  }

  visitExpressionStatement(element: ExpressionStatementNode): void {
    // Check variables/function calls within the expression
    element.expression.accept(this); // Visit expression
  }

  // Method to clear state for reuse
  public reset(): void {
    this.errors = [];
    this.scopeStack = [new Set<string>()]; // Re-initialize with global scope (using Set as originally intended)
  }
}
