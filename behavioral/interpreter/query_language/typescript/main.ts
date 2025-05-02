import { QueryEngine } from './query_language';

function main() {
  // Sample data - list of employees
  const employees = [
    { name: "John", age: 34, department: "Engineering", salary: 85000 },
    { name: "Sarah", age: 29, department: "Marketing", salary: 72000 },
    { name: "Michael", age: 41, department: "Engineering", salary: 110000 },
    { name: "Emma", age: 27, department: "HR", salary: 65000 },
    { name: "Robert", age: 36, department: "Finance", salary: 95000 },
    { name: "Lisa", age: 32, department: "Marketing", salary: 78000 },
    { name: "David", age: 45, department: "Engineering", salary: 120000 },
    { name: "Jessica", age: 31, department: "HR", salary: 68000 },
  ];
  
  // Create query engine
  const queryEngine = new QueryEngine();
  
  // Example 1: Find all engineers over 35
  const query1 = "department = Engineering AND age > 35";
  const result1 = queryEngine.filter(employees, query1);
  
  console.log("Engineers over 35:");
  result1.forEach(employee => {
    console.log(`  - ${employee.name}: ${employee.age} years old, ${employee.department}`);
  });
  
  // Example 2: Find marketing employees or anyone with salary over 100k
  const query2 = "department = Marketing OR salary > 100000";
  const result2 = queryEngine.filter(employees, query2);
  
  console.log("\nMarketing employees or high earners:");
  result2.forEach(employee => {
    console.log(`  - ${employee.name}: ${employee.department}, $${employee.salary}`);
  });
  
  // Example 3: Complex query with parentheses
  const query3 = "(department = Engineering OR department = HR) AND age < 35";
  const result3 = queryEngine.filter(employees, query3);
  
  console.log("\nYoung engineers or HR employees:");
  result3.forEach(employee => {
    console.log(`  - ${employee.name}: ${employee.age} years old, ${employee.department}`);
  });
}

// Execute the main function
main();
