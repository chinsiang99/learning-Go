package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit successfully", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance().String()
		want := Bitcoin(10).String()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw successfully", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(80)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance().String()
		want := Bitcoin(70).String()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
