# Go Types

### Booleans
- `bool`
### Integer Types
- `int8` `–128 to 127`
- `int16` `–32768 to 32767`
- `int32` `–2147483648 to 2147483647`
- `int64` `–9223372036854775808 to 9223372036854775807`
- `uint8` `0 to 255`
- `uint16` `0 to 65535`
- `uint32` `0 to 4294967295`
- `uint64` `0 to 18446744073709551615`
### Floating Point Types
- `float32` `IEEE-754 32-bit floating point numbers`
- `float64` `IEEE-754 64-bit floating point numbers`
### Integers Operators
- `+`, `-`, `*`, `/`, and `%`
- `+=`, `-=`, `*=`, `/=`, and `%=`
- `==`, `!=`, `>`, `>=`, `<`, and `<=`

### Const Variables
In many languages, constants are always written in all uppercase letters, with words separated by underscores 
(names like INDEX_COUNTER or NUMBER_TRIES). Go does not follow this pattern. This is because Go uses the case of the first letter in the name of a package-level declaration to determine if the item is accessible outside the package. 