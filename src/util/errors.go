package util

import (
	"fmt"
)

var panicFunc = func(v interface{}) {
	panic(v)
}

func PanicFunc(f func(v interface{})) {
	panicFunc = f
}

func Must(args ...interface{}) {
	for _, arg := range args {
		if err, ok := arg.(error); ok && err != nil {
			panicFunc(err)
			return
		}
	}
}

func selectVal(rule func(interface{}) bool, args ...interface{}) interface{} {
	var val interface{}
	found := false
	for _, arg := range args {
		if arg == nil {
			continue
		}
		if rule(arg) {
			val = arg
			found = true
			continue
		}
		Must(arg)
	}
	if !found {
		panic(fmt.Errorf("no value found"))
	}
	return val
}

func SelectNotNil(args ...interface{}) interface{} {
	return selectVal(func(arg interface{}) bool {
		return true
	}, args...)
}

func SelectError(args ...interface{}) error {
	return selectVal(func(arg interface{}) bool {
		if _, ok := arg.(error); ok {
			return true
		}
		return false
	}, args...).(error)
}

func SelectString(args ...interface{}) string {
	return selectVal(func(arg interface{}) bool {
		if str, ok := arg.(string); ok {
			return len(str) > 0
		}
		return false
	}, args...).(string)
}

func SelectAnyString(args ...interface{}) string {
	return selectVal(func(arg interface{}) bool {
		if _, ok := arg.(string); ok {
			return true
		}
		return false
	}, args...).(string)
}

func SelectBool(args ...interface{}) bool {
	return selectVal(func(arg interface{}) bool {
		if v, ok := arg.(bool); ok {
			return v
		}
		return false
	}, args...).(bool)
}

func SelectByteSlice(args ...interface{}) []byte {
	return selectVal(func(arg interface{}) bool {
		if v, ok := arg.([]byte); ok {
			return len(v) > 0
		}
		return false
	}, args...).([]byte)
}

func SelectAnyByteSlice(args ...interface{}) []byte {
	return selectVal(func(arg interface{}) bool {
		if _, ok := arg.([]byte); ok {
			return true
		}
		return false
	}, args...).([]byte)
}
