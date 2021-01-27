package pointer

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	balance := 10
	if got != balance {
		t.Errorf("got %d balance %d", got, balance)
	}
}
