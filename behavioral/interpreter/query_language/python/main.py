from query_language import QueryEngine

def main():
    # Sample data - list of employees
    employees = [
        {"name": "John", "age": 34, "department": "Engineering", "salary": 85000},
        {"name": "Sarah", "age": 29, "department": "Marketing", "salary": 72000},
        {"name": "Michael", "age": 41, "department": "Engineering", "salary": 110000},
        {"name": "Emma", "age": 27, "department": "HR", "salary": 65000},
        {"name": "Robert", "age": 36, "department": "Finance", "salary": 95000},
        {"name": "Lisa", "age": 32, "department": "Marketing", "salary": 78000},
        {"name": "David", "age": 45, "department": "Engineering", "salary": 120000},
        {"name": "Jessica", "age": 31, "department": "HR", "salary": 68000},
    ]
    
    # Create query engine
    query_engine = QueryEngine()
    
    # Example 1: Find all engineers over 35
    query1 = "department = Engineering AND age > 35"
    result1 = query_engine.filter(employees, query1)
    
    print("Engineers over 35:")
    for employee in result1:
        print(f"  - {employee['name']}: {employee['age']} years old, {employee['department']}")
    
    # Example 2: Find marketing employees or anyone with salary over 100k
    query2 = "department = Marketing OR salary > 100000"
    result2 = query_engine.filter(employees, query2)
    
    print("\nMarketing employees or high earners:")
    for employee in result2:
        print(f"  - {employee['name']}: {employee['department']}, ${employee['salary']}")
    
    # Example 3: Complex query with parentheses
    query3 = "(department = Engineering OR department = HR) AND age < 35"
    result3 = query_engine.filter(employees, query3)
    
    print("\nYoung engineers or HR employees:")
    for employee in result3:
        print(f"  - {employee['name']}: {employee['age']} years old, {employee['department']}")


if __name__ == "__main__":
    main()
