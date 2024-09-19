package wallet

import "fmt"


type Bitcoin float64


// Implementando a interface Stringer (define como a struct sera printada)
// podemos usar %s nos t.Errorf agr passando a struct inteira
func (b Bitcoin) String() string {
  return fmt.Sprintf("%.2f BTC", b)
}


type Wallet struct {
  balance Bitcoin
}

// Precisamos usar ponteiros para alterar o estado dessa struct
func (w *Wallet) Deposit(value Bitcoin) {
  w.balance += value
}

func (w *Wallet) Balance() Bitcoin {
  return w.balance
}
