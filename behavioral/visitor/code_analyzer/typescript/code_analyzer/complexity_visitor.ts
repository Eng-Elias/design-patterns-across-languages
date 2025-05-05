import { Visitor } from "./visitor";
import {
  FunctionDefinitionNode,
  VariableDeclarationNode,
  AssignmentStatementNode,
  IfStatementNode,
  ExpressionStatementNode,
} from "./ast_nodes";

// Calculates a simple cyclomatic complexity score (counting If statements).
export class ComplexityVisitor implements Visitor {
  private complexityScore: number = 0; // Base complexity starts at 0

  getComplexity(): number {
    return this.complexityScore;
  }

  reset(): void {
    // Allow resetting for analyzing multiple functions independently
    this.complexityScore = 0; // Reset to 0
  }

  visitFunctionDefinition(element: FunctionDefinitionNode): void {
    this.complexityScore += 1; // Count function definition
    // Explicitly traverse the function body
    element.body.forEach(stmt => stmt.accept(this));
  }

  visitVariableDeclaration(element: VariableDeclarationNode): void {
    // Declarations don't add complexity
    // Still visit initializer if needed for nested complexity (e.g., lambda functions)
    element.initializer?.accept(this);
  }

  visitAssignmentStatement(element: AssignmentStatementNode): void {
    // Assignments don't add complexity
    // Visit value if needed
    element.value.accept(this);
  }

  visitIfStatement(element: IfStatementNode): void {
    this.complexityScore += 1; // Each 'if' adds a decision point/path
    // Explicitly traverse the condition and branches
    element.condition.accept(this);
    element.thenBranch.forEach(stmt => stmt.accept(this));
    element.elseBranch?.forEach(stmt => stmt.accept(this)); // Visit else branch if exists
  }

  visitExpressionStatement(element: ExpressionStatementNode): void {
    // Simple expressions don't add complexity in this model
    // Visit expression if needed
    element.expression.accept(this);
  }
}
