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
	

	{decimal.NewFromFloat(0.1), "สิบสตางค์"},
	{decimal.NewFromFloat(0.01), "หนึ่งสตางค์"},
	{decimal.NewFromFloat(0.21), "ยี่สิบเอ็ดสตางค์"},

	{decimal.NewFromFloat(6.1), "หกบาทสิบสตางค์"},
	//depend on the requirement on translating this part
	{decimal.NewFromFloat(6.0), "หกบาทถ้วน"},
	{decimal.NewFromFloat(6.00), "หกบาทถ้วน"},
	{decimal.NewFromFloat(6.), "หกบาทถ้วน"},

	{decimal.NewFromFloat(1.004), "หนึ่งบาทถ้วน"},
	{decimal.NewFromFloat(1.005), "หนึ่งบาทหนึ่งสตางค์"},
	{decimal.NewFromFloat(1.105), "หนึ่งบาทสิบเอ็ดสตางค์"},
	{decimal.NewFromFloat(1.205), "หนึ่งบาทยี่สิบเอ็ดสตางค์"},

	{decimal.NewFromFloat(829.02), "แปดร้อยยี่สิบเก้าบาทสองสตางค์"},
	{decimal.NewFromFloat(829.03), "แปดร้อยยี่สิบเก้าบาทสามสตางค์"},
	{decimal.NewFromFloat(829.3), "แปดร้อยยี่สิบเก้าบาทสามสิบสตางค์"},

	{decimal.NewFromFloat(1829.2), "หนึ่งพันแปดร้อยยี่สิบเก้าบาทยี่สิบสตางค์"},
	{decimal.NewFromFloat(1829.20), "หนึ่งพันแปดร้อยยี่สิบเก้าบาทยี่สิบสตางค์"},
	{decimal.NewFromFloat(1829.21), "หนึ่งพันแปดร้อยยี่สิบเก้าบาทยี่สิบเอ็ดสตางค์"},
	{decimal.NewFromFloat(1829.22), "หนึ่งพันแปดร้อยยี่สิบเก้าบาทยี่สิบสองสตางค์"},

	{decimal.NewFromFloat(12.34), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.340), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.341), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.342), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.343), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.344), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.345), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.346), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.347), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.348), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.349), "สิบสองบาทสามสิบห้าสตางค์"},

	{decimal.NewFromFloat(3214789.08), "สามล้านสองแสนหนึ่งหมื่นสี่พันเจ็ดร้อยแปดสิบเก้าบาทแปดสตางค์"},
}

func TestText(t *testing.T) {
	passCount := 0
	failCount := 0
	for _, tt := range testTable {
		result := DecimalToThaiCurrencyText(tt.d)
		if result != tt.expected {
			t.Errorf("%s = actual %s; want %s", tt.d.String(), result, tt.expected)
			failCount++
		} else {
			passCount++
		}
	}
    fmt.Printf("\n✅ Passed: %d | ❌ Failed: %d\n", passCount, failCount)

}