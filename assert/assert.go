package assert

import "reflect"

type Ter interface {
	Fail()
	Errorf(format string, args ...interface{})
}

func Fail(t Ter, msg ...string) {
	if len(msg) > 0 {
		t.Errorf(msg[0])
	} else {
		t.Fail()
	}
}

func AssertEqual(t Ter, obj1, obj2 interface{}, msg ...string) {
	if !reflect.DeepEqual(obj1, obj2) {
		Fail(t, msg...)
	}
}

func AssertNotEqual(t Ter, obj1, obj2 interface{}, msg ...string) {
	if reflect.DeepEqual(obj1, obj2) {
		Fail(t, msg...)
	}
}

func AssertNil(t Ter, obj1 interface{}, msg ...string) {
	if obj1 != nil {
		Fail(t, msg...)
	}
}

func AssertNotNil(t Ter, obj1 interface{}, msg ...string) {
	if obj1 == nil {
		Fail(t, msg...)
	}
}

func AssertLessThan(t Ter, obj1, obj2 interface{}, msg ...string) {
	if !lessThan(obj1, obj2) {
		Fail(t, msg...)
	}
}

func AssertLessThanOrEqual(t Ter, obj1, obj2 interface{}, msg ...string) {
	if !lessThanOrEqual(obj1, obj2) {
		Fail(t, msg...)
	}
}

func AssertGreaterThan(t Ter, obj1, obj2 interface{}, msg ...string) {
	if !greaterThan(obj1, obj2) {
		Fail(t, msg...)
	}
}

func AssertGreaterThanOrEqual(t Ter, obj1, obj2 interface{}, msg ...string) {
	if !greaterThanOrEqual(obj1, obj2) {
		Fail(t, msg...)
	}
}

func AssertTrue(t Ter, obj1 bool, msg ...string) {
	if !obj1 {
		Fail(t, msg...)
	}
}

func AssertFalse(t Ter, obj1 bool, msg ...string) {
	if obj1 {
		Fail(t, msg...)
	}
}

func lessThan(obj1, obj2 interface{}) bool {
	t1 := reflect.TypeOf(obj1)
	t2 := reflect.TypeOf(obj2)
	if t1 != t2 {
		return false
	}
	switch v1 := obj1.(type) {
	case int:
		return v1 < obj2.(int)
	case int8:
		return v1 < obj2.(int8)
	case int16:
		return v1 < obj2.(int16)
	case int32:
		return v1 < obj2.(int32)
	case int64:
		return v1 < obj2.(int64)
	case float32:
		return v1 < obj2.(float32)
	case float64:
		return v1 < obj2.(float64)

	}
	return false
}

func lessThanOrEqual(obj1, obj2 interface{}) bool {
	t1 := reflect.TypeOf(obj1)
	t2 := reflect.TypeOf(obj2)
	if t1 != t2 {
		return false
	}
	switch v1 := obj1.(type) {
	case int:
		return v1 <= obj2.(int)
	case int8:
		return v1 <= obj2.(int8)
	case int16:
		return v1 <= obj2.(int16)
	case int32:
		return v1 <= obj2.(int32)
	case int64:
		return v1 <= obj2.(int64)
	case float32:
		return v1 <= obj2.(float32)
	case float64:
		return v1 <= obj2.(float64)

	}
	return false
}

func greaterThan(obj1, obj2 interface{}) bool {
	t1 := reflect.TypeOf(obj1)
	t2 := reflect.TypeOf(obj2)
	if t1 != t2 {
		return false
	}
	switch v1 := obj1.(type) {
	case int:
		return v1 > obj2.(int)
	case int8:
		return v1 > obj2.(int8)
	case int16:
		return v1 > obj2.(int16)
	case int32:
		return v1 > obj2.(int32)
	case int64:
		return v1 > obj2.(int64)
	case float32:
		return v1 > obj2.(float32)
	case float64:
		return v1 > obj2.(float64)

	}
	return false
}

func greaterThanOrEqual(obj1, obj2 interface{}) bool {
	t1 := reflect.TypeOf(obj1)
	t2 := reflect.TypeOf(obj2)
	if t1 != t2 {
		return false
	}
	switch v1 := obj1.(type) {
	case int:
		return v1 >= obj2.(int)
	case int8:
		return v1 >= obj2.(int8)
	case int16:
		return v1 >= obj2.(int16)
	case int32:
		return v1 >= obj2.(int32)
	case int64:
		return v1 >= obj2.(int64)
	case float32:
		return v1 >= obj2.(float32)
	case float64:
		return v1 >= obj2.(float64)

	}
	return false
}
