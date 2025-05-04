import { Order } from "./order";

/**
 * State interface defining actions possible on an order.
 */
export interface IOrderState {
  order: Order; // Reference back to the context

  addItem(item: string): void;
  processPayment(): void;
  ship(): void;
  deliver(): void;
  cancel(): void;
  toString(): string; // For descriptive state name
}

// --- Concrete States --- //

export class NewOrderState implements IOrderState {
  constructor(public order: Order) {}

  addItem(item: string): void {
    console.log(`Adding item '${item}' to the order.`);
    this.order.items.push(item);
  }

  processPayment(): void {
    if (this.order.items.length === 0) {
      console.log("Cannot process payment: Order is empty.");
      return;
    }
    console.log("Processing payment...");
    // Simulate payment success
    console.log("Payment successful. Transitioning to PendingPaymentState.");
    this.order.setState(new PendingPaymentState(this.order));
  }

  ship(): void {
    console.log("Cannot ship: Order payment not processed yet.");
  }

  deliver(): void {
    console.log("Cannot deliver: Order not shipped yet.");
  }

  cancel(): void {
    console.log("Cancelling the new order.");
    this.order.setState(new CancelledState(this.order));
  }

  toString(): string {
    return "NewOrderState";
  }
}

export class PendingPaymentState implements IOrderState {
  constructor(public order: Order) {}

  addItem(item: string): void {
    console.log("Cannot add items: Order payment has been processed.");
  }

  processPayment(): void {
    console.log("Payment already processed for this order.");
  }

  ship(): void {
    console.log("Shipping the order...");
    // Simulate shipping success
    console.log("Order shipped. Transitioning to ShippedState.");
    this.order.setState(new ShippedState(this.order));
  }

  deliver(): void {
    console.log("Cannot deliver: Order not shipped yet.");
  }

  cancel(): void {
    console.log("Cancelling the order. Refunding payment...");
    // Simulate refund
    this.order.setState(new CancelledState(this.order));
  }

  toString(): string {
    return "PendingPaymentState";
  }
}

export class ShippedState implements IOrderState {
  constructor(public order: Order) {}

  addItem(item: string): void {
    console.log("Cannot add items: Order has been shipped.");
  }

  processPayment(): void {
    console.log("Payment already processed.");
  }

  ship(): void {
    console.log("Order already shipped.");
  }

  deliver(): void {
    console.log("Delivering the order...");
    // Simulate delivery success
    console.log("Order delivered. Transitioning to DeliveredState.");
    this.order.setState(new DeliveredState(this.order));
  }

  cancel(): void {
    console.log("Cannot cancel: Order has already been shipped.");
  }

  toString(): string {
    return "ShippedState";
  }
}

export class DeliveredState implements IOrderState {
  constructor(public order: Order) {}

  addItem(item: string): void {
    console.log("Cannot add items: Order has been delivered.");
  }

  processPayment(): void {
    console.log("Payment already processed.");
  }

  ship(): void {
    console.log("Order already shipped and delivered.");
  }

  deliver(): void {
    console.log("Order already delivered.");
  }

  cancel(): void {
    console.log("Cannot cancel: Order has already been delivered.");
  }

  toString(): string {
    return "DeliveredState";
  }
}

export class CancelledState implements IOrderState {
  constructor(public order: Order) {}

  addItem(item: string): void {
    console.log("Cannot add items: Order is cancelled.");
  }

  processPayment(): void {
    console.log("Cannot process payment: Order is cancelled.");
  }

  ship(): void {
    console.log("Cannot ship: Order is cancelled.");
  }

  deliver(): void {
    console.log("Cannot deliver: Order is cancelled.");
  }

  cancel(): void {
    console.log("Order is already cancelled.");
  }

  toString(): string {
    return "CancelledState";
  }
}
