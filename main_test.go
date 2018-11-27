package sure

import (
	"testing"
)

var in interface{}

func TestBool(t *testing.T) {
	in = 0
	if Bool(in) {
		t.Fail()
	}
	in = true
	if Bool(in) {
		return
	}
	t.Fail()
}
func TestInt(t *testing.T) {
	in = "okay"
	if Int(in) != 0 {
		t.Fail()
	}
	in = 687
	if Int(in) == 687 {
		return
	}
	t.Fail()
}
func TestInt64(t *testing.T) {
	in = int(977)
	if Int64(in) != 0 {
		t.Fail()
	}
	in = int64(89)
	if Int64(in) == 89 {
		return
	}
	t.Fail()
}
func TestString(t *testing.T) {
	in = true
	if String(in) != "" {
		t.Fail()
	}
	in = "Hello"
	if String(in) == "Hello" {
		return
	}
	t.Fail()
}
func TestFloat32(t *testing.T) {
	in = 5
	if Float32(in) != 0.00 {
		t.Fail()
	}
	in = float32(0.056)
	if Float32(in) == 0.056 {
		return
	}
	t.Fail()
}
func TestFloat64(t *testing.T) {
	in = 5
	if Float64(in) != 0.00 {
		t.Fail()
	}
	in = 8.988
	if Float64(in) == 8.988 {
		return
	}
	t.Fail()
}
func TestHasSet(t *testing.T) {
	in = 887
	if HasSet(in) {
		t.Fail()
	}
	in = float32(0.056)
	if HasSet(in) == true {
		return
	}
	t.Fail()
}
