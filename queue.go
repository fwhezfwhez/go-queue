package Queue

import (
	"fmt"
	"sync"
)

type Queue struct{
	Data []interface{}
}
const (
	DEFAULT_SIZE = 10
)
func NewEmpty() *Queue{
	return &Queue{Data:make([]interface{},0)}
}
func New(size int) *Queue{
	return &Queue{Data:make([]interface{},size)}
}

//a queue's real head
func (q *Queue) Head()(interface{},int){
	if len(q.Data)==0{
		return nil,-1
	}
	return q.Data[0],0
}

//a queue's real tail
func (q *Queue) Tail()(interface{},int){
	if len(q.Data)==0{
		return nil,-1
	}
	return q.Data[len(q.Data)-1],len(q.Data)-1
}

//the first not nil value
func (q *Queue) ValidHead() (interface{},int){
	if len(q.Data)==0{
		return nil,-1
	}
	for i:=0;i<len(q.Data);i++{
		if q.Data[i]!=nil {
			return q.Data[i],i
		}
	}
	return nil,-1
}

//the last not nil value
func (q *Queue) ValidTail() (interface{},int){
	if len(q.Data)==0{
		return nil,-1
	}
	for i:=len(q.Data)-1;i>=0;i--{
		if q.Data[i]!=nil {
			return q.Data[i],i
		}
	}
	return nil,-1
}
//push a value in queue
func (q *Queue) Push(data interface{}){
	if rs,_:=q.ValidTail();rs!=nil||len(q.Data)==0{
		q.Data = append(q.Data,data)
	}else{
		q.Data[len(q.Data)-1] = data
	}
}

//pop a value in queue
func (q *Queue) Pop()interface{}{
	rs,index := q.ValidHead()
	q.Data = q.Data[index+1:]
	return rs
}

//print this queue
func (q *Queue) Print(){
	fmt.Println("<-out",q.Data,"<-in")
}

//push a data  routine safe
func (q *Queue) SafePush(data interface{}){
	m := sync.Mutex{}
	m.Lock()
	defer m.Unlock()
	q.Push(data)
}

//Pop a data routine safe
func (q *Queue) SafePop() interface{}{
	m := sync.Mutex{}
	m.Lock()
	defer m.Unlock()
	return q.Pop()
}

// a queue's real length
func (q *Queue) Length() int{
	return len(q.Data)
}

// trip it's nil value at head and tail and return its valid length
func (q *Queue) ValidLength() int {
	begin:=-1
	end:=-1
	for i:=0;i<len(q.Data);i++{
		if q.Data[i] !=nil{
			begin = i
			break
		}
	}
	if begin == -1{
		return 0
	}
	for i:=len(q.Data)-1;i>=0;i--{
		if q.Data[i]!=nil{
			end = i
			break
		}
	}
	return len(q.Data[begin:end+1])
}