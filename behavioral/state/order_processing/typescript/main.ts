import { Order } from "./order";

/**
 * Demonstrates the State pattern with Order processing.
 */
function main() {
  console.log("--- TypeScript State Pattern Demo: Order Processing ---");

  const order = new Order("ORD123");
  console.log(`\nCurrent Status: ${order.toString()}`);

  console.log("\n# --- Attempting actions in New state ---");
  order.addItem("Laptop");
  order.addItem("Mouse");
  console.log(`Current Status: ${order.toString()}`);
  order.ship(); // Invalid action
  order.processPayment();
  console.log(`Current Status: ${order.toString()}`);

  console.log("\n# --- Attempting actions in PendingPayment state ---");
  order.addItem("Keyboard"); // Invalid action
  order.deliver(); // Invalid action
  order.ship();
  console.log(`Current Status: ${order.toString()}`);

  console.log("\n# --- Attempting actions in Shipped state ---");
  order.processPayment(); // Invalid action
  order.cancel(); // Invalid action
  order.deliver();
  console.log(`Current Status: ${order.toString()}`);

  console.log("\n# --- Attempting actions in Delivered state ---");
  order.addItem("Monitor"); // Invalid action
  order.ship(); // Invalid action
  order.cancel(); // Invalid action
  console.log(`Current Status: ${order.toString()}`);

  console.log("\n# --- Demonstrating Cancellation (from New) ---");
  const order2 = new Order("ORD456");
  order2.addItem("Webcam");
  console.log(`Current Status: ${order2.toString()}`);
  order2.cancel();
  console.log(`Current Status: ${order2.toString()}`);
  order2.addItem("Microphone"); // Invalid action
  order2.processPayment(); // Invalid action

  console.log("\n--- Demo Finished ---");
}

// Run the main function
main();
