from abc import ABC, abstractmethod

class Visitor(ABC):
    """Visitor interface declaring visit methods for each concrete element type."""
    @abstractmethod
    def visit_function_definition(self, element):
        pass

    @abstractmethod
    def visit_variable_declaration(self, element):
        pass

    @abstractmethod
    def visit_assignment_statement(self, element):
        pass

    @abstractmethod
    def visit_if_statement(self, element):
        pass

    @abstractmethod
    def visit_expression_statement(self, element):
        pass
