
# jpnumber

`jpnumber` is a Go package that provides a function to convert integers into their Japanese number representations. It supports numbers up to 99,999,999,999,999,999 (99 quadrillion).

## Features

- Converts integers to Japanese number representation using kanji characters.
- Handles numbers up to 99,999,999,999,999,999 (max value: 99 quadrillion).
- Outputs Japanese numbers in the standard format, omitting the '一' in '一十' and other conventions.

## Installation

To install the package, simply run:

```bash
go get github.com/Yamaguchi-Katsuya/jpnumber.git
```

## Usage

To convert an integer into its Japanese number representation, use the `ToJPNumber` function. Here’s how you can use it:

```go
package main

import (
	"fmt"
	
	"github.com/Yamaguchi-Katsuya/jpnumber"
)

func main() {
	num := 123456789
	result, err := jpnumber.ToJPNumber(num)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Japanese number:", result)
	}
}
```

### Example Outputs

```go
ToJPNumber(0) → "零"
ToJPNumber(1) → "一"
ToJPNumber(10) → "十"
ToJPNumber(100) → "百"
ToJPNumber(1000) → "千"
ToJPNumber(123456789) → "一億二千三百四十五万六千七百八十九"
ToJPNumber(99999999999999999) → "九十九京九千九百九十九兆九千九百九十九億九千九百九十九万九千九百九十九"
```

## Error Handling

- If the input number exceeds the maximum allowed value (99,999,999,999,999,999), the function will return an error.
- If the input number is zero, it returns the Japanese word for zero ("零").

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
