package constant_test

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	fmt.Println(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
