package thaibath

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

func DecimalToThaiCurrencyText(d decimal.Decimal, rounding ...string) (res string) {
	fmt.Println("Decimal", d.Round(2).String())
	dstr := roundingMethod(d, rounding).String()
	fmt.Println("Decimal Round: ", dstr)
	parts := strings.Split(dstr, ".")
	// 6. = "หกบาทถ้วน"
	hasFraction := len(parts) == 2 && len(parts[1]) != 0

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
		isEd := true // 0.21 = "ยี่สิบเอ็ดสตางค์"
		for i, c := range parts[1] {
			decimalPlace := i + 1
			if decimalPlace == 1 && string(c) == "0" {
				isEd = false
			}
			//skip 0 fraction on 1st decimal place
			if decimalPlace == 1 && string(c) == "0" {
				continue
			}
			//fraction
			if decimalPlace == 1 && string(c) == "2" {
				res += twoPos1
			} else if decimalPlace == 1 && string(c) == "1" {
				res += "" //skip fraction
			} else if decimalPlace == 2 && string(c) == "1" && isEd {
				res += onePos0
			} else {
				res += thaiNumeral[string(c)]
			}
			//fraction place
			if decimalPlace == 1 && string(c) != "0" {
				res += thaiPowerOfTen[1]
			}
		}
		res += fractionSuffix
	} else {
		res += noFractionSuffix
	}

	return res
}


func roundingMethod(d decimal.Decimal, method []string) decimal.Decimal {
	if len(method) == 0 {
		return d.Round(2)
	}
	switch method[0] {
	case "U":
		return d.RoundUp(2)
	case "D":
		return d.RoundDown(2)
	default:
		return d.Round(2)
	}
}
