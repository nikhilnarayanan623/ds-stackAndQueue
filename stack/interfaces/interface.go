package interfaces

type StackInterface interface {
	// Stack folow LIFO Last In First Out
	Push(data int) // to insert a new value to the Stack
	MultiplePush() // to push multiple vlues to stack from user data
	Pop()          // to remove data from Stack
	DisplayOrder() // to display all vlaues in the stack
}
