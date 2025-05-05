import {
  FunctionDefinitionNode,
  VariableDeclarationNode,
  AssignmentStatementNode,
  IfStatementNode,
  ExpressionStatementNode,
  ExpressionNode,
  StatementNode,
} from "./ast_nodes"; // Adjust path as necessary
import { ComplexityVisitor } from "./complexity_visitor";
import { SyntaxCheckVisitor } from "./syntax_check_visitor";
import { PrettyPrintVisitor } from "./pretty_print_visitor";

describe("Code Analyzer Visitors", () => {
  let ast: FunctionDefinitionNode;

  // Set up a common AST for testing different visitors - same as Python example
  beforeEach(() => {
    ast = new FunctionDefinitionNode(
      "testFunc",
      ["a"], // Parameter 'a'
      [
        new VariableDeclarationNode("b", "number", new ExpressionNode("a + 1")),
        new IfStatementNode(
          new ExpressionNode("a > 0"), // Condition
          [
            // Then branch
            new AssignmentStatementNode("b", new ExpressionNode("10")),
          ],
          [
            // Else branch
            new AssignmentStatementNode("b", new ExpressionNode("20")),
          ]
        ),
        new AssignmentStatementNode("c", new ExpressionNode("b")), // Use undeclared 'c'
        new ExpressionStatementNode(new ExpressionNode("console.log(b)")), // Simulate print
      ]
    );
  });

  test("PrettyPrintVisitor should format AST correctly", () => {
    const printer = new PrettyPrintVisitor();
    // Manually adjust accept calls if PrettyPrintVisitor needs explicit indent control
    // For simplicity, we use the standard accept and check output contains expected parts
    ast.accept(printer);
    const output = printer.getOutput();
    console.log(`--- Pretty Print Test Output ---\n${output}`); // Log for visual check

    expect(output).toContain("Function: testFunc(a)");
    // Indentation makes exact string matching hard, usetoContain
    expect(output).toContain("Declare: b: number = a + 1");
    expect(output).toContain("If (a > 0):");
    expect(output).toContain("Assign: b = 10");
    expect(output).toContain("Else:");
    expect(output).toContain("Assign: b = 20");
    expect(output).toContain("Assign: c = b");
    expect(output).toContain("ExprStmt: console.log(b)");
  });

  test("ComplexityVisitor should calculate correct complexity", () => {
    const visitor = new ComplexityVisitor();
    ast.accept(visitor);
    // Base complexity 1 + 1 for the IfStatement = 2
    expect(visitor.getComplexity()).toBe(2);
  });

  test("SyntaxCheckVisitor should detect undeclared variable usage", () => {
    const visitor = new SyntaxCheckVisitor();
    ast.accept(visitor);

    const errors = visitor.getErrors();
    console.log(`--- Syntax Undeclared Test Errors ---\n${errors}`);

    expect(errors).toHaveLength(1);
    expect(errors[0]).toContain("Variable 'c' used before declaration");
  });

  test("SyntaxCheckVisitor should detect variable redeclaration", () => {
    const redeclareAst = new FunctionDefinitionNode(
      "redeclareFunc",
      [],
      [
        new VariableDeclarationNode("x", "number"),
        new VariableDeclarationNode("x", "string"), // Redeclaration
      ]
    );
    const visitor = new SyntaxCheckVisitor();

    redeclareAst.accept(visitor);

    const errors = visitor.getErrors();
    console.log(`--- Syntax Redeclaration Test Errors ---\n${errors}`);

    expect(errors.length).toBeGreaterThanOrEqual(1); // Could be 1 or more depending on exact logic
    expect(
      errors.some((e) => e.includes("Identifier 'x' already declared"))
    ).toBe(true);
  });

  test("SyntaxCheckVisitor should pass for correct code", () => {
    const okAst = new FunctionDefinitionNode(
      "okFunc",
      ["p"],
      [
        new VariableDeclarationNode(
          "v",
          "boolean",
          new ExpressionNode("p == 0")
        ),
        new IfStatementNode(new ExpressionNode("p > 10"), [
          new AssignmentStatementNode("v", new ExpressionNode("true")),
        ]),
        new AssignmentStatementNode("v", new ExpressionNode("false")),
      ]
    );
    const visitor = new SyntaxCheckVisitor();

    okAst.accept(visitor);

    const errors = visitor.getErrors();
    console.log(`--- Syntax OK Test Errors ---\n${errors}`);

    expect(errors).toHaveLength(0);
  });
});
