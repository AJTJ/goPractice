## Exported names

- capitalized names like `Pi` are exported.
- Thus, when importing a package, you must refer to its exported (capitalized) names.

## Some Basics

### variables

`var i bool`

- declares a list of variables with type being last
- generally for variable declaration at package or function level

`var i int = 1`

- variable with initializer
- type can be omitted since it takes the type of the initializer

`:=`

- var with implicit type
- only available inside a function (since at package level every statement begins with a keyword)

`i := 42` // int
`f := 3.142` // float64
`g := 0.867 + 0.5i` // complex128

`const`

- same as var, but can't be changed with `:=`

### Other Variable notes

 Variables declared without an explicit initial value are given their zero value.

The zero value is:

    0 for numeric types,
    false for the boolean type, and
    "" (the empty string) for strings.


### type conversion and inference

`X(y)`
type conversion expression converts `y` to type `X`
`var i int = 42`
`var f float64 = float64(i)`
`var u uint = uint(f)`

`var i int`
`j := i` // j is an int as well, then.

### functions

`func`

```
func add(x, y int) int {
  return x + y
}
```

- here, it takes two "arguments" of "int" type, and returns one result of "int" type

### TYPES

`string`
| string |

`byte`
| byte | (uint8) |
`rune`
| rune | (int32) |

`float`
| float32 | float64 |

`int`
| uint | uint8 | uint16 | uint32 | uint64 |
| uintptr |
| int | int8 | int16 | int32 | int64 |

`complex`
|complex64 |complex 128 |

represents a Unicode code point

## Structs
```
type Vertex struct {
	X int
	Y int
}
```
```
d := Vertex{1, 2}
```

## Arrays

The type [n]T is an array of n values of type T. 

```
var a [2]string
```
```
twoWords := [2]string{"blah", "woof"}
```
- array length is part of its type