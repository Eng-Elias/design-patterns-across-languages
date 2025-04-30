// Utility function for async delay (simulating I/O)
const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

// --- Subject Interface --- //

export interface IImage {
    /**
     * Subject Interface: Declares the common interface for RealSubject and Proxy.
     */
    display(): Promise<void>; // Display might involve loading, so make it async
    getFilename(): string;
}

// --- Real Subject --- //

export class RealImage implements IImage {
    /**
     * RealSubject: Defines the real object that the proxy represents.
     * Loading the image is simulated here as an async operation.
     */
    private filename: string;

    // Private constructor ensures it's typically created via a factory or proxy
    private constructor(filename: string) {
        this.filename = filename;
        // Loading happens immediately in the 'create' method
    }

    // Static factory method to simulate async loading during creation
    public static async create(filename: string): Promise<RealImage> {
        const image = new RealImage(filename);
        await image.loadFromDisk();
        return image;
    }

    private async loadFromDisk(): Promise<void> {
        /**
         * Private helper method to simulate loading the image data asynchronously.
         */
        console.log(`Loading image: '${this.filename}' from disk... (Simulating async delay)`);
        await delay(1500); // Simulate time-consuming I/O
        console.log(`Finished loading image: '${this.filename}'`);
    }

    public async display(): Promise<void> {
        /**
         * Displays the image (after it has been loaded).
         */
        console.log(`Displaying image: '${this.filename}'`);
        // In a real scenario, this might render to a canvas, etc.
    }

    public getFilename(): string {
        /**
         * Returns the filename.
         */
        return this.filename;
    }
}

// --- Proxy --- //

export class ProxyImage implements IImage {
    /**
     * Proxy: Maintains a reference that lets the proxy access the real subject.
     * Implements the same interface as the RealSubject.
     * Controls access to the real subject and is responsible for its creation (lazy loading).
     */
    private filename: string;
    private realImage: RealImage | null = null; // Reference to RealImage, initially null

    constructor(filename: string) {
        this.filename = filename;
        console.log(`ProxyImage created for: '${this.filename}' (Real image not loaded yet)`);
    }

    public async display(): Promise<void> {
        /**
         * Handles the display request.
         * Loads the RealImage asynchronously only if it hasn't been loaded yet (Lazy Initialization).
         */
        if (this.realImage === null) {
            console.log(`Proxy for '${this.filename}': Real image needs loading.`);
            // Lazy initialization: Create and load the RealImage object only when needed
            this.realImage = await RealImage.create(this.filename);
        } else {
            console.log(`Proxy for '${this.filename}': Real image already loaded.`);
        }

        // Delegate the display call to the RealImage object
        // No need to await here as the display logic in RealImage itself is synchronous
        // If RealImage.display() were async, we would await here.
        await this.realImage.display();
    }

    public getFilename(): string {
        /**
         * Returns the filename without loading the real image.
         */
        return this.filename;
    }

    // Optional: Method to check if loaded without triggering load
    public isLoaded(): boolean {
        /**
         * Checks if the real image has been loaded.
         */
        return this.realImage !== null;
    }
}
