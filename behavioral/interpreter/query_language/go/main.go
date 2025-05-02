package main

import (
	"fmt"
	"interpreter_query_language/query_language" // Use the module path from go.mod
)

func main() {
	// Sample data - list of employees (matches Python example)
	employees := []map[string]interface{}{
		{"name": "John", "age": 34, "department": "Engineering", "salary": 85000},
		{"name": "Sarah", "age": 29, "department": "Marketing", "salary": 72000},
		{"name": "Michael", "age": 41, "department": "Engineering", "salary": 110000},
		{"name": "Emma", "age": 27, "department": "HR", "salary": 65000},
		{"name": "Robert", "age": 36, "department": "Finance", "salary": 95000},
		{"name": "Lisa", "age": 32, "department": "Marketing", "salary": 78000},
		{"name": "David", "age": 45, "department": "Engineering", "salary": 120000},
		{"name": "Jessica", "age": 31, "department": "HR", "salary": 68000},
	}

	// Create a new query engine
	engine := query_language.NewQueryEngine()

	// --- Query 1: Engineers over 35 --- 
	query1 := "department = Engineering AND age > 35"
	result1 := engine.Filter(employees, query1)
	fmt.Println("Engineers over 35:")
	if len(result1) == 0 {
		fmt.Println("  No matching records found.")
	} else {
		for _, employee := range result1 {
			fmt.Printf("  - %s: %d years old, %s\n", employee["name"], employee["age"], employee["department"])
		}
	}

	// --- Query 2: Marketing employees or high earners --- 
	query2 := "department = Marketing OR salary > 100000"
	result2 := engine.Filter(employees, query2)
	fmt.Println("\nMarketing employees or high earners:")
	if len(result2) == 0 {
		fmt.Println("  No matching records found.")
	} else {
		for _, employee := range result2 {
			fmt.Printf("  - %s: %s, $%d\n", employee["name"], employee["department"], employee["salary"])
		}
	}

	// --- Query 3: Young engineers or HR employees --- 
	query3 := "(department = Engineering OR department = HR) AND age < 35"
	result3 := engine.Filter(employees, query3)
	fmt.Println("\nYoung engineers or HR employees:")
	if len(result3) == 0 {
		fmt.Println("  No matching records found.")
	} else {
		for _, employee := range result3 {
			fmt.Printf("  - %s: %d years old, %s\n", employee["name"], employee["age"], employee["department"])
		}
	}
}