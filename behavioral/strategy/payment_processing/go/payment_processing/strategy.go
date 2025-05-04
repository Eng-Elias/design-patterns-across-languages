package payment_processing

import (
	"fmt"
	"log"
)

// PaymentStrategy defines the interface for payment methods.
type PaymentStrategy interface {
	Pay(amount float64)
}

// --- Concrete Strategies ---

// CreditCardPayment implements PaymentStrategy for credit cards.
type CreditCardPayment struct {
	CardNumber string // Store masked number for safety in logs/output
	ExpiryDate string
	CVV        string
}

// NewCreditCardPayment creates a new CreditCardPayment strategy.
func NewCreditCardPayment(cardNumber, expiryDate, cvv string) *CreditCardPayment {
	// Mask card number for logging, show only last 4 digits
	maskedNumber := "****-****-****-" + cardNumber[len(cardNumber)-4:]
	log.Printf("Initialized Credit Card Payment for card ending in %s", cardNumber[len(cardNumber)-4:])
	return &CreditCardPayment{
		CardNumber: maskedNumber,
		ExpiryDate: expiryDate,
		CVV:        cvv,
	}
}

func (cc *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Processing $%.2f using Credit Card: %s\n", amount, cc.CardNumber)
	// Simulate processing
	fmt.Println("Credit Card payment successful.")
	log.Printf("Successfully processed $%.2f via Credit Card %s", amount, cc.CardNumber)
}

// PayPalPayment implements PaymentStrategy for PayPal.
type PayPalPayment struct {
	Email string
}

// NewPayPalPayment creates a new PayPalPayment strategy.
func NewPayPalPayment(email string) *PayPalPayment {
	log.Printf("Initialized PayPal Payment for email: %s", email)
	return &PayPalPayment{Email: email}
}

func (pp *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Processing $%.2f using PayPal: %s\n", amount, pp.Email)
	// Simulate processing
	fmt.Println("PayPal payment successful.")
	log.Printf("Successfully processed $%.2f via PayPal account %s", amount, pp.Email)
}

// BitcoinPayment implements PaymentStrategy for Bitcoin.
type BitcoinPayment struct {
	WalletAddress string
}

// NewBitcoinPayment creates a new BitcoinPayment strategy.
func NewBitcoinPayment(walletAddress string) *BitcoinPayment {
	// Log a shortened address
	logMsg := walletAddress
	if len(walletAddress) > 10 { // Basic check for length
		logMsg = fmt.Sprintf("%s...%s", walletAddress[:5], walletAddress[len(walletAddress)-4:])
	}
	log.Printf("Initialized Bitcoin Payment for wallet: %s", logMsg)
	return &BitcoinPayment{WalletAddress: walletAddress}
}

func (btc *BitcoinPayment) Pay(amount float64) {
	fmt.Printf("Processing $%.2f equivalent in BTC to wallet: %s\n", amount, btc.WalletAddress)
	// Simulate processing
	fmt.Println("Bitcoin payment initiated (waiting for confirmation).")
	log.Printf("Initiated Bitcoin payment of $%.2f to wallet %s", amount, btc.WalletAddress)
}
