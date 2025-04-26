/** @type {import('ts-jest').JestConfigWithTsJest} */
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'node',
  roots: ['<rootDir>'], // Specify the root directory for tests
  testMatch: [
    '**/?(*.)+(spec|test).ts?(x)' // Standard Jest pattern for test files
  ],
  moduleFileExtensions: ['ts', 'tsx', 'js', 'jsx', 'json', 'node'],
  collectCoverage: true, // Optional: enable coverage reports
  coverageDirectory: 'coverage', // Optional: specify coverage directory
  coverageProvider: 'v8', // Optional: use V8 for coverage
};
