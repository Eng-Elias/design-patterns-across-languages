package main

import (
	"fmt"

	// Use the module path defined in go.mod
	code_analyzer "visitor_pattern_code_analyzer_go/code_analyzer"
)

func main() {
	fmt.Println("--- Visitor Pattern: Code Analysis Demo (Go) ---")

	// 1. Construct a simple AST (Object Structure) - Same as Python/TS example
	// Declare astRoot as the Node interface type
	var astRoot code_analyzer.Node = &code_analyzer.FunctionDefinitionNode{
		Name:       "calculate",
		Parameters: []string{"x", "y"},
		Body: []code_analyzer.StatementNode{ // Slice of interfaces
			&code_analyzer.VariableDeclarationNode{Name: "result", TypeHint: "int", Initializer: &code_analyzer.ExpressionNode{Representation: "0"}},
			&code_analyzer.VariableDeclarationNode{Name: "temp", TypeHint: "int", Initializer: nil}, // No initializer
			&code_analyzer.IfStatementNode{
				Condition: &code_analyzer.ExpressionNode{Representation: "x > y"},
				ThenBranch: []code_analyzer.StatementNode{
					&code_analyzer.AssignmentStatementNode{TargetVariable: "result", Value: &code_analyzer.ExpressionNode{Representation: "x"}},
					&code_analyzer.AssignmentStatementNode{TargetVariable: "temp", Value: &code_analyzer.ExpressionNode{Representation: "1"}},
				},
				ElseBranch: []code_analyzer.StatementNode{
					&code_analyzer.AssignmentStatementNode{TargetVariable: "result", Value: &code_analyzer.ExpressionNode{Representation: "y"}},
					&code_analyzer.AssignmentStatementNode{TargetVariable: "temp", Value: &code_analyzer.ExpressionNode{Representation: "0"}},
				},
			},
			&code_analyzer.ExpressionStatementNode{Expression: &code_analyzer.ExpressionNode{Representation: "fmt.Println(result)"}}, // Simulate print
			// Uncomment below to introduce a syntax error for the demo
			// &code_analyzer.AssignmentStatementNode{TargetVariable: "z", Value: &code_analyzer.ExpressionNode{Representation: "temp"}},
		},
	}

	// 2. Create Visitors
	prettyPrinter := code_analyzer.NewPrettyPrintVisitor("  ")
	complexityChecker := code_analyzer.NewComplexityVisitor()
	// Syntax checker state needs resetting before use
	syntaxChecker := code_analyzer.NewSyntaxCheckVisitor()

	// --- Apply Visitors ---

	fmt.Println("\n--- Applying PrettyPrintVisitor ---")
	prettyPrinter.Reset() // Reset state before use
	// Manually manage indent for root if needed, or rely on Visit method logic
	astRoot.Accept(prettyPrinter) // Node controls traversal
	fmt.Println(prettyPrinter.GetOutput())
	// Note: Indentation logic primarily resides within the visitor's visit methods

	fmt.Println("\n--- Applying ComplexityVisitor ---")
	complexityChecker.Reset() // Reset state before use
	astRoot.Accept(complexityChecker)
	fmt.Printf("Calculated Complexity: %d\n", complexityChecker.GetComplexity())

	fmt.Println("\n--- Applying SyntaxCheckVisitor ---")
	syntaxChecker.Reset() // Reset state before use

	// Manually manage scope entry/exit around the accept call for the root function node
	if fnNode, ok := astRoot.(*code_analyzer.FunctionDefinitionNode); ok {
		fnNode.Accept(syntaxChecker) // Visit body statement
	} else {
		// Fallback if root isn't a function
		astRoot.Accept(syntaxChecker)
	}

	errors := syntaxChecker.GetErrors()
	if len(errors) > 0 {
		fmt.Println("Syntax Errors Found:")
		for _, err := range errors {
			fmt.Printf("- %s\n", err)
		}
	} else {
		fmt.Println("No syntax errors found (based on current AST).")
		fmt.Println("(Uncomment the 'z = temp' line in main.go to see an error)")
	}

	fmt.Println("\n--- Demo Finished ---")
}
