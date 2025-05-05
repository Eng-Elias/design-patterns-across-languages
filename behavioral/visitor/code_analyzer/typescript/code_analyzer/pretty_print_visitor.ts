import { Visitor } from "./visitor";
import {
  FunctionDefinitionNode,
  VariableDeclarationNode,
  AssignmentStatementNode,
  IfStatementNode,
  ExpressionStatementNode,
} from "./ast_nodes";

// Prints the AST nodes in a structured, indented format.
export class PrettyPrintVisitor implements Visitor {
  private indentLevel: number = 0;
  private output: string = "";
  private readonly indentChar: string;

  constructor(indentChar: string = "  ") {
    this.indentChar = indentChar;
  }

  private indent(): string {
    return this.indentChar.repeat(this.indentLevel);
  }

  getOutput(): string {
    return this.output;
  }

  visitFunctionDefinition(element: FunctionDefinitionNode): void {
    this.output += `${this.indent()}Function: ${element.name}(${element.parameters.join(", ")})\n`;
    this.indentLevel++;
    // Explicitly traverse body with indentation
    element.body.forEach((stmt) => stmt.accept(this));
    this.indentLevel--; // Decrease indent after visiting body
  }

  visitVariableDeclaration(element: VariableDeclarationNode): void {
    const initStr = element.initializer
      ? ` = ${element.initializer.representation}`
      : "";
    this.output += `${this.indent()}Declare: ${element.name}: ${element.typeHint}${initStr}\n`;
    // Initializer traversal via its own accept
    element.initializer?.accept(this);
  }

  visitAssignmentStatement(element: AssignmentStatementNode): void {
    this.output += `${this.indent()}Assign: ${element.targetVariable} = ${element.value.representation}\n`;
    // Value traversal via its own accept
    element.value.accept(this);
  }

  visitIfStatement(element: IfStatementNode): void {
    this.output += `${this.indent()}If (${element.condition.representation}):\n`;
    // Visit condition first (does not usually affect indent)
    // element.condition.accept(this); // Typically not visited by pretty printer unless complex

    this.indentLevel++;
    // Visit 'then' branch
    element.thenBranch.forEach((stmt) => stmt.accept(this));
    this.indentLevel--;

    if (element.elseBranch) {
      this.output += `${this.indent()}Else:\n`;
      this.indentLevel++;
      // Visit 'else' branch
      element.elseBranch.forEach((stmt) => stmt.accept(this));
      this.indentLevel--;
    }
  }

  visitExpressionStatement(element: ExpressionStatementNode): void {
    this.output += `${this.indent()}ExprStmt: ${element.expression.representation}\n`;
    // Expression traversal via its own accept
    element.expression.accept(this);
  }

  // Reset indent level manually if needed between visiting different top-level nodes
  resetIndent(): void {
    this.indentLevel = 0;
  }
}
