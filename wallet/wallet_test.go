package wallet

import "testing"


func TestWallet(t *testing.T) {
  wallet := Wallet{}
  wallet.Deposit(10)

  got := wallet.Balance()
  want := Bitcoin(10.0)

  if got != want {
    t.Errorf("Expect %s but got %s", want, got)
  }

}
