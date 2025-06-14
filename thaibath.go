package thaibath

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

func decimalToThaiCurrencyText(d decimal.Decimal) (res string) {
	dstr := d.String()
	parts := strings.Split(dstr, ".")
	hasFraction := len(parts) == 2

	//integer part
	// 11,111,121,456,780.45   3214789.080
	for i, c := range parts[0] {
		pos := len(parts[0]) - i - 1

		if pos%6 == 1 && string(c) == "1" {
			res += ""
		} else if pos%6 == 0 && string(c) == "1" && len(parts[0]) > 1 {
			res += onePos0
		} else if pos%6 == 1 && string(c) == "2" {
			res += twoPos1
		} else {
			res += thaiNumeral[string(c)]
		}

		if pos/6 > 0 && pos%6 == 0 {
			res += strings.Repeat(thaiPowerOfTen[0], pos/6)
		} else if pos%6 != 0 {
			res += thaiPowerOfTen[(pos)%6]
		}
	}

	// fraction
	if hasFraction {
		res += thaiCurrency
		for i, c := range parts[1] {
			res += thaiNumeral[string(c)]
			fmt.Println(i, string(c))
		}
		res += fractionSuffix
	} else {
		res += thaiCurrency + noFractionSuffix
	}

	return res
}

// func validate(){
// 	errors.New(fmt.Sprintf("Invalid decimal format: %s", dstr))
// }
