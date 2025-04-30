import { ProxyImage, IImage } from './image_proxy';

async function main(): Promise<void> {
    /**
     * Demonstrates the Virtual Proxy pattern for lazy loading of images in TypeScript.
     */

    console.log("--- Creating image proxies --- ");
    // Creating proxy objects doesn't load the actual images yet
    const image1: IImage = new ProxyImage("landscape_high_res_A.png");
    const image2: IImage = new ProxyImage("portrait_high_res_B.jpg");
    const image3: IImage = new ProxyImage("icon_vector_C.svg"); // Filename only, loading is simulated

    console.log("\n--- Accessing filenames (should not load) --- ");
    console.log(`Image 1 filename: ${image1.getFilename()}`);
    console.log(`Image 2 filename: ${image2.getFilename()}`);

    // Check loaded status (optional method)
    // Type assertion needed if IImage doesn't have isLoaded
    if (image1 instanceof ProxyImage) {
        console.log(`Image 1 loaded status: ${image1.isLoaded()}`); // false
    }

    console.log("\n--- Requesting display for image 1 (will trigger loading) --- ");
    let startTime = Date.now();
    await image1.display();
    let endTime = Date.now();
    console.log(`(Display took ${(endTime - startTime)} ms)`);

    if (image1 instanceof ProxyImage) {
        console.log(`Image 1 loaded status: ${image1.isLoaded()}`); // true
    }

    console.log("\n--- Requesting display for image 1 AGAIN (should be fast) --- ");
    startTime = Date.now();
    await image1.display(); // The real image is already loaded, delegation is fast
    endTime = Date.now();
    console.log(`(Display took ${(endTime - startTime)} ms)`);

    console.log("\n--- Requesting display for image 2 (will trigger loading) --- ");
    startTime = Date.now();
    await image2.display();
    endTime = Date.now();
    console.log(`(Display took ${(endTime - startTime)} ms)`);

    console.log("\n--- Requesting display for image 3 (will trigger loading) --- ");
    startTime = Date.now();
    await image3.display();
    endTime = Date.now();
    console.log(`(Display took ${(endTime - startTime)} ms)`);

    console.log("\nDemo finished.");
}

// Execute the async main function
main().catch(error => {
    console.error("An error occurred:", error);
});
