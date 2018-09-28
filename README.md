# queue
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/fwhezfwhez/go-queue)

queue realization in go

## What is different from 'container/list'
| Property| queue | container/list|
|:----| :---|:--|
| allow raw type| yes | no|
| concurrently safe| yes | no |
| time_out management| yes | no|

## Benchmark
| function| b.N | result|
|:---|:--|:--|
|Push()|10000000|169 ns/op|
|Pop()|	10000000|172 ns/op|
|SafePush()|10000000|229 ns/op|
|SafePop()| 10000000| 234 ns/op|
|TPop()| 5000000 |295 ns/op|
|TPush()| 5000000 | 221 ns/op|
|SafeTPop()| 5000000 | 405 ns/op|
|SafeTPush()| 5000000 | 301 ns/op|

## Start
`go get github/fwhezfwhez/go-queue`

## Example

```go
package  main
import (
	queue "github.com/fwhezfwhez/go-queue"
	"fmt"
)
func main() {
    //初始化,init
    q:= queue.NewEmpty()
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
    tq := queue.TimeQueueWithTimeStep(10*time.Second, 50, 1*time.Nanosecond)
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
```
result
```
<-out [5 4] <-in
5
<-out [4] <-in
1
4<-out [6] <-in
time supervisor starts
init:
<-out [{5 1538111699} {6 1538111699}] <-in
start time spying, data in the queue can stay for 10s
after 5s:
<-out [{5 1538111699} {6 1538111699}] <-in
after 14s
<-out [] <-in
--- PASS: TestReadME (14.00s)
```