package thaibath

import (
	"strings"

	"github.com/shopspring/decimal"
)

func DecimalToThaiCurrencyText(d decimal.Decimal, rounding ...string) (res string) {

	dstr := roundingMethod(d, rounding[0]).String()
	parts := strings.Split(dstr, ".")
	hasFraction := 

	//integer part
	// 11,111,121,456,780.45   3214789.080
	for i, c := range parts[0] {
		pos := len(parts[0]) - i - 1

		//numerical
		if pos%6 == 1 && string(c) == "1" {
			res += ""
		} else if pos%6 == 0 && string(c) == "1" && len(parts[0]) > 1 {
			res += onePos0
		} else if pos%6 == 1 && string(c) == "2" {
			res += twoPos1
		} else {
			res += thaiNumeral[string(c)]
		}

		//numerical place
		if pos/6 > 0 && pos%6 == 0 {
			res += strings.Repeat(thaiPowerOfTen[0], pos/6)
		} else if pos%6 != 0 {
			res += thaiPowerOfTen[(pos)%6]
		}
	}
	res += thaiCurrency

	// fraction
	if hasFraction {
		for i, c := range parts[1] {
			//skip 0 on 1st decimal place
			if i == 0 && string(c) == "0" {
				continue
			}
			//fraction
			if i == 0 && string(c) == "2" {
				res += twoPos1
			} else if i == 1 && string(c) == "1" {
				res += onePos0
			} else {
				res += thaiNumeral[string(c)]
			}
			//fraction place
			if i == 0 && string(c) != "0" {
				res += thaiPowerOfTen[1]
			}
		}
		res += fractionSuffix
	} else {
		res += noFractionSuffix
	}

	return res
}

func roundingMethod(d decimal.Decimal, method ...string) decimal.Decimal {
	if len(method) == 0 {
		return d.RoundDown(2)
	}
	switch method[0] {
	case "U":
		return d.RoundUp(2)
	case "H":
		return d.Round(2)
	default:
		return d.RoundDown(2)
	}
}

func hasFraction() bool {
	if len(parts) == 2 && parts[2][0] != "0" {
		return
	}
	return false
}
// func validate(){
// 	errors.New(fmt.Sprintf("Invalid decimal format: %s", dstr))
// }
