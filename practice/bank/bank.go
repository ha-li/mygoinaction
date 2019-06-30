package bank

import (
	"fmt"
)

type Account struct {
	Number string
	Holder string
}

type Bank struct {
	Name string
	Accounts []Account
}

type Banker interface {
	Transact(account *Account) error
}

// a map of register bankers for managing your account
var banks = make (map[string]Banker)

func (b Bank) Transact(account *Account) error {
	fmt.Println("doing a transaciton")
	return nil
}

// register a bank into the map of all banks
func Register(bankName string, bank Banker) {
    if _, exists := banks[bankName]; exists {
    	fmt.Println("Bank already exists")
	} else {
		banks[bankName] = bank
	}
}

func GetBanker(bank string) Banker {
	return banks[bank]
}