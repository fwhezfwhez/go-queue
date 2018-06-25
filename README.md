a good queue realization in go

api list is introduced in interface.go

**start**

go get github/fwhezfwhez/go-queue

**example**

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
}
```
