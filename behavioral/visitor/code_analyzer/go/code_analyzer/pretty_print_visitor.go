package code_analyzer

import (
	"fmt"
	"strings"
)

// PrettyPrintVisitor prints the AST in an indented format.
type PrettyPrintVisitor struct {
	indentLevel int
	output      strings.Builder // Efficiently build the output string
	indentChar  string
}

// NewPrettyPrintVisitor creates a new PrettyPrintVisitor.
func NewPrettyPrintVisitor(indentChar string) *PrettyPrintVisitor {
	if indentChar == "" {
		indentChar = "  " // Default indent
	}
	return &PrettyPrintVisitor{
		indentChar: indentChar,
	}
}

func (ppv *PrettyPrintVisitor) indent() string {
	return strings.Repeat(ppv.indentChar, ppv.indentLevel)
}

// GetOutput returns the formatted string.
func (ppv *PrettyPrintVisitor) GetOutput() string {
	return ppv.output.String()
}

// Reset clears the output and resets indentation.
func (ppv *PrettyPrintVisitor) Reset() {
	ppv.output.Reset()
	ppv.indentLevel = 0
}

func (ppv *PrettyPrintVisitor) VisitFunctionDefinition(element *FunctionDefinitionNode) {
	ppv.output.WriteString(fmt.Sprintf("%sFunction: %s(%s)\n", ppv.indent(), element.Name, strings.Join(element.Parameters, ", ")))
	ppv.indentLevel++
	// Explicitly traverse body within correct indent level
	for _, stmt := range element.Body {
		stmt.Accept(ppv)
	}
	ppv.indentLevel-- // Decrease indent after traversing body
}

func (ppv *PrettyPrintVisitor) VisitVariableDeclaration(element *VariableDeclarationNode) {
	initStr := ""
	if element.Initializer != nil {
		initStr = fmt.Sprintf(" = %s", element.Initializer.Representation)
	}
	ppv.output.WriteString(fmt.Sprintf("%sDeclare: %s: %s%s\n", ppv.indent(), element.Name, element.TypeHint, initStr))
}

func (ppv *PrettyPrintVisitor) VisitAssignmentStatement(element *AssignmentStatementNode) {
	valStr := "nil"
	if element.Value != nil {
		valStr = element.Value.Representation
	}
	ppv.output.WriteString(fmt.Sprintf("%sAssign: %s = %s\n", ppv.indent(), element.TargetVariable, valStr))
}

func (ppv *PrettyPrintVisitor) VisitIfStatement(element *IfStatementNode) {
	condStr := "<no condition>" // Default if nil
	if element.Condition != nil {
		condStr = element.Condition.Representation
	}
	ppv.output.WriteString(fmt.Sprintf("%sIf (%s):\n", ppv.indent(), condStr))

	ppv.indentLevel++
	// Explicitly traverse 'then' branch
	for _, stmt := range element.ThenBranch {
		stmt.Accept(ppv)
	}
	ppv.indentLevel--

	if len(element.ElseBranch) > 0 {
		ppv.output.WriteString(fmt.Sprintf("%sElse:\n", ppv.indent()))
		ppv.indentLevel++
		// Explicitly traverse 'else' branch
		for _, stmt := range element.ElseBranch {
			stmt.Accept(ppv)
		}
		ppv.indentLevel--
	}
}

func (ppv *PrettyPrintVisitor) VisitExpressionStatement(element *ExpressionStatementNode) {
	exprStr := "nil"
	if element.Expression != nil {
		exprStr = element.Expression.Representation
	}
	ppv.output.WriteString(fmt.Sprintf("%sExprStmt: %s\n", ppv.indent(), exprStr))
}

// --- Manual Indent Helpers (if needed for external control) ---
func (ppv *PrettyPrintVisitor) IncreaseIndent() {
    ppv.indentLevel++;
}
func (ppv *PrettyPrintVisitor) DecreaseIndent() {
    if ppv.indentLevel > 0 {
        ppv.indentLevel--;
    }
}
