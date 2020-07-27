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
