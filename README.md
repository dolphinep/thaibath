# thaibath
Convert Thai decimal number into readable th currency text


# Pre-caution :pencil2: 🌟
- **Supports 2 Decimal Places:** since thai currency is limited to 2 decimal places
- **Custom Rounding Options:** user can select to use round up/half/down method by passing U/H/D in a second parameter default is half round up method ex. 2.3 or 2.30 or 2.308 = "สองบาทสามสิบเอ็ดสตางค์"
- **Special Case Handling:** Understand the nuances of Thai currency text, like "ยี่สิบเอ็ดสตางค์." <ins>this may different from other lib, it's depending on the requirements</ins>
- **Zero Value Support:** Even zeros are gracefully handled, returning "ศูนย์บาทถ้วน."
- **Zero Fraction will be ignored** 3.0, 3.00, 3., 3 = "สามบาทถ้วน"


## How It Works 🎯

Using `thaibath`, convert your decimals to Thai currency text effortlessly. 

### Installation

```sh
go install github.com/dolphinep/thaibath@latest
```

### Usage

Just pass the decimal number and optional rounding method:

```go
package main

import (
    "fmt"
    "github.com/yourusername/thaibath"
)

func main() {
    fmt.Println(thaibath.FromDecimal(decimal.NewFromFloat(2.30))) // Output: สองบาทสามสิบเอ็ดสตางค์
    fmt.Println(thaibath.FromDecimal(decimal.NewFromFloat(21.237), "D")) // Output: ยี่สิบเอ็ดบาทยี่สิบสามสตางค์
}
```

## Examples 📊

- **Basic Conversion:** `2.30` → "สองบาทสามสิบสตางค์"
- **Custom Rounding:** `21.21` → "ยี่สิบเอ็ดบาทยี่สิบเอ็ดสตางค์"
- **More examples on thaibath_test.go**

## Testing
```
go test
```


---
