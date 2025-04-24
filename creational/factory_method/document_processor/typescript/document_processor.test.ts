import * as fs from "fs";
import * as path from "path";
// Corrected imports for the refactored classes from the src directory
import {
  TextProcessor,
  JSONProcessor,
  HTMLProcessor,
} from "./document_processor";

const TEST_OUTPUT_DIR = "temp_test_output";
const TEST_FILENAME_BASE = "test_doc";
const TEST_TITLE = "TS Test Document Title";
const TEST_CONTENT = ["TypeScript Line 1.", "Another line for TS."];

beforeAll(async () => {
  try {
    await fs.promises.mkdir(TEST_OUTPUT_DIR, { recursive: true });
  } catch (err) {
    console.error("Failed to create test directory:", err);
    throw err;
  }
});

afterAll(async () => {
  try {
    await fs.promises.rm(TEST_OUTPUT_DIR, { recursive: true, force: true });
  } catch (err) {
    console.error("Failed to remove test directory:", err);
  }
});

describe("DocumentProcessor Factory Method - Refactored TypeScript", () => {
  test("TextProcessor should create and save a correct .txt file", async () => {
    const processor = new TextProcessor();
    await processor.processAndSave(
      TEST_TITLE,
      TEST_CONTENT,
      TEST_OUTPUT_DIR,
      TEST_FILENAME_BASE
    );

    const expectedFilepath = path.join(
      TEST_OUTPUT_DIR,
      `${TEST_FILENAME_BASE}.txt`
    );

    await expect(
      fs.promises.access(expectedFilepath, fs.constants.F_OK)
    ).resolves.toBeUndefined();

    const fileContent = await fs.promises.readFile(expectedFilepath, "utf-8");
    expect(fileContent).toContain(`Title: ${TEST_TITLE}`);
    expect(fileContent).toContain(TEST_CONTENT[0]);
    expect(fileContent).toContain(TEST_CONTENT[1]);
  });

  test("JSONProcessor should create and save a correct .json file", async () => {
    const processor = new JSONProcessor();
    await processor.processAndSave(
      TEST_TITLE,
      TEST_CONTENT,
      TEST_OUTPUT_DIR,
      TEST_FILENAME_BASE
    );

    const expectedFilepath = path.join(
      TEST_OUTPUT_DIR,
      `${TEST_FILENAME_BASE}.json`
    );

    await expect(
      fs.promises.access(expectedFilepath, fs.constants.F_OK)
    ).resolves.toBeUndefined();

    const fileContent = await fs.promises.readFile(expectedFilepath, "utf-8");
    const jsonData = JSON.parse(fileContent);

    expect(jsonData).toHaveProperty("title", TEST_TITLE);
    expect(jsonData).toHaveProperty("content");
    expect(jsonData.content).toEqual(TEST_CONTENT);
  });

  test("HTMLProcessor should create and save a correct .html file", async () => {
    const processor = new HTMLProcessor();
    await processor.processAndSave(
      TEST_TITLE,
      TEST_CONTENT,
      TEST_OUTPUT_DIR,
      TEST_FILENAME_BASE
    );

    const expectedFilepath = path.join(
      TEST_OUTPUT_DIR,
      `${TEST_FILENAME_BASE}.html`
    );

    await expect(
      fs.promises.access(expectedFilepath, fs.constants.F_OK)
    ).resolves.toBeUndefined();

    const fileContent = await fs.promises.readFile(expectedFilepath, "utf-8");
    expect(fileContent).toContain(`<title>${TEST_TITLE}</title>`);
    expect(fileContent).toContain(`<h1>${TEST_TITLE}</h1>`);
    expect(fileContent).toContain(`<p>${TEST_CONTENT[0]}</p>`);
    expect(fileContent).toContain(`<p>${TEST_CONTENT[1]}</p>`);
  });
});
