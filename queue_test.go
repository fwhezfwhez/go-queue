package Queue

import (
	"fmt"
	"testing"
	"time"
)

var q *Queue

func Init() {
	//q=New(5)
	q = TimeQueue(10*time.Second, 500)

	//q.Data[0]=1
	//q.Data[1]=2
	q.Push(3)
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
	q.Pop()
}
func TestQueue_ValidHead(t *testing.T) {
	Init()
	t.Log(q.ValidHead())
	t.Log(q.Data)
}
func TestQueue_Pop(t *testing.T) {
	Init()
	t.Log(q.Pop())
	t.Log(q.Data)
}

func TestQueue_Push(t *testing.T) {
	Init()
	q.Push(5)

	q.Push(8)
	q.Push(8)
	q.Push(10)
	q.Print()
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
func TestQueue_InversePop(t *testing.T) {
	Init()
	q.Push(3)
	q.Push(4)
	q.Push(5)
	//q.Push(nil)
	t.Log(q.InversePop())
	q.Print()
}

func TestTimeRemoveQueue(t *testing.T) {
	tq := TimeQueue(2*time.Second, 50)
	tq.Push(5)
	time.Sleep(1 * time.Second)
	tq.Push(6)
	tq.Print()

	time.Sleep(4 * time.Second)
	tq.timingRemove()
	tq.Print()
}

func TestTimeQueue(t *testing.T) {
	tq := TimeQueue(10*time.Second, 50)
	tq.StartTimeSpying()
	tq.Push(5)
	tq.Push(6)
	tq.Push(6)

	fmt.Println("init:")
	tq.Print()

	time.Sleep(5 * time.Second)
	fmt.Println("after 5s:")
	tq.Print()

	time.Sleep(9 * time.Second)
	fmt.Println("after 14s")
	tq.Print()
}

func TestReadME(t *testing.T) {
	//初始化,init
	q:= NewEmpty()
	//压入,push
	q.Push(5)
	q.Push(4)
	//打印,print
	q.Print()
	//出列,pop
	fmt.Println(q.Pop())
	//打印,print
	q.Print()
	//长度,len
	fmt.Println(q.Length())
	//并发安全压入,currently safe push
	q.SafePush(6)
	//并发安全出列,currently safe pop
	fmt.Print(q.SafePop())
	q.Print()

	// time queue
	tq := TimeQueueWithTimeStep(10*time.Second, 50, 1*time.Nanosecond)
	tq.StartTimeSpying()
	tq.TPush(5)
	tq.SafeTPush(6)

	fmt.Println("init:")
	tq.Print()

	time.Sleep(5 * time.Second)
	fmt.Println("after 5s:")
	tq.Print()

	time.Sleep(9 * time.Second)
	fmt.Println("after 14s")
	tq.Print()
}

func BenchmarkQueue_Push(b *testing.B) {
	q := NewCap(0)
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
}
func BenchmarkQueue_Pop(b *testing.B) {
	q := NewCap(0)
	for i := 0; i <b.N; i++ {
		q.Push(i)
	}
	for i:=0;i<b.N;i++{
		q.Pop()
	}
}
func BenchmarkQueue_SafePush(b *testing.B) {
	q := NewCap(5000)
	for i := 0; i < b.N; i++ {
		q.SafePush(i)
	}
}

func BenchmarkQueue_SafePop(b *testing.B) {
	q := NewCap(0)
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	for i:=0;i<b.N;i++{
		q.SafePop()
	}
}

func BenchmarkQueue_TPop(b *testing.B) {
	q := TimeQueue(10*time.Second,1000)
	for i := 0; i <b.N; i++ {
		q.TPush(i)
	}
	for i:=0;i<b.N;i++{
		q.TPop()
	}
}

func BenchmarkQueue_TPush(b *testing.B) {
	q := TimeQueue(10*time.Second,10)
	for i := 0; i <b.N; i++ {
		q.TPush(i)
	}
}

func BenchmarkQueue_SafeTPop(b *testing.B) {
	q := TimeQueue(10*time.Second,1000)
	for i := 0; i <b.N; i++ {
		q.SafeTPush(i)
	}
	for i:=0;i<b.N;i++{
		q.SafeTPop()
	}
}

func BenchmarkQueue_SafeTPush(b *testing.B) {
	q := TimeQueue(10*time.Second,1)
	for i := 0; i <b.N; i++ {
		q.SafeTPush(i)
	}
}