package bank

func init () {
	check := Account {
		"0001",
		"boa",
	}
	accounts := []Account {check}

	boa := Bank {
		"boa",
		 accounts,
	}

	Register ( "boa", boa)
}
