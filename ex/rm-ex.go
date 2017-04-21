package main

import (
	"github.com/davidwalter0/go-deepcopy"

	"fmt"
)

type Z struct {
	I int
	Y float64
}

type MZ map[string]Z

type Inner struct {
	SSM []MZ
}

// S test struct
type S struct {
	I    int
	Y    float64
	Zvar Z
	M    map[int]string
	S    []int
	StSl []Z
	In   Inner
}

var Copy = deepcopy.Copy

func main() {
	m := map[int]string{1: "one", 2: "two"}
	slice := []int{0, 1, 2, 3}
	StSl := []Z{Z{I: 1, Y: 3.1415926}, Z{I: 2, Y: 2.71828}}
	var mz1 = MZ{"one": Z{1, 2}, "two": Z{3, 4}}
	var mz2 = MZ{"three": Z{5, 6}, "four": Z{7, 8}}
	var inner = Inner{SSM: []MZ{mz1, mz2}}
	var src = S{I: 1, Y: 3.1415, Zvar: Z{I: 2, Y: 2.71828}, M: m, S: slice, StSl: StSl, In: inner}

	var dst = Copy(&src).(S)
	fmt.Printf("%v %T %p %v %T %p\n", src.StSl, src.StSl, &src.StSl, dst.StSl, dst.StSl, &dst.StSl)
	fmt.Printf("Input\n%v %T %p\nOutput\n%v %T %p\n", src.In, src.In, &src.In, dst.In, dst.In, &dst.In)
	fmt.Printf("Input\n%v %T %p\nOutput\n%v %T %p\n", src, src, &src, dst, dst, &dst)
}
