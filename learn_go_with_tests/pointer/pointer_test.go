package pointer

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s balance %s", got, want)
		}
	}
	assertError := func(t *testing.T, err error, want error) {
		if err == nil {
			t.Fatal("want an error but didn't get one")
		}
		if err != want {
			t.Errorf("got err '%s', want '%s'", err, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw too much money", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, InsufficientFundsError)
	})

}
