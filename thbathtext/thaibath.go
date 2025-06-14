package thbathtext

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

/*
- this function will round the decimal into 2 decimal place then split into 2 part 1.integer 2.fraction
- In order to convert Thai currency into text, each digits compose of 1.number 2.place. ex 432 = 400(สี่+ร้อย) + 30(สาม+สิบ) + 2(สอง)
- There are many specials cases for example 21 = 20 ("ยี่สิบ") + 1 ("เอ็ด"), 1 = 1(หนึ่ง)
*/
func FromDecimal(d decimal.Decimal, rounding ...string) (res string) {
	dstr := roundingMethod(d, rounding).String()
	parts := strings.Split(dstr, ".")
	if len(parts[0]) > 15 {
		fmt.Println("Decimal with digits lager than 15 digt may not precise")
	}
	hasFraction := len(parts) == 2 && len(parts[1]) != 0

	//1. integer part
	buildIntegerText(parts[0], &res)

	if res != "" {
		res += thaiBath
	}

	//2. fraction part
	if hasFraction {
		buildFractionText(parts[1], &res)
		res += fractionSuffix
	} else {
		if res == "" {
			res += thaiNumeral["0"] + thaiBath
		}
		res += noFractionSuffix
	}

	return res
}

func buildIntegerText(part string, res *string) {
	for i, c := range part {
		pos := len(part) - i - 1

		//1.1 numerical
		if pos%6 == 1 && string(c) == "1" {
			*res += ""
		} else if string(c) == "0" {
			*res += ""
		} else if pos%6 == 0 && string(c) == "1" && len(part) > 1 {
			*res += onePos0
		} else if pos%6 == 1 && string(c) == "2" {
			*res += twoPos1
		} else {
			*res += thaiNumeral[string(c)]
		}

		//1.2 numerical place
		if pos/6 > 0 && pos%6 == 0 {
			*res += strings.Repeat(thaiPowerOfTen[0], pos/6)
		} else if string(c) == "0" {
			*res += ""
		} else if pos%6 != 0 {
			*res += thaiPowerOfTen[(pos)%6]
		}
	}
}

func buildFractionText(part string, res *string) {
	isEd := true // ed = "เอ็ด" = onePos0
	for i, c := range part {
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
			*res += twoPos1
		} else if decimalPlace == 1 && string(c) == "1" {
			*res += "" //skip fraction
		} else if decimalPlace == 2 && string(c) == "1" && isEd {
			*res += onePos0
		} else {
			*res += thaiNumeral[string(c)]
		}
		//2.2 fraction place
		if decimalPlace == 1 && string(c) != "0" {
			*res += thaiPowerOfTen[1]
		}
	}
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
