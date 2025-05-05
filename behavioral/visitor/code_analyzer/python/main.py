# Import from the code_analyzer package
from code_analyzer.ast_nodes import (
    FunctionDefinitionNode,
    VariableDeclarationNode,
    AssignmentStatementNode,
    IfStatementNode,
    ExpressionStatementNode,
    ExpressionNode,
    Node # Base node type if needed
)
from code_analyzer.complexity_visitor import ComplexityVisitor
from code_analyzer.syntax_check_visitor import SyntaxCheckVisitor
from code_analyzer.pretty_print_visitor import PrettyPrintVisitor
# No need to import Visitor interface itself for the demo

# --- Client Code: Demo ---
if __name__ == "__main__":
    print("--- Visitor Pattern: Code Analysis Demo (Python) ---")

    # 1. Construct a simple AST (Object Structure)
    # Simulating a function like:
    # func calculate(x: int, y: int) {
    #   var result: int = 0
    #   var temp: int // Declaration without init
    #   if (x > y) {
    #     result = x
    #     temp = 1
    #   } else {
    #     result = y
    #     temp = 0
    #   }
    #   print(result) // Simulated by ExpressionStatement
    #   # Undeclared variable usage for SyntaxCheckVisitor demo
    #   # z = temp
    # }
    ast_root: Node = FunctionDefinitionNode(
        name="calculate",
        parameters=["x", "y"],
        body=[
            VariableDeclarationNode(name="result", type_hint="int", initializer=ExpressionNode("0")),
            VariableDeclarationNode(name="temp", type_hint="int"), # No initializer
            IfStatementNode(
                condition=ExpressionNode("x > y"),
                then_branch=[
                    AssignmentStatementNode(target_variable="result", value=ExpressionNode("x")),
                    AssignmentStatementNode(target_variable="temp", value=ExpressionNode("1")),
                ],
                else_branch=[
                    AssignmentStatementNode(target_variable="result", value=ExpressionNode("y")),
                    AssignmentStatementNode(target_variable="temp", value=ExpressionNode("0")),
                ]
            ),
            ExpressionStatementNode(expression=ExpressionNode("print(result)")),
            # Uncomment below to introduce a syntax error for the demo
            # AssignmentStatementNode(target_variable="z", value=ExpressionNode("temp"))
        ]
    )

    # 2. Create Visitors
    pretty_printer = PrettyPrintVisitor()
    complexity_checker = ComplexityVisitor()
    syntax_checker = SyntaxCheckVisitor()

    # --- Apply Visitors ---

    print("\n--- Applying PrettyPrintVisitor ---")
    # Reset visitor state if it were stateful (PrettyPrint usually isn't, but good practice)
    # pretty_printer.reset() # Assuming a reset method if needed
    ast_root.accept(pretty_printer)
    print(pretty_printer.get_output())

    print("\n--- Applying ComplexityVisitor ---")
    complexity_checker.reset() # Reset state before use
    ast_root.accept(complexity_checker)
    print(f"Calculated Complexity: {complexity_checker.get_complexity()}")

    print("\n--- Applying SyntaxCheckVisitor ---")
    syntax_checker.reset() # Reset state before use

    # Manually manage scope entry/exit around the accept call for the root function node
    if isinstance(ast_root, FunctionDefinitionNode):
        syntax_checker.enter_function_scope(ast_root.name, ast_root.parameters)
        # Explicitly iterate and Accept body statements *within* the scope
        for stmt in ast_root.body:
            stmt.accept(syntax_checker)
        syntax_checker.exit_function_scope()
    else:
        # Fallback if root isn't a function
        ast_root.accept(syntax_checker)

    errors = syntax_checker.get_errors()
    if errors:
        print("Syntax Errors Found:")
        for error in errors:
            print(f"- {error}")
    else:
        print("No syntax errors found (based on current AST).")
        print("(Uncomment the 'z = temp' line in main.py to see an error)")


    print("\n--- Demo Finished ---")
