package Queue

import "testing"

var q *Queue
func Init(){
	q=New(2)
	//q.Data[0]=1
	//q.Data[1]=2
	//q.Push(3)
}
func TestQueue_Tail(t *testing.T) {
	Init()
	t.Log(q.Tail())
	t.Log(q.Data)
}
func TestQueue_Head(t *testing.T) {
	Init()
	t.Log(q.Head())
	t.Log(q.Data)
}

func TestQueue_ValidTail(t *testing.T) {
	Init()
	t.Log(q.ValidTail())
	t.Log(q.Data)
}
func TestQueue_ValidHead(t *testing.T) {
	Init()
	t.Log(q.ValidHead())
	t.Log(q.Data)
}
func TestQueue_Pop(t *testing.T){
	Init()
	q.Pop()
	t.Log(q.Data)
}

func TestQueue_Print(t *testing.T) {
	Init()
	q.Push(2)
	q.Push(3)
	q.Print()
	q.Pop()
	q.Print()
	q.Pop()
	q.Print()
	q.Pop()
	q.Print()
	q.Pop()
}

func TestQueue_Length(t *testing.T) {
	Init()
	q.Push(3)
	q.Push(nil)
	q.Print()
	t.Log(q.Length())
	t.Log(q.ValidLength())
}