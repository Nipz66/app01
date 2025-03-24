package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	var invenstmentAmount float64
	expectetdReturnRate, years := 5.5, 10.0

	fmt.Scan(&invenstmentAmount)

	futureValue := invenstmentAmount * math.Pow(1+expectetdReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
