package thbathtext

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

var testTable = []struct {
	d        decimal.Decimal
	expected string
}{
	{decimal.NewFromFloat(0), "ศูนย์บาทถ้วน"},

	{decimal.NewFromInt(1.0), "หนึ่งบาทถ้วน"},
	{decimal.NewFromInt(2), "สองบาทถ้วน"},
	{decimal.NewFromInt(12), "สิบสองบาทถ้วน"},
	{decimal.NewFromInt(21), "ยี่สิบเอ็ดบาทถ้วน"},
	{decimal.NewFromInt(112), "หนึ่งร้อยสิบสองบาทถ้วน"},
	{decimal.NewFromInt(121), "หนึ่งร้อยยี่สิบเอ็ดบาทถ้วน"},

	{decimal.NewFromFloat(0.1), "สิบสตางค์"},
	{decimal.NewFromFloat(0.01), "หนึ่งสตางค์"},
	{decimal.NewFromFloat(0.21), "ยี่สิบเอ็ดสตางค์"},

	{decimal.NewFromFloat(6.1), "หกบาทสิบสตางค์"},
	//depend on the requirement on converting this part
	{decimal.NewFromFloat(6.0), "หกบาทถ้วน"},
	{decimal.NewFromFloat(6.00), "หกบาทถ้วน"},
	{decimal.NewFromFloat(6.), "หกบาทถ้วน"},

	{decimal.NewFromFloat(1.004), "หนึ่งบาทถ้วน"},
	{decimal.NewFromFloat(1.005), "หนึ่งบาทหนึ่งสตางค์"},
	{decimal.NewFromFloat(1.105), "หนึ่งบาทสิบเอ็ดสตางค์"},
	{decimal.NewFromFloat(1.205), "หนึ่งบาทยี่สิบเอ็ดสตางค์"},

	{decimal.NewFromFloat(21.21), "ยี่สิบเอ็ดบาทยี่สิบเอ็ดสตางค์"},
	{decimal.NewFromFloat(21.01), "ยี่สิบเอ็ดบาทหนึ่งสตางค์"},
	{decimal.NewFromFloat(21.1), "ยี่สิบเอ็ดบาทสิบสตางค์"},
	{decimal.NewFromFloat(20.1), "ยี่สิบบาทสิบสตางค์"},

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
	{decimal.NewFromFloat(12300457802.08), "หนึ่งหมื่นสองพันสามร้อยล้านสี่แสนห้าหมื่นเจ็ดพันแปดร้อยสองบาทแปดสตางค์"},
	{decimal.NewFromFloat(42312300457802.08), "สี่สิบสองล้านล้านสามแสนหนึ่งหมื่นสองพันสามร้อยล้านสี่แสนห้าหมื่นเจ็ดพันแปดร้อยสองบาทแปดสตางค์"}, // 42,312,300,457,802.08

	{decimal.NewFromFloat(-1007.65), "ลบหนึ่งพันเจ็ดบาทหกสิบห้าสตางค์"},
	{decimal.NewFromFloat(80_003_000_300), "แปดหมื่นสามล้านสามร้อยบาทถ้วน"},
	{decimal.NewFromFloat(3000003000300), "สามล้านล้านสามล้านสามร้อยบาทถ้วน"},
	{decimal.NewFromFloat(3000001000000), "สามล้านล้านหนึ่งล้านบาทถ้วน"},
	{decimal.NewFromFloat(5000000000000), "ห้าล้านล้านบาทถ้วน"},
	{decimal.NewFromFloat(5_040_000_000_000), "ห้าล้านล้านสี่หมื่นล้านบาทถ้วน"},

}

var testTableForRoundDown = []struct {
	d        decimal.Decimal
	expected string
}{
	{decimal.NewFromFloat(12.34), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.340), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.341), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.342), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.343), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.344), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.345), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.346), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.347), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.348), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.349), "สิบสองบาทสามสิบสี่สตางค์"},
}

var testTableForRoundUp = []struct {
	d        decimal.Decimal
	expected string
}{
	{decimal.NewFromFloat(12.34), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.340), "สิบสองบาทสามสิบสี่สตางค์"},
	{decimal.NewFromFloat(12.341), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.342), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.343), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.344), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.345), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.346), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.347), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.348), "สิบสองบาทสามสิบห้าสตางค์"},
	{decimal.NewFromFloat(12.349), "สิบสองบาทสามสิบห้าสตางค์"},
}
func TestText(t *testing.T) {
	passCount := 0
	failCount := 0
	for _, tt := range testTable {
		result := FromDecimal(tt.d)
		if result != tt.expected {
			t.Errorf("%s = actual %s; want %s", tt.d.StringFixedBank(2), result, tt.expected)
			failCount++
		} else {
			passCount++
		}
	}
	fmt.Printf("\n✅ Passed: %d | ❌ Failed: %d\n", passCount, failCount)

}

func TestTextRoundDown(t *testing.T) {
	passCount := 0
	failCount := 0
	for _, tt := range testTableForRoundDown {
		result := FromDecimal(tt.d, "D")
		if result != tt.expected {
			t.Errorf("%s = actual %s; want %s", tt.d.StringFixedBank(2), result, tt.expected)
			failCount++
		} else {
			passCount++
		}
	}
	fmt.Printf("\n✅ Passed: %d | ❌ Failed: %d\n", passCount, failCount)

}

func TestTextRoundUp(t *testing.T) {
	passCount := 0
	failCount := 0
	for _, tt := range testTableForRoundUp {
		result := FromDecimal(tt.d, "U")
		if result != tt.expected {
			t.Errorf("%s = actual %s; want %s", tt.d.StringFixedBank(2), result, tt.expected)
			failCount++
		} else {
			passCount++
		}
	}
	fmt.Printf("\n✅ Passed: %d | ❌ Failed: %d\n", passCount, failCount)

}

func BenchmarkThaiBathText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testTable {
			_ = FromDecimal(tt.d)
		}
	}
}