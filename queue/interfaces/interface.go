package interfaces

type QueueInterface interface {
	//Queue folow FIFO method First In First Out
	Enqueue(data int) // to insert a value
	MultipleEnqueue() // to inser multiple values to queue by geting values from user
	Dequeue()         // to delete a value
	DisplayOrder()    // to display all values in the Queue
}
