package bank

type defaultBank struct {}

func init () {
	var bank defaultBank
	Register ( "default", bank)
}

func (b defaultBank) Transact(account *Account) error {
   return nil
}
