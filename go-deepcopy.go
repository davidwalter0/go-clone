package main

import (
	// "encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
)

func init() {
	if false {
		fmt.Println()
		log.Println()
		os.Exit(0)
	}
}

var ErrInvalidArgPointerRequired = errors.New("Argument must be a pointer")

// Copy deep recursive copy of object
func Copy(from interface{}) (to interface{}) {
	return RecursiveDeepCopy(from).Interface()
}

// RecursiveDeepCopy from return to duplicate copy of any arbitrary object
func RecursiveDeepCopy(from interface{}) (toVal reflect.Value) {
	if from == nil {
		return toVal
	}

	var T = reflect.TypeOf(from)
	if T.Kind() == reflect.Ptr {
		T = T.Elem()
	}

	toVal = reflect.New(T)
	if toVal.Kind() == reflect.Ptr {
		toVal = toVal.Elem()
	}

	fromVal := reflect.ValueOf(from)
	fmt.Printf("fromVal  : %v %T %v %v\n", fromVal, fromVal, fromVal, fromVal.Kind())
	if fromVal.Kind() == reflect.Ptr {
		if fromVal.IsNil() {
			fromVal = reflect.New(T)
		}
		fromVal = fromVal.Elem()
	}
	fmt.Printf("fromVal  : %v %T %v %v\n", fromVal, fromVal, fromVal, fromVal.Kind())

	switch toVal.Kind() {
	case reflect.String:
		toVal.SetString(fromVal.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		toVal.SetInt(fromVal.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		toVal.SetUint(fromVal.Uint())
	case reflect.Bool:
		toVal.SetBool(fromVal.Bool())
	case reflect.Float32, reflect.Float64:
		toVal.SetFloat(fromVal.Float())
	case reflect.Slice:
		toVal = reflect.MakeSlice(fromVal.Type(), 0, 0)
		var arg interface{}
		for i := 0; i < fromVal.Len(); i++ {
			arg = Pointerize(fromVal.Index(i))
			toVal = reflect.Append(toVal, RecursiveDeepCopy(arg))
		}
	case reflect.Map:
		toVal = reflect.MakeMap(fromVal.Type())
		valType := fromVal.Type().Elem()
		keyType := fromVal.Type().Key()
		for _, key := range fromVal.MapKeys() {
			keyTo := reflect.New(keyType).Elem()
			valTo := reflect.New(valType).Elem()
			keyTo.Set(RecursiveDeepCopy(Pointerize(key)))
			tmp := fromVal.MapIndex(key)
			// fmt.Printf("Debugging %v %T %v %T %v\n",
			// 	keyTo.Interface(), keyTo.Interface(), tmp, tmp, tmp.Kind())
			value := Pointerize(tmp)
			// fmt.Printf("Debugging %v %T %v %T\n",
			// 	keyTo.Interface(), keyTo.Interface(), value, value)
			valTo.Set(RecursiveDeepCopy(value))
			// fmt.Printf("** %v %T %v %T\n", keyTo.Interface(), keyTo.Interface(), valTo.Interface(), valTo.Interface())
			toVal.SetMapIndex(keyTo, valTo)
			// fmt.Printf("Map: value %v type: %T\n", toVal.Interface(), toVal.Interface())
		}

	case reflect.Struct:
		element := fromVal
		elementType := element.Type()
		for i := 0; i < elementType.NumField(); i++ {
			fmt.Printf("** %v %T %v\n", element.Field(i), element.Field(i), element.Field(i).Kind())
			field := Pointerize(element.Field(i))
			set := toVal.Field(i)
			fmt.Printf("field: %v %T %v\n", field, field, field)
			fmt.Printf("set  : %v %T %v %v\n", set, set, set, set.Kind())
			if set.IsValid() {
				if set.Kind() == reflect.Ptr && set.IsNil() && set.CanSet() {
					// toVal.Field(i).Set(RecursiveDeepCopy(field))
					v := RecursiveDeepCopy(field)
					set.Set(reflect.ValueOf(v.Addr().Interface()))
				} else {
					// toVal.Field(i).Set(RecursiveDeepCopy(field))
					set.Set(RecursiveDeepCopy(field))
				}
			}
		}
	}
	// fmt.Printf("<< %v %T\n", toVal, toVal)
	return toVal
}

// Pointerize return proper interface type
func Pointerize(in reflect.Value) (arg interface{}) {
	fmt.Printf("** %v %T %v\n", in, in, in)
	switch in.Type().Kind() {
	case reflect.Ptr:
		// arg = in.Elem().Interface()
		arg = in.Interface()
	case reflect.String:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fallthrough
	case reflect.Bool:
		fallthrough
	case reflect.Float32, reflect.Float64:
		fallthrough
	case reflect.Map, reflect.Struct, reflect.Slice:
		arg = in.Interface()
	default:
		arg = in.Addr().Interface()
	}
	fmt.Printf("** %v %T %v\n", arg, arg, arg)
	fmt.Printf("** %v %T %v\n", in, in, in)
	return
}

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
	P           *int
	I           int
	Y           float64
	Zvar        Z
	M           map[int]string
	S           []int
	StructSlice []Z
	In          Inner
}

func main() {
	var i = 1
	// var mz1 MZ = MZ{"one": Z{1, 2}, "two": Z{3, 4}}
	// var mz2 MZ = MZ{"three": Z{5, 6}, "four": Z{7, 8}}
	// var inner Inner = Inner{SSM: []MZ{mz1, mz2}}
	var pi = 3
	m := map[int]string{1: "one", 2: "two"}
	slice := []int{0, 1, 2, 3}
	StructSlice := []Z{Z{I: 1, Y: 3.1415926}, Z{I: 2, Y: 2.71828}}
	var s S = S{P: &pi, I: 1, Y: 3.1415, Zvar: Z{I: 2, Y: 2.71828}, M: m, S: slice, StructSlice: StructSlice}
	// var s S = S{I: 1, Y: 3.1415, Zvar: Z{I: 2, Y: 2.71828}, M: m, S: slice, StructSlice: StructSlice, In: inner}
	// var s S = S{1, 3.1415, Z{2, 2.71828}}
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
	fmt.Printf("%v %T %p %v %T %p\n", s.StructSlice, s.StructSlice, &s.StructSlice,
		z.StructSlice, z.StructSlice, &z.StructSlice)

	var mz1 MZ = MZ{"one": Z{1, 2}, "two": Z{3, 4}}
	var mz2 MZ = MZ{"three": Z{5, 6}, "four": Z{7, 8}}
	var inner Inner = Inner{SSM: []MZ{mz1, mz2}}

	// m := map[int]string{1: "one", 2: "two"}
	// slice := []int{0, 1, 2, 3}
	// StructSlice := []Z{Z{I: 1, Y: 3.1415926}, Z{I: 2, Y: 2.71828}}
	var qin S = S{I: 1, Y: 3.1415, Zvar: Z{I: 2, Y: 2.71828}, M: m, S: slice, StructSlice: StructSlice, In: inner}

	var qout = Copy(&qin).(S)
	fmt.Printf("%v %T %p %v %T %p\n", qin.StructSlice, qin.StructSlice, &qin.StructSlice, qout.StructSlice, qout.StructSlice, &qout.StructSlice)
	fmt.Printf("Input\n%v %T %p\nOutput\n%v %T %p\n", qin.In, qin.In, &qin.In, qout.In, qout.In, &qout.In)
	fmt.Printf("Input\n%v %T %p\nOutput\n%v %T %p\n", qin, qin, &qin, qout, qout, &qout)
	fmt.Printf("Input\n%v %T %p\nOutput\n%v %T %p\n", qin.P, qin.P, &qin.P, qout.P, qout.P, &qout.P)
	fmt.Printf("Input\n%v %T %p\nOutput\n%v %T %p\n", qin.P, qin.P, qin.P, *qout.P, *qout.P, qout.P)
}
