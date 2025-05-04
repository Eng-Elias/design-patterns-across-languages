package payment_processing

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// --- Test Setup: Capture stdout ONLY ---

// captureStdout captures standard output during the execution of testFunc.
func captureStdout(t *testing.T, testFunc func()) string {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe() // Create a pipe
	os.Stdout = w         // Redirect stdout to the pipe writer

	outC := make(chan string)
	// Copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r) // Read from pipe reader
		outC <- buf.String()
	}()

	t.Cleanup(func() {
		w.Close() // Close the writer, signaling EOF to the reader goroutine
		os.Stdout = originalStdout // Restore original stdout
	})

	testFunc() // Execute the function that prints to stdout

	// Ensure the writer is closed BEFORE waiting for the output
	// (Cleanup will handle this, but explicit close can be clearer)
	w.Close()
	out := <-outC // Wait for the reader goroutine to finish

	return out
}

// TestMain sets up global log flags (optional, keeps logs clean if they appear).
func TestMain(m *testing.M) {
	// Optional: Keep logs clean if they show up in test runner, but we don't assert them.
	log.SetOutput(io.Discard) // Discard log output entirely for tests
	log.SetFlags(0)
	code := m.Run()
	// Restore default logging behavior if necessary (usually not needed)
	// log.SetOutput(os.Stderr)
	// log.SetFlags(log.LstdFlags)
	os.Exit(code)
}

// --- Unit Tests (Asserting STDOUT only) ---

func TestCreditCardPayment(t *testing.T) {
	output := captureStdout(t, func() {
		strategy := NewCreditCardPayment("1111222233334444", "01/26", "555")
		context := NewPaymentContext(strategy)
		context.ProcessPayment(150.00)
	})

	// Check processing messages printed by ProcessPayment (stdout)
	assert.Contains(t, output, "Processing $150.00 using Credit Card: ****-****-****-4444", "Should contain CC processing message")
	assert.Contains(t, output, "Credit Card payment successful.", "Should contain success message")
}

func TestPayPalPayment(t *testing.T) {
	output := captureStdout(t, func() {
		strategy := NewPayPalPayment("test@domain.com")
		context := NewPaymentContext(strategy)
		context.ProcessPayment(75.25)
	})

	// Check processing messages printed by ProcessPayment (stdout)
	assert.Contains(t, output, "Processing $75.25 using PayPal: test@domain.com", "Should contain PayPal processing message")
	assert.Contains(t, output, "PayPal payment successful.", "Should contain success message")
}

func TestBitcoinPayment(t *testing.T) {
	output := captureStdout(t, func() {
		strategy := NewBitcoinPayment("3AnotherBitcoinAddressExample")
		context := NewPaymentContext(strategy)
		context.ProcessPayment(300.00)
	})

	// Check processing messages printed by ProcessPayment (stdout)
	assert.Contains(t, output, "Processing $300.00 equivalent in BTC to wallet: 3AnotherBitcoinAddressExample", "Should contain Bitcoin processing message")
	assert.Contains(t, output, "Bitcoin payment initiated", "Should contain initiation message")
}

func TestStrategySwitching(t *testing.T) {
	output := captureStdout(t, func() {
		creditCard := NewCreditCardPayment("5555666677778888", "02/27", "666")
		paypal := NewPayPalPayment("switch@test.org")
		bitcoin := NewBitcoinPayment("1SwitchAddressExample")

		context := NewPaymentContext(creditCard)

		context.ProcessPayment(10.00)

		context.SetStrategy(paypal)
		context.ProcessPayment(20.00)

		context.SetStrategy(bitcoin)
		context.ProcessPayment(30.00)
	})

	// Check processing messages printed by ProcessPayment calls (stdout)
	assert.Contains(t, output, "Processing $10.00 using Credit Card: ****-****-****-8888")
	assert.Contains(t, output, "Processing $20.00 using PayPal: switch@test.org")
	assert.Contains(t, output, "Processing $30.00 equivalent in BTC to wallet: 1SwitchAddressExample")

	// Check if the last payment processed was Bitcoin (using captured STDOUT)
	lines := strings.Split(strings.TrimSpace(output), "\n")
	lastProcessingLine := ""
	for i := len(lines) - 1; i >= 0; i-- {
        line := strings.TrimSpace(lines[i])
		// Find the last line that starts with "Processing $"
		if strings.HasPrefix(line, "Processing $") {
			lastProcessingLine = line
			break
		}
	}

	// Ensure we actually found a processing line before asserting
	if lastProcessingLine == "" {
		t.Fatalf("Could not find any 'Processing $' line in STDOUT part of output:\n%s", output)
	}

	assert.Contains(t, lastProcessingLine, "BTC", "The last payment processed should be Bitcoin")
	assert.Contains(t, lastProcessingLine, "$30.00", "The last payment amount should be $30.00")
}
