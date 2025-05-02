package query_language

import (
	"reflect"
	"testing"
)

// Shared context for expression tests
var testContext = Context{
	"name":       "John",
	"age":        30,
	"department": "Engineering",
	"active":     true,
}

// Shared data for engine tests
var testData = []map[string]interface{}{
	{"name": "John", "age": 30, "department": "Engineering"},
	{"name": "Jane", "age": 25, "department": "Marketing"},
	{"name": "Bob", "age": 35, "department": "Engineering"},
	{"name": "Alice", "age": 28, "department": "HR"},
}

func TestExpressions(t *testing.T) {
	t.Run("EqualsExpression", func(t *testing.T) {
		exprTrue := EqualsExpression{Variable: "name", Value: "John"}
		if !exprTrue.Interpret(testContext) {
			t.Errorf("Expected 'name = John' to be true, got false")
		}

		exprFalse := EqualsExpression{Variable: "name", Value: "Jane"}
		if exprFalse.Interpret(testContext) {
			t.Errorf("Expected 'name = Jane' to be false, got true")
		}

		// Test type conversion (int vs string)
		exprIntStr := EqualsExpression{Variable: "age", Value: "30"}
		if !exprIntStr.Interpret(testContext) {
			t.Errorf("Expected 'age = \"30\"' to be true, got false")
		}
	})

	t.Run("GreaterThanExpression", func(t *testing.T) {
		exprTrue := GreaterThanExpression{Variable: "age", Value: 25}
		if !exprTrue.Interpret(testContext) {
			t.Errorf("Expected 'age > 25' to be true, got false")
		}

		exprFalse := GreaterThanExpression{Variable: "age", Value: 30}
		if exprFalse.Interpret(testContext) {
			t.Errorf("Expected 'age > 30' to be false, got true")
		}
		
		// Test type conversion (int vs string)
		exprIntStr := GreaterThanExpression{Variable: "age", Value: "25"} 
		if !exprIntStr.Interpret(testContext) {
			t.Errorf("Expected 'age > \"25\"' to be true, got false")
		}
	})

	t.Run("LessThanExpression", func(t *testing.T) {
		exprTrue := LessThanExpression{Variable: "age", Value: 35}
		if !exprTrue.Interpret(testContext) {
			t.Errorf("Expected 'age < 35' to be true, got false")
		}

		exprFalse := LessThanExpression{Variable: "age", Value: 30}
		if exprFalse.Interpret(testContext) {
			t.Errorf("Expected 'age < 30' to be false, got true")
		}

		// Test type conversion (int vs string)
		exprIntStr := LessThanExpression{Variable: "age", Value: "35"}
		if !exprIntStr.Interpret(testContext) {
			t.Errorf("Expected 'age < \"35\"' to be true, got false")
		}
	})

	t.Run("AndExpression", func(t *testing.T) {
		left := EqualsExpression{Variable: "name", Value: "John"}
		right := GreaterThanExpression{Variable: "age", Value: 25}
		exprTrue := AndExpression{Left: &left, Right: &right}
		if !exprTrue.Interpret(testContext) {
			t.Errorf("Expected 'name = John AND age > 25' to be true, got false")
		}

		rightFalse := GreaterThanExpression{Variable: "age", Value: 35}
		exprFalse := AndExpression{Left: &left, Right: &rightFalse}
		if exprFalse.Interpret(testContext) {
			t.Errorf("Expected 'name = John AND age > 35' to be false, got true")
		}
	})

	t.Run("OrExpression", func(t *testing.T) {
		leftFalse := EqualsExpression{Variable: "name", Value: "Jane"}
		rightTrue := GreaterThanExpression{Variable: "age", Value: 25}
		exprTrue := OrExpression{Left: &leftFalse, Right: &rightTrue}
		if !exprTrue.Interpret(testContext) {
			t.Errorf("Expected 'name = Jane OR age > 25' to be true, got false")
		}

		rightFalse := GreaterThanExpression{Variable: "age", Value: 35}
		exprFalse := OrExpression{Left: &leftFalse, Right: &rightFalse}
		if exprFalse.Interpret(testContext) {
			t.Errorf("Expected 'name = Jane OR age > 35' to be false, got true")
		}
	})

	t.Run("NotExpression", func(t *testing.T) {
		innerFalse := EqualsExpression{Variable: "name", Value: "Jane"}
		exprTrue := NotExpression{Expression: &innerFalse}
		if !exprTrue.Interpret(testContext) {
			t.Errorf("Expected 'NOT name = Jane' to be true, got false")
		}

		innerTrue := EqualsExpression{Variable: "name", Value: "John"}
		exprFalse := NotExpression{Expression: &innerTrue}
		if exprFalse.Interpret(testContext) {
			t.Errorf("Expected 'NOT name = John' to be false, got true")
		}
	})
}

func TestQueryParser(t *testing.T) {
	parser := QueryParser{}

	tests := []struct {
		query string
		expected bool
	}{
		{"name = John", true},
		{"name = Jane", false},
		{"age > 25", true},
		{"age > 30", false},
		{"age < 35", true},
		{"age < 30", false},
		{"name = John AND age > 25", true},
		{"name = John AND age > 30", false},
		{"name = Jane OR age > 25", true},
		{"name = Jane OR age > 30", false},
		{"NOT name = Jane", true},
		{"NOT name = John", false},
		{"(name = John AND age > 25) OR department = HR", true},
		{"(name = Jane AND age > 30) OR department = Marketing", false},
		{"active = true", true}, // Test boolean value
		{"department = Engineering", true}, // Test simple equality
	}

	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			expr := parser.Parse(tt.query)
			result := expr.Interpret(testContext)
			if result != tt.expected {
				t.Errorf("Query '%s': expected %v, got %v", tt.query, tt.expected, result)
			}
		})
	}
}

func TestQueryEngine(t *testing.T) {
	engine := NewQueryEngine()

	tests := []struct {
		name          string
		query         string
		expectedNames []string // Just check names for simplicity
	}{
		{
			name:          "FilterEngineers",
			query:         "department = Engineering",
			expectedNames: []string{"John", "Bob"},
		},
		{
			name:          "FilterAgeOver30",
			query:         "age > 30",
			expectedNames: []string{"Bob"},
		},
		{
			name:          "FilterComplexAND",
			query:         "department = Engineering AND age > 30",
			expectedNames: []string{"Bob"},
		},
		{
			name:          "FilterComplexOR",
			query:         "department = HR OR department = Marketing",
			expectedNames: []string{"Jane", "Alice"},
		},
		{
			name:          "FilterComplexParentheses",
			query:         "(department = Engineering AND age < 35) OR (department = Marketing AND age < 30)",
			expectedNames: []string{"John", "Jane"},
		},
		{
			name:          "FilterNoResults",
			query:         "age > 40",
			expectedNames: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := engine.Filter(testData, tt.query)
			resultNames := make([]string, len(result))
			for i, item := range result {
				resultNames[i] = item["name"].(string)
			}

			if !reflect.DeepEqual(resultNames, tt.expectedNames) {
				t.Errorf("Query '%s': expected names %v, got %v", tt.query, tt.expectedNames, resultNames)
			}
		})
	}
}