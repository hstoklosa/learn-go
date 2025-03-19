**General Verbs:**

- **`%v`**: The default format for printing values.
- **`%+v`**: Like `%v` but adds field names when printing structs.
- **`%#v`**: A Go-syntax representation of the value.
- **`%T`**: Prints the type of the value.
- **`%%`**: A literal percent sign.

**Booleans:**

- **`%t`**: Prints the boolean value as `true` or `false`.

**Integers:**

- **`%b`**: Binary representation.
- **`%c`**: The character represented by the corresponding Unicode code point.
- **`%d`**: Decimal representation.
- **`%o`**: Octal representation.
- **`%q`**: Single-quoted character literal, safely escaped if necessary.
- **`%x`**: Hexadecimal representation with lower-case letters.
- **`%X`**: Hexadecimal representation with upper-case letters.

**Floating-Point and Complex Numbers:**

- **`%b`**: For floating-point numbers, it prints a representation of the number in binary exponent format (e.g., `-123456p-78`).
- **`%e`**: Scientific notation (e.g., `-1.234567e+08`).
- **`%E`**: Scientific notation with an uppercase `E` (e.g., `-1.234567E+08`).
- **`%f`**: Decimal point but no exponent (e.g., `123.456`).
- **`%F`**: Same as `%f`.
- **`%g`**: Uses `%e` or `%f` whichever is more compact for the given value.
- **`%G`**: Like `%g` but uses `%E` or `%F`.

**Strings:**

- **`%s`**: The uninterpreted string or slice of bytes.
- **`%q`**: A double-quoted string safely escaped with Go syntax.
- **`%x`**: A hex dump of the string, with each byte represented as two lower-case hex digits.
- **`%X`**: A hex dump with upper-case hex digits.

**Pointers:**

- **`%p`**: Base 16 notation, with leading `0x` for pointers.
