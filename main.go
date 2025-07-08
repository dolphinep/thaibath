package main

import (
	"fmt"
	"github.com/dolphinep/thaibath/thbathtext"
	"github.com/shopspring/decimal"
)
func main(){
	inputs := []decimal.Decimal{
		// decimal.NewFromFloat(1234),
		// decimal.NewFromFloat(33333.75),
		// decimal.NewFromFloat(1.004),
		// decimal.NewFromFloat(1.005),
		// decimal.NewFromFloat(21.01),
		// decimal.NewFromFloat(171.1),
		// decimal.NewFromFloat(3214789.08),
		// decimal.NewFromFloat(42312300457802.08),
		

		// decimal.NewFromFloat(1_000_000),

        decimal.NewFromFloat(-1007.65),
        decimal.NewFromFloat(80_003_000_300),
        decimal.NewFromFloat(3000003000300),
        decimal.NewFromFloat(3000001000000),
        decimal.NewFromFloat(5000000000000),
        decimal.NewFromFloat(5_040_000_000_000),
	}
	for _, input := range inputs {
		fmt.Println(thbathtext.FromDecimal(input))
	}
}