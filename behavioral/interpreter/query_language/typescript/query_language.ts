/**
 * Interface for all expression types in the query language
 */
export interface Expression {
  interpret(context: Record<string, any>): boolean;
}

/**
 * Terminal expression representing a literal value
 */
export class LiteralExpression implements Expression {
  private value: any;

  constructor(value: any) {
    this.value = value;
  }

  interpret(_context: Record<string, any>): boolean {
    return Boolean(this.value);
  }
}

/**
 * Terminal expression representing a variable in the context
 */
export class VariableExpression implements Expression {
  private name: string;

  constructor(name: string) {
    this.name = name;
  }

  interpret(context: Record<string, any>): boolean {
    return Boolean(context[this.name]);
  }

  getValue(context: Record<string, any>): any {
    return context[this.name];
  }
}

/**
 * Non-terminal expression for equality comparison
 */
export class EqualsExpression implements Expression {
  private variable: string;
  private value: any;

  constructor(variable: string, value: any) {
    this.variable = variable;
    this.value = value;
  }

  interpret(context: Record<string, any>): boolean {
    const contextValue = context[this.variable];
    
    // Handle null/undefined comparison explicitly
    if (contextValue == null && this.value == null) {
      return true;
    }
    if (contextValue == null || this.value == null) {
      return false; // One is null/undefined, the other isn't
    }

    // Attempt type coercion for comparison
    try {
      // If context is numeric, try converting query value to that type
      if (typeof contextValue === 'number') {
        const numValue = Number(this.value);
        return !isNaN(numValue) && contextValue === numValue;
      } 
      // If query value is numeric, try converting context value to that type
      else if (typeof this.value === 'number') {
          const numContextValue = Number(contextValue);
          return !isNaN(numContextValue) && numContextValue === this.value;
      }
      // Handle boolean comparison (case-insensitive string check for query value)
      else if (typeof contextValue === 'boolean') {
        if (typeof this.value === 'boolean') {
          return contextValue === this.value;
        } else if (typeof this.value === 'string') {
          return contextValue === (this.value.toLowerCase() === 'true');
        }
      }
      // Add more specific type comparisons if needed
    } catch (e) {
      // Fallback to string comparison if coercion fails or types are incompatible
    }

    // Default/fallback: string comparison
    return String(contextValue) === String(this.value);
  }
}

/**
 * Non-terminal expression for greater than comparison
 */
export class GreaterThanExpression implements Expression {
  private variable: string;
  private value: any;

  constructor(variable: string, value: any) {
    this.variable = variable;
    this.value = value;
  }

  interpret(context: Record<string, any>): boolean {
    const contextValue = context[this.variable];
    
    // Comparison usually makes sense only for numeric types
    try {
      const numContextValue = Number(contextValue);
      const numValue = Number(this.value);
      // Check if both conversions were successful
      if (!isNaN(numContextValue) && !isNaN(numValue)) {
          return numContextValue > numValue;
      }
    } catch (e) { 
      // Ignore conversion errors
    }
    return false; // Cannot compare if not convertible to numbers
  }
}

/**
 * Non-terminal expression for less than comparison
 */
export class LessThanExpression implements Expression {
  private variable: string;
  private value: any;

  constructor(variable: string, value: any) {
    this.variable = variable;
    this.value = value;
  }

  interpret(context: Record<string, any>): boolean {
    const contextValue = context[this.variable];

    // Comparison usually makes sense only for numeric types
    try {
      const numContextValue = Number(contextValue);
      const numValue = Number(this.value);
      // Check if both conversions were successful
      if (!isNaN(numContextValue) && !isNaN(numValue)) {
          return numContextValue < numValue;
      }
    } catch (e) {
      // Ignore conversion errors
    }
    return false; // Cannot compare if not convertible to numbers
  }
}

/**
 * Non-terminal expression for logical AND
 */
export class AndExpression implements Expression {
  private left: Expression;
  private right: Expression;

  constructor(left: Expression, right: Expression) {
    this.left = left;
    this.right = right;
  }

  interpret(context: Record<string, any>): boolean {
    return this.left.interpret(context) && this.right.interpret(context);
  }
}

/**
 * Non-terminal expression for logical OR
 */
export class OrExpression implements Expression {
  private left: Expression;
  private right: Expression;

  constructor(left: Expression, right: Expression) {
    this.left = left;
    this.right = right;
  }

  interpret(context: Record<string, any>): boolean {
    return this.left.interpret(context) || this.right.interpret(context);
  }
}

/**
 * Non-terminal expression for logical NOT
 */
export class NotExpression implements Expression {
  private expression: Expression;

  constructor(expression: Expression) {
    this.expression = expression;
  }

  interpret(context: Record<string, any>): boolean {
    return !this.expression.interpret(context);
  }
}

/**
 * Parser for the query language, converting strings to expression trees
 */
export class QueryParser {

  private findSplitIndex(query: string, operator: string): number {
    let level = 0;
    const opLen = operator.length;
    for (let i = 0; i <= query.length - opLen; i++) {
      if (query[i] === '(') {
        level++;
      } else if (query[i] === ')') {
        level--;
      } else if (level === 0 && query.substring(i, i + opLen) === operator) {
        return i;
      }
    }
    return -1;
  }

  private tryConvertValue(valueStr: string): any {
    valueStr = valueStr.trim();
    // Remove potential quotes
    if ((valueStr.startsWith('"') && valueStr.endsWith('"')) ||
        (valueStr.startsWith("'") && valueStr.endsWith("'"))) {
      valueStr = valueStr.substring(1, valueStr.length - 1);
    }
    // Check boolean
    if (valueStr.toLowerCase() === 'true') return true;
    if (valueStr.toLowerCase() === 'false') return false;
    // Check numeric
    const numVal = Number(valueStr);
    if (!isNaN(numVal)) return numVal;
    
    return valueStr; // Keep as string
  }

  parse(query: string): Expression {
    query = query.trim();
    if (!query) {
      // Handle empty query - returning Literal(false) or throwing error
      return new LiteralExpression(false); // Placeholder
    }

    // 1. Handle parentheses if the entire query is enclosed
    if (query.startsWith('(') && query.endsWith(')')) {
        let level = 0;
        let match = true;
        // Check if the parentheses are matching outer ones
        for (let i = 0; i < query.length; i++) {
            const char = query[i];
            if (char === '(') level++;
            else if (char === ')') level--;
            if (level === 0 && i < query.length - 1) { // Closed before the end
                match = false;
                break;
            }
        }
        // Ensure level is 0 at the end and matching
        if (match && level === 0) {
            return this.parse(query.substring(1, query.length - 1));
        } else if (level !== 0) {
             // Handle mismatched parentheses - returning Literal(false) or throwing error
             return new LiteralExpression(false); // Placeholder
        }
        // If not matching outer parentheses, proceed
    }

    // 2. Handle OR (lowest precedence)
    const orIndex = this.findSplitIndex(query, ' OR ');
    if (orIndex !== -1) {
        const left = this.parse(query.substring(0, orIndex));
        const right = this.parse(query.substring(orIndex + 4)); // length of ' OR '
        return new OrExpression(left, right);
    }

    // 3. Handle AND (next precedence)
    const andIndex = this.findSplitIndex(query, ' AND ');
    if (andIndex !== -1) {
        const left = this.parse(query.substring(0, andIndex));
        const right = this.parse(query.substring(andIndex + 5)); // length of ' AND '
        return new AndExpression(left, right);
    }

    // 4. Handle NOT (prefix operator)
    if (query.startsWith('NOT ')) {
      const expression = this.parse(query.substring(4));
      return new NotExpression(expression);
    }

    // 5. Handle comparison expressions
    const compOps = [' = ', ' > ', ' < ']; // Order matters if extending
    for (const op of compOps) {
        const opIndex = this.findSplitIndex(query, op);
        if (opIndex !== -1) {
            const variable = query.substring(0, opIndex).trim();
            const valueStr = query.substring(opIndex + op.length).trim();
            const value = this.tryConvertValue(valueStr);

            switch (op) {
                case ' = ': return new EqualsExpression(variable, value);
                case ' > ': return new GreaterThanExpression(variable, value);
                case ' < ': return new LessThanExpression(variable, value);
            }
        }
    }

    // 6. Handle simple variable/literal
    // Check if it's a known boolean literal string just in case
    if (query.toLowerCase() === 'true') return new LiteralExpression(true);
    if (query.toLowerCase() === 'false') return new LiteralExpression(false);

    // Assume it's a variable name if not parsed otherwise
    return new VariableExpression(query);
  }
}

/**
 * Engine to run queries against collections of data
 */
export class QueryEngine {
  private parser: QueryParser;

  constructor() {
    this.parser = new QueryParser();
  }

  filter(data: Record<string, any>[], query: string): Record<string, any>[] {
    const expression = this.parser.parse(query);
    return data.filter(item => expression.interpret(item));
  }
}
