package pointers

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance().String()
	want := Bitcoin(10).String()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
