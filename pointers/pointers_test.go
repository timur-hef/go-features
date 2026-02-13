package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	money := 10.0

	wallet.Deposit(&money, nil)

	got := wallet.Balance()
	want := float64(10)

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}
func TestWalletBitcoin(t *testing.T) {
	t.Run("deposit bitcoins", func(t *testing.T) {
		wallet := Wallet{}
		bitcoins := Bitcoin(10)
		wallet.Deposit(nil, &bitcoins)

		got := wallet.BalanceBitcoin()
		want := Bitcoin(10.0)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw bitcoins", func(t *testing.T) {
		wallet := Wallet{BitcoinFunds: Bitcoin(20)}
		bitcoins := Bitcoin(10)
		wallet.Withdraw(nil, &bitcoins)

		got := wallet.BalanceBitcoin()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{BitcoinFunds: Bitcoin(10)}
		bitcoins := Bitcoin(20)
		err := wallet.Withdraw(nil, &bitcoins)
		fmt.Print(err)

		if err == nil {
			t.Error("withdraw insufficient money")
		}
	})
}
