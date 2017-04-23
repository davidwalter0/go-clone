package clone

import (
	"fmt"
	"testing"
)

func compareIntSlice(t *testing.T, src, dst []int) {
	if debug {
		fmt.Println(src, dst)
	}
	for i, v := range src {
		if v != dst[i] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[i])
		}
	}
}

func compareZSlice(t *testing.T, src, dst []Z) {
	if debug {
		fmt.Println(src, dst)
	}
	for i, v := range src {
		if v != dst[i] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[i])
		}
	}
}

func compareBoolSlice(t *testing.T, src, dst []bool) {
	if debug {
		fmt.Println(src, dst)
	}
	for i, v := range src {
		if v != dst[i] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[i])
		}
	}
}

func compareMapStringBool(t *testing.T, src, dst map[string]bool) {
	if debug {
		fmt.Println(src, dst)
	}
	for k, v := range src {
		if v != dst[k] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[k])
		}
	}
}
func compareMapIntString(t *testing.T, src, dst map[int]string) {
	if debug {
		fmt.Println(src, dst)
	}
	for k, v := range src {
		if v != dst[k] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[k])
		}
	}
}

func compareMapOfZStruct(t *testing.T, src, dst map[string]Z) {
	if debug {
		fmt.Println(src, dst)
	}
	for k, v := range src {
		if v != dst[k] {
			t.Errorf("Test_Int: %v %v %v", src, dst, v == dst[k])
		}
	}
}

func compareInt(t *testing.T, src, dst int) {
	if debug {
		fmt.Println(src, dst)
	}
	if src != dst {
		t.Errorf("Test_Int: %v %v %v", src, dst, src == dst)
	}
}

func compareFloat(t *testing.T, src, dst float64) {
	if debug {
		fmt.Println(src, dst)
	}
	if src != dst {
		t.Errorf("Test_Int: %v %v %v", src, dst, src == dst)
	}
}

func compareString(t *testing.T, src, dst string) {
	if debug {
		fmt.Println(src, dst)
	}
	if src != dst {
		t.Errorf("Test_Int: %v %v %v", src, dst, src == dst)
	}
}

func compareBool(t *testing.T, src, dst bool) {
	if debug {
		fmt.Println(src, dst)
	}
	if src != dst {
		t.Errorf("Test_Int: %v %v %v", src, dst, src == dst)
	}
}

func compareInnerStruct(t *testing.T, src, dst Inner) {
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

func compareZ(t *testing.T, src, dst Z) {
	if debug {
		fmt.Println(src, dst)
	}
	if src != dst {
		t.Errorf("Test_Int: %v %v %v", src, dst, src == dst)
	}
}
