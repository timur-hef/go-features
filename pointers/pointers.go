package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

// implementing Stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

type Wallet struct {
	Funds        float64
	BitcoinFunds Bitcoin
}

func (w *Wallet) Deposit(money *float64, bitcoin *Bitcoin) {
	if money != nil {
		w.Funds += *money
	}
	if bitcoin != nil {
		w.BitcoinFunds += *bitcoin
	}
}

func (w *Wallet) Withdraw(money *float64, bitcoin *Bitcoin) error {
	if money != nil {
		if w.Funds < *money {
			return errors.New("not enough money to withdraw")
		}
		w.Funds -= *money
	}
	if bitcoin != nil {
		if w.BitcoinFunds < *bitcoin {
			return errors.New("not enough bitcoins to withdraw")
		}
		w.BitcoinFunds -= *bitcoin
	}

	return nil
}

func (w *Wallet) Balance() float64 {
	return w.Funds
}

func (w *Wallet) BalanceBitcoin() Bitcoin {
	return w.BitcoinFunds
}
