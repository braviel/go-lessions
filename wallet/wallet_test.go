package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bcoin(10))
		assertBalance(t, wallet, Bcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bcoin(20)}
		err := wallet.Withdraw(Bcoin(10))
		assertBalance(t, wallet, Bcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw Insufficent funds", func(t *testing.T) {
		startingBalance := Bcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected error but did not get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got an error but did not expected one: ", got.Error())
	}
}
