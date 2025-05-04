import { PaymentContext } from "./context";
import { CreditCardPayment, PayPalPayment, BitcoinPayment } from "./strategy";

// Mock console.log to capture output
let consoleOutput: string[] = [];
const originalLog = console.log;

beforeEach(() => {
  // Reset captured output and mock console.log before each test
  consoleOutput = [];
  console.log = (message: any, ...optionalParams: any[]) => {
    // Format message similar to how console.log does
    const formattedMessage = [message, ...optionalParams]
      .map((arg) =>
        typeof arg === "object" ? JSON.stringify(arg) : String(arg)
      )
      .join(" ");
    consoleOutput.push(formattedMessage);
    // originalLog(message, ...optionalParams); // Uncomment to see output during test run
  };
});

afterEach(() => {
  // Restore original console.log after each test
  console.log = originalLog;
});

describe("Payment Processing Strategy Pattern", () => {
  test("should process payment using Credit Card strategy", () => {
    const strategy = new CreditCardPayment("1111222233334444", "01/28", "987");
    const context = new PaymentContext(strategy);
    context.processPayment(150.0);

    // Check console output
    const outputString = consoleOutput.join("\n");
    expect(outputString).toContain(
      "Processing $150.00 using Credit Card: ****-****-****-4444"
    );
    expect(outputString).toContain("Credit Card payment successful.");
  });

  test("should process payment using PayPal strategy", () => {
    const strategy = new PayPalPayment("test@domain.com");
    const context = new PaymentContext(strategy);
    context.processPayment(75.25);

    const outputString = consoleOutput.join("\n");
    expect(outputString).toContain(
      "Processing $75.25 using PayPal: test@domain.com"
    );
    expect(outputString).toContain("PayPal payment successful.");
  });

  test("should process payment using Bitcoin strategy", () => {
    const strategy = new BitcoinPayment("3AnotherBitcoinAddressExample");
    const context = new PaymentContext(strategy);
    context.processPayment(300.0);

    const outputString = consoleOutput.join("\n");
    expect(outputString).toContain(
      "Processing $300.00 equivalent in BTC to wallet: 3AnotherBitcoinAddressExample"
    );
    expect(outputString).toContain("Bitcoin payment initiated");
  });

  test("should switch strategies and process payments correctly", () => {
    const creditCard = new CreditCardPayment(
      "9999888877776666",
      "11/26",
      "555"
    );
    const paypal = new PayPalPayment("switch@test.org");
    const bitcoin = new BitcoinPayment("1SwitchAddressExample");

    // Start with Credit Card
    const context = new PaymentContext(creditCard);
    context.processPayment(10.0);

    // Switch to PayPal
    context.setStrategy(paypal);
    context.processPayment(20.0);

    // Switch to Bitcoin
    context.setStrategy(bitcoin);
    context.processPayment(30.0);

    const outputString = consoleOutput.join("\n");

    // Check initialization and switching logs
    expect(outputString).toContain(
      "Initialized Credit Card Payment for card ending in 6666"
    );
    expect(outputString).toContain(
      "Payment context initialized with strategy: CreditCardPayment"
    );
    expect(outputString).toContain(
      "Changing payment strategy from CreditCardPayment to PayPalPayment"
    );
    expect(outputString).toContain(
      "Initialized PayPal Payment for email: switch@test.org"
    ); // Note: Init logs happen when strategy created
    expect(outputString).toContain(
      "Changing payment strategy from PayPalPayment to BitcoinPayment"
    );
    expect(outputString).toContain(
      "Initialized Bitcoin Payment for wallet: 1Swit...mple"
    ); // Note: Init logs happen when strategy created

    // Check payment processing logs
    expect(outputString).toContain(
      "Processing $10.00 using Credit Card: ****-****-****-6666"
    );
    expect(outputString).toContain(
      "Processing $20.00 using PayPal: switch@test.org"
    );
    expect(outputString).toContain(
      "Processing $30.00 equivalent in BTC to wallet: 1SwitchAddressExample"
    );

    // Check final state (last payment was Bitcoin)
    const lastProcessingLine = consoleOutput
      .filter((line) => line.startsWith("Processing $"))
      .pop();
    expect(lastProcessingLine).toContain("BTC");
    expect(lastProcessingLine).toContain("$30.00");
  });
});
