package pointers_test

import (
	"github.com/vicgrygorchyk/example_hello/pointers"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Test Deposit", func (t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(10))

		got := wallet.Balance()
		expected := pointers.Bitcoin(10)
	
		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})

	t.Run("Test Withdrawn", func (t *testing.T) {
		wallet := pointers.Wallet{}
		wallet.Deposit(pointers.Bitcoin(20))
		wallet.Withdraw(pointers.Bitcoin(10))

		got := wallet.Balance()
		expected := pointers.Bitcoin(10)
	
		if got != expected {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	})
}