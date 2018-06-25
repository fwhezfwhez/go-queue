package Queue
type QueueI interface{
	//the first element of the queue
	Head()(interface{},int)
	//the last element of the queue
	Tail()(interface{},int)
	//the first not nil element
	ValidHead() (interface{},int)
	//the last not nil element
	ValidTail() (interface{},int)
	//push or add an element to a queue's tail
	Push(data interface{})
	//pop an element from head
	Pop()interface{}
	//pop an element from tail
	InversePop()interface{}
	//print the elements in queue
	Print()
	//push an element in queue's tail thread/routine safe
	SafePush(data interface{})
	//pop an element in queue's head thread/routine safe
	SafePop() interface{}
	//a queue's real length
	Length() int
	//queue's length after trim nil
	ValidLength() int
}