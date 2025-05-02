package query_language

import (
	"strconv"
	"strings"
)

// Context represents the data against which expressions are evaluated
type Context map[string]interface{}

// Expression is the interface for all expressions in the query language
type Expression interface {
	Interpret(context Context) bool
}

// LiteralExpression represents a literal value
type LiteralExpression struct {
	Value interface{}
}

// Interpret returns the boolean value of the literal
func (e *LiteralExpression) Interpret(context Context) bool {
	switch v := e.Value.(type) {
	case bool:
		return v
	case string:
		return v != ""
	case int:
		return v != 0
	default:
		return e.Value != nil
	}
}

// VariableExpression represents a variable in the context
type VariableExpression struct {
	Name string
}

// Interpret returns the boolean value of the variable in the context
func (e *VariableExpression) Interpret(context Context) bool {
	value, exists := context[e.Name]
	if !exists {
		return false
	}

	switch v := value.(type) {
	case bool:
		return v
	case string:
		return v != ""
	case int:
		return v != 0
	default:
		return value != nil
	}
}

// GetValue returns the actual value of the variable from the context
func (e *VariableExpression) GetValue(context Context) interface{} {
	return context[e.Name]
}

// EqualsExpression represents an equality comparison
type EqualsExpression struct {
	Variable string
	Value    interface{}
}

// Interpret checks if the variable equals the value
func (e *EqualsExpression) Interpret(context Context) bool {
	value, exists := context[e.Variable]
	if !exists {
		return false
	}

	switch v := value.(type) {
	case string:
		if strValue, ok := e.Value.(string); ok {
			return v == strValue
		}
	case int:
		if intValue, ok := e.Value.(int); ok {
			return v == intValue
		}
		if strValue, ok := e.Value.(string); ok {
			if intValue, err := strconv.Atoi(strValue); err == nil {
				return v == intValue
			}
		}
	case float64:
		if floatValue, ok := e.Value.(float64); ok {
			return v == floatValue
		}
		if strValue, ok := e.Value.(string); ok {
			if floatValue, err := strconv.ParseFloat(strValue, 64); err == nil {
				return v == floatValue
			}
		}
	case bool:
		if boolValue, ok := e.Value.(bool); ok {
			return v == boolValue
		}
		if strValue, ok := e.Value.(string); ok {
			if strValue == "true" {
				return v == true
			} else if strValue == "false" {
				return v == false
			}
		}
	}

	// Fall back to string comparison
	return strings.EqualFold(toString(value), toString(e.Value))
}

// GreaterThanExpression represents a greater than comparison
type GreaterThanExpression struct {
	Variable string
	Value    interface{}
}

// Interpret checks if the variable is greater than the value
func (e *GreaterThanExpression) Interpret(context Context) bool {
	value, exists := context[e.Variable]
	if !exists {
		return false
	}

	switch v := value.(type) {
	case int:
		if intValue, ok := e.Value.(int); ok {
			return v > intValue
		}
		if strValue, ok := e.Value.(string); ok {
			if intValue, err := strconv.Atoi(strValue); err == nil {
				return v > intValue
			}
		}
	case float64:
		if floatValue, ok := e.Value.(float64); ok {
			return v > floatValue
		}
		if strValue, ok := e.Value.(string); ok {
			if floatValue, err := strconv.ParseFloat(strValue, 64); err == nil {
				return v > floatValue
			}
		}
	}

	// Can't compare for greater than
	return false
}

// LessThanExpression represents a less than comparison
type LessThanExpression struct {
	Variable string
	Value    interface{}
}

// Interpret checks if the variable is less than the value
func (e *LessThanExpression) Interpret(context Context) bool {
	value, exists := context[e.Variable]
	if !exists {
		return false
	}

	switch v := value.(type) {
	case int:
		if intValue, ok := e.Value.(int); ok {
			return v < intValue
		}
		if strValue, ok := e.Value.(string); ok {
			if intValue, err := strconv.Atoi(strValue); err == nil {
				return v < intValue
			}
		}
	case float64:
		if floatValue, ok := e.Value.(float64); ok {
			return v < floatValue
		}
		if strValue, ok := e.Value.(string); ok {
			if floatValue, err := strconv.ParseFloat(strValue, 64); err == nil {
				return v < floatValue
			}
		}
	}

	// Can't compare for less than
	return false
}

// AndExpression represents a logical AND of two expressions
type AndExpression struct {
	Left  Expression
	Right Expression
}

// Interpret returns true if both the left and right expressions are true
func (e *AndExpression) Interpret(context Context) bool {
	return e.Left.Interpret(context) && e.Right.Interpret(context)
}

// OrExpression represents a logical OR of two expressions
type OrExpression struct {
	Left  Expression
	Right Expression
}

// Interpret returns true if either the left or right expression is true
func (e *OrExpression) Interpret(context Context) bool {
	return e.Left.Interpret(context) || e.Right.Interpret(context)
}

// NotExpression represents a logical NOT of an expression
type NotExpression struct {
	Expression Expression
}

// Interpret returns true if the expression is false
func (e *NotExpression) Interpret(context Context) bool {
	return !e.Expression.Interpret(context)
}

// QueryParser converts query strings into expression trees
type QueryParser struct{}

// findSplitIndex finds the index of the operator outside parentheses.
func (p *QueryParser) findSplitIndex(query, operator string) int {
	level := 0
	opLen := len(operator)
	for i := 0; i <= len(query)-opLen; i++ {
		if query[i] == '(' {
			level++
		} else if query[i] == ')' {
			level--
		} else if level == 0 && query[i:i+opLen] == operator {
			return i
		}
	}
	return -1
}

// tryConvertValue attempts to convert string value to bool, int, float64, or keeps as string.
func tryConvertValue(valueStr string) interface{} {
	valueStr = strings.TrimSpace(valueStr)
	// Remove potential quotes
	if (strings.HasPrefix(valueStr, "\"") && strings.HasSuffix(valueStr, "\"")) ||
	   (strings.HasPrefix(valueStr, "'") && strings.HasSuffix(valueStr, "'")) {
		valueStr = valueStr[1 : len(valueStr)-1]
	}
	// Check boolean
	if strings.EqualFold(valueStr, "true") { return true }
	if strings.EqualFold(valueStr, "false") { return false }
	// Check numeric
	if intVal, err := strconv.Atoi(valueStr); err == nil {
		return intVal
	}
	if floatVal, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return floatVal
	}
	return valueStr // Keep as string
}

// Parse parses a query string into an expression tree respecting precedence and parentheses.
func (p *QueryParser) Parse(query string) Expression {
	query = strings.TrimSpace(query)
	if query == "" {
		// Return a placeholder or handle error appropriately
		// For now, returning a literal false might prevent crashes but isn't ideal.
		return &LiteralExpression{Value: false}
	}

	// 1. Handle parentheses if the entire query is enclosed
	if strings.HasPrefix(query, "(") && strings.HasSuffix(query, ")") {
		level := 0
		match := true
		// Check if the parentheses are matching outer ones
		for i, char := range query {
			if char == '(' {
				level++
			} else if char == ')' {
				level--
			}
			if level == 0 && i < len(query)-1 { // Closed before the end
				match = false
				break
			}
		}
		// Ensure level is 0 at the end and matching
		if match && level == 0 {
			return p.Parse(query[1 : len(query)-1])
		} else if level != 0 {
			// Handle mismatched parentheses error, perhaps return LiteralExpression(false)
			return &LiteralExpression{Value: false} // Placeholder for error
		}
		// If not matching outer parentheses, proceed
	}

	// 2. Handle OR (lowest precedence)
	orIndex := p.findSplitIndex(query, " OR ")
	if orIndex != -1 {
		left := p.Parse(query[:orIndex])
		right := p.Parse(query[orIndex+4:]) // len(" OR ") == 4
		return &OrExpression{Left: left, Right: right}
	}

	// 3. Handle AND (next precedence)
	andIndex := p.findSplitIndex(query, " AND ")
	if andIndex != -1 {
		left := p.Parse(query[:andIndex])
		right := p.Parse(query[andIndex+5:]) // len(" AND ") == 5
		return &AndExpression{Left: left, Right: right}
	}

	// 4. Handle NOT (prefix operator)
	if strings.HasPrefix(query, "NOT ") {
		expression := p.Parse(query[4:])
		return &NotExpression{Expression: expression}
	}

	// 5. Handle comparison expressions
	compOps := []string{" = ", " > ", " < "} // Order matters if extending (e.g., >=)
	for _, op := range compOps {
		opIndex := p.findSplitIndex(query, op)
		if opIndex != -1 {
			variable := strings.TrimSpace(query[:opIndex])
			valueStr := strings.TrimSpace(query[opIndex+len(op):])
			value := tryConvertValue(valueStr)

			switch op {
			case " = ":
				return &EqualsExpression{Variable: variable, Value: value}
			case " > ":
				return &GreaterThanExpression{Variable: variable, Value: value}
			case " < ":
				return &LessThanExpression{Variable: variable, Value: value}
			}
		}
	}

	// 6. Handle simple variable/literal
	// Check if it's a known boolean literal string just in case
	if strings.EqualFold(query, "true") { return &LiteralExpression{Value: true} }
	if strings.EqualFold(query, "false") { return &LiteralExpression{Value: false} }

	// Assume it's a variable name if not parsed otherwise
	return &VariableExpression{Name: query}
}

// QueryEngine provides methods to filter data using queries
type QueryEngine struct {
	parser *QueryParser
}

// NewQueryEngine creates a new QueryEngine
func NewQueryEngine() *QueryEngine {
	return &QueryEngine{
		parser: &QueryParser{},
	}
}

// Filter filters a slice of data using a query expression
func (e *QueryEngine) Filter(data []map[string]interface{}, query string) []map[string]interface{} {
	expression := e.parser.Parse(query)
	result := make([]map[string]interface{}, 0)

	for _, item := range data {
		context := Context(item)
		if expression.Interpret(context) {
			result = append(result, item)
		}
	}

	return result
}

// Helper function to convert a value to string
func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch value := v.(type) {
	case string:
		return value
	case int:
		return strconv.Itoa(value)
	case float64:
		return strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		return ""
	}
}
