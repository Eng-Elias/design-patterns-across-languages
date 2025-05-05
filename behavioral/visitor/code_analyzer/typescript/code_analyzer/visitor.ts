import {
  FunctionDefinitionNode,
  VariableDeclarationNode,
  AssignmentStatementNode,
  IfStatementNode,
  ExpressionStatementNode,
} from "./ast_nodes"; // Import concrete node types

// Visitor interface declaring visit methods for each concrete element type.
export interface Visitor {
  visitFunctionDefinition(element: FunctionDefinitionNode): void;
  visitVariableDeclaration(element: VariableDeclarationNode): void;
  visitAssignmentStatement(element: AssignmentStatementNode): void;
  visitIfStatement(element: IfStatementNode): void;
  visitExpressionStatement(element: ExpressionStatementNode): void;
}
