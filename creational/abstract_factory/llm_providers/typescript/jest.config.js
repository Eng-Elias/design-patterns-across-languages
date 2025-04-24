/** @type {import('ts-jest').JestConfigWithTsJest} */
module.exports = {
  preset: 'ts-jest',             // Use ts-jest preset
  testEnvironment: 'node',       // Environment for testing (Node.js)
  roots: ['<rootDir>'],          // Specify the root directory (current directory)
  testMatch: [
    '**/*.test.ts'           // Look for files ending in .test.ts
  ],
  moduleFileExtensions: ['ts', 'js', 'json', 'node'], // File extensions Jest should look for
  transform: {
    '^.+\.ts$': [
      'ts-jest',
      {
        tsconfig: 'tsconfig.json' // Use our tsconfig for transformation
      }
    ]
  },
  // Optional: If you use path aliases like @/ in tsconfig.json
  // moduleNameMapper: {
  //   '^@/(.*)$': '<rootDir>/src/$1' // Adjust if you use a src/ directory
  // }
};
