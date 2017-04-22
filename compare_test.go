package clone

import (
	"fmt"
	"testing"
)

func CompareIntSlice(t *testing.T, src, dst []int) {
	if debug {
		fmt.Println(src, dst)
	}
	for i, v := range src {
		if v != dst[i] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[i])
		}
	}
}

func CompareZSlice(t *testing.T, src, dst []Z) {
	if debug {
		fmt.Println(src, dst)
	}
	for i, v := range src {
		if v != dst[i] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[i])
		}
	}
}

func CompareBoolSlice(t *testing.T, src, dst []bool) {
	if debug {
		fmt.Println(src, dst)
	}
	for i, v := range src {
		if v != dst[i] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[i])
		}
	}
}

func CompareMapStringBool(t *testing.T, src, dst map[string]bool) {
	if debug {
		fmt.Println(src, dst)
	}
	for k, v := range src {
		if v != dst[k] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[k])
		}
	}
}
func CompareMapIntString(t *testing.T, src, dst map[int]string) {
	if debug {
		fmt.Println(src, dst)
	}
	for k, v := range src {
		if v != dst[k] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[k])
		}
	}
}

func CompareMapOfZStruct(t *testing.T, src, dst map[string]Z) {
	if debug {
		fmt.Println(src, dst)
	}
	for k, v := range src {
		if v != dst[k] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[k])
		}
	}
}

func CompareInt(t *testing.T, src, dst int) {
	if debug {
		fmt.Println(src, dst)
	}
	if src != dst {
		t.Errorf("Test_Int: %v %v %v", src, dst, src == dst)
	}
}

func CompareFloat(t *testing.T, src, dst float64) {
	if debug {
		fmt.Println(src, dst)
	}
	if src != dst {
		t.Errorf("Test_Int: %v %v %v", src, dst, src == dst)
	}
}

func CompareString(t *testing.T, src, dst string) {
	if debug {
		fmt.Println(src, dst)
	}
	if src != dst {
		t.Errorf("Test_Int: %v %v %v", src, dst, src == dst)
	}
}

func CompareBool(t *testing.T, src, dst bool) {
	if debug {
		fmt.Println(src, dst)
	}
	if src != dst {
		t.Errorf("Test_Int: %v %v %v", src, dst, src == dst)
	}
}

func CompareInnerStruct(t *testing.T, src, dst Inner) {
	if debug {
		fmt.Println(src, dst)
	}
	for i, s := range src.SSM {
		for k, v := range s {
			if v != dst.SSM[i][k] {
				t.Errorf("Test_Int: %v %v", src, dst)
			}
		}
	}
}

func CompareZ(t *testing.T, src, dst Z) {
	if debug {
		fmt.Println(src, dst)
	}
	if src != dst {
		t.Errorf("Test_Int: %v %v %v", src, dst, src == dst)
	}
}
