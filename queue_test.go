package Queue

import (
	"testing"
	"fmt"
)

var q *Queue
func Init(){
	q=NewEmpty()
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
	t.Log(q.Pop())
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
func TestQueue_SafePush(t *testing.T) {
	Init()
	q.Push(3)
	q.SafePush(4)
	q.SafePush(5)
	q.Print()
	fmt.Println(q.Pop())
	fmt.Println(q.SafePop())
	q.Print()
}