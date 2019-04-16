package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timespent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timespent(slowFun)
	t.Log(tsSF(10))
}