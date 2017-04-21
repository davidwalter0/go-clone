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
	// return ReflectionObjectCopy(from).Interface()
	return ReflectionObjectCopy(from).Interface()
}

// ReflectionObjectCopy from return to duplicate copy of any arbitrary object
func ReflectionObjectCopy(from interface{}) (toVal reflect.Value) {
	fmt.Printf("%v %T\n", from, from)

	// T := reflect.TypeOf(from).Elem()
	// to := reflect.New(T)
	// fromVal := reflect.ValueOf(from).Elem()
	// toVal = reflect.ValueOf(to)

	var T = reflect.TypeOf(from)
	if T.Kind() == reflect.Ptr {
		T = T.Elem()
	}

	fmt.Printf("%v %T\n", T, T)

	toVal = reflect.New(T)
	if toVal.Kind() == reflect.Ptr {
		toVal = toVal.Elem()
	}

	fmt.Printf(">> %v %T\n", toVal, toVal)
	fromVal := reflect.ValueOf(from)
	if fromVal.Kind() == reflect.Ptr {
		fromVal = fromVal.Elem()
	}

	fmt.Printf("%v %T\n", fromVal, fromVal)
	fmt.Printf("** Kind: %v value:[%v] type: %T\n", toVal.Kind(), toVal.Interface(), toVal.Interface())
	// toVal = reflect.ValueOf(to).Elem()
	// fmt.Printf("%v %T\n", toVal, toVal)

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

			if reflect.TypeOf(fromVal.Index(i)).Kind() == reflect.Ptr {
				arg = fromVal.Index(i).Interface()
			} else {
				arg = fromVal.Index(i).Addr().Interface()
			}
			fmt.Printf("arg %v %T\n", arg, arg)
			// os.Exit(1)
			// fmt.Println("asdfasdf", toVal)
			// if reflect.TypeOf(arg).Kind() == reflect.Ptr {
			// 	toVal = reflect.Append(toVal, ReflectionObjectCopy(reflect.ValueOf(arg).Addr().Interface()))
			// } else {
			toVal = reflect.Append(toVal, ReflectionObjectCopy(arg))
			// }
		}
	case reflect.Map:
		toVal = reflect.MakeMap(fromVal.Type())
		valType := fromVal.Type().Elem()
		keyType := fromVal.Type().Key()
		for _, key := range fromVal.MapKeys() {
			keyTo := reflect.New(keyType).Elem()
			valTo := reflect.New(valType).Elem()
			keyTo.Set(ReflectionObjectCopy(key.Interface()))
			value := Pointerize(fromVal.MapIndex(key))
			fmt.Printf("** %v %T %v %T\n", keyTo.Interface(), keyTo.Interface(), value, value)
			valTo.Set(ReflectionObjectCopy(value))
			fmt.Printf("** %v %T %v %T\n", keyTo.Interface(), keyTo.Interface(), valTo.Interface(), valTo.Interface())
			toVal.SetMapIndex(keyTo, valTo)
			fmt.Printf("Map: value %v type: %T\n", toVal.Interface(), toVal.Interface())
		}

	case reflect.Struct:

		if reflect.TypeOf(from).Kind() != reflect.Ptr {
			panic(ErrInvalidArgPointerRequired)
		} else {

			element := fromVal
			fmt.Println("element", element)
			elementType := element.Type()
			for i := 0; i < elementType.NumField(); i++ {
				// structField := elementType.Field(i)
				fmt.Printf("element %v %T %p\n", &element, element, &element)
				// ptr := element.Field(i).Addr().Interface()
				// fmt.Printf("ptr     %v %T %p\n", ptr, ptr, ptr)
				// toVal.Field(i).Set(ReflectionObjectCopy(ptr))
				ptr := Pointerize(element.Field(i))
				fmt.Printf("ptr     %v %T %p\n", ptr, ptr, ptr)
				toVal.Field(i).Set(ReflectionObjectCopy(ptr))
			}
		}
	}
	fmt.Printf("<< %v %T\n", toVal, toVal)
	return toVal
}

// Pointerize return proper interface type
func Pointerize(in reflect.Value) (arg interface{}) {
	fmt.Printf("Pointerize: %v %T %v\n", in, in, in.Kind())
	switch in.Type().Kind() {
	case reflect.Ptr:
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
		arg = in.Interface()
	case reflect.Map, reflect.Struct, reflect.Slice:
		arg = in.Addr().Interface()
	default:

		arg = in.Addr().Interface()
	}
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
	I           int
	Y           float64
	Zvar        Z
	M           map[int]string
	S           []int
	StructSlice []Z
	// In          Inner
}

func main() {
	var i = 1
	// var mz1 MZ = MZ{"one": Z{1, 2}, "two": Z{3, 4}}
	// var mz2 MZ = MZ{"three": Z{5, 6}, "four": Z{7, 8}}
	// var inner Inner = Inner{SSM: []MZ{mz1, mz2}}

	m := map[int]string{1: "one", 2: "two"}
	slice := []int{0, 1, 2, 3}
	StructSlice := []Z{Z{I: 1, Y: 3.1415926}, Z{I: 2, Y: 2.71828}}
	var s S = S{I: 1, Y: 3.1415, Zvar: Z{I: 2, Y: 2.71828}, M: m, S: slice, StructSlice: StructSlice}
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
	// fmt.Printf("%v %T %p %v %T %p\n", s.In, s.In, &s.In,
	// 	z.In, z.In, &z.In)
}
