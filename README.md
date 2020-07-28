# Go
Examples and exercises using Golang

## Go Characteristics
* Strong and statically typed.
* Excellent community
* Key features:
  * Simplicity.
  * Fast compile times.
  * Has its own garbage collector.
  * Compiled to standalone binaries.

## Variable's declaration
```go
var i int // var variable_name type
k := 99 // variable_name := value
var I int // Package variable
```
## Primitives
### Boolean
- Values are true or false.
- The default value is false.
### Numeric Types
#### Integers
- Signed Integers
  - int type has varying size, from 8-bit to 64-bit.
- Unsigned Integers
  - From 8-bit to 32-bit.
- Can perform arithmetic operations as well as bitwise operations.
- Default value: 0
#### Floating point numbers
- They follow IEEE-754 standard.
- There are two versions; 32-bit and 64-bit.
- Can perform just arithmetic operations.
- Default value: 0
#### Complex numbers
- 64 and 128-bit versions.
- They have built-in functions.
- Default value: 0+0i.
### Text types
#### String
- UTF-8.
- They are immutable.
- Can be concatenated and can be converted to [] byte.
#### Rune
- UTF-32.
- It is an alias for int32.

## Constants
### Naming convention
In other languages, it is common to name a constant like this: MY_CONST.
If we use this kind of notation in Go, the variable will be taken as a 'package variable'.
Hence, the keyword 'const' is used and the name of the constant should be like any other variable.
Also, it is not possible to create a constant with a run-time value:
```go
const myConst float64 = math.Sin(1.57) // Not possible. math.Sin(x) is executed in run-time.
```
### Typed constants
```go
const myConst int = 42
```
### Untyped constants
The type is inferred.
```go
const myConst = 42
```
### Enumerated constants
```go
const ( 
  a = iota
  b
  c
)
```
### Enumeration expressions
It assign the value dynamicallly.
```go
const ( 
  a = iota + 5
  b
  c
)
```

## Arrays
They are continuous in memory.
```go
grades := [3]int{97, 85, 70} // Array with 3 elements
moreGrades := [...]int{100, 55, 88, 96}
```
## Slices
- Are referenced types.
```go
a := []int{2, 4, 6, 8, 10, 12, 14}
b := a[3:] // Slice from 4th element to end.
c := a[:5] // Slice first 5 elements

x := make([] int, 3) // Making a slice (type, size)
y := make([] int, 3, 100) // Making a slice (type, size, capacity)
```
- Iterating using loops
```go
s := []int{0, 1, 2, 3}
for k, v := range s {
  fmt.Println(k, v)
}
```
## Maps
- Key / Value (map[keyType]valueType).
- Can be created with make function ( make(map[string]int) )
- Referenced types.

## Structs
- No referenced types.
```go
type Doctor struct {
  number int
  actorName string
  companions [] string
}
aDoctor := Doctor {
  number: 3,
  actorName: "Jon Pertwee",
  companions: []string {"Liz Shaw", "Jo Grant"}
}
// aDoctor.actorName --> Access to struct's properties.
```
### Tags
`required max:"100"`

## Control Flow
### If statement
```go
if true {
  // Code
}
```
### Switch
```go
switch 2:
  case 1:
    // Code 1
  case 2:
    // Code 2
  default: 
    // Default code
```

## Loops
### For
- Normal for
```go
for i := 0; i < 5; i +2 {
  // Loop code
}
```
- Using two variables
```go
for i, j := 0, 1; i < 5; i, j = i+1, j+2 {
  fmt.Println(j)
}
```
- Do - while
```go
for i < 5 {
  fmt.Println(i)
  i++
}
```
- Infinite loop
```go
i := 0
for {
  fmt.Println(i)
  i++
  if i == 5 {
    break
  }
}
```
## Defer
- Executes the passed function after the last statement of the function, but right before it actually returns.
- Last in - First Out
- It takes the value of the variables when the 'defer' is; if the variable is updated, it retains the original value.
```go
fmt.Println("Start")
defer fmt.Println("Middle")
fmt.Println("End")
```

## Panic
- In Go, exceptions do not exist.
```go
fmt.Println("Start")
panic("Something wrong happened")
fmt.Println("End")
```

## Recover
```go
fmt.Println("Start")
defer func() {
	if err := recover(); err != nil {
    log.Println("Error: ", err)
    panic(err) // If the error is not controlable.
	}
}()
panic("Something wrong happened")
fmt.Println("End")
```
## Pointers
- A pointer holds the memory location.
- No pointer arithmetic, by default. It is posible using 'unssafe' package.
```go
var a int = 42
var b *int = &a
fmt.Println(a, b) // Prints address.
fmt.Println(a, *b) // Prints value.
* b = 100 // Updating value. (Dereferrencing pointers.)
```
## Functions
```go
func sayMessage(message string) {
  fmt.Println(message)
}

func sum(values ...int) int { // Variable number of args
  result := 0
  for _, v := range values {
    result += v
  }
  return result
}

func main() {
  sayMessage("Hello Go!")
}
```

## Interfaces
The way interfaces are implemented in Go, is one of the reasons why Go applications tend to be maintainable and scalable.
Interfaces do not describe data; instead they describe behaviour. It stores method definitions.

```go
package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Writer interface {
 	Write([] byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
```
### Composing interfaces
```go
type Writer interface {
 	Write([] byte) (int, error)
}
type Closer interface {
 	Close() error
}
type WriterCloser interface {
  Writer
  Closer
}
```

### Best Practices
- Use many, but small interfaces.
  - Single method interfaces are some of the most powerful and flexible.
- Do not export interfaces for types that will be consumed.
- Do export interfaces for types that will be used by package.
- Design functions and methods to receive interfaces whenever possible.

## Go Routines
- A routine enables to create efficent and highly concurrent applications.
- A Go Routine is an abstraction of a Thread.
- If you are writing a library, it is recommended to do not have routines.
- It is posible to check race conditions at compile time with:
  ```bash
    go run --race myProgram.go
  ```
- A routine is created using the keyword `go`
```go
go myRoutine
```