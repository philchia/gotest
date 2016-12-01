package assert

import "testing"

type testStatus int8

const (
	empty testStatus = iota
	fail
	errorf
)

type fakeT struct {
	status testStatus
}

func (t *fakeT) Fail() {
	t.status = fail
}

func (t *fakeT) Errorf(format string, args ...interface{}) {
	t.status = errorf
}

func TestFail(t *testing.T) {
	type args struct {
		t   Ter
		msg []string
	}
	tests := []struct {
		name string
		args args
		want testStatus
	}{
		{
			"case1",
			args{
				new(fakeT),
				nil,
			},
			fail,
		},
		{
			"case2",
			args{
				new(fakeT),
				[]string{"failed"},
			},
			errorf,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fail(tt.args.t, tt.args.msg...)
			if tt.args.t.(*fakeT).status != tt.want {
				t.Fail()
			}
		})
	}
}

func TestAssertEqual(t *testing.T) {
	type args struct {
		t    Ter
		obj1 interface{}
		obj2 interface{}
		msg  []string
	}
	tests := []struct {
		name string
		args args
		want testStatus
	}{
		{
			"case1",
			args{
				new(fakeT),
				1,
				1,
				nil,
			},
			empty,
		},
		{
			"case2",
			args{
				new(fakeT),
				2,
				1,
				nil,
			},
			fail,
		},
		{
			"case3",
			args{
				new(fakeT),
				2,
				3,
				[]string{"failed"},
			},
			errorf,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Equal(tt.args.t, tt.args.obj1, tt.args.obj2, tt.args.msg...)
			if tt.args.t.(*fakeT).status != tt.want {
				t.Fail()
			}
		})
	}
}

func TestAssertNotEqual(t *testing.T) {
	type args struct {
		t    Ter
		obj1 interface{}
		obj2 interface{}
		msg  []string
	}
	tests := []struct {
		name string
		args args
		want testStatus
	}{
		{
			"case1",
			args{
				new(fakeT),
				1,
				1,
				nil,
			},
			fail,
		},
		{
			"case2",
			args{
				new(fakeT),
				2,
				1,
				nil,
			},
			empty,
		},
		{
			"case3",
			args{
				new(fakeT),
				2,
				2,
				[]string{"failed"},
			},
			errorf,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NotEqual(tt.args.t, tt.args.obj1, tt.args.obj2, tt.args.msg...)
			if tt.args.t.(*fakeT).status != tt.want {
				t.Fail()
			}
		})
	}
}

func TestAssertNil(t *testing.T) {
	type args struct {
		t    Ter
		obj1 interface{}
		msg  []string
	}
	tests := []struct {
		name string
		args args
		want testStatus
	}{
		{
			"case1",
			args{
				new(fakeT),
				&struct{}{},
				nil,
			},
			fail,
		},
		{
			"case2",
			args{
				new(fakeT),
				&struct{}{},
				[]string{"fail"},
			},
			errorf,
		},
		{
			"case3",
			args{
				new(fakeT),
				nil,
				nil,
			},
			empty,
		},
		{
			"case3",
			args{
				new(fakeT),
				2,
				nil,
			},
			fail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Nil(tt.args.t, tt.args.obj1, tt.args.msg...)
			if tt.args.t.(*fakeT).status != tt.want {
				t.Fail()
			}
		})
	}
}

func TestAssertNotNil(t *testing.T) {
	type args struct {
		t    Ter
		obj1 interface{}
		msg  []string
	}
	tests := []struct {
		name string
		args args
		want testStatus
	}{
		{
			"case1",
			args{
				new(fakeT),
				&struct{}{},
				nil,
			},
			empty,
		},
		{
			"case2",
			args{
				new(fakeT),
				nil,
				[]string{"fail"},
			},
			errorf,
		},
		{
			"case3",
			args{
				new(fakeT),
				2,
				nil,
			},
			empty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NotNil(tt.args.t, tt.args.obj1, tt.args.msg...)
			if tt.args.t.(*fakeT).status != tt.want {
				t.Fail()
			}
		})
	}
}

func TestAssertLessThan(t *testing.T) {
	type args struct {
		t    Ter
		obj1 interface{}
		obj2 interface{}
		msg  []string
	}
	tests := []struct {
		name string
		args args
		want testStatus
	}{
		{
			"case1",
			args{
				new(fakeT),
				nil,
				nil,
				[]string{"fail"},
			},
			errorf,
		},
		{
			"case2",
			args{
				new(fakeT),
				nil,
				2,
				[]string{"fail"},
			},
			errorf,
		},
		{
			"case3",
			args{
				new(fakeT),
				3,
				2,
				nil,
			},
			fail,
		},
		{
			"case4",
			args{
				new(fakeT),
				1,
				2,
				nil,
			},
			empty,
		},
		{
			"case5",
			args{
				new(fakeT),
				2,
				2,
				nil,
			},
			fail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LessThan(tt.args.t, tt.args.obj1, tt.args.obj2, tt.args.msg...)
			if tt.args.t.(*fakeT).status != tt.want {
				t.Fail()
			}
		})
	}
}

func TestAssertLessThanOrEqual(t *testing.T) {
	type args struct {
		t    Ter
		obj1 interface{}
		obj2 interface{}
		msg  []string
	}
	tests := []struct {
		name string
		args args
		want testStatus
	}{
		{
			"case1",
			args{
				new(fakeT),
				nil,
				nil,
				[]string{"fail"},
			},
			empty,
		},
		{
			"case2",
			args{
				new(fakeT),
				nil,
				2,
				[]string{"fail"},
			},
			errorf,
		},
		{
			"case3",
			args{
				new(fakeT),
				3,
				2,
				[]string{"fail"},
			},
			errorf,
		},
		{
			"case4",
			args{
				new(fakeT),
				3,
				2,
				nil,
			},
			fail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LessThanOrEqual(tt.args.t, tt.args.obj1, tt.args.obj2, tt.args.msg...)
		})
	}
}

func TestAssertGreaterThan(t *testing.T) {
	type args struct {
		t    Ter
		obj1 interface{}
		obj2 interface{}
		msg  []string
	}
	tests := []struct {
		name string
		args args
		want testStatus
	}{
		{
			"case1",
			args{
				new(fakeT),
				nil,
				nil,
				[]string{"fail"},
			},
			errorf,
		},
		{
			"case2",
			args{
				new(fakeT),
				nil,
				2,
				[]string{"fail"},
			},
			errorf,
		},
		{
			"case3",
			args{
				new(fakeT),
				3,
				2,
				[]string{"fail"},
			},
			empty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GreaterThan(tt.args.t, tt.args.obj1, tt.args.obj2, tt.args.msg...)
			if tt.args.t.(*fakeT).status != tt.want {
				t.Fail()
			}
		})
	}
}

func TestAssertGreaterThanOrEqual(t *testing.T) {
	type args struct {
		t    Ter
		obj1 interface{}
		obj2 interface{}
		msg  []string
	}
	tests := []struct {
		name string
		args args
		want testStatus
	}{
		{
			"case1",
			args{
				new(fakeT),
				nil,
				nil,
				[]string{"fail"},
			},
			errorf,
		},
		{
			"case2",
			args{
				new(fakeT),
				nil,
				2,
				[]string{"fail"},
			},
			errorf,
		},
		{
			"case3",
			args{
				new(fakeT),
				3,
				2,
				[]string{"fail"},
			},
			empty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GreaterThanOrEqual(tt.args.t, tt.args.obj1, tt.args.obj2, tt.args.msg...)
			if tt.args.t.(*fakeT).status != tt.want {
				t.Fail()
			}
		})
	}
}

func TestAssertTrue(t *testing.T) {
	type args struct {
		t    Ter
		obj1 bool
		msg  []string
	}
	tests := []struct {
		name string
		args args
		want testStatus
	}{
		{
			"case1",
			args{
				new(fakeT),
				false,
				nil,
			},
			fail,
		},
		{
			"case2",
			args{
				new(fakeT),
				true,
				nil,
			},
			empty,
		},
		{
			"case3",
			args{
				new(fakeT),
				false,
				[]string{"fail"},
			},
			errorf,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			True(tt.args.t, tt.args.obj1, tt.args.msg...)
		})
	}
}

func TestAssertFalse(t *testing.T) {
	type args struct {
		t    Ter
		obj1 bool
		msg  []string
	}
	tests := []struct {
		name string
		args args
		want testStatus
	}{
		{
			"case1",
			args{
				new(fakeT),
				false,
				nil,
			},
			empty,
		},
		{
			"case2",
			args{
				new(fakeT),
				true,
				nil,
			},
			fail,
		},
		{
			"case3",
			args{
				new(fakeT),
				true,
				[]string{"fail"},
			},
			errorf,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			False(tt.args.t, tt.args.obj1, tt.args.msg...)
			if tt.args.t.(*fakeT).status != tt.want {
				t.Fail()
			}
		})
	}
}

func Test_lessThan(t *testing.T) {
	type args struct {
		obj1 interface{}
		obj2 interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{
				1,
				2,
			},
			true,
		},
		{
			"case2",
			args{
				2,
				2,
			},
			false,
		},
		{
			"case3",
			args{
				2,
				1,
			},
			false,
		},
		{
			"case4",
			args{
				2,
				1.5,
			},
			false,
		},
		{
			"case5",
			args{
				int8(2),
				int8(1),
			},
			false,
		},
		{
			"case6",
			args{
				int16(2),
				int16(2),
			},
			false,
		},
		{
			"case7",
			args{
				int32(2),
				int32(2),
			},
			false,
		},
		{
			"case8",
			args{
				int64(2),
				int64(2),
			},
			false,
		},
		{
			"case9",
			args{
				float32(2),
				float32(2),
			},
			false,
		},
		{
			"case10",
			args{
				float64(2),
				float64(2),
			},
			false,
		},
		{
			"case11",
			args{
				"2",
				"3",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lessThan(tt.args.obj1, tt.args.obj2); got != tt.want {
				t.Errorf("lessThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lessThanOrEqual(t *testing.T) {
	type args struct {
		obj1 interface{}
		obj2 interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{
				1,
				2,
			},
			true,
		},
		{
			"case2",
			args{
				2,
				2,
			},
			true,
		},
		{
			"case3",
			args{
				2,
				1,
			},
			false,
		},
		{
			"case4",
			args{
				2,
				1.5,
			},
			false,
		},
		{
			"case5",
			args{
				int8(2),
				int8(1),
			},
			false,
		},
		{
			"case6",
			args{
				int16(2),
				int16(2),
			},
			true,
		},
		{
			"case7",
			args{
				int32(2),
				int32(2),
			},
			true,
		},
		{
			"case8",
			args{
				int64(2),
				int64(2),
			},
			true,
		},
		{
			"case9",
			args{
				float32(2),
				float32(2),
			},
			true,
		},
		{
			"case10",
			args{
				float64(2),
				float64(2),
			},
			true,
		},
		{
			"case11",
			args{
				"2",
				"3",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lessThanOrEqual(tt.args.obj1, tt.args.obj2); got != tt.want {
				t.Errorf("lessThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greaterThan(t *testing.T) {
	type args struct {
		obj1 interface{}
		obj2 interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{
				1,
				2,
			},
			false,
		},
		{
			"case2",
			args{
				2,
				2,
			},
			false,
		},
		{
			"case3",
			args{
				2,
				1,
			},
			true,
		},
		{
			"case4",
			args{
				2,
				1.5,
			},
			false,
		},
		{
			"case5",
			args{
				int8(2),
				int8(1),
			},
			true,
		},
		{
			"case6",
			args{
				int16(2),
				int16(2),
			},
			false,
		},
		{
			"case7",
			args{
				int32(2),
				int32(2),
			},
			false,
		},
		{
			"case8",
			args{
				int64(2),
				int64(2),
			},
			false,
		},
		{
			"case9",
			args{
				float32(2),
				float32(2),
			},
			false,
		},
		{
			"case10",
			args{
				float64(2),
				float64(2),
			},
			false,
		},
		{
			"case11",
			args{
				"2",
				"3",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greaterThan(tt.args.obj1, tt.args.obj2); got != tt.want {
				t.Errorf("greaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greaterThanOrEqual(t *testing.T) {
	type args struct {
		obj1 interface{}
		obj2 interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case1",
			args{
				1,
				2,
			},
			false,
		},
		{
			"case2",
			args{
				2,
				2,
			},
			true,
		},
		{
			"case3",
			args{
				2,
				1,
			},
			true,
		},
		{
			"case4",
			args{
				2,
				1.5,
			},
			false,
		},
		{
			"case5",
			args{
				int8(2),
				int8(1),
			},
			true,
		},
		{
			"case6",
			args{
				int16(2),
				int16(2),
			},
			true,
		},
		{
			"case7",
			args{
				int32(2),
				int32(2),
			},
			true,
		},
		{
			"case8",
			args{
				int64(2),
				int64(2),
			},
			true,
		},
		{
			"case9",
			args{
				float32(2),
				float32(2),
			},
			true,
		},
		{
			"case10",
			args{
				float64(2),
				float64(2),
			},
			true,
		},
		{
			"case11",
			args{
				"2",
				"3",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greaterThanOrEqual(tt.args.obj1, tt.args.obj2); got != tt.want {
				t.Errorf("greaterThanOrEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNil(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantIsnil bool
	}{
		{
			"case1",
			args{
				nil,
			},
			true,
		},
		{
			"case2",
			args{
				2,
			},
			false,
		},
		{
			"case3",
			args{
				&struct{}{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsnil := isNil(tt.args.i); gotIsnil != tt.wantIsnil {
				t.Errorf("isNil() = %v, want %v", gotIsnil, tt.wantIsnil)
			}
		})
	}
}
