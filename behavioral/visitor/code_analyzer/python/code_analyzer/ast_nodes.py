from abc import ABC, abstractmethod
from .visitor import Visitor
from typing import List, Optional

# --- Element Interface ---
class Node(ABC):
    """Element interface that declares an accept method."""
    @abstractmethod
    def accept(self, visitor: Visitor):
        pass

# --- Concrete Elements (AST Nodes) ---
class ExpressionNode(Node):
    """Represents a simple code expression (e.g., literal, variable, binary op)."""
    def __init__(self, representation: str):
        self.representation = representation # Simplified representation for the demo

    def accept(self, visitor: Visitor):
        if hasattr(visitor, 'visit_expression'):
            visitor.visit_expression(self)


class StatementNode(Node):
    """Base class for statement nodes."""
    pass


class FunctionDefinitionNode(Node):
    """Represents a function definition."""
    def __init__(self, name: str, parameters: List[str], body: List[StatementNode]):
        self.name = name
        self.parameters = parameters
        self.body = body

    def accept(self, visitor: Visitor):
        visitor.visit_function_definition(self)


class VariableDeclarationNode(StatementNode):
    """Represents a variable declaration (e.g., var x: int = 5)."""
    def __init__(self, name: str, type_hint: str, initializer: Optional[ExpressionNode] = None):
        self.name = name
        self.type_hint = type_hint
        self.initializer = initializer

    def accept(self, visitor: Visitor):
        visitor.visit_variable_declaration(self)


class AssignmentStatementNode(StatementNode):
    """Represents an assignment (e.g., x = 10)."""
    def __init__(self, target_variable: str, value: ExpressionNode):
        self.target_variable = target_variable
        self.value = value

    def accept(self, visitor: Visitor):
        visitor.visit_assignment_statement(self)


class IfStatementNode(StatementNode):
    """Represents an if(-else) statement."""
    def __init__(self, condition: ExpressionNode, then_branch: List[StatementNode], else_branch: Optional[List[StatementNode]] = None):
        self.condition = condition
        self.then_branch = then_branch
        self.else_branch = else_branch

    def accept(self, visitor: Visitor):
        visitor.visit_if_statement(self)


class ExpressionStatementNode(StatementNode):
    """Represents a statement that consists of just an expression (e.g., a function call)."""
    def __init__(self, expression: ExpressionNode):
        self.expression = expression

    def accept(self, visitor: Visitor):
        visitor.visit_expression_statement(self)
