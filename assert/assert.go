package assert

import "reflect"

// Ter represent a *testing.T
type Ter interface {
	Fail()
	Errorf(format string, args ...interface{})
}

// Fail will call t.Fail if msg is empty, else call t.Errorf
func Fail(t Ter, msg ...string) {
	if len(msg) > 0 {
		t.Errorf(msg[0])
	} else {
		t.Fail()
	}
}

// Equal will compare obj1 and obj2, and will fail if obj1 != obj2 or obj1 has different type with obj2
func Equal(t Ter, obj1, obj2 interface{}, msg ...string) {
	if !reflect.DeepEqual(obj1, obj2) {
		Fail(t, msg...)
	}
}

// NotEqual will compare obj1 and obj2, and will fail if obj1 == obj2
func NotEqual(t Ter, obj1, obj2 interface{}, msg ...string) {
	if reflect.DeepEqual(obj1, obj2) {
		Fail(t, msg...)
	}
}

// Nil will compare obj1 and nil, and will fail if obj1 != nil
func Nil(t Ter, obj1 interface{}, msg ...string) {
	if !isNil(obj1) {
		Fail(t, msg...)
	}
}

// NotNil will compare obj1 and nil, and will fail if obj1 == nil
func NotNil(t Ter, obj1 interface{}, msg ...string) {
	if isNil(obj1) {
		Fail(t, msg...)
	}
}

// LessThan will compare obj1 and obj2, and will fail if obj1 >= obj2 or obj1 has different type with obj2
func LessThan(t Ter, obj1, obj2 interface{}, msg ...string) {
	if !lessThan(obj1, obj2) {
		Fail(t, msg...)
	}
}

// LessThanOrEqual will compare obj1 and obj2, and will fail if obj1 > obj2 or obj1 has different type with obj2
func LessThanOrEqual(t Ter, obj1, obj2 interface{}, msg ...string) {
	if !lessThanOrEqual(obj1, obj2) {
		Fail(t, msg...)
	}
}

// GreaterThan will compare obj1 and obj2, and will fail if obj1 <= obj2 or obj1 has different type with obj2
func GreaterThan(t Ter, obj1, obj2 interface{}, msg ...string) {
	if !greaterThan(obj1, obj2) {
		Fail(t, msg...)
	}
}

// GreaterThanOrEqual will compare obj1 and obj2, and will fail if obj1 < obj2 or obj1 has different type with obj2
func GreaterThanOrEqual(t Ter, obj1, obj2 interface{}, msg ...string) {
	if !greaterThanOrEqual(obj1, obj2) {
		Fail(t, msg...)
	}
}

// True wil fail if obj1 is not true
func True(t Ter, obj1 bool, msg ...string) {
	if !obj1 {
		Fail(t, msg...)
	}
}

// False wil fail if obj1 is true
func False(t Ter, obj1 bool, msg ...string) {
	if obj1 {
		Fail(t, msg...)
	}
}

// lessThan will compare obj1 and obj2, and will return false if obj1 >= obj2 or obj1 has different type with obj2
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

// lessThanOrEqual will compare obj1 and obj2, and will return false if obj1 > obj2 or obj1 has different type with obj2
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

// greaterThan will compare obj1 and obj2, and will return false if obj1 <= obj2 or obj1 has different type with obj2
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

// greaterThanOrEqual will compare obj1 and obj2, and will return false if obj1 < obj2 or obj1 has different type with obj2
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

func isNil(i interface{}) (isnil bool) {
	defer func() {
		if r := recover(); r != nil {
			isnil = false
		}
	}()
	return i == nil || reflect.ValueOf(i).IsNil()
}
