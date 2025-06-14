package thaibath

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

/*
- this function will round the decimal into 2 decimal place then split into 2 part 1.integer 2.fraction
- In order to convert Thai currency into text, each digits compose of 1.number 2.place. ex 432 = 400(สี่+ร้อย) + 30(สาม+สิบ) + 2(สอง)
- There are many specials cases for example 21 = 20 ("ยี่สิบ") + 1 ("เอ็ด"), 1 = 1(หนึ่ง)
*/
func DecimalToThaiCurrencyText(d decimal.Decimal, rounding ...string) (res string) {
	dstr := roundingMethod(d, rounding).String()
	fmt.Println("Decimal Round: ", dstr)
	parts := strings.Split(dstr, ".")
	hasFraction := len(parts) == 2 && len(parts[1]) != 0

	//integer part
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

	//ignore converting 0 if it's the only number ex. 0.02 = "สองสตางค์"
	if res == thaiNumeral["0"] {
		res = ""
	} else {
		res += thaiBath
	}

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
		if res == "" {
			res += thaiNumeral["0"] + thaiBath
		}
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
