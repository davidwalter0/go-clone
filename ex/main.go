package main

import (
	"github.com/davidwalter0/go-clone"

	"fmt"
)

// Z test struct
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
	P    *int
	I    int
	Y    float64
	Zvar Z
	M    map[int]string
	S    []int
	StSl []Z
	In   Inner
}

var Copy = clone.Copy

func main() {
	var i = 1
	var pi = 3
	m := map[int]string{1: "one", 2: "two"}
	slice := []int{0, 1, 2, 3}
	StSl := []Z{Z{I: 1, Y: 3.1415926}, Z{I: 2, Y: 2.71828}}
	var s S = S{P: &pi, I: 1, Y: 3.1415, Zvar: Z{I: 2, Y: 2.71828}, M: m, S: slice, StSl: StSl}
	var y int = Copy(i).(int)
	fmt.Printf("%v %T %p %v %T %p\n", &i, i, &i, &y, y, &y)
	var z = Copy(&s).(S)
	fmt.Printf("%v %T %p %v %T %p\n", &s, s, &s, &z, z, &z)
	fmt.Printf("%v %T %p %v %T %p\n", s.I, s.I, &s.I, z.I, z.I, &z.I)
	fmt.Printf("%v %T %p %v %T %p\n", s.Y, s.Y, &s.Y, z.Y, z.Y, &z.Y)
	fmt.Printf("%v %T %p %v %T %p\n", s.Zvar.I, s.Zvar.I, &s.Zvar.I,
		z.Zvar.I, z.Zvar.I, &z.Zvar.I)
	fmt.Printf("%v %T %p %v %T %p\n", s.Zvar.Y, s.Zvar.Y, &s.Zvar.Y,
		z.Zvar.Y, z.Zvar.Y, &z.Zvar.Y)
	fmt.Printf("%v %T %p %v %T %p\n", s.M, s.M, &s.M,
		z.M, z.M, &z.M)
	fmt.Printf("%v %T %p %v %T %p\n", s.S, s.S, &s.S,
		z.S, z.S, &z.S)
	fmt.Printf("%v %T %p %v %T %p\n", s.StSl, s.StSl, &s.StSl,
		z.StSl, z.StSl, &z.StSl)

	var mz1 MZ = MZ{"one": Z{1, 2}, "two": Z{3, 4}}
	var mz2 MZ = MZ{"three": Z{5, 6}, "four": Z{7, 8}}
	var inner Inner = Inner{SSM: []MZ{mz1, mz2}}
	var qin S = S{I: 1, Y: 3.1415, Zvar: Z{I: 2, Y: 2.71828}, M: m, S: slice, StSl: StSl, In: inner}

	var qout = Copy(&qin).(S)
	fmt.Printf("%v %T %p %v %T %p\n", qin.StSl, qin.StSl, &qin.StSl, qout.StSl, qout.StSl, &qout.StSl)
	fmt.Printf("Input\n%v %T %p\nOutput\n%v %T %p\n", qin.In, qin.In, &qin.In, qout.In, qout.In, &qout.In)
	fmt.Printf("Input\n%v %T %p\nOutput\n%v %T %p\n", qin, qin, &qin, qout, qout, &qout)
	fmt.Printf("Input\n%v %T %p\nOutput\n%v %T %p\n", qin.P, qin.P, &qin.P, qout.P, qout.P, &qout.P)
	fmt.Printf("Input\n%v %T %p\nOutput\n%v %T %p\n", qin.P, qin.P, qin.P, *qout.P, *qout.P, qout.P)
}
