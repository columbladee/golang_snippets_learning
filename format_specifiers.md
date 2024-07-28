## General
| Specifier | Description                              |
|-----------|------------------------------------------|
| `%v`      | Default format                           |
| `%+v`     | Include field names in structs           |
| `%#v`     | Go-syntax representation of the value    |
| `%T`      | Type of the value                        |
| `%%`      | Literal `%` character                    |

## Boolean
| Specifier | Description                              |
|-----------|------------------------------------------|
| `%t`      | The word `true` or `false`               |

## Integer Types
| Specifier | Description                              |
|-----------|------------------------------------------|
| `%b`      | Base 2 (binary)                          |
| `%c`      | Character represented by the corresponding Unicode code point |
| `%d`      | Base 10 (decimal)                        |
| `%o`      | Base 8 (octal)                           |
| `%O`      | Base 8 with leading 0o                   |
| `%q`      | Single-quoted character literal          |
| `%x`      | Base 16 (hexadecimal, lowercase)         |
| `%X`      | Base 16 (hexadecimal, uppercase)         |
| `%U`      | Unicode format: U+1234                   |

## Floating-Point and Complex Types
| Specifier | Description                              |
|-----------|------------------------------------------|
| `%b`      | Decimalless scientific notation with exponent a power of two |
| `%e`      | Scientific notation (e.g., -1234.456e+78)|
| `%E`      | Scientific notation (e.g., -1234.456E+78)|
| `%f`      | Decimal point, no exponent               |
| `%F`      | Synonym for `%f`                         |
| `%g`      | Use `%e` for large exponents, `%f` otherwise |
| `%G`      | Use `%E` for large exponents, `%F` otherwise |
| `%x`      | Hexadecimal notation (lowercase)         |
| `%X`      | Hexadecimal notation (uppercase)         |

## String and Slice of Bytes
| Specifier | Description                              |
|-----------|------------------------------------------|
| `%s`      | String or slice of bytes                 |
| `%q`      | Double-quoted string with Go syntax      |
| `%x`      | Base 16, with two characters per byte (lowercase) |
| `%X`      | Base 16, with two characters per byte (uppercase) |

## Pointer
| Specifier | Description                              |
|-----------|------------------------------------------|
| `%p`      | Base 16 notation, with leading 0x        |

## Width and Precision
- `%[width]d`: Minimum width
- `%.precision`: Precision for floating-point numbers
- `%[width].[precision]f`: Width and precision together

## Examples
```go
fmt.Printf("%d", 123)             // Decimal integer
fmt.Printf("%x", 123)             // Hexadecimal (lowercase)
fmt.Printf("%X", 123)             // Hexadecimal (uppercase)
fmt.Printf("%f", 123.456)         // Floating point
fmt.Printf("%9.2f", 123.456)      // Width 9, precision 2
fmt.Printf("%s", "string")        // String
fmt.Printf("%q", "string")        // Double-quoted string
fmt.Printf("%p", &a)              // Pointer
