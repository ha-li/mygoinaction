package main

import (
	"github.com/mygoinaction/practice/bank"
)

func main () {
	//var d defaultBank
	d := bank.GetBanker("default")
	account := bank.Account {
		"00128",
		"system",
	}
	d.Transact(&account)

	// since the interface is on a value, need
	// to dereference the pointer
	dp := bank.GetBankerPointer("default")
	(*dp).Transact(&account)

	//
    bb := new(bank.Bank)
	(*bb).Transact(&account)

}
