package funky

import "testing"

func TestSome(t *testing.T) {
	expected := 12
	v := SomeOf(expected)
	if v.IsEmpty() {
		t.Errorf("expected IsEmpty() to return false")
	}
	if !v.IsPresent() {
		t.Errorf("expected IsPresent() to return true")
	}
	if v.Value() != expected {
		t.Errorf("expected Value() to return %d", expected)
	}
	if v.OrGet(10) != expected {
		t.Errorf("expected OrGet() to return %d", expected)
	}
	var executed bool = false
	v.IfPresent(func(value int) {
		if value != expected {
			t.Errorf("expected parameter to IfPresent to be equal %d", expected)
		}
		executed = true
	})
	if !executed {
		t.Errorf("expected function passed to IfPresent to execute")
	}
}

func TestNone(t *testing.T) {
	v := NoneOf[int]()
	if !v.IsEmpty() {
		t.Errorf("expected IsEmpty() to return true")
	}
	if v.IsPresent() {
		t.Errorf("expected IsPresent() to return false")
	}
	expected := 10
	if v.OrGet(expected) != expected {
		t.Errorf("expected OrGet() to return %d", expected)
	}
	var executed bool = false
	v.IfPresent(func(value int) {
		t.Errorf("expected parameter to IfPresent to be equal %d", expected)
		executed = true
	})
	if executed {
		t.Errorf("expected function passed to IfPresent to not execute")
	}
}

func TestNone_Value(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	v := NoneOf[string]()
	v.Value()
	t.Errorf("expected Value() to panic")
}
