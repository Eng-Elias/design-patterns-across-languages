from .visitor import Visitor
from .ast_nodes import (
    FunctionDefinitionNode,
    VariableDeclarationNode,
    AssignmentStatementNode,
    IfStatementNode,
    ExpressionStatementNode,
    ExpressionNode, # Needed for assignment RHS check (though simplified here)
)
from typing import List, Set

class SyntaxCheckVisitor(Visitor):
    """Performs simple syntax checks, like tracking declared variables."""
    def __init__(self):
        self._errors: List[str] = []
        # Stack of sets to manage variable scopes
        self._scope_stack: List[Set[str]] = [set()] # Start with global scope

    def get_errors(self) -> List[str]:
        return self._errors

    def _add_error(self, message: str):
        self._errors.append(message)

    def _push_scope(self):
        self._scope_stack.append(set())

    def _pop_scope(self):
        if len(self._scope_stack) > 1: # Avoid popping the global scope
            self._scope_stack.pop()

    def _declare_variable(self, name: str, node_info="variable"):
        current_scope = self._scope_stack[-1]
        if name in current_scope:
            self._add_error(f"Syntax Error: Identifier '{name}' already declared as a {node_info} in this scope.")
        else:
            current_scope.add(name)

    def _is_declared(self, name: str) -> bool:
        # Check from current scope outwards
        for scope in reversed(self._scope_stack):
            if name in scope:
                return True
        return False

    # --- Visitor Methods ---

    def visit_function_definition(self, element: FunctionDefinitionNode):
        # Scope entry/exit should be handled externally by the caller
        # around the accept() call for the FunctionDefinitionNode.
        # This method now focuses only on traversing the body.
        for stmt in element.body:
             stmt.accept(self)
        # No scope pop here - done externally by exit_function_scope

    def visit_variable_declaration(self, element: VariableDeclarationNode):
        # Declare in the current scope
        self._declare_variable(element.name)
        # Check initializer if present (e.g., ensure variables in expression are declared)
        if element.initializer:
             # Simplified: We don't deeply check expression variables here
             pass
        # No need to push/pop scope here

    def visit_assignment_statement(self, element: AssignmentStatementNode):
        # Check if the target variable exists
        if not self._is_declared(element.target_variable):
            self._add_error(f"Syntax Error: Variable '{element.target_variable}' used before declaration.")
        # Check RHS expression variables if needed (simplified)
        pass

    def visit_if_statement(self, element: IfStatementNode):
        # Check condition variables (simplified)
        # Visit the 'then' branch within its own scope
        self.enter_block_scope()
        for stmt in element.then_branch:
            stmt.accept(self) # Visit children within the 'then' scope
        self.exit_block_scope()

        if element.else_branch:
            # Visit the 'else' branch within its own scope
            self.enter_block_scope()
            for stmt in element.else_branch:
                stmt.accept(self) # Visit children within the 'else' scope
            self.exit_block_scope()

    def visit_expression_statement(self, element: ExpressionStatementNode):
        # Check variables/function calls within the expression (simplified)
        pass

    # --- Scope Management Helpers ---
    def enter_function_scope(self, function_name: str, parameters: list[str]):
        # Declare function name in the current (likely global) scope
        self._declare_variable(function_name, "function")
        # Push a new scope for the function body
        self._push_scope()
        # Declare parameters in the new function scope
        for param in parameters:
            self._declare_variable(param, "parameter")

    def exit_function_scope(self):
        self._pop_scope()

    def enter_block_scope(self):
        self._push_scope()

    def exit_block_scope(self):
        self._pop_scope()

    def reset(self):
        """Resets the visitor's state (errors and scope stack)."""
        self._errors = []
        self._scope_stack = [set()] # Initialize with global scope

    def get_errors(self):
        return self._errors
