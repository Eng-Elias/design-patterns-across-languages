import { IOrderState, NewOrderState } from "./state";

/**
 * The Context class representing an order.
 */
export class Order {
  public items: string[] = [];
  private state: IOrderState;
  public readonly orderId: string;

  constructor(orderId: string) {
    this.orderId = orderId;
    // Default state is NewOrderState
    this.state = new NewOrderState(this);
    console.log(
      `Order ${this.orderId} created in state: ${this.state.toString()}`
    );
  }

  /**
   * Allows changing the state.
   */
  public setState(state: IOrderState): void {
    console.log(
      `Order ${
        this.orderId
      }: Transitioning from ${this.state.toString()} to ${state.toString()}`
    );
    this.state = state;
  }

  /**
   * Returns the current state.
   */
  public getState(): IOrderState {
    return this.state;
  }

  public addItem(item: string): void {
    this.state.addItem(item);
  }

  public processPayment(): void {
    this.state.processPayment();
  }

  public ship(): void {
    this.state.ship();
  }

  public deliver(): void {
    this.state.deliver();
  }

  public cancel(): void {
    this.state.cancel();
  }

  public toString(): string {
    return `Order [ID: ${
      this.orderId
    }, State: ${this.state.toString()}, Items: [${this.items.join(", ")}]]`;
  }
}
