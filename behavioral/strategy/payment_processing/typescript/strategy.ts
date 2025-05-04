/**
 * Interface for payment strategies.
 */
export interface IPaymentStrategy {
  pay(amount: number): void;
}

/**
 * Concrete strategy for Credit Card payments.
 */
export class CreditCardPayment implements IPaymentStrategy {
  private cardNumber: string;
  private expiryDate: string;
  private cvv: string;

  constructor(cardNumber: string, expiryDate: string, cvv: string) {
    this.cardNumber = cardNumber;
    this.expiryDate = expiryDate;
    this.cvv = cvv;
    // Mask card number for logging, show only last 4 digits
    console.log(
      `Initialized Credit Card Payment for card ending in ${this.cardNumber.slice(
        -4
      )}`
    );
  }

  pay(amount: number): void {
    const maskedNumber = `****-****-****-${this.cardNumber.slice(-4)}`;
    console.log(
      `Processing $${amount.toFixed(2)} using Credit Card: ${maskedNumber}`
    );
    // Simulate payment processing logic
    console.log("Credit Card payment successful.");
  }
}

/**
 * Concrete strategy for PayPal payments.
 */
export class PayPalPayment implements IPaymentStrategy {
  private email: string;

  constructor(email: string) {
    this.email = email;
    console.log(`Initialized PayPal Payment for email: ${this.email}`);
  }

  pay(amount: number): void {
    console.log(`Processing $${amount.toFixed(2)} using PayPal: ${this.email}`);
    // Simulate PayPal API call
    console.log("PayPal payment successful.");
  }
}

/**
 * Concrete strategy for Bitcoin payments.
 */
export class BitcoinPayment implements IPaymentStrategy {
  private walletAddress: string;

  constructor(walletAddress: string) {
    this.walletAddress = walletAddress;
    console.log(
      `Initialized Bitcoin Payment for wallet: ${this.walletAddress.substring(
        0,
        5
      )}...${this.walletAddress.slice(-4)}`
    );
  }

  pay(amount: number): void {
    console.log(
      `Processing $${amount.toFixed(2)} equivalent in BTC to wallet: ${
        this.walletAddress
      }`
    );
    // Simulate Bitcoin transaction logic
    console.log("Bitcoin payment initiated (waiting for confirmation).");
  }
}
