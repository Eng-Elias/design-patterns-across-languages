import { Order } from "./order";
import {
  NewOrderState,
  PendingPaymentState,
  ShippedState,
  DeliveredState,
  CancelledState,
} from "./state";

// Mock console.log to prevent test output clutter
beforeAll(() => {
  jest.spyOn(console, "log").mockImplementation(() => {});
});

afterAll(() => {
  jest.restoreAllMocks();
});

describe("Order Processing State Pattern", () => {
  let order: Order;

  beforeEach(() => {
    // Create a new order instance for each test
    order = new Order("TEST1");
  });

  test("should initialize in NewOrderState", () => {
    expect(order.getState()).toBeInstanceOf(NewOrderState);
    expect(order.items).toEqual([]);
  });

  test("should transition correctly from NewOrderState", () => {
    // Add item -> stays New
    order.addItem("Book");
    expect(order.getState()).toBeInstanceOf(NewOrderState);
    expect(order.items).toEqual(["Book"]);

    // Process payment -> transitions to PendingPayment
    order.processPayment();
    expect(order.getState()).toBeInstanceOf(PendingPaymentState);

    // Reset and test cancel
    order = new Order("TEST2");
    order.cancel();
    expect(order.getState()).toBeInstanceOf(CancelledState);
  });

  test("should transition correctly from PendingPaymentState", () => {
    order.addItem("Pen");
    order.processPayment(); // Now in PendingPaymentState
    expect(order.getState()).toBeInstanceOf(PendingPaymentState);

    // Ship -> transitions to Shipped
    order.ship();
    expect(order.getState()).toBeInstanceOf(ShippedState);

    // Reset and test cancel
    order = new Order("TEST3");
    order.addItem("Paper");
    order.processPayment(); // PendingPaymentState
    order.cancel();
    expect(order.getState()).toBeInstanceOf(CancelledState);
  });

  test("should transition correctly from ShippedState", () => {
    order.addItem("Stapler");
    order.processPayment();
    order.ship(); // Now in ShippedState
    expect(order.getState()).toBeInstanceOf(ShippedState);

    // Deliver -> transitions to Delivered
    order.deliver();
    expect(order.getState()).toBeInstanceOf(DeliveredState);
  });

  test("should remain in DeliveredState after reaching it", () => {
    order.addItem("Eraser");
    order.processPayment();
    order.ship();
    order.deliver(); // Now in DeliveredState
    expect(order.getState()).toBeInstanceOf(DeliveredState);
    const currentState = order.getState();

    // Try all actions
    order.addItem("Ruler");
    order.processPayment();
    order.ship();
    order.deliver();
    order.cancel();
    expect(order.getState()).toBeInstanceOf(DeliveredState);
    expect(order.getState()).toBe(currentState); // Should be the same instance
  });

  test("should remain in CancelledState after reaching it", () => {
    order.cancel(); // Now in CancelledState
    expect(order.getState()).toBeInstanceOf(CancelledState);
    const currentState = order.getState();

    // Try all actions
    order.addItem("Tape");
    order.processPayment();
    order.ship();
    order.deliver();
    order.cancel();
    expect(order.getState()).toBeInstanceOf(CancelledState);
    expect(order.getState()).toBe(currentState);
  });

  test("should handle invalid actions in each state correctly", () => {
    // New state
    expect(order.getState()).toBeInstanceOf(NewOrderState);
    order.ship(); // Invalid
    order.deliver(); // Invalid
    expect(order.getState()).toBeInstanceOf(NewOrderState); // No state change
    order.processPayment(); // Invalid (no items)
    expect(order.getState()).toBeInstanceOf(NewOrderState);
    order.addItem("Marker");
    order.processPayment(); // Valid -> PendingPayment
    expect(order.getState()).toBeInstanceOf(PendingPaymentState);

    // PendingPayment state
    order.addItem("Board"); // Invalid
    order.processPayment(); // Invalid (already paid)
    order.deliver(); // Invalid
    expect(order.getState()).toBeInstanceOf(PendingPaymentState);
    order.ship(); // Valid -> Shipped
    expect(order.getState()).toBeInstanceOf(ShippedState);

    // Shipped state
    order.addItem("Projector"); // Invalid
    order.processPayment(); // Invalid
    order.ship(); // Invalid (already shipped)
    order.cancel(); // Invalid
    expect(order.getState()).toBeInstanceOf(ShippedState);
    order.deliver(); // Valid -> Delivered
    expect(order.getState()).toBeInstanceOf(DeliveredState);
  });
});
