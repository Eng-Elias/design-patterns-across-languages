import {
  EqualsExpression,
  GreaterThanExpression,
  LessThanExpression,
  AndExpression,
  OrExpression,
  NotExpression,
  QueryParser,
  QueryEngine
} from './query_language';

describe('Expression Tests', () => {
  const context = { name: "John", age: 30, department: "Engineering", active: true };

  test('EqualsExpression should evaluate equality correctly', () => {
    const expr1 = new EqualsExpression("name", "John");
    expect(expr1.interpret(context)).toBe(true);
    
    const expr2 = new EqualsExpression("name", "Jane");
    expect(expr2.interpret(context)).toBe(false);
  });

  test('GreaterThanExpression should evaluate greater than correctly', () => {
    const expr1 = new GreaterThanExpression("age", 25);
    expect(expr1.interpret(context)).toBe(true);
    
    const expr2 = new GreaterThanExpression("age", 30);
    expect(expr2.interpret(context)).toBe(false);
  });

  test('LessThanExpression should evaluate less than correctly', () => {
    const expr1 = new LessThanExpression("age", 35);
    expect(expr1.interpret(context)).toBe(true);
    
    const expr2 = new LessThanExpression("age", 30);
    expect(expr2.interpret(context)).toBe(false);
  });

  test('AndExpression should combine expressions with AND logic', () => {
    const left1 = new EqualsExpression("name", "John");
    const right1 = new GreaterThanExpression("age", 25);
    const expr1 = new AndExpression(left1, right1);
    expect(expr1.interpret(context)).toBe(true);
    
    const left2 = new EqualsExpression("name", "John");
    const right2 = new GreaterThanExpression("age", 35);
    const expr2 = new AndExpression(left2, right2);
    expect(expr2.interpret(context)).toBe(false);
  });

  test('OrExpression should combine expressions with OR logic', () => {
    const left1 = new EqualsExpression("name", "Jane");
    const right1 = new GreaterThanExpression("age", 25);
    const expr1 = new OrExpression(left1, right1);
    expect(expr1.interpret(context)).toBe(true);
    
    const left2 = new EqualsExpression("name", "Jane");
    const right2 = new GreaterThanExpression("age", 35);
    const expr2 = new OrExpression(left2, right2);
    expect(expr2.interpret(context)).toBe(false);
  });

  test('NotExpression should negate expression result', () => {
    const expr1 = new NotExpression(new EqualsExpression("name", "Jane"));
    expect(expr1.interpret(context)).toBe(true);
    
    const expr2 = new NotExpression(new EqualsExpression("name", "John"));
    expect(expr2.interpret(context)).toBe(false);
  });
});

describe('QueryParser Tests', () => {
  const parser = new QueryParser();
  const context = { name: "John", age: 30, department: "Engineering", active: true };

  test('should parse equals expressions', () => {
    const expr = parser.parse("name = John");
    expect(expr.interpret(context)).toBe(true);
  });

  test('should parse greater than expressions', () => {
    const expr = parser.parse("age > 25");
    expect(expr.interpret(context)).toBe(true);
  });

  test('should parse less than expressions', () => {
    const expr = parser.parse("age < 35");
    expect(expr.interpret(context)).toBe(true);
  });

  test('should parse AND expressions', () => {
    const expr = parser.parse("name = John AND age > 25");
    expect(expr.interpret(context)).toBe(true);
  });

  test('should parse OR expressions', () => {
    const expr = parser.parse("name = Jane OR age > 25");
    expect(expr.interpret(context)).toBe(true);
  });

  test('should parse NOT expressions', () => {
    const expr = parser.parse("NOT name = Jane");
    expect(expr.interpret(context)).toBe(true);
  });

  test('should parse expressions with parentheses', () => {
    const expr = parser.parse("(name = John AND age > 25) OR department = HR");
    expect(expr.interpret(context)).toBe(true);
  });
});

describe('QueryEngine Tests', () => {
  const engine = new QueryEngine();
  const data = [
    { name: "John", age: 30, department: "Engineering" },
    { name: "Jane", age: 25, department: "Marketing" },
    { name: "Bob", age: 35, department: "Engineering" },
    { name: "Alice", age: 28, department: "HR" }
  ];

  test('should filter data correctly with simple expressions', () => {
    // Test filtering for engineers
    const result1 = engine.filter(data, "department = Engineering");
    expect(result1.length).toBe(2);
    expect(result1[0].name).toBe("John");
    expect(result1[1].name).toBe("Bob");
  });

  test('should filter data correctly with comparison expressions', () => {
    // Test filtering for people over 30
    const result = engine.filter(data, "age > 30");
    expect(result.length).toBe(1);
    expect(result[0].name).toBe("Bob");
  });

  test('should filter data correctly with AND expressions', () => {
    // Test complex filtering with AND
    const result = engine.filter(data, "department = Engineering AND age > 30");
    expect(result.length).toBe(1);
    expect(result[0].name).toBe("Bob");
  });

  test('should filter data correctly with OR expressions', () => {
    // Test complex filtering with OR
    const result = engine.filter(data, "department = HR OR department = Marketing");
    expect(result.length).toBe(2);
    expect(result[0].name).toBe("Jane");
    expect(result[1].name).toBe("Alice");
  });
});