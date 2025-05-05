package code_analyzer

// ComplexityVisitor calculates a simple cyclomatic complexity score.
type ComplexityVisitor struct {
	complexityScore int
}

// NewComplexityVisitor creates a new ComplexityVisitor.
func NewComplexityVisitor() *ComplexityVisitor {
	return &ComplexityVisitor{complexityScore: 0} // Start with base complexity 0
}

// GetComplexity returns the calculated score.
func (cv *ComplexityVisitor) GetComplexity() int {
	return cv.complexityScore
}

// Reset sets the score back to 0.
func (cv *ComplexityVisitor) Reset() {
	cv.complexityScore = 0
}

// VisitFunctionDefinition - Increments score and traverses body.
func (cv *ComplexityVisitor) VisitFunctionDefinition(element *FunctionDefinitionNode) {
	cv.complexityScore++ // Count function definition
	// Explicitly traverse body
	for _, stmt := range element.Body {
		stmt.Accept(cv)
	}
}

// VisitVariableDeclaration - No complexity added.
func (cv *ComplexityVisitor) VisitVariableDeclaration(element *VariableDeclarationNode) {
	// Traversal happens via element.Accept()
}

// VisitAssignmentStatement - No complexity added.
func (cv *ComplexityVisitor) VisitAssignmentStatement(element *AssignmentStatementNode) {
	// Traversal happens via element.Accept()
}

// VisitIfStatement - Increments complexity score.
func (cv *ComplexityVisitor) VisitIfStatement(element *IfStatementNode) {
	cv.complexityScore++ // Each 'if' adds a decision point
	// Explicitly traverse branches
	for _, stmt := range element.ThenBranch {
		stmt.Accept(cv)
	}
	if len(element.ElseBranch) > 0 {
		for _, stmt := range element.ElseBranch {
			stmt.Accept(cv)
		}
	}
}

// VisitExpressionStatement - No complexity added.
func (cv *ComplexityVisitor) VisitExpressionStatement(element *ExpressionStatementNode) {
	// Traversal happens via element.Accept()
}
