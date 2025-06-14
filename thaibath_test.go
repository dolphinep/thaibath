package thaibath

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

var testTable = []struct {
	d        decimal.Decimal
	expected string
}{
	{decimal.NewFromInt(2), "สองบาทถ้วน"},
	{decimal.NewFromInt(12), "สิบสองบาทถ้วน"},
	{decimal.NewFromInt(112), "หนึ่งร้อยสิบสองบาทถ้วน"},
	{decimal.NewFromInt(121), "หนึ่งร้อยยี่สิบเอ็ดบาทถ้วน"},
	{decimal.NewFromFloat(829.03), "แปดร้อยยี่สิบเก้าบาทสามสตางค์"},
	{decimal.NewFromFloat(829.3), "แปดร้อยยี่สิบเก้าบาทสามสิบสตางค์"},
	{decimal.NewFromFloat(112.34), "หนึ่งร้อยสิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(3214789.08), "สามล้านสองแสนหนึ่งหมื่นสี่พันเจ็ดร้อยแปดสิบเก้าบาทแปดสตางค์"},
}

func TestText(t *testing.T) {

	passCount := 0
	failCount := 0
	for _, tt := range testTable {
		result := decimalToThaiCurrencyText(tt.d)
		if result != tt.expected {
			t.Errorf("%s = actual %s; want %s", tt.d.String(), result, tt.expected)
			failCount++
		} else {
			passCount++
		}
	}
    fmt.Printf("\n✅ Passed: %d | ❌ Failed: %d\n", passCount, failCount)

}