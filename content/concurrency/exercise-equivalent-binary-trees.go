// +build no-build OMIT

package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)


// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t.Left!=nil {
		Walk(t.Left, ch)
	}
	ch<-t.Value
	if t.Right!=nil {
		Walk(t.Right, ch)
	}
}


// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	return true
}

func main() {
	ch:=make(chan int,10)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
