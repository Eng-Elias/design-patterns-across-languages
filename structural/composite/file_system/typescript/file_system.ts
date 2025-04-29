// --- Component Interface (using an abstract class) ---
export abstract class FileSystemComponent {
  protected name: string;

  constructor(name: string) {
    this.name = name;
  }

  public getName(): string {
    return this.name;
  }

  /**
   * Returns the size of the component in bytes.
   */
  public abstract getSize(): number;

  /**
   * Displays the component's structure.
   */
  public abstract display(indent?: string): void;

  // Optional: Default implementations for child management (throw error)
  public add(component: FileSystemComponent): void {
    throw new Error("Cannot add to this component");
  }

  public remove(component: FileSystemComponent): void {
    throw new Error("Cannot remove from this component");
  }

  public getChild(index: number): FileSystemComponent {
    throw new Error("Cannot get child from this component");
  }
}

// --- Leaf Class ---
export class File extends FileSystemComponent {
  private size: number;

  constructor(name: string, size: number) {
    super(name);
    this.size = size;
  }

  public getSize(): number {
    return this.size;
  }

  public display(indent: string = ""): void {
    console.log(`${indent}- ${this.getName()} (${this.getSize()} bytes)`);
  }
}

// --- Composite Class ---
export class Directory extends FileSystemComponent {
  private children: FileSystemComponent[] = [];

  constructor(name: string) {
    super(name);
  }

  public add(component: FileSystemComponent): void {
    this.children.push(component);
  }

  public remove(component: FileSystemComponent): void {
    const index = this.children.indexOf(component);
    if (index > -1) {
      this.children.splice(index, 1);
    } else {
      throw new Error(
        `Component '${component.getName()}' not found in directory '${this.getName()}'`
      );
    }
  }

  public getChild(index: number): FileSystemComponent {
    if (index < 0 || index >= this.children.length) {
      throw new RangeError(
        `Index ${index} out of bounds for directory '${this.getName()}'`
      );
    }
    return this.children[index];
  }

  public getSize(): number {
    let totalSize = 0;
    for (const child of this.children) {
      totalSize += child.getSize();
    }
    return totalSize;
  }

  public display(indent: string = ""): void {
    console.log(`${indent}+ ${this.getName()} (${this.getSize()} bytes total)`);
    for (const child of this.children) {
      child.display(indent + "  ");
    }
  }
}
