import { CreditCardPayment, PayPalPayment, BitcoinPayment } from "./strategy";
import { PaymentContext } from "./context";

function main(): void {
  console.log("--- TypeScript Strategy Pattern Demo: Payment Processing ---");

  // Create concrete strategies
  const creditCard = new CreditCardPayment("1234567890123456", "12/25", "123");
  const paypal = new PayPalPayment("user@example.com");
  const bitcoin = new BitcoinPayment("1AbcDeFgHiJkLmNoPqRsTuVwXyZ"); // Example Bitcoin address

  // Create context with initial strategy (Credit Card)
  const context = new PaymentContext(creditCard);

  // Process payment using the current strategy
  context.processPayment(100.0);

  // Change strategy to PayPal
  context.setStrategy(paypal);
  context.processPayment(50.5);

  // Change strategy to Bitcoin
  context.setStrategy(bitcoin);
  context.processPayment(250.75);

  console.log("\n--- Demo Finished ---");
}

// Run the demo
main();
