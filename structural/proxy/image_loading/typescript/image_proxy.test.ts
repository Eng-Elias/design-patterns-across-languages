import {
  describe,
  beforeEach,
  beforeAll,
  afterAll,
  test,
  expect,
  jest,
} from "@jest/globals";
import { ProxyImage, RealImage, IImage } from "./image_proxy";

// Mock the console.log to prevent test output noise
const originalConsoleLog = console.log;
beforeAll(() => {
  console.log = jest.fn();
});

afterAll(() => {
  console.log = originalConsoleLog;
});

// Create a mock implementation of RealImage
class MockRealImage implements IImage {
  private filename: string;

  constructor(filename: string) {
    this.filename = filename;
    console.log(`Loading image: '${filename}' from disk... (Simulating async delay)`);
    console.log(`Finished loading image: '${filename}'`);
  }

  async display(): Promise<void> {
    console.log(`Displaying image: '${this.filename}'`);
  }

  getFilename(): string {
    return this.filename;
  }
}

describe("ProxyImage", () => {
  const filename = "test_image.jpg";
  let proxy: ProxyImage;
  // Let TypeScript infer the spy type
  let createSpy: jest.SpiedFunction<typeof RealImage.create>;

  beforeEach(() => {
    jest.clearAllMocks();
    // Spy on the actual RealImage.create and provide the mock implementation
    createSpy = jest.spyOn(RealImage, 'create')
                    // Cast the mock result type to satisfy the original signature
                    .mockImplementation(async (filename: string) => new MockRealImage(filename) as any);
    proxy = new ProxyImage(filename);
  });

  test("proxy initialization does not load real image", () => {
    expect(proxy).toBeDefined();
    expect(proxy.getFilename()).toBe(filename);
    expect(proxy.isLoaded()).toBe(false);
    expect(createSpy).not.toHaveBeenCalled();
  });

  test("proxy loads real image on first display", async () => {
    expect(proxy.isLoaded()).toBe(false);
    expect(createSpy).not.toHaveBeenCalled();

    await proxy.display();

    expect(createSpy).toHaveBeenCalledTimes(1);
    expect(createSpy).toHaveBeenCalledWith(filename);
    expect(proxy.isLoaded()).toBe(true);
  });

  test("proxy reuses loaded real image on subsequent displays", async () => {
    await proxy.display();
    expect(createSpy).toHaveBeenCalledTimes(1);

    createSpy.mockClear();

    await proxy.display();

    expect(createSpy).not.toHaveBeenCalled();
    expect(proxy.isLoaded()).toBe(true);
  });

  test("proxy get_filename does not load real image", () => {
    const retrievedFilename = proxy.getFilename();

    expect(retrievedFilename).toBe(filename);
    expect(createSpy).not.toHaveBeenCalled();
    expect(proxy.isLoaded()).toBe(false);
  });

  test("real image loading simulation", async () => {
    const consoleSpy = jest.spyOn(console, "log");

    // Create and load the real image by calling the spied-upon method
    // This will execute the mockImplementation provided to the spy
    const realImage = await RealImage.create(filename); 

    // Verify loading messages from MockRealImage constructor
    expect(consoleSpy).toHaveBeenCalledWith(`Loading image: '${filename}' from disk... (Simulating async delay)`);
    expect(consoleSpy).toHaveBeenCalledWith(`Finished loading image: '${filename}'`);

    consoleSpy.mockClear();

    await realImage.display();
    expect(consoleSpy).toHaveBeenCalledWith(`Displaying image: '${filename}'`);
    expect(realImage.getFilename()).toBe(filename);
  });
});
