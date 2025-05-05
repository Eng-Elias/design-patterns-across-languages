import { Visitor } from "./visitor";

// --- Element Interface ---
export interface Node {
  // Element interface that declares an accept method.
  accept(visitor: Visitor): void;
}

// --- Concrete Elements (AST Nodes) ---

// Represents a simple code expression (e.g., literal, variable, binary op).
export class ExpressionNode implements Node {
  // Simplified representation for the demo
  constructor(public representation: string) {}

  accept(visitor: Visitor): void {
    // If visitors need to specifically handle expressions, they might implement visitExpression
    // visitor.visitExpression?.(this); // Use optional chaining if method is optional
    // Keep simple: traversal usually happens from statements containing expressions
  }
}

// Base class/interface for statement nodes. Using an abstract class for potential common methods.
export abstract class StatementNode implements Node {
  abstract accept(visitor: Visitor): void;
}

// Represents a function definition.
export class FunctionDefinitionNode implements Node {
  constructor(
    public name: string,
    public parameters: string[],
    public body: StatementNode[]
  ) {}

  accept(visitor: Visitor): void {
    visitor.visitFunctionDefinition(this);
    // Traversal is now the responsibility of the visitor's visit method.
  }
}

// Represents a variable declaration (e.g., var x: int = 5).
export class VariableDeclarationNode extends StatementNode {
  constructor(
    public name: string,
    public typeHint: string,
    public initializer?: ExpressionNode
  ) {
    super();
  }

  accept(visitor: Visitor): void {
    visitor.visitVariableDeclaration(this);
    this.initializer?.accept(visitor); // Visit initializer if it exists
  }
}

// Represents an assignment (e.g., x = 10).
export class AssignmentStatementNode extends StatementNode {
  constructor(public targetVariable: string, public value: ExpressionNode) {
    super();
  }

  accept(visitor: Visitor): void {
    visitor.visitAssignmentStatement(this);
    this.value.accept(visitor);
  }
}

// Represents an if(-else) statement.
export class IfStatementNode extends StatementNode {
  constructor(
    public condition: ExpressionNode,
    public thenBranch: StatementNode[],
    public elseBranch?: StatementNode[] // Optional else branch
  ) {
    super();
  }

  accept(visitor: Visitor): void {
    visitor.visitIfStatement(this);
    // Traversal is now the responsibility of the visitor's visit method.
  }
}

// Represents a statement that consists of just an expression (e.g., a function call).
export class ExpressionStatementNode extends StatementNode {
  constructor(public expression: ExpressionNode) {
    super();
  }

  accept(visitor: Visitor): void {
    visitor.visitExpressionStatement(this);
    this.expression.accept(visitor);
  }
}
