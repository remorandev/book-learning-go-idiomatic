# Composite Types

## Arrays

Like most programming languages, Go has arrays. However, arrays are rarely used directly in Go.

## Slices

A slice is a sequence of values. Each element in a slice is assigned to consecutive memory locations, which makes it
quick to read or write these values. The length of a slice is the number of consecutive memory locations that have been
assigned a value. Every slice also has a capacity, which is the number of consecutive memory locations reserved. This
can be larger than the length.

## Maps

Go provides a built-in data type for situations where you want to associate one value to another. The map type is
written as map[keyType]valueType.

A map returns the zero value if you ask for the value associated with a key that’s not in the map. This is handy when
implementing things like the totalWins counter you saw earlier. However, you sometimes do need to find out if a key is
in a map. Go provides the comma ok idiom to tell the difference between a key that’s associated with a zero value and a
key that’s not in the map.

#### Comparing Maps

Go 1.21 added a package to the standard library called maps that contains helper functions for working with maps.

#### Using Maps as Sets

Many languages include a set in their standard library. A set is a data type that ensures there is at most one of a
value, but doesn’t guarantee that the values are in any particular order.

Go doesn’t include a set, but you can use a map to simulate some of its features.

```go
stringSet := map[string]bool{}
```

Some people prefer to use struct{} for the value when a map is being used to implement a set. The advantage is that
an empty struct uses zero bytes, while a boolean uses one byte.

```go
stringSet := map[string]struct{}{}
```

### Structs

A struct type is defined with the keyword type, the name of the struct type, the keyword struct, and a pair of
braces ({}). Within the braces, you list the fields in the struct. Just as you put the variable name first and the
variable type second in a var declaration, you put the struct field name first and the struct field type second. Also
note that unlike in map literals, no commas separate the fields in a struct declaration. You can define a struct type
inside or outside of a function.

If you already know an object-oriented language, you might be wondering about the difference between classes and
structs. The difference is simple: Go doesn’t have classes, because it doesn’t have inheritance. This doesn’t mean that
Go doesn’t have some of the features of object-oriented languages it just does things a little differently.