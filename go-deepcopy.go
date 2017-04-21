package deepcopy

import (
	// "encoding/json"
	// "errors"
	"fmt"
	// "log"
	// "os"
	"reflect"
)

// func init() {
// 	if false {
// 		fmt.Println()
// 		log.Println()
// 		os.Exit(0)
// 	}
// }
var debug = false

// var ErrInvalidArgPointerRequired = errors.New("Argument must be a pointer")

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
	if debug {
		fmt.Printf("fromVal  : %v %T %v %v\n", fromVal, fromVal, fromVal, fromVal.Kind())
	}
	if fromVal.Kind() == reflect.Ptr {
		if fromVal.IsNil() {
			fromVal = reflect.New(T)
		}
		fromVal = fromVal.Elem()
	}
	if debug {
		fmt.Printf("fromVal  : %v %T %v %v\n", fromVal, fromVal, fromVal, fromVal.Kind())
	}
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
			if debug {
				fmt.Printf("** %v %T %v\n",
					element.Field(i), element.Field(i), element.Field(i).Kind())
			}
			field := Pointerize(element.Field(i))
			set := toVal.Field(i)
			if debug {
				fmt.Printf("field: %v %T %v\n", field, field, field)
				fmt.Printf("set  : %v %T %v %v\n", set, set, set, set.Kind())
			}
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
	if debug {
		fmt.Printf("** %v %T %v\n", in, in, in)
	}
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
	if debug {
		fmt.Printf("** %v %T %v\n", arg, arg, arg)
		fmt.Printf("** %v %T %v\n", in, in, in)
	}
	return

}
