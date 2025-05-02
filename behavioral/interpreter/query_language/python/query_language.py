from abc import ABC, abstractmethod
from typing import Dict, List, Any


class Expression(ABC):
    """Abstract base class for all expressions in our query language."""
    
    @abstractmethod
    def interpret(self, context: Dict[str, Any]) -> bool:
        """Interpret the expression against the given context."""
        pass


class LiteralExpression(Expression):
    """Terminal expression representing a literal value."""
    
    def __init__(self, value: Any):
        self.value = value
    
    def interpret(self, context: Dict[str, Any]) -> bool:
        # Literals simply evaluate to themselves, but they're rarely used directly
        # This could be used for constant expressions like "TRUE" or "FALSE"
        return bool(self.value)


class VariableExpression(Expression):
    """Terminal expression representing a variable in the context."""
    
    def __init__(self, name: str):
        self.name = name
    
    def interpret(self, context: Dict[str, Any]) -> bool:
        # Variables evaluate to their value in the context
        # This is primarily used internally by other expressions
        return bool(context.get(self.name, False))
    
    def get_value(self, context: Dict[str, Any]) -> Any:
        """Get the actual value from the context."""
        return context.get(self.name)


class EqualsExpression(Expression):
    """Non-terminal expression for equality comparison."""
    
    def __init__(self, variable: str, value: Any):
        self.variable = variable
        self.value = value
    
    def interpret(self, context: Dict[str, Any]) -> bool:
        """Check if the variable equals the value, attempting type coercion."""
        context_value = context.get(self.variable)
        
        # Handle None comparison explicitly
        if context_value is None and self.value is None:
            return True
        if context_value is None or self.value is None:
            return False # One is None, the other isn't
        
        # Attempt type coercion for comparison
        try:
            # If context is numeric, try converting query value to that type
            if isinstance(context_value, (int, float)):
                if isinstance(self.value, (int, float)):
                    return context_value == self.value
                else:
                    return context_value == type(context_value)(self.value)
            # If query value is numeric, try converting context value to that type
            elif isinstance(self.value, (int, float)):
                 return type(self.value)(context_value) == self.value
            # Handle boolean comparison (case-insensitive string check)
            elif isinstance(context_value, bool):
                if isinstance(self.value, bool):
                    return context_value == self.value
                elif isinstance(self.value, str):
                    return context_value == (self.value.lower() == 'true')
            # Add more specific type comparisons if needed
        except (ValueError, TypeError):
            # Fallback to string comparison if coercion fails or types are incompatible
            pass 
            
        # Default/fallback: direct comparison (or string comparison)
        return str(context_value) == str(self.value)


class GreaterThanExpression(Expression):
    """Non-terminal expression for greater than comparison."""
    
    def __init__(self, variable: str, value: Any):
        self.variable = variable
        self.value = value
    
    def interpret(self, context: Dict[str, Any]) -> bool:
        """Check if the variable is greater than the value, attempting numeric conversion."""
        context_value = context.get(self.variable)
        
        # Comparison only makes sense for numeric types usually
        try:
            # Ensure both can be treated as numbers (float for generality)
            num_context_value = float(context_value) 
            num_value = float(self.value)
            return num_context_value > num_value
        except (ValueError, TypeError, TypeError):
            return False # Cannot compare if not convertible to numbers


class LessThanExpression(Expression):
    """Non-terminal expression for less than comparison."""
    
    def __init__(self, variable: str, value: Any):
        self.variable = variable
        self.value = value
    
    def interpret(self, context: Dict[str, Any]) -> bool:
        """Check if the variable is less than the value, attempting numeric conversion."""
        context_value = context.get(self.variable)
        
        # Comparison only makes sense for numeric types usually
        try:
            # Ensure both can be treated as numbers (float for generality)
            num_context_value = float(context_value)
            num_value = float(self.value)
            return num_context_value < num_value
        except (ValueError, TypeError):
            return False # Cannot compare if not convertible to numbers


class AndExpression(Expression):
    """Non-terminal expression for logical AND."""
    
    def __init__(self, left: Expression, right: Expression):
        self.left = left
        self.right = right
    
    def interpret(self, context: Dict[str, Any]) -> bool:
        # Both expressions must be true
        return self.left.interpret(context) and self.right.interpret(context)


class OrExpression(Expression):
    """Non-terminal expression for logical OR."""
    
    def __init__(self, left: Expression, right: Expression):
        self.left = left
        self.right = right
    
    def interpret(self, context: Dict[str, Any]) -> bool:
        # At least one expression must be true
        return self.left.interpret(context) or self.right.interpret(context)


class NotExpression(Expression):
    """Non-terminal expression for logical NOT."""
    
    def __init__(self, expression: Expression):
        self.expression = expression
    
    def interpret(self, context: Dict[str, Any]) -> bool:
        # Negate the result of the expression
        return not self.expression.interpret(context)


class QueryParser:
    """Parser for the query language, converting strings to expression trees."""
    
    def _find_split_index(self, query: str, operator: str) -> int:
        """Finds the index of the operator outside parentheses."""
        level = 0
        # Iterate checking for operator match at the current level (0)
        op_len = len(operator)
        for i in range(len(query) - op_len + 1):
            if query[i] == '(': level += 1
            elif query[i] == ')': level -= 1
            elif level == 0 and query[i:i+op_len] == operator:
                return i
        return -1

    def _try_convert_value(self, value_str: str) -> Any:
        """Tries to convert string value to int, float, bool, or keeps as string."""
        value_str = value_str.strip()
        # Remove potential quotes
        if (value_str.startswith('"') and value_str.endswith('"')) or \
           (value_str.startswith("'") and value_str.endswith("'")):
            return value_str[1:-1]
        # Check boolean
        if value_str.lower() == 'true': return True
        if value_str.lower() == 'false': return False
        # Check numeric
        try: return int(value_str)
        except ValueError:
            try: return float(value_str)
            except ValueError: return value_str # Keep as string

    def parse(self, query: str) -> Expression:
        """Parse a query string into an expression tree respecting precedence and parentheses."""
        query = query.strip()
        if not query: 
            raise ValueError("Cannot parse empty query")

        # 1. Handle parentheses if the entire query is enclosed
        if query.startswith('(') and query.endswith(')'):
            level = 0
            match = True
            # Check if the parentheses are matching outer ones
            for i, char in enumerate(query): 
                if char == '(': level += 1
                elif char == ')': level -= 1
                if level == 0 and i < len(query) - 1: # Closed before the end
                    match = False
                    break
            # Ensure level is 0 at the end and matching
            if match and level == 0: 
                return self.parse(query[1:-1])
            elif level != 0:
                 raise ValueError(f"Mismatched parentheses in query: {query}")
            # If not matching outer parentheses, proceed

        # 2. Handle OR (lowest precedence)
        or_index = self._find_split_index(query, ' OR ')
        if or_index != -1:
            left = self.parse(query[:or_index])
            right = self.parse(query[or_index + 4:]) # len(' OR ') == 4
            return OrExpression(left, right)

        # 3. Handle AND (next precedence)
        and_index = self._find_split_index(query, ' AND ')
        if and_index != -1:
            left = self.parse(query[:and_index])
            right = self.parse(query[and_index + 5:]) # len(' AND ') == 5
            return AndExpression(left, right)

        # 4. Handle NOT (prefix operator)
        if query.startswith('NOT '):
            expression = self.parse(query[4:])
            return NotExpression(expression)

        # 5. Handle comparison expressions (higher precedence)
        # Use find_split_index to avoid splitting inside potential nested function calls if added later
        comp_ops = [' = ', ' > ', ' < '] # Order matters if extending (e.g., >=)
        for op in comp_ops:
            op_index = self._find_split_index(query, op)
            if op_index != -1:
                variable = query[:op_index].strip()
                value_str = query[op_index + len(op):].strip()
                value = self._try_convert_value(value_str)
                
                if op == ' = ': return EqualsExpression(variable, value)
                if op == ' > ': return GreaterThanExpression(variable, value)
                if op == ' < ': return LessThanExpression(variable, value)
        
        # 6. Handle simple variable/literal (treat as variable for now)
        # Could add logic to detect quoted literals if needed
        # Check if it's a known boolean literal string just in case
        if query.lower() == 'true': return LiteralExpression(True)
        if query.lower() == 'false': return LiteralExpression(False)
        
        # Assume it's a variable name if not parsed otherwise
        # We could validate variable names here if needed (e.g., regex)
        return VariableExpression(query)


class QueryEngine:
    """Engine to run queries against collections of data."""
    
    def __init__(self):
        self.parser = QueryParser()
    
    def filter(self, data: List[Dict[str, Any]], query: str) -> List[Dict[str, Any]]:
        """Filter a list of data items using a query expression."""
        expression = self.parser.parse(query)
        return [item for item in data if expression.interpret(item)]
