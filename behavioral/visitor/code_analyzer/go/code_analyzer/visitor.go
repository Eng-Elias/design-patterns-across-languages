package code_analyzer

// Visitor interface declares visit methods for each concrete element type.
// Using interface{} for element allows different node types, but requires type assertion in implementations.
type Visitor interface {
	VisitFunctionDefinition(element *FunctionDefinitionNode)
	VisitVariableDeclaration(element *VariableDeclarationNode)
	VisitAssignmentStatement(element *AssignmentStatementNode)
	VisitIfStatement(element *IfStatementNode)
	VisitExpressionStatement(element *ExpressionStatementNode)
}
