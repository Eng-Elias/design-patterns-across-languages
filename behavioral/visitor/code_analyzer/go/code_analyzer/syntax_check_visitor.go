package code_analyzer

import "fmt"

// SyntaxCheckVisitor performs simple syntax checks using scopes.
type SyntaxCheckVisitor struct {
	errors     []string
	scopeStack []map[string]string // Slice of maps acting as a stack of scopes [varName] -> "type" (e.g., "variable", "function", "parameter")
}

// NewSyntaxCheckVisitor creates a new SyntaxCheckVisitor.
func NewSyntaxCheckVisitor() *SyntaxCheckVisitor {
	return &SyntaxCheckVisitor{
		errors:     []string{},
		scopeStack: []map[string]string{make(map[string]string)}, // Start with global scope
	}
}

// GetErrors returns the list of found syntax errors.
func (scv *SyntaxCheckVisitor) GetErrors() []string {
	return scv.errors
}

func (scv *SyntaxCheckVisitor) addError(format string, args ...interface{}) {
	scv.errors = append(scv.errors, fmt.Sprintf(format, args...))
}

func (scv *SyntaxCheckVisitor) pushScope() {
	scv.scopeStack = append(scv.scopeStack, make(map[string]string))
}

func (scv *SyntaxCheckVisitor) popScope() {
	if len(scv.scopeStack) > 1 { // Avoid popping the global scope
		scv.scopeStack = scv.scopeStack[:len(scv.scopeStack)-1] // Slice off the last element
	}
}

// declareVariable adds a variable to the current scope.
func (scv *SyntaxCheckVisitor) declareVariable(name string, nodeType string) {
	if len(scv.scopeStack) == 0 {
		scv.addError("Internal Error: Scope stack is empty during declaration.")
		return
	}
	currentScope := scv.scopeStack[len(scv.scopeStack)-1]
	if existingType, exists := currentScope[name]; exists {
		scv.addError("Syntax Error: Identifier '%s' already declared as a %s in this scope.", name, existingType)
	} else {
		currentScope[name] = nodeType
	}
}

// isDeclared checks if a variable exists in any scope, starting from the innermost.
func (scv *SyntaxCheckVisitor) isDeclared(name string) bool {
	for i := len(scv.scopeStack) - 1; i >= 0; i-- {
		if _, exists := scv.scopeStack[i][name]; exists {
			return true
		}
	}
	return false
}

// --- Visitor Methods ---

// VisitFunctionDefinition declares the function, pushes scope, declares params, traverses body, pops scope.
func (scv *SyntaxCheckVisitor) VisitFunctionDefinition(element *FunctionDefinitionNode) {
	scv.declareVariable(element.Name, "function") // Declare func name in outer scope
	scv.pushScope()                              // Push scope for function body
	for _, param := range element.Parameters {
		scv.declareVariable(param, "parameter") // Declare params in new scope
	}
	// Explicitly traverse the body within the function's scope
	for _, stmt := range element.Body {
		stmt.Accept(scv)
	}
	scv.popScope() // Pop function scope *after* traversing body
}

// VisitVariableDeclaration declares the variable in the current scope.
func (scv *SyntaxCheckVisitor) VisitVariableDeclaration(element *VariableDeclarationNode) {
	scv.declareVariable(element.Name, "variable") // Or use element.TypeHint if needed
	// Check initializer expression if necessary (simplified here)
}

// VisitAssignmentStatement checks if the target variable is declared.
func (scv *SyntaxCheckVisitor) VisitAssignmentStatement(element *AssignmentStatementNode) {
	if !scv.isDeclared(element.TargetVariable) {
		scv.addError("Syntax Error: Variable '%s' used before declaration.", element.TargetVariable)
	}
	// Check RHS expression if necessary (simplified here)
}

// VisitIfStatement manages scope and traverses branches.
func (scv *SyntaxCheckVisitor) VisitIfStatement(element *IfStatementNode) {
	// Check condition expression if necessary (simplified here)
	// element.Condition.Accept(scv) // If expressions needed checking

	scv.pushScope() // Push scope for 'then' branch
	// Explicitly traverse the 'then' branch within its scope
	for _, stmt := range element.ThenBranch {
		stmt.Accept(scv)
	}
	scv.popScope() // Pop scope *after* 'then' branch

	if len(element.ElseBranch) > 0 {
		scv.pushScope() // Push scope for 'else' branch
		// Explicitly traverse the 'else' branch within its scope
		for _, stmt := range element.ElseBranch {
			stmt.Accept(scv)
		}
		scv.popScope() // Pop scope *after* 'else' branch
	}
}

// VisitExpressionStatement - Potential checks inside the expression (simplified here).
func (scv *SyntaxCheckVisitor) VisitExpressionStatement(element *ExpressionStatementNode) {
	// Check expression details if needed
}

// Reset clears the visitor's state (errors and scopes).
func (scv *SyntaxCheckVisitor) Reset() {
	scv.errors = []string{}
	scv.scopeStack = []map[string]string{make(map[string]string)} // Initialize with global scope
}
