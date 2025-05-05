import {
  Node,
  FunctionDefinitionNode,
  VariableDeclarationNode,
  AssignmentStatementNode,
  IfStatementNode,
  ExpressionStatementNode,
  ExpressionNode,
} from "./code_analyzer/ast_nodes";
import { ComplexityVisitor } from "./code_analyzer/complexity_visitor";
import { SyntaxCheckVisitor } from "./code_analyzer/syntax_check_visitor";
import { PrettyPrintVisitor } from "./code_analyzer/pretty_print_visitor";

// --- Client Code: Demo ---
function main() {
  console.log("--- Visitor Pattern: Code Analysis Demo (TypeScript) ---");

  // 1. Construct a simple AST (Object Structure) - Same as Python example
  const astRoot: Node = new FunctionDefinitionNode(
    "calculate",
    ["x", "y"], // Parameters
    [
      new VariableDeclarationNode("result", "number", new ExpressionNode("0")),
      new VariableDeclarationNode("temp", "number"), // No initializer
      new IfStatementNode(
        new ExpressionNode("x > y"), // Condition
        [
          // Then branch
          new AssignmentStatementNode("result", new ExpressionNode("x")),
          new AssignmentStatementNode("temp", new ExpressionNode("1")),
        ],
        [
          // Else branch
          new AssignmentStatementNode("result", new ExpressionNode("y")),
          new AssignmentStatementNode("temp", new ExpressionNode("0")),
        ]
      ),
      new ExpressionStatementNode(new ExpressionNode("console.log(result)")), // Simulate print
      // Uncomment below to introduce a syntax error for the demo
      // new AssignmentStatementNode("z", new ExpressionNode("temp"))
    ]
  );

  // 2. Create Visitors
  const prettyPrinter = new PrettyPrintVisitor();
  const complexityChecker = new ComplexityVisitor();
  // Syntax checker needs careful scope management due to accept() traversal
  const syntaxChecker = new SyntaxCheckVisitor();

  // --- Apply Visitors ---

  console.log("\n--- Applying PrettyPrintVisitor ---");
  astRoot.accept(prettyPrinter); // Node controls traversal
  console.log(prettyPrinter.getOutput());
  // Note: Indentation logic primarily resides within the visitor's visit methods

  console.log("\n--- Applying ComplexityVisitor ---");
  complexityChecker.reset(); // Reset state before use
  astRoot.accept(complexityChecker); // Node controls traversal
  console.log(`Calculated Complexity: ${complexityChecker.getComplexity()}`);

  console.log("\n--- Applying SyntaxCheckVisitor ---");
  syntaxChecker.reset(); // Reset state before use
  astRoot.accept(syntaxChecker);

  const errors = syntaxChecker.getErrors();
  if (errors.length > 0) {
    console.log("Syntax Errors Found:");
    errors.forEach((error: string) => console.log(`- ${error}`));
  } else {
    console.log("No syntax errors found (based on current AST).");
    console.log("(Uncomment the 'z = temp' line in main.ts to see an error)");
  }

  console.log("\n--- Demo Finished ---");
}

// Run the main function
main();
