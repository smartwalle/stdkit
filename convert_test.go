package stdkit_test

import (
	"errors"
	"math"
	"strconv"
	"testing"

	"github.com/smartwalle/stdkit"
)

type stringerValue struct {
	value string
}

func (v *stringerValue) String() string {
	if v == nil {
		panic("nil stringer")
	}
	return v.value
}

func TestBool(t *testing.T) {
	var tests = []struct {
		v interface{}
		r bool
	}{
		{"true", true},
		{"t", true},
		{"1", true},
		{"0", false},
		{"2", false},
		{"false", false},
		{0, false},
		{2, true},
		{-2, true},
		{1, true},
	}

	for _, tt := range tests {
		if actual, _ := stdkit.Bool(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 bool, 期望获得 %v, 实际获得  %v", tt.v, tt.r, actual)
		}
	}
}

func TestFloat64(t *testing.T) {
	var tests = []struct {
		v interface{}
		r float64
	}{
		{"9.9", 9.9},
		{"9.99", 9.99},
		{"9.91", 9.91},
		{"9.90", 9.90},
		{"9.901", 9.901},
		{"9.001", 9.001},
	}

	for _, tt := range tests {
		if actual, _ := stdkit.Float64(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 float64, 期望获得 %f, 实际获得  %f", tt.v, tt.r, actual)
		}
	}
}

func TestFloat32(t *testing.T) {
	var tests = []struct {
		v interface{}
		r float32
	}{
		{"9.9", 9.9},
		{"9.99", 9.99},
		{"9.91", 9.91},
		{"9.90", 9.90},
		{"9.901", 9.901},
		{"9.001", 9.001},
		{float64(math.MaxFloat32), math.MaxFloat32},
		{math.Inf(1), float32(math.Inf(1))},
		{math.Inf(-1), float32(math.Inf(-1))},
	}

	for _, tt := range tests {
		if actual, err := stdkit.Float32(tt.v); err != nil || actual != tt.r {
			t.Errorf("把 %v 转换为 float32, 期望获得 %f nil, 实际获得 %f %v", tt.v, tt.r, actual, err)
		}
	}
}

func TestFloat32RejectOutOfRange(t *testing.T) {
	type float64Alias float64

	var tests = []struct {
		name string
		v    interface{}
	}{
		{"Float64 over MaxFloat32", math.MaxFloat32 * 2},
		{"Float64 under negative MaxFloat32", -math.MaxFloat32 * 2},
		{"String over MaxFloat32", "1e1000"},
		{"Float64 alias over MaxFloat32", float64Alias(math.MaxFloat32 * 2)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := stdkit.Float32(tt.v); err == nil {
				t.Fatalf("期望获得错误, 实际获得 nil")
			}
		})
	}
}

func TestEnsureFloat32UseDefaultOnOutOfRange(t *testing.T) {
	if actual := stdkit.EnsureFloat32(math.MaxFloat32*2, 9); actual != 9 {
		t.Fatalf("EnsureFloat32 溢出时应返回默认值 9, 实际获得 %f", actual)
	}
}

func FuzzFloat32RejectFiniteOutOfRange(f *testing.F) {
	for _, seed := range []float64{-math.MaxFloat32 * 2, -math.MaxFloat32, -1, 0, 1, math.MaxFloat32, math.MaxFloat32 * 2, math.Inf(1), math.Inf(-1), math.NaN()} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, value float64) {
		var actual, err = stdkit.Float32(value)
		if !math.IsNaN(value) && !math.IsInf(value, 0) && (value > math.MaxFloat32 || value < -math.MaxFloat32) {
			if !errors.Is(err, stdkit.ErrOutOfRange) {
				t.Fatalf("Float32(%f) 期望获得 ErrOutOfRange, 实际获得 %v", value, err)
			}
			return
		}
		if err != nil {
			t.Fatalf("Float32(%f) 期望获得 nil error, 实际获得 %v", value, err)
		}
		if math.IsNaN(value) {
			if !math.IsNaN(float64(actual)) {
				t.Fatalf("Float32(NaN) 期望获得 NaN, 实际获得 %f", actual)
			}
			return
		}
		if actual != float32(value) {
			t.Fatalf("Float32(%f) 期望获得 %f, 实际获得 %f", value, float32(value), actual)
		}
	})
}

var benchmarkFloat32Value interface{} = float64(127.9)
var benchmarkFloat32Result float32

func BenchmarkFloat32(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var actual, err = stdkit.Float32(benchmarkFloat32Value)
		if err != nil {
			b.Fatal(err)
		}
		benchmarkFloat32Result = actual
	}
}

func TestInt(t *testing.T) {
	var tests = []struct {
		v interface{}
		r int
	}{
		{1, 1},
		{9, 9},
		{9999, 9999},
		{uint64(9999), 9999},
		{int64(9999), 9999},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{-2, -2},
		{"1", 1},
		{"-1", -1},
		{"999999", 999999},
	}

	for _, tt := range tests {
		if actual, _ := stdkit.Int(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 int, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestInt64(t *testing.T) {
	var tests = []struct {
		v interface{}
		r int64
	}{
		{1, 1},
		{9, 9},
		{9999, 9999},
		{uint64(9999), 9999},
		{int64(9999), 9999},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{"1", 1},
		{"999999", 999999},
		{"", 0},
		{"3414416614257328130", 3414416614257328130},
	}

	for _, tt := range tests {
		if actual, _ := stdkit.Int64(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 int64, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestUint(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint
	}{
		{"1", 1},
		{"999999", 999999},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
	}

	for _, tt := range tests {
		if actual, _ := stdkit.Uint(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint32, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestUint8(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint8
	}{
		{"1", 1},
		{"65", 65},
		{100, 100},
		{99, 99},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(99), 99},
		{uint8(math.MaxUint8), math.MaxUint8},
	}

	for _, tt := range tests {
		if actual, _ := stdkit.Uint8(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint8, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestUint16(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint16
	}{
		{"1", 1},
		{"65535", 65535},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
		{uint16(math.MaxUint16), math.MaxUint16},
	}

	for _, tt := range tests {
		if actual, _ := stdkit.Uint16(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint16, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestUint32(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint32
	}{
		{"1", 1},
		{"999999", 999999},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
		{uint32(math.MaxUint32), math.MaxUint32},
	}

	for _, tt := range tests {
		if actual, _ := stdkit.Uint32(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint32, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestUint64(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint64
	}{
		{"1", 1},
		{"999999", 999999},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
		{uint64(math.MaxUint64), math.MaxUint64},
	}

	for _, tt := range tests {
		if actual, _ := stdkit.Uint64(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 Uint64, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestIntegerConversionsRejectDecimalStrings(t *testing.T) {
	var tests = []struct {
		name string
		err  func() error
	}{
		{"Int decimal string", func() error {
			_, err := stdkit.Int("999.999")
			return err
		}},
		{"Int malformed decimal string", func() error {
			_, err := stdkit.Int("1.2.3")
			return err
		}},
		{"Int trailing decimal string", func() error {
			_, err := stdkit.Int("123.")
			return err
		}},
		{"Uint decimal string", func() error {
			_, err := stdkit.Uint("999.1")
			return err
		}},
		{"Uint8 decimal string", func() error {
			_, err := stdkit.Uint8("9.999")
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.err(); err == nil {
				t.Fatalf("期望获得错误, 实际获得 nil")
			}
		})
	}
}

func TestIntegerConversionsRejectOutOfRange(t *testing.T) {
	var tests = []struct {
		name string
		err  func() error
	}{
		{"Int from uint overflow", func() error {
			_, err := stdkit.Int(uint64(math.MaxInt) + 1)
			return err
		}},
		{"Int8 overflow", func() error {
			_, err := stdkit.Int8(128)
			return err
		}},
		{"Int8 underflow", func() error {
			_, err := stdkit.Int8(-129)
			return err
		}},
		{"Int8 from uint overflow", func() error {
			_, err := stdkit.Int8(uint(128))
			return err
		}},
		{"Int8 from float overflow", func() error {
			_, err := stdkit.Int8(128.0)
			return err
		}},
		{"Int64 from uint64 overflow", func() error {
			_, err := stdkit.Int64(uint64(math.MaxInt64) + 1)
			return err
		}},
		{"Int64 from positive infinity", func() error {
			_, err := stdkit.Int64(math.Inf(1))
			return err
		}},
		{"Int64 from NaN", func() error {
			_, err := stdkit.Int64(math.NaN())
			return err
		}},
		{"Uint from negative int", func() error {
			_, err := stdkit.Uint(-1)
			return err
		}},
		{"Uint8 from negative int", func() error {
			_, err := stdkit.Uint8(-1)
			return err
		}},
		{"Uint8 overflow", func() error {
			_, err := stdkit.Uint8(256)
			return err
		}},
		{"Uint8 from uint overflow", func() error {
			_, err := stdkit.Uint8(uint16(256))
			return err
		}},
		{"Uint8 from float overflow", func() error {
			_, err := stdkit.Uint8(256.0)
			return err
		}},
		{"Uint64 from negative int", func() error {
			_, err := stdkit.Uint64(int64(-1))
			return err
		}},
		{"Uint64 from negative float", func() error {
			_, err := stdkit.Uint64(-1.0)
			return err
		}},
		{"Uintptr from negative int", func() error {
			_, err := stdkit.Uintptr(-1)
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.err(); !errors.Is(err, stdkit.ErrOutOfRange) {
				t.Fatalf("期望获得 ErrOutOfRange, 实际获得 %v", err)
			}
		})
	}
}

func TestIntegerConversionsAcceptRangeBoundaries(t *testing.T) {
	if actual, err := stdkit.Int8(math.MinInt8); err != nil || actual != math.MinInt8 {
		t.Fatalf("把 MinInt8 转换为 int8, 期望获得 %d nil, 实际获得 %d %v", math.MinInt8, actual, err)
	}
	if actual, err := stdkit.Int8(math.MaxInt8); err != nil || actual != math.MaxInt8 {
		t.Fatalf("把 MaxInt8 转换为 int8, 期望获得 %d nil, 实际获得 %d %v", math.MaxInt8, actual, err)
	}
	if actual, err := stdkit.Uint8(math.MaxUint8); err != nil || actual != math.MaxUint8 {
		t.Fatalf("把 MaxUint8 转换为 uint8, 期望获得 %d nil, 实际获得 %d %v", math.MaxUint8, actual, err)
	}
	if actual, err := stdkit.Uint64(uint64(math.MaxUint64)); err != nil || actual != math.MaxUint64 {
		t.Fatalf("把 MaxUint64 转换为 uint64, 期望获得 %d nil, 实际获得 %d %v", uint64(math.MaxUint64), actual, err)
	}
}

func TestSignedIntegerConversionsTruncateFloatBeforeRangeCheck(t *testing.T) {
	var tests = []struct {
		name string
		v    interface{}
		r    int8
	}{
		{"MinInt8 fraction", -128.9, math.MinInt8},
		{"MaxInt8 fraction", 127.9, math.MaxInt8},
		{"Negative fraction", -1.9, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual, err = stdkit.Int8(tt.v)
			if err != nil || actual != tt.r {
				t.Fatalf("把 %v 转换为 int8, 期望获得 %d nil, 实际获得 %d %v", tt.v, tt.r, actual, err)
			}
		})
	}
}

func TestUnsignedIntegerConversionsTruncateFloatBeforeRangeCheck(t *testing.T) {
	var tests = []struct {
		name string
		v    interface{}
		r    uint8
	}{
		{"Negative fraction", -0.9, 0},
		{"Negative small fraction", -0.1, 0},
		{"MaxUint8 fraction", 255.9, math.MaxUint8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual, err = stdkit.Uint8(tt.v)
			if err != nil || actual != tt.r {
				t.Fatalf("把 %v 转换为 uint8, 期望获得 %d nil, 实际获得 %d %v", tt.v, tt.r, actual, err)
			}
		})
	}
}

func TestEnsureIntegerConversionsUseDefaultOnOutOfRange(t *testing.T) {
	if actual := stdkit.EnsureUint(-1, 10); actual != 10 {
		t.Fatalf("EnsureUint 溢出时应返回默认值 10, 实际获得 %d", actual)
	}
	if actual := stdkit.EnsureInt8(128, 9); actual != 9 {
		t.Fatalf("EnsureInt8 溢出时应返回默认值 9, 实际获得 %d", actual)
	}
}

func FuzzInt8AndUint8RejectOutOfRange(f *testing.F) {
	for _, seed := range []int64{-129, -128, -1, 0, 1, 127, 128, 255, 256} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, value int64) {
		var actualInt8, errInt8 = stdkit.Int8(value)
		if value < math.MinInt8 || value > math.MaxInt8 {
			if !errors.Is(errInt8, stdkit.ErrOutOfRange) {
				t.Fatalf("Int8(%d) 期望获得 ErrOutOfRange, 实际获得 %v", value, errInt8)
			}
		} else if errInt8 != nil || actualInt8 != int8(value) {
			t.Fatalf("Int8(%d) 期望获得 %d nil, 实际获得 %d %v", value, int8(value), actualInt8, errInt8)
		}

		var actualUint8, errUint8 = stdkit.Uint8(value)
		if value < 0 || value > math.MaxUint8 {
			if !errors.Is(errUint8, stdkit.ErrOutOfRange) {
				t.Fatalf("Uint8(%d) 期望获得 ErrOutOfRange, 实际获得 %v", value, errUint8)
			}
		} else if errUint8 != nil || actualUint8 != uint8(value) {
			t.Fatalf("Uint8(%d) 期望获得 %d nil, 实际获得 %d %v", value, uint8(value), actualUint8, errUint8)
		}
	})
}

func FuzzUint8TruncateFloatBeforeRangeCheck(f *testing.F) {
	for _, seed := range []float64{-1.1, -1, -0.9, -0.1, 0, 1.9, math.MaxUint8, math.MaxUint8 + 0.9, math.MaxUint8 + 1, math.Inf(1), math.Inf(-1), math.NaN()} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, value float64) {
		var actual, err = stdkit.Uint8(value)
		if math.IsNaN(value) || math.IsInf(value, 0) {
			if !errors.Is(err, stdkit.ErrOutOfRange) {
				t.Fatalf("Uint8(%f) 期望获得 ErrOutOfRange, 实际获得 %v", value, err)
			}
			return
		}

		var truncated = math.Trunc(value)
		if truncated < 0 || truncated > math.MaxUint8 {
			if !errors.Is(err, stdkit.ErrOutOfRange) {
				t.Fatalf("Uint8(%f) 期望获得 ErrOutOfRange, 实际获得 %v", value, err)
			}
			return
		}
		if err != nil || actual != uint8(truncated) {
			t.Fatalf("Uint8(%f) 期望获得 %d nil, 实际获得 %d %v", value, uint8(truncated), actual, err)
		}
	})
}

func FuzzIntegerConversionsRejectDecimalStrings(f *testing.F) {
	for _, seed := range []string{"1", "-1", "0", "999999", "999.999", "1.2.3", "123.", ".123", ""} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, value string) {
		_, err := stdkit.Int(value)
		if _, parseErr := strconv.ParseInt(value, 10, 64); parseErr != nil {
			if err == nil {
				t.Fatalf("Int(%q) 期望获得错误, 实际获得 nil", value)
			}
			return
		}
		if err != nil {
			t.Fatalf("Int(%q) 期望获得 nil error, 实际获得 %v", value, err)
		}
	})
}

var benchmarkConvertValue interface{} = int64(127)
var benchmarkConvertResult int8

func BenchmarkInt8(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var actual, err = stdkit.Int8(benchmarkConvertValue)
		if err != nil {
			b.Fatal(err)
		}
		benchmarkConvertResult = actual
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		v interface{}
		r string
	}{
		{"1", "1"},
		{1, "1"},
		{1.1, "1.1"},
		{true, "true"},
		{&stringerValue{value: "stringer"}, "stringer"},
		{int64(3414416614257328130), "3414416614257328130"},
		{[]byte("Hello世界"), "Hello世界"},
		{[]rune("Hello世界"), "Hello世界"},
	}

	for _, tt := range tests {
		if actual, _ := stdkit.String(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 string, 期望获得 %v, 实际获得  %v", tt.v, tt.r, actual)
		}
	}
}

func TestStringTypedNilStringer(t *testing.T) {
	var value *stringerValue

	var actual, err = stdkit.String(value)
	if err != nil || actual != "" {
		t.Fatalf("把 typed nil fmt.Stringer 转换为 string, 期望获得空字符串 nil, 实际获得 %q %v", actual, err)
	}
}
