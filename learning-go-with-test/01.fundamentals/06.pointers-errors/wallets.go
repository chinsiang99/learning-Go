package pointers

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(money Bitcoin) {
	w.balance = w.balance + money
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
