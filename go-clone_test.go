package clone

import (
	"testing"
)

var copy = Copy

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

func Test_Bool(t *testing.T) {
	var src = true
	dst := copy(src).(bool)
	compareBool(t, src, dst)
}

func Test_Int(t *testing.T) {
	var src = 2
	dst := copy(src).(int)
	compareInt(t, src, dst)
}

func Test_Float(t *testing.T) {
	var src = 2.71828
	dst := copy(src).(float64)
	compareFloat(t, src, dst)
}

func Test_String(t *testing.T) {
	var src = "string"
	dst := copy(src).(string)
	compareString(t, src, dst)
}

func Test_BoolSlice(t *testing.T) {
	type T []bool
	var src = T{true, false, true}
	var dst = copy(src).(T)
	compareBoolSlice(t, src, dst)
}

func Test_MapStringBool(t *testing.T) {
	type T map[string]bool
	var src = T{"a": true, "b": false, "c": true}
	var dst = copy(src).(T)
	compareMapStringBool(t, src, dst)
}

func Test_MapOfZStruct(t *testing.T) {
	type T map[string]Z
	StSlice := []Z{Z{I: 1, Y: 3.1415926}, Z{I: 2, Y: 2.71828}}
	var src = T{"a": StSlice[0], "b": StSlice[1]}
	var dst = copy(src).(T)
	compareMapOfZStruct(t, src, dst)
}

func Test_InnerStruct(t *testing.T) {
	var mz1 = MZ{"one": Z{1, 2}, "two": Z{3, 4}}
	var mz2 = MZ{"three": Z{5, 6}, "four": Z{7, 8}}
	var src = Inner{SSM: []MZ{mz1, mz2}}
	dst := copy(src).(Inner)
	compareInnerStruct(t, src, dst)
}

func Test_Struct(t *testing.T) {
	var StSl = []Z{Z{I: 1, Y: 3.1415926}, Z{I: 2, Y: 2.71828}}
	var slice = []int{0, 1, 2, 3}
	var m = map[int]string{1: "one", 2: "two"}
	var mz1 = MZ{"one": Z{1, 2}, "two": Z{3, 4}}
	var mz2 = MZ{"three": Z{5, 6}, "four": Z{7, 8}}
	var inner = Inner{SSM: []MZ{mz1, mz2}}
	var qin = S{I: 1, Y: 3.1415, Zvar: Z{I: 2, Y: 2.71828}, M: m, S: slice, StSl: StSl, In: inner}
	var qout = Copy(&qin).(S)
	compareInt(t, qin.I, qout.I)
	compareFloat(t, qin.Y, qout.Y)
	compareZ(t, qin.Zvar, qout.Zvar)
	compareMapIntString(t, qin.M, qout.M)
	compareIntSlice(t, qin.S, qout.S)
	compareZSlice(t, qin.StSl, qout.StSl)
	compareInnerStruct(t, qin.In, qout.In)
}
