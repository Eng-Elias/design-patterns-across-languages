import unittest

# Use absolute imports from the 'code_analyzer' package
from code_analyzer.ast_nodes import (
    FunctionDefinitionNode,
    VariableDeclarationNode,
    AssignmentStatementNode,
    IfStatementNode,
    ExpressionStatementNode,
    ExpressionNode
)
from code_analyzer.complexity_visitor import ComplexityVisitor
from code_analyzer.syntax_check_visitor import SyntaxCheckVisitor
from code_analyzer.pretty_print_visitor import PrettyPrintVisitor

class TestCodeAnalyzerVisitors(unittest.TestCase):

    def setUp(self):
        """Set up a common AST for testing different visitors."""
        # AST for:
        # func testFunc(a) {
        #   var b: int = a + 1
        #   if (a > 0) {
        #     b = 10
        #   } else {
        #     b = 20
        #   }
        #   c = b // Use undeclared 'c'
        #   print(b)
        # }
        self.ast = FunctionDefinitionNode(
            name="testFunc",
            parameters=["a"], # Assume param 'a' is declared by function definition
            body=[
                VariableDeclarationNode(name="b", type_hint="int", initializer=ExpressionNode("a + 1")),
                IfStatementNode(
                    condition=ExpressionNode("a > 0"),
                    then_branch=[
                        AssignmentStatementNode(target_variable="b", value=ExpressionNode("10"))
                    ],
                    else_branch=[
                        AssignmentStatementNode(target_variable="b", value=ExpressionNode("20"))
                    ]
                ),
                 AssignmentStatementNode(target_variable="c", value=ExpressionNode("b")), # Use undeclared 'c'
                 ExpressionStatementNode(expression=ExpressionNode("print(b)"))
            ]
        )

        self.ast_ok = FunctionDefinitionNode(
            name="okFunc",
            parameters=["p"],
            body=[
                VariableDeclarationNode(name="v", type_hint="bool", initializer=ExpressionNode("p == 0")),
                IfStatementNode(
                     condition=ExpressionNode("p > 10"),
                     then_branch=[AssignmentStatementNode(target_variable="v", value=ExpressionNode("True"))]
                ),
                AssignmentStatementNode(target_variable="v", value=ExpressionNode("False"))
            ]
        )

        self.ast_undeclared = FunctionDefinitionNode(
            name="undeclaredFunc",
            parameters=[],
            body=[
                AssignmentStatementNode(target_variable="z", value=ExpressionNode("10"))
            ]
        )

        self.ast_redeclare = FunctionDefinitionNode(
            name="redeclareFunc",
            parameters=[],
            body=[
                VariableDeclarationNode(name="a", type_hint="int"),
                # Redeclaration in the same scope (function body)
                VariableDeclarationNode(name="a", type_hint="string")
            ]
        )

    def test_pretty_print_visitor(self):
        """Test the output format of the PrettyPrintVisitor."""
        printer = PrettyPrintVisitor()
        self.ast.accept(printer)
        output = printer.get_output()
        print(f"--- Pretty Print Test Output ---\n{output}") # Print for visual check

        # Basic checks for structure - exact indentation might be fragile to test strictly
        self.assertIn("Function: testFunc(a)", output)
        self.assertIn("Declare: b: int = a + 1", output)
        self.assertIn("If (a > 0):", output)
        self.assertIn("Assign: b = 10", output)
        self.assertIn("Else:", output)
        self.assertIn("Assign: b = 20", output)
        self.assertIn("Assign: c = b", output)
        self.assertIn("ExprStmt: print(b)", output)
        # Check indentation roughly
        self.assertGreater(output.find("Declare: b"), output.find("Function: testFunc"))
        self.assertGreater(output.find("Assign: b = 10"), output.find("If (a > 0):"))
        self.assertGreater(output.find("Assign: b = 20"), output.find("Else:"))


    def test_complexity_visitor(self):
        """Test the complexity calculation."""
        visitor = ComplexityVisitor()
        self.ast.accept(visitor)
        # Base complexity 1 + 1 for the IfStatement = 2
        self.assertEqual(visitor.get_complexity(), 2, "Complexity should be 2")

    def test_syntax_check_visitor_ok(self):
        """Test syntax check on valid code."""
        visitor = SyntaxCheckVisitor()
        visitor.reset()
        # Manage scope externally for the root function node
        visitor.enter_function_scope(self.ast_ok.name, self.ast_ok.parameters)
        self.ast_ok.accept(visitor) # Visitor traverses body within scope
        visitor.exit_function_scope()

        errors = visitor.get_errors()
        print(f"--- Syntax OK Test Errors ---\n{errors}")
        self.assertEqual(len(errors), 0, "Should be no errors for valid AST")

    def test_syntax_check_visitor_undeclared(self):
        """Test detection of using an undeclared variable."""
        visitor = SyntaxCheckVisitor()
        visitor.reset()
        # Manage scope externally
        visitor.enter_function_scope(self.ast_undeclared.name, self.ast_undeclared.parameters)
        self.ast_undeclared.accept(visitor) # Visitor traverses body within scope
        visitor.exit_function_scope()

        errors = visitor.get_errors()
        print(f"--- Syntax Undeclared Test Errors ---\n{errors}")
        self.assertEqual(len(errors), 1, f"Expected 1 error, got {len(errors)}: {errors}")
        if errors:
            self.assertIn("'z' used before declaration", errors[0])

    def test_syntax_check_visitor_redeclaration(self):
        """Test detection of redeclaring a variable in the same scope."""
        visitor = SyntaxCheckVisitor()
        visitor.reset()
        # Manage scope externally
        visitor.enter_function_scope(self.ast_redeclare.name, self.ast_redeclare.parameters)
        self.ast_redeclare.accept(visitor) # Visitor traverses body within scope
        visitor.exit_function_scope()

        errors = visitor.get_errors()
        print(f"--- Syntax Redeclaration Test Errors ---\n{errors}")
        self.assertEqual(len(errors), 1, f"Expected 1 error for redeclaration, got {len(errors)}: {errors}")
        if errors:
            # Update the expected error string to match the actual output
            self.assertIn("Identifier 'a' already declared", errors[0])


if __name__ == '__main__':
    # This allows running the tests directly
    # To run from parent dir: python -m code_analyzer.test_code_analyzer
    unittest.main()
