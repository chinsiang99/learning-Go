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
		withdrawTestCases := []struct {
			name        string
			startingBal Bitcoin
			withdrawAmt Bitcoin
			wantBal     Bitcoin
		}{
			{name: "withdraw 10 from 80", startingBal: Bitcoin(80), withdrawAmt: Bitcoin(10), wantBal: Bitcoin(70)},
			{name: "withdraw entire balance", startingBal: Bitcoin(50), withdrawAmt: Bitcoin(50), wantBal: Bitcoin(0)},
			{name: "withdraw nothing", startingBal: Bitcoin(20), withdrawAmt: Bitcoin(0), wantBal: Bitcoin(20)},
		}

		for _, tc := range withdrawTestCases {
			t.Run(tc.name, func(t *testing.T) {
				wallet := Wallet{balance: tc.startingBal}

				wallet.Withdraw(tc.withdrawAmt)

				got := wallet.Balance().String()
				want := tc.wantBal.String()

				if got != want {
					t.Errorf("got %s want %s", got, want)
				}
			})
		}
	})
}
