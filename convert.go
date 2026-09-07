package stdkit

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

var ErrUnsupportedValue = errors.New("unsupported value")
var ErrOutOfRange = errors.New("value out of range")

func EnsureBool(value interface{}, defaultValue bool) bool {
	var rValue, err = Bool(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Bool(value interface{}) (bool, error) {
	switch rValue := value.(type) {
	case int:
		return rValue != 0, nil
	case int8:
		return rValue != 0, nil
	case int16:
		return rValue != 0, nil
	case int32:
		return rValue != 0, nil
	case int64:
		return rValue != 0, nil
	case uint:
		return rValue != 0, nil
	case uint8:
		return rValue != 0, nil
	case uint16:
		return rValue != 0, nil
	case uint32:
		return rValue != 0, nil
	case uint64:
		return rValue != 0, nil
	case uintptr:
		return rValue != 0, nil
	case float32:
		return rValue != 0, nil
	case float64:
		return rValue != 0, nil
	case bool:
		return rValue, nil
	case string:
		var nValue, err = strconv.ParseBool(rValue)
		return nValue, err
	default:
		var refValue = reflect.ValueOf(value)
		if !refValue.IsValid() {
			return false, nil
		}
		var refKind = refValue.Kind()

		switch refKind {
		case reflect.Ptr:
			if refValue.IsNil() {
				return false, nil
			}
			return Bool(refValue.Elem().Interface())
		case reflect.Bool:
			return Bool(refValue.Bool())
		case reflect.String:
			return Bool(refValue.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return Bool(refValue.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return Bool(refValue.Uint())
		case reflect.Float32, reflect.Float64:
			return Bool(refValue.Float())
		default:
			return false, ErrUnsupportedValue
		}
	}
}

func EnsureFloat32(value interface{}, defaultValue float32) float32 {
	var rValue, err = Float32(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Float32(value interface{}) (float32, error) {
	switch rValue := value.(type) {
	case int:
		return float32(rValue), nil
	case int8:
		return float32(rValue), nil
	case int16:
		return float32(rValue), nil
	case int32:
		return float32(rValue), nil
	case int64:
		return float32(rValue), nil
	case uint:
		return float32(rValue), nil
	case uint8:
		return float32(rValue), nil
	case uint16:
		return float32(rValue), nil
	case uint32:
		return float32(rValue), nil
	case uint64:
		return float32(rValue), nil
	case uintptr:
		return float32(rValue), nil
	case float32:
		return rValue, nil
	case float64:
		return checkedFloat32(rValue)
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseFloat(rValue, 32)
		return float32(nValue), err
	default:
		var refValue = reflect.ValueOf(value)
		if !refValue.IsValid() {
			return 0, nil
		}
		var refKind = refValue.Kind()

		switch refKind {
		case reflect.Ptr:
			if refValue.IsNil() {
				return 0, nil
			}
			return Float32(refValue.Elem().Interface())
		case reflect.Bool:
			return Float32(refValue.Bool())
		case reflect.String:
			return Float32(refValue.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return Float32(refValue.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return Float32(refValue.Uint())
		case reflect.Float32:
			return Float32(refValue.Float())
		case reflect.Float64:
			return checkedFloat32(refValue.Float())
		default:
			return 0, ErrUnsupportedValue
		}
	}
}

func EnsureFloat64(value interface{}, defaultValue float64) float64 {
	var rValue, err = Float64(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Float64(value interface{}) (float64, error) {
	switch rValue := value.(type) {
	case int:
		return float64(rValue), nil
	case int8:
		return float64(rValue), nil
	case int16:
		return float64(rValue), nil
	case int32:
		return float64(rValue), nil
	case int64:
		return float64(rValue), nil
	case uint:
		return float64(rValue), nil
	case uint8:
		return float64(rValue), nil
	case uint16:
		return float64(rValue), nil
	case uint32:
		return float64(rValue), nil
	case uint64:
		return float64(rValue), nil
	case uintptr:
		return float64(rValue), nil
	case float32:
		return float64(rValue), nil
	case float64:
		return rValue, nil
	case bool:
		if rValue {
			return 1, nil
		}
		return 0, nil
	case string:
		var nValue, err = strconv.ParseFloat(rValue, 64)
		return nValue, err
	default:
		var refValue = reflect.ValueOf(value)
		if !refValue.IsValid() {
			return 0, nil
		}
		var refKind = refValue.Kind()

		switch refKind {
		case reflect.Ptr:
			if refValue.IsNil() {
				return 0, nil
			}
			return Float64(refValue.Elem().Interface())
		case reflect.Bool:
			return Float64(refValue.Bool())
		case reflect.String:
			return Float64(refValue.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return Float64(refValue.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return Float64(refValue.Uint())
		case reflect.Float32, reflect.Float64:
			return Float64(refValue.Float())
		default:
			return 0, ErrUnsupportedValue
		}
	}
}

func EnsureInt(value interface{}, defaultValue int) int {
	var rValue, err = Int(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Int(value interface{}) (int, error) {
	var nValue, err = int64Value(value, math.MinInt, math.MaxInt)
	if err != nil {
		return 0, err
	}
	return int(nValue), nil
}

func EnsureInt8(value interface{}, defaultValue int8) int8 {
	var rValue, err = Int8(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Int8(value interface{}) (int8, error) {
	var nValue, err = int64Value(value, math.MinInt8, math.MaxInt8)
	if err != nil {
		return 0, err
	}
	return int8(nValue), nil
}

func EnsureInt16(value interface{}, defaultValue int16) int16 {
	var rValue, err = Int16(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Int16(value interface{}) (int16, error) {
	var nValue, err = int64Value(value, math.MinInt16, math.MaxInt16)
	if err != nil {
		return 0, err
	}
	return int16(nValue), nil
}

func EnsureInt32(value interface{}, defaultValue int32) int32 {
	var rValue, err = Int32(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Int32(value interface{}) (int32, error) {
	var nValue, err = int64Value(value, math.MinInt32, math.MaxInt32)
	if err != nil {
		return 0, err
	}
	return int32(nValue), nil
}

func EnsureInt64(value interface{}, defaultValue int64) int64 {
	var rValue, err = Int64(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Int64(value interface{}) (int64, error) {
	return int64Value(value, math.MinInt64, math.MaxInt64)
}

func EnsureUint(value interface{}, defaultValue uint) uint {
	var rValue, err = Uint(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Uint(value interface{}) (uint, error) {
	var nValue, err = uint64Value(value, math.MaxUint)
	if err != nil {
		return 0, err
	}
	return uint(nValue), nil
}

func EnsureUint8(value interface{}, defaultValue uint8) uint8 {
	var rValue, err = Uint8(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Uint8(value interface{}) (uint8, error) {
	var nValue, err = uint64Value(value, math.MaxUint8)
	if err != nil {
		return 0, err
	}
	return uint8(nValue), nil
}

func EnsureUint16(value interface{}, defaultValue uint16) uint16 {
	var rValue, err = Uint16(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Uint16(value interface{}) (uint16, error) {
	var nValue, err = uint64Value(value, math.MaxUint16)
	if err != nil {
		return 0, err
	}
	return uint16(nValue), nil
}

func EnsureUint32(value interface{}, defaultValue uint32) uint32 {
	var rValue, err = Uint32(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Uint32(value interface{}) (uint32, error) {
	var nValue, err = uint64Value(value, math.MaxUint32)
	if err != nil {
		return 0, err
	}
	return uint32(nValue), nil
}

func EnsureUint64(value interface{}, defaultValue uint64) uint64 {
	var rValue, err = Uint64(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Uint64(value interface{}) (uint64, error) {
	return uint64Value(value, math.MaxUint64)
}

func EnsureUintptr(value interface{}, defaultValue uintptr) uintptr {
	var rValue, err = Uintptr(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func Uintptr(value interface{}) (uintptr, error) {
	var nValue, err = uint64Value(value, math.MaxUint)
	if err != nil {
		return 0, err
	}
	return uintptr(nValue), nil
}

func EnsureString(value interface{}, defaultValue string) string {
	var rValue, err = String(value)
	if err != nil {
		return defaultValue
	}
	return rValue
}

func String(value interface{}) (string, error) {
	switch rValue := value.(type) {
	case fmt.Stringer:
		var refValue = reflect.ValueOf(rValue)
		if refValue.Kind() == reflect.Ptr && refValue.IsNil() {
			return "", nil
		}
		return rValue.String(), nil
	case int:
		return strconv.FormatInt(int64(rValue), 10), nil
	case int8:
		return strconv.FormatInt(int64(rValue), 10), nil
	case int16:
		return strconv.FormatInt(int64(rValue), 10), nil
	case int32:
		return strconv.FormatInt(int64(rValue), 10), nil
	case int64:
		return strconv.FormatInt(rValue, 10), nil
	case uint:
		return strconv.FormatUint(uint64(rValue), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(rValue), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(rValue), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(rValue), 10), nil
	case uint64:
		return strconv.FormatUint(rValue, 10), nil
	case uintptr:
		return strconv.FormatUint(uint64(rValue), 10), nil
	case float32:
		return strconv.FormatFloat(float64(rValue), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(rValue, 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(rValue), nil
	case string:
		return rValue, nil
	case []byte:
		return string(rValue), nil
	case []rune:
		return string(rValue), nil
	default:
		var refValue = reflect.ValueOf(value)
		if !refValue.IsValid() {
			return "", nil
		}
		var refKind = refValue.Kind()

		switch refKind {
		case reflect.Ptr:
			if refValue.IsNil() {
				return "", nil
			}
			return String(refValue.Elem().Interface())
		case reflect.Bool:
			return String(refValue.Bool())
		case reflect.String:
			return String(refValue.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return String(refValue.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return String(refValue.Uint())
		case reflect.Float32, reflect.Float64:
			return String(refValue.Float())
		default:
			return "", ErrUnsupportedValue
		}
	}
}

func int64Value(value interface{}, min int64, max int64) (int64, error) {
	switch rValue := value.(type) {
	case int:
		return checkedInt64(int64(rValue), min, max)
	case int8:
		return checkedInt64(int64(rValue), min, max)
	case int16:
		return checkedInt64(int64(rValue), min, max)
	case int32:
		return checkedInt64(int64(rValue), min, max)
	case int64:
		return checkedInt64(rValue, min, max)
	case uint:
		return checkedUint64ToInt64(uint64(rValue), max)
	case uint8:
		return checkedUint64ToInt64(uint64(rValue), max)
	case uint16:
		return checkedUint64ToInt64(uint64(rValue), max)
	case uint32:
		return checkedUint64ToInt64(uint64(rValue), max)
	case uint64:
		return checkedUint64ToInt64(rValue, max)
	case uintptr:
		return checkedUint64ToInt64(uint64(rValue), max)
	case float32:
		return checkedFloat64ToInt64(float64(rValue), min, max)
	case float64:
		return checkedFloat64ToInt64(rValue, min, max)
	case bool:
		if rValue {
			return checkedInt64(1, min, max)
		}
		return checkedInt64(0, min, max)
	case string:
		var nValue, err = strconv.ParseInt(rValue, 10, 64)
		if err != nil {
			return 0, err
		}
		return checkedInt64(nValue, min, max)
	default:
		var refValue = reflect.ValueOf(value)
		if !refValue.IsValid() {
			return 0, nil
		}

		switch refValue.Kind() {
		case reflect.Ptr:
			if refValue.IsNil() {
				return 0, nil
			}
			return int64Value(refValue.Elem().Interface(), min, max)
		case reflect.Bool:
			return int64Value(refValue.Bool(), min, max)
		case reflect.String:
			return int64Value(refValue.String(), min, max)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return checkedInt64(refValue.Int(), min, max)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return checkedUint64ToInt64(refValue.Uint(), max)
		case reflect.Float32, reflect.Float64:
			return checkedFloat64ToInt64(refValue.Float(), min, max)
		default:
			return 0, ErrUnsupportedValue
		}
	}
}

func uint64Value(value interface{}, max uint64) (uint64, error) {
	switch rValue := value.(type) {
	case int:
		return checkedInt64ToUint64(int64(rValue), max)
	case int8:
		return checkedInt64ToUint64(int64(rValue), max)
	case int16:
		return checkedInt64ToUint64(int64(rValue), max)
	case int32:
		return checkedInt64ToUint64(int64(rValue), max)
	case int64:
		return checkedInt64ToUint64(rValue, max)
	case uint:
		return checkedUint64(uint64(rValue), max)
	case uint8:
		return checkedUint64(uint64(rValue), max)
	case uint16:
		return checkedUint64(uint64(rValue), max)
	case uint32:
		return checkedUint64(uint64(rValue), max)
	case uint64:
		return checkedUint64(rValue, max)
	case uintptr:
		return checkedUint64(uint64(rValue), max)
	case float32:
		return checkedFloat64ToUint64(float64(rValue), max)
	case float64:
		return checkedFloat64ToUint64(rValue, max)
	case bool:
		if rValue {
			return checkedUint64(1, max)
		}
		return checkedUint64(0, max)
	case string:
		var nValue, err = strconv.ParseUint(rValue, 10, 64)
		if err != nil {
			return 0, err
		}
		return checkedUint64(nValue, max)
	default:
		var refValue = reflect.ValueOf(value)
		if !refValue.IsValid() {
			return 0, nil
		}

		switch refValue.Kind() {
		case reflect.Ptr:
			if refValue.IsNil() {
				return 0, nil
			}
			return uint64Value(refValue.Elem().Interface(), max)
		case reflect.Bool:
			return uint64Value(refValue.Bool(), max)
		case reflect.String:
			return uint64Value(refValue.String(), max)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return checkedInt64ToUint64(refValue.Int(), max)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return checkedUint64(refValue.Uint(), max)
		case reflect.Float32, reflect.Float64:
			return checkedFloat64ToUint64(refValue.Float(), max)
		default:
			return 0, ErrUnsupportedValue
		}
	}
}

func checkedInt64(value int64, min int64, max int64) (int64, error) {
	if value < min || value > max {
		return 0, ErrOutOfRange
	}
	return value, nil
}

func checkedUint64(value uint64, max uint64) (uint64, error) {
	if value > max {
		return 0, ErrOutOfRange
	}
	return value, nil
}

func checkedUint64ToInt64(value uint64, max int64) (int64, error) {
	if value > uint64(max) {
		return 0, ErrOutOfRange
	}
	return int64(value), nil
}

func checkedInt64ToUint64(value int64, max uint64) (uint64, error) {
	if value < 0 || uint64(value) > max {
		return 0, ErrOutOfRange
	}
	return uint64(value), nil
}

func checkedFloat64ToInt64(value float64, min int64, max int64) (int64, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, ErrOutOfRange
	}
	value = math.Trunc(value)
	if value < float64(min) || value >= float64(max)+1 {
		return 0, ErrOutOfRange
	}
	return int64(value), nil
}

func checkedFloat64ToUint64(value float64, max uint64) (uint64, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, ErrOutOfRange
	}
	value = math.Trunc(value)
	if value < 0 {
		return 0, ErrOutOfRange
	}
	var maxExclusive = float64(max) + 1
	if max == math.MaxUint64 {
		maxExclusive = 18446744073709551616
	}
	if value >= maxExclusive {
		return 0, ErrOutOfRange
	}
	return uint64(value), nil
}

func checkedFloat32(value float64) (float32, error) {
	if !math.IsNaN(value) && !math.IsInf(value, 0) && (value > math.MaxFloat32 || value < -math.MaxFloat32) {
		return 0, ErrOutOfRange
	}
	return float32(value), nil
}
