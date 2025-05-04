/** @type {import('ts-jest').JestConfigWithTsJest} */
module.exports = {
  preset: "ts-jest",
  testEnvironment: "node",
  testMatch: ["**/?(*.)+(spec|test).ts"], // Look for .test.ts or .spec.ts files
  moduleFileExtensions: ["ts", "js", "json", "node"],
  roots: ["."], // Look for tests in the current directory
};
