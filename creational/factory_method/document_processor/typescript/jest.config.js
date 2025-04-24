/** @type {import('ts-jest').JestConfigWithTsJest} */
module.exports = {
  preset: "ts-jest",
  testEnvironment: "node",
  roots: ["<rootDir>"], // Specify the root directory for tests
  testMatch: [
    "**/?(*.)+(spec|test).ts?(x)", // Standard pattern for Jest test files
  ],
  moduleFileExtensions: ["ts", "tsx", "js", "jsx", "json", "node"],
  transform: {
    "^.+.tsx?$": [
      "ts-jest",
      {
        // ts-jest configuration options
        tsconfig: "tsconfig.json",
      },
    ],
  },
};
