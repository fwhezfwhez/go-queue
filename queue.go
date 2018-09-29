package Queue

import (
	"errorX"
	"fmt"
	"sync"
	"time"
)

const (
	DEFAULT_SIZE = 10
)

type Queue struct {
	Data  []interface{}
	Mutex sync.Mutex

	timeSpy     bool //whether drop data out of time. timeSpy should be set only at beginning and unmodifiable
	ExpireAfter time.Duration
	flag        chan int //use this to stop time spy for a time queue
	cap         int  //queue extends minimum step
	timeStep    time.Duration  //how frequency is the spy routine runs by default 10s
}
type TimeWrapper struct {
	Data      interface{}
	CreatedAt int64
}

// cap refers to its expand scope.When array of data is full,the array will automatically expand to add a cap amount
func TimeQueue(expireAfter time.Duration, cap int) *Queue {
	return &Queue{
		Data:        make([]interface{}, 0, cap),
		timeSpy:     true,
		ExpireAfter: expireAfter,
		cap : cap,
		timeStep: 10*time.Second,
	}
}

// cap refers to its expand scope.When array of data is full,the array will automatically expand to add a cap amount
func TimeQueueWithTimeStep(expireAfter time.Duration, cap int, tsp time.Duration) *Queue {
	return &Queue{
		Data:        make([]interface{}, 0, cap),
		timeSpy:     true,
		ExpireAfter: expireAfter,
		cap : cap,
		timeStep: tsp,
	}
}

func NewEmpty() *Queue {
	return &Queue{Data: make([]interface{}, 0, DEFAULT_SIZE)}
}
func New(size int) *Queue {
	return &Queue{Data: make([]interface{}, size, 2*size)}
}

func NewCap(cap int) *Queue {
	return &Queue{Data: make([]interface{}, 0, cap),cap:cap}
}

//a queue's real head
func (q *Queue) Head() (interface{}, int) {
	if len(q.Data) == 0 {
		return nil, -1
	}
	if !q.timeSpy {
		return q.Data[0], 0
	} else {
		wrapper, index := q.Data[0], 0
		return wrapper.(TimeWrapper).Data, index
	}
}

//a queue's real head
func (q *Queue) SafeHead() (interface{}, int) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	return q.Head()
}

//a queue's real tail
func (q *Queue) Tail() (interface{}, int) {
	if len(q.Data) == 0 {
		return nil, -1
	}
	if !q.timeSpy {
		return q.Data[len(q.Data)-1], len(q.Data) - 1
	} else {
		wrapper, index := q.Data[len(q.Data)-1], len(q.Data)-1
		return wrapper.(TimeWrapper).Data, index
	}
}

func (q *Queue) SafeTail() (interface{}, int) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	return q.Tail()
}

//the first not nil value
func (q *Queue) ValidHead() (interface{}, int) {
	if len(q.Data) == 0 {
		return nil, -1
	}
	for i := 0; i < len(q.Data); i++ {
		if q.Data[i] != nil {

			if !q.timeSpy {
				return q.Data[i], i
			} else {
				return q.Data[i].(TimeWrapper).Data, i
			}

		}
	}
	return nil, -1
}

func (q *Queue) SafeValidHead()(interface{}, int){
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	return q.ValidHead()
}

func (q *Queue) THead() (*TimeWrapper, int, error) {
	if len(q.Data) == 0 {
		return nil, -1, errorx.NewFromString("empty queue")
	}

	if q.timeSpy {
		tw := q.Data[0].(TimeWrapper)
		return &tw, 0, nil
	} else {
		return nil, -1, errorx.NewFromString("THead only use in a time queue,got by TimeQueue(time.Duration,int)")
	}
}

func (q *Queue) SafeTHead()(interface{}, int, error){
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	return q.THead()
}

//the last not nil value
func (q *Queue) ValidTail() (interface{}, int) {
	if len(q.Data) == 0 {
		return nil, -1
	}
	for i := len(q.Data) - 1; i >= 0; i-- {
		if q.Data[i] != nil {
			if !q.timeSpy {
				return q.Data[i], i
			} else {
				return q.Data[i].(TimeWrapper).Data, i
			}
		}
	}
	return nil, -1
}

//push a value in queue
func (q *Queue) Push(data interface{}) {
	if q.timeSpy {
		data = TimeWrapper{
			Data:      data,
			CreatedAt: time.Now().Unix(),
		}
	}
	if rs, _ := q.ValidTail(); rs != nil || len(q.Data) == 0 {
		q.Data = append(q.Data, data)
	} else {
		q.Data[len(q.Data)-1] = data
	}
	_, i := q.ValidHead()
	q.Data = q.Data[i:]
}

func (q *Queue) TPush(data interface{}){
	q.Push(data)
}

//pop a value in queue
func (q *Queue) Pop() interface{} {
	rs, index := q.ValidHead()
	if len(q.Data)>1 {
		q.Data = q.Data[index+1:]
	}else{
		q.Data = make([]interface{},0, q.cap)
	}

	return rs
}

// Pop a timewarpper for a time queue
func (q *Queue) TPop() (*TimeWrapper, int, error) {
	if !q.timeSpy {
		return nil, -1, errorx.NewFromString("q'timeSpy is false,make sure queue from TimeQueue(time.Duration.int) ")
	}
	rs, index, er := q.THead()
	if er != nil {
		return nil, -1, errorx.Wrap(er)
	}

	if len(q.Data)>1 {
		q.Data = q.Data[index+1:]
	}else{
		q.Data = make([]interface{},0, q.cap)
	}

	return rs, index, nil
}
func (q *Queue) InversePop() interface{} {
	rs, index := q.ValidTail()
	q.Data = q.Data[:index]
	return rs
}

//print this queue
func (q *Queue) Print() {
	fmt.Println("<-out", q.Data, "<-in")
}

//push a data  routine safe
func (q *Queue) SafePush(data interface{}) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.Push(data)
}

//push a data  into time queue routine safe
func (q *Queue) SafeTPush(data interface{}) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	q.TPush(data)
}

//Pop a data routine safe
func (q *Queue) SafePop() interface{} {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	return q.Pop()
}


//Pop a data from a time queue routine safe
func (q *Queue) SafeTPop() (*TimeWrapper,int,error) {
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	return q.TPop()
}

// a queue's real length
func (q *Queue) Length() int {
	return len(q.Data)
}

func (q *Queue) SafeLength() int{
	q.Mutex.Lock()
	defer q.Mutex.Unlock()
	return len(q.Data)
}

// trip it's nil value at head and tail and return its valid length
func (q *Queue) ValidLength() int {
	begin := -1
	end := -1
	for i := 0; i < len(q.Data); i++ {
		if q.Data[i] != nil {
			begin = i
			break
		}
	}
	if begin == -1 {
		return 0
	}
	for i := len(q.Data) - 1; i >= 0; i-- {
		if q.Data[i] != nil {
			end = i
			break
		}
	}
	return len(q.Data[begin : end+1])
}

// start to spy on queue's time-out data and throw it
func (q *Queue) StartTimeSpying() {
	fmt.Println("time supervisor starts")
	go q.startTimeSpying()
}

// detail of StartTimeSpying function
func (q *Queue) startTimeSpying() error {
	var err = make(chan string, 0)
	go func(queue *Queue, er chan string) {
		fmt.Println("start time spying, data in the queue can stay for " + q.ExpireAfter.String())
		for {
			if queue.timeSpy == false {
				err <- "spying routine stops because: queue's timeSpy is false, make sure the queue is definition by q=TimeQueue(time.Duration,int)"
				return
			}
			select {
			case <-queue.flag:
				fmt.Println("time spy executing stops")
				return
			default:
				fmt.Print()
			}
			ok,er:=queue.timingRemove()
			if er!=nil{
				err <- er.(errorx.Error).StackTrace()
			}
			if ok {
				time.Sleep(queue.timeStep)
			}
		}
	}(q, err)
	select {
	case msg := <-err:
		fmt.Println("time spy supervisor accidentally stops because: ",msg)
		return errorx.NewFromString(msg)
	case <-q.flag:
		fmt.Println("time spy supervisor stops")
		return nil
	}
}

//stop supervisor and execution of time spying
func (q *Queue) StopTimeSpying() {
	close(q.flag)
}

// remove those time-out data
func (q *Queue) timingRemove() (bool,error) {
	if len(q.Data) <1 {
		return true,nil
	}
	head, index, er := q.THead()
	if er != nil {
		return false, errorx.Wrap(er)
	}
	if index < 0 {
		return false, errorx.NewFromString("queue'length goes 0")
	}
	now := time.Now().Unix()
	created := time.Unix(head.CreatedAt, 0)
	//fmt.Println("now:",now)
	//fmt.Println("expire:",created.Add(q.ExpireAfter).Unix())
	if created.Add(q.ExpireAfter).Unix() < now {
		// out of time
		_,_,e := q.TPop()
		if e!=nil {
			return false, errorx.Wrap(e)
		}
		if len(q.Data) >0 {
			return q.timingRemove()
		}else{
			return true,nil
		}
	} else{
		return true ,nil
	}
}
