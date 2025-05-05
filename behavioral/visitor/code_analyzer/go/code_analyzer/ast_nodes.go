package code_analyzer

// --- Element Interface ---
// Node defines the interface that all AST components must implement.
type Node interface {
	Accept(v Visitor)
}

// --- Concrete Elements (AST Nodes) ---

// ExpressionNode represents a simple code expression.
// Using a struct for concrete types.
type ExpressionNode struct {
	Representation string // Simplified representation for the demo
}

// Accept for ExpressionNode (often not directly visited from Visitor interface)
func (e *ExpressionNode) Accept(v Visitor) {
	// If visitor has VisitExpression:
	// v.VisitExpression(e)
}

// StatementNode serves as a marker interface or base for statements (Go doesn't have abstract classes)
// We can use embedding or just have statement types implement Node directly.
// Using an interface allows slices of different statement types.
type StatementNode interface {
	Node // Embed the base Node interface
	isStatementNode() // Dummy method to act like a 'marker'
}

// FunctionDefinitionNode represents a function definition.
type FunctionDefinitionNode struct {
	Name       string
	Parameters []string
	Body       []StatementNode // Slice of StatementNode interfaces
}

func (f *FunctionDefinitionNode) Accept(v Visitor) {
	v.VisitFunctionDefinition(f)
	// Node no longer controls traversal
}

func (f *FunctionDefinitionNode) isStatementNode() {} // Implement marker method (though func def isn't usually a statement *within* a block)

// VariableDeclarationNode represents a variable declaration.
type VariableDeclarationNode struct {
	Name        string
	TypeHint    string
	Initializer *ExpressionNode // Pointer allows for optional initializer (nil)
}

func (vd *VariableDeclarationNode) Accept(v Visitor) {
	v.VisitVariableDeclaration(vd)
	// Node no longer controls traversal
}

func (vd *VariableDeclarationNode) isStatementNode() {} // Implement marker method

// AssignmentStatementNode represents an assignment.
type AssignmentStatementNode struct {
	TargetVariable string
	Value          *ExpressionNode // Use pointer for consistency if needed, or value type if always present
}

func (as *AssignmentStatementNode) Accept(v Visitor) {
	v.VisitAssignmentStatement(as)
	// Node no longer controls traversal
}

func (as *AssignmentStatementNode) isStatementNode() {} // Implement marker method

// IfStatementNode represents an if(-else) statement.
type IfStatementNode struct {
	Condition  *ExpressionNode
	ThenBranch []StatementNode
	ElseBranch []StatementNode // Slice can be empty or nil if no else branch
}

func (is *IfStatementNode) Accept(v Visitor) {
	v.VisitIfStatement(is)
	// Node no longer controls traversal
}

func (is *IfStatementNode) isStatementNode() {} // Implement marker method

// ExpressionStatementNode represents a statement consisting of just an expression.
type ExpressionStatementNode struct {
	Expression *ExpressionNode
}

func (es *ExpressionStatementNode) Accept(v Visitor) {
	v.VisitExpressionStatement(es)
	// Node no longer controls traversal
}

func (es *ExpressionStatementNode) isStatementNode() {} // Implement marker method
