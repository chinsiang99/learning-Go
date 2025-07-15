package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, got string, want string) {
		t.Helper()
		// got := wallet.Balance().String()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit successfully", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance().String()
		want := Bitcoin(10).String()

		assertBalance(t, got, want)
	})

	t.Run("withdraw operations", func(t *testing.T) {
		withdrawTestCases := []struct {
			name        string
			startingBal Bitcoin
			withdrawAmt Bitcoin
			wantBal     Bitcoin
			expectErr   error
		}{
			{name: "withdraw 10 from 80", startingBal: Bitcoin(80), withdrawAmt: Bitcoin(10), wantBal: Bitcoin(70), expectErr: nil},
			{name: "withdraw entire balance", startingBal: Bitcoin(50), withdrawAmt: Bitcoin(50), wantBal: Bitcoin(0), expectErr: nil},
			{name: "withdraw nothing", startingBal: Bitcoin(20), withdrawAmt: Bitcoin(0), wantBal: Bitcoin(20), expectErr: nil},
			{name: "withdraw insufficient funds", startingBal: Bitcoin(10), withdrawAmt: Bitcoin(15), wantBal: Bitcoin(10), expectErr: ErrInsufficientFunds},
		}

		for _, tc := range withdrawTestCases {
			t.Run(tc.name, func(t *testing.T) {
				// since they dont share state, we can run it parallelly
				t.Parallel()
				wallet := Wallet{balance: tc.startingBal}

				err := wallet.Withdraw(tc.withdrawAmt)

				if err != tc.expectErr {
					t.Errorf("got error %q, want %q", err, tc.expectErr)
				}

				got := wallet.Balance().String()
				want := tc.wantBal.String()

				if got != want {
					t.Errorf("got %s want %s", got, want)
				}
			})
		}
	})
}
