package transaction

import (
	"fmt"
	"github.com/mygoinaction/practice/bank"
)

type Transaction struct {
	To       bank.Account
    From     bank.Account
	Amount   float64
}

func Transact(amount float64) {

	fmt.Println("doing a transaction")
	defaultBank := bank.GetBanker("default")
	boa := bank.GetBanker("boa")

	fmt.Println (defaultBank)
	fmt.Println (boa)
}