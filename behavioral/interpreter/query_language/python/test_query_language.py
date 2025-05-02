import unittest
from query_language import (
    EqualsExpression, 
    GreaterThanExpression, 
    LessThanExpression,
    AndExpression, 
    OrExpression, 
    NotExpression,
    QueryParser,
    QueryEngine
)


class TestExpressions(unittest.TestCase):
    def setUp(self):
        self.context = {"name": "John", "age": 30, "department": "Engineering", "active": True}
    
    def test_equals_expression(self):
        expr = EqualsExpression("name", "John")
        self.assertTrue(expr.interpret(self.context))
        
        expr = EqualsExpression("name", "Jane")
        self.assertFalse(expr.interpret(self.context))
    
    def test_greater_than_expression(self):
        expr = GreaterThanExpression("age", 25)
        self.assertTrue(expr.interpret(self.context))
        
        expr = GreaterThanExpression("age", 30)
        self.assertFalse(expr.interpret(self.context))
    
    def test_less_than_expression(self):
        expr = LessThanExpression("age", 35)
        self.assertTrue(expr.interpret(self.context))
        
        expr = LessThanExpression("age", 30)
        self.assertFalse(expr.interpret(self.context))
    
    def test_and_expression(self):
        left = EqualsExpression("name", "John")
        right = GreaterThanExpression("age", 25)
        expr = AndExpression(left, right)
        self.assertTrue(expr.interpret(self.context))
        
        left = EqualsExpression("name", "John")
        right = GreaterThanExpression("age", 35)
        expr = AndExpression(left, right)
        self.assertFalse(expr.interpret(self.context))
    
    def test_or_expression(self):
        left = EqualsExpression("name", "Jane")
        right = GreaterThanExpression("age", 25)
        expr = OrExpression(left, right)
        self.assertTrue(expr.interpret(self.context))
        
        left = EqualsExpression("name", "Jane")
        right = GreaterThanExpression("age", 35)
        expr = OrExpression(left, right)
        self.assertFalse(expr.interpret(self.context))
    
    def test_not_expression(self):
        expr = NotExpression(EqualsExpression("name", "Jane"))
        self.assertTrue(expr.interpret(self.context))
        
        expr = NotExpression(EqualsExpression("name", "John"))
        self.assertFalse(expr.interpret(self.context))


class TestQueryParser(unittest.TestCase):
    def setUp(self):
        self.parser = QueryParser()
        self.context = {"name": "John", "age": 30, "department": "Engineering", "active": True}
    
    def test_parse_equals(self):
        expr = self.parser.parse("name = John")
        self.assertTrue(expr.interpret(self.context))
    
    def test_parse_greater_than(self):
        expr = self.parser.parse("age > 25")
        self.assertTrue(expr.interpret(self.context))
    
    def test_parse_less_than(self):
        expr = self.parser.parse("age < 35")
        self.assertTrue(expr.interpret(self.context))
    
    def test_parse_and(self):
        expr = self.parser.parse("name = John AND age > 25")
        self.assertTrue(expr.interpret(self.context))
    
    def test_parse_or(self):
        expr = self.parser.parse("name = Jane OR age > 25")
        self.assertTrue(expr.interpret(self.context))
    
    def test_parse_not(self):
        expr = self.parser.parse("NOT name = Jane")
        self.assertTrue(expr.interpret(self.context))
    
    def test_parse_parentheses(self):
        expr = self.parser.parse("(name = John AND age > 25) OR department = HR")
        self.assertTrue(expr.interpret(self.context))


class TestQueryEngine(unittest.TestCase):
    def setUp(self):
        self.engine = QueryEngine()
        self.data = [
            {"name": "John", "age": 30, "department": "Engineering"},
            {"name": "Jane", "age": 25, "department": "Marketing"},
            {"name": "Bob", "age": 35, "department": "Engineering"},
            {"name": "Alice", "age": 28, "department": "HR"}
        ]
    
    def test_filter(self):
        # Test filtering for engineers
        result = self.engine.filter(self.data, "department = Engineering")
        self.assertEqual(len(result), 2)
        self.assertEqual(result[0]["name"], "John")
        self.assertEqual(result[1]["name"], "Bob")
        
        # Test filtering for people over 30
        result = self.engine.filter(self.data, "age > 30")
        self.assertEqual(len(result), 1)
        self.assertEqual(result[0]["name"], "Bob")
        
        # Test complex filtering with AND
        result = self.engine.filter(self.data, "department = Engineering AND age > 30")
        self.assertEqual(len(result), 1)
        self.assertEqual(result[0]["name"], "Bob")
        
        # Test complex filtering with OR
        result = self.engine.filter(self.data, "department = HR OR department = Marketing")
        self.assertEqual(len(result), 2)
        self.assertEqual(result[0]["name"], "Jane")
        self.assertEqual(result[1]["name"], "Alice")


if __name__ == "__main__":
    unittest.main()
