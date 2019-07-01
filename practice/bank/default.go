package bank


// an empty struct, should never need this default bank
// empty structs use zero bytes when created,
// use default structs when you need a type, but not any state
type defaultBank struct {}


// init methods get called when the package is imported
func init () {
	var bank defaultBank
	Register ( "default", bank)
}

// implement the interface but don't do anything
// we just need this so that it can go into the map of Banker
// in general you want to associate a method to a pointer of the type,
// so that you can change the objects state, but since this is
// a defaultBank with has zero bytes and no state, using a
// value reference is preferred because there are no states to change
func (b defaultBank) Transact(account *Account) error {
   return nil
}


