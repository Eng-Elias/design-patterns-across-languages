from image_proxy import ProxyImage, Image
import time

def main() -> None:
    """Demonstrates the Virtual Proxy pattern for lazy loading of images."""

    print("--- Creating image proxies --- ")
    # Creating proxy objects doesn't load the actual images yet
    image1: Image = ProxyImage("photo_high_res_001.jpg")
    image2: Image = ProxyImage("photo_high_res_002.jpg")
    image3: Image = ProxyImage("document_scan_003.png")

    print("\n--- Accessing filenames (should not load) --- ")
    print(f"Image 1 filename: {image1.get_filename()}")
    print(f"Image 2 filename: {image2.get_filename()}")

    # Check loaded status (optional method)
    if isinstance(image1, ProxyImage):
        print(f"Image 1 loaded status: {image1.is_loaded()}") # False

    print("\n--- Requesting display for image 1 (will trigger loading) --- ")
    start_time = time.time()
    image1.display()
    end_time = time.time()
    print(f"(Display took {end_time - start_time:.2f} seconds)")

    if isinstance(image1, ProxyImage):
        print(f"Image 1 loaded status: {image1.is_loaded()}") # True

    print("\n--- Requesting display for image 1 AGAIN (should be fast) --- ")
    start_time = time.time()
    image1.display() # The real image is already loaded, delegation is fast
    end_time = time.time()
    print(f"(Display took {end_time - start_time:.2f} seconds)")

    print("\n--- Requesting display for image 2 (will trigger loading) --- ")
    start_time = time.time()
    image2.display()
    end_time = time.time()
    print(f"(Display took {end_time - start_time:.2f} seconds)")

    print("\n--- Requesting display for image 3 (will trigger loading) --- ")
    start_time = time.time()
    image3.display()
    end_time = time.time()
    print(f"(Display took {end_time - start_time:.2f} seconds)")

    print("\nDemo finished.")

if __name__ == "__main__":
    main()
