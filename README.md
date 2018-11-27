# sure
--
    import "github.com/anikhasibul/sure"

Package sure helps you to use interfaces with more confidence. It returns zero
value when the confidence is not so good.

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
v, ok = anInterface.(int) if !ok {

    // hndle zero value

} 
// use `v` 
```

## Usage

```go
var PanicWhenWrong = false
```
PanicWhenWrong - when true it panics when the interface doesn't match with the
prefered type. Default is false

#### func  Bool

```go
func Bool(value interface{}) bool
```
Bool returns bool

#### func  Float32

```go
func Float32(value interface{}) float32
```
Float32 returns float32

#### func  Float64

```go
func Float64(value interface{}) float64
```
Float64 returns float64

#### func  HasSet

```go
func HasSet(value interface{}) bool
```
HasSet returns false if the interface is nil

#### func  Int

```go
func Int(value interface{}) int
```
Int returns int

#### func  Int64

```go
func Int64(value interface{}) int64
```
Int64 returns int64

#### func  String

```go
func String(value interface{}) string
```
String returns string
