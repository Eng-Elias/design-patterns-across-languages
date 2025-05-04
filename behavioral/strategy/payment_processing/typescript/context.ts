import { IPaymentStrategy } from "./strategy";

/**
 * The Context defines the interface of interest to clients.
 * It maintains a reference to one of the Strategy objects.
 */
export class PaymentContext {
  private strategy: IPaymentStrategy;

  constructor(strategy: IPaymentStrategy) {
    this.strategy = strategy;
    console.log(
      `Payment context initialized with strategy: ${strategy.constructor.name}`
    );
  }

  /**
   * Allows changing the strategy at runtime.
   */
  public setStrategy(strategy: IPaymentStrategy): void {
    console.log(
      `Changing payment strategy from ${this.strategy.constructor.name} to ${strategy.constructor.name}`
    );
    this.strategy = strategy;
  }

  /**
   * Delegates the payment processing to the current strategy.
   */
  public processPayment(amount: number): void {
    console.log(`\nAttempting to process payment of $${amount.toFixed(2)}...`);
    this.strategy.pay(amount);
  }
}
