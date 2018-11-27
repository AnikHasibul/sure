// Package sure helps you to use interfaces with more confidence.
// It returns zero value when the confidence is not so good.
/*
It helps you to code like this:

```
if sure.Bool(value){
	//..
}
```

instead of

```
if v,ok := value.(bool); ok {
	//...
}
```

With `sure`:

```
var v = sure.Int(anInterface)
```

Without `sure`:

```
var (
	v int
	ok bool
)
v, ok = anInterface.(int)
if !ok {
	// hndle zero value
}
// use `v`
```
*/
package sure

// PanicWhenWrong - when true it panics when the interface doesn't match with the prefered type. Default is false
var PanicWhenWrong = false

func panicOrNot(msg string) {
	if PanicWhenWrong {
		panic(msg)
	}
}

// Bool returns bool
func Bool(value interface{}) bool {
	if v, ok := value.(bool); ok {
		return v
	}
	panicOrNot("value is not bool")
	return false
}

// Int returns int
func Int(value interface{}) int {
	if v, ok := value.(int); ok {
		return v
	}
	panicOrNot("value is not int")
	return 0
}

// Int64 returns int64
func Int64(value interface{}) int64 {
	if v, ok := value.(int64); ok {
		return v
	}
	panicOrNot("value is not int64")
	return 0
}

// String returns string
func String(value interface{}) string {
	if v, ok := value.(string); ok {
		return v
	}
	panicOrNot("value is not string")
	return ""
}

// Float32 returns float32
func Float32(value interface{}) float32 {
	if v, ok := value.(float32); ok {
		return v
	}
	panicOrNot("value is not float32")
	return 0.00
}

// Float64 returns float64
func Float64(value interface{}) float64 {
	if v, ok := value.(float64); ok {
		return v
	}
	panicOrNot("value is not float64")
	return 0.00
}

// HasSet returns false if the interface is nil
func HasSet(value interface{}) bool {
	if value != nil {
		return true
	}
	panicOrNot("value is not string")
	return false
}
