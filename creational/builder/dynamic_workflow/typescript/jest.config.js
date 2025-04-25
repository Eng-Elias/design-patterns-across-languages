/** @type {import('ts-jest').JestConfigWithTsJest} */
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  roots: ['.'], // Look for tests in the current directory
  testMatch: [
    '**/*.test.ts' // Pattern to find test files
  ],
  collectCoverage: true, // Optional: Enable code coverage reporting
  coverageDirectory: 'coverage', // Optional: Directory for coverage reports
  coverageProvider: 'v8' // Optional: Coverage provider
};
