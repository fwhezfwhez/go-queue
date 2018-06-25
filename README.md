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
	//初始化
	q:= queue.NewEmpty()
	//压入
	q.Push(5)
	q.Push(4)
	//打印
	q.Print()
	//出列
	fmt.Println(q.Pop())
	//打印
	q.Print()
	//长度
	fmt.Println(q.Length())
	//并发安全压入
	q.SafePush(6)
	//并发安全出列
	fmt.Print(q.SafePop())
	q.Print()
}
```
