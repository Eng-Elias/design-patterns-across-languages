package code_analyzer

import (
	"strings"
	"testing"
)

// Helper function to create the common AST used in tests
func createTestAST() *FunctionDefinitionNode {
	return &FunctionDefinitionNode{
		Name:       "testFunc",
		Parameters: []string{"a"},
		Body: []StatementNode{
			&VariableDeclarationNode{Name: "b", TypeHint: "int", Initializer: &ExpressionNode{Representation: "a + 1"}},
			&IfStatementNode{
				Condition: &ExpressionNode{Representation: "a > 0"},
				ThenBranch: []StatementNode{
					&AssignmentStatementNode{TargetVariable: "b", Value: &ExpressionNode{Representation: "10"}},
				},
				ElseBranch: []StatementNode{
					&AssignmentStatementNode{TargetVariable: "b", Value: &ExpressionNode{Representation: "20"}},
				},
			},
			&AssignmentStatementNode{TargetVariable: "c", Value: &ExpressionNode{Representation: "b"}}, // Use undeclared 'c'
			&ExpressionStatementNode{Expression: &ExpressionNode{Representation: "fmt.Println(b)"}}, // Simulate print
		},
	}
}

func TestPrettyPrintVisitor(t *testing.T) {
	ast := createTestAST()
	printer := NewPrettyPrintVisitor("  ")

	ast.Accept(printer) // Use standard accept
	output := printer.GetOutput()
	t.Logf("--- Pretty Print Test Output ---\n%s", output) // Log for visual check

	// Use Contains for basic checks due to potential indent variations
	if !strings.Contains(output, "Function: testFunc(a)") {
		t.Errorf("Expected output to contain 'Function: testFunc(a)'")
	}
	if !strings.Contains(output, "Declare: b: int = a + 1") {
		t.Errorf("Expected output to contain 'Declare: b: int = a + 1'")
	}
	if !strings.Contains(output, "If (a > 0):") {
		t.Errorf("Expected output to contain 'If (a > 0):'")
	}
	if !strings.Contains(output, "Assign: b = 10") {
		t.Errorf("Expected output to contain 'Assign: b = 10'")
	}
	if !strings.Contains(output, "Else:") {
		t.Errorf("Expected output to contain 'Else:'")
	}
	if !strings.Contains(output, "Assign: b = 20") {
		t.Errorf("Expected output to contain 'Assign: b = 20'")
	}
	if !strings.Contains(output, "Assign: c = b") {
        t.Errorf("Expected output to contain 'Assign: c = b'")
    }
	if !strings.Contains(output, "ExprStmt: fmt.Println(b)") {
		t.Errorf("Expected output to contain 'ExprStmt: fmt.Println(b)'")
	}
}

func TestComplexityVisitor(t *testing.T) {
	ast := createTestAST()
	visitor := NewComplexityVisitor()
	ast.Accept(visitor)

	expectedComplexity := 2 // Base 1 + 1 for IfStatement
	if visitor.GetComplexity() != expectedComplexity {
		t.Errorf("Expected complexity %d, got %d", expectedComplexity, visitor.GetComplexity())
	}
}

func TestSyntaxCheckVisitor_Undeclared(t *testing.T) {
	visitor := NewSyntaxCheckVisitor()
	ast := createTestAST()

	ast.Accept(visitor)

	errors := visitor.GetErrors()
	t.Logf("--- Syntax Undeclared Test Errors ---\n%v", errors)

	if len(errors) != 1 {
		t.Fatalf("Expected 1 error, got %d: %v", len(errors), errors)
	}
	if !strings.Contains(errors[0], "Variable 'c' used before declaration") {
		t.Errorf("Expected error message for undeclared 'c', got: %s", errors[0])
	}
}

func TestSyntaxCheckVisitor_Redeclaration(t *testing.T) {
	astRedeclare := &FunctionDefinitionNode{
		Name:       "redeclareFunc",
		Parameters: []string{},
		Body: []StatementNode{
			&VariableDeclarationNode{Name: "x", TypeHint: "int"},
			&VariableDeclarationNode{Name: "x", TypeHint: "string"}, // Redeclaration
		},
	}
	visitor := NewSyntaxCheckVisitor()

	astRedeclare.Accept(visitor)

	errors := visitor.GetErrors()
	t.Logf("--- Syntax Redeclaration Test Errors ---\n%v", errors)

	if len(errors) == 0 {
        t.Fatalf("Expected at least 1 error for redeclaration, got 0")
    }

    found := false
    for _, err := range errors {
        if strings.Contains(err, "Identifier 'x' already declared") {
            found = true
            break
        }
    }
    if !found {
		 t.Errorf("Expected error message for redeclaring 'x', got: %v", errors)
	}
}

func TestSyntaxCheckVisitor_OK(t *testing.T) {
    astOK := &FunctionDefinitionNode{
        Name:       "okFunc",
        Parameters: []string{"p"},
        Body: []StatementNode{
            &VariableDeclarationNode{Name: "v", TypeHint: "bool", Initializer: &ExpressionNode{Representation: "p == 0"}},
            &IfStatementNode{
                 Condition: &ExpressionNode{Representation: "p > 10"},
                 ThenBranch: []StatementNode{
                     &AssignmentStatementNode{TargetVariable: "v", Value: &ExpressionNode{Representation: "true"}},
                 },
                 // No Else branch
            },
            &AssignmentStatementNode{TargetVariable: "v", Value: &ExpressionNode{Representation: "false"}},
        },
    }
    visitor := NewSyntaxCheckVisitor()

    astOK.Accept(visitor)

	errors := visitor.GetErrors()
	t.Logf("--- Syntax OK Test Errors ---\n%v", errors)

	if len(errors) != 0 {
		t.Errorf("Expected 0 errors for correct code, got %d: %v", len(errors), errors)
	}
}
