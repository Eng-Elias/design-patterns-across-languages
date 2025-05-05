from .visitor import Visitor
from .ast_nodes import FunctionDefinitionNode, VariableDeclarationNode, AssignmentStatementNode, IfStatementNode, ExpressionStatementNode, ExpressionNode, Node

class PrettyPrintVisitor(Visitor):
    """Prints a textual representation of the AST with indentation, controlling traversal."""
    def __init__(self, indent_char="  "):
        self._output = ""
        self._indent_char = indent_char
        self._indent_level = 0

    def _indent(self):
        return self._indent_char * self._indent_level

    def reset(self):
        self._output = ""
        self._indent_level = 0

    def get_output(self):
        return self._output.strip() # Remove any trailing newline

    def visit_function_definition(self, element: FunctionDefinitionNode):
        params = ", ".join(element.parameters)
        self._output += f"{self._indent()}Function: {element.name}({params})\n"
        self._indent_level += 1
        for stmt in element.body:
            stmt.accept(self) # Visitor explicitly traverses children
        self._indent_level -= 1

    def visit_variable_declaration(self, element: VariableDeclarationNode):
        init_str = ""
        if element.initializer:
            # If initializer is complex, it might need its own accept call
            # For simple ExpressionNode, just accessing representation is ok
             init_str = f" = {element.initializer.representation}"
            # If initializer could be another statement/complex expr:
            # element.initializer.accept(self) # Would need careful output structuring
        type_hint = f": {element.type_hint}" if element.type_hint else ""
        self._output += f"{self._indent()}Declare: {element.name}{type_hint}{init_str}\n"

    def visit_assignment_statement(self, element: AssignmentStatementNode):
        # Similar to declaration, handle value access/visit
        value_repr = element.value.representation
        # If value could be complex:
        # element.value.accept(self)
        self._output += f"{self._indent()}Assign: {element.target_variable} = {value_repr}\n"

    def visit_if_statement(self, element: IfStatementNode):
        # Visit condition (if needed, e.g., if it's not just a simple string)
        # element.condition.accept(self)
        self._output += f"{self._indent()}If ({element.condition.representation}):\n"
        self._indent_level += 1
        for stmt in element.then_branch:
            stmt.accept(self) # Visitor explicitly traverses children
        self._indent_level -= 1

        if element.else_branch:
            self._output += f"{self._indent()}Else:\n"
            self._indent_level += 1
            for stmt in element.else_branch:
                stmt.accept(self) # Visitor explicitly traverses children
            self._indent_level -= 1

    def visit_expression_statement(self, element: ExpressionStatementNode):
        # Visit expression
        # element.expression.accept(self)
        self._output += f"{self._indent()}ExprStmt: {element.expression.representation}\n"
