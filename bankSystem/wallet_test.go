package bankSystem

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	const initialBalance Bitcoin = 20

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(50)
		var want Bitcoin = 50.0

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{initialBalance}
		err := wallet.Withdraw(10)
		want := Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{initialBalance}
		err := wallet.Withdraw(75)

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, initialBalance)
	})

}

// Helpers
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if !errors.Is(got, want) {
		t.Errorf("got %q, want %q", got, want)

	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
