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