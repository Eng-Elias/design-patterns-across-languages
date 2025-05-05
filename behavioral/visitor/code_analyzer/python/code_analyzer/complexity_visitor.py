from .visitor import Visitor
from .ast_nodes import ( # Import specific node types
    FunctionDefinitionNode,
    VariableDeclarationNode,
    AssignmentStatementNode,
    IfStatementNode,
    ExpressionStatementNode,
    # ExpressionNode - not directly needed for complexity calc here
)

class ComplexityVisitor(Visitor):
    """Calculates a simple complexity score (e.g., number of decision points)."""
    def __init__(self):
        self._complexity_score = 0 # Start complexity score at 0

    def get_complexity(self) -> int:
        return self._complexity_score

    def reset(self): # Allow resetting for analyzing multiple functions independently
        self._complexity_score = 0

    def visit_function_definition(self, element: FunctionDefinitionNode):
        # Count the function definition itself as 1 unit of complexity
        self._complexity_score += 1
        # Explicitly traverse the body
        for stmt in element.body:
            stmt.accept(self)

    def visit_variable_declaration(self, element: VariableDeclarationNode):
        # Declarations don't add complexity
        pass # Traversal into initializer happens via Node.accept

    def visit_assignment_statement(self, element: AssignmentStatementNode):
        # Assignments don't add complexity
        pass # Traversal into value happens via Node.accept

    def visit_if_statement(self, element: IfStatementNode):
        self._complexity_score += 1 # Increment for the if statement itself
        # Explicitly visit the 'then' branch
        for stmt in element.then_branch:
            stmt.accept(self)
        # Explicitly visit the 'else' branch if it exists
        if element.else_branch:
             # Don't increment complexity again for 'else', just visit contents
            for stmt in element.else_branch:
                stmt.accept(self)

    def visit_expression_statement(self, element: ExpressionStatementNode):
        # Simple expressions don't add complexity in this model
        pass # Traversal into expression happens via Node.accept
