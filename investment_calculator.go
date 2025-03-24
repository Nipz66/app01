package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	var invenstmentAmount float64
	expectetdReturnRate := 5.5
	var years float64

	fmt.Print("invenstment Amount = ")
	fmt.Scan(&invenstmentAmount)

	fmt.Print("Year = ")
	fmt.Scan(&years)

	futureValue := invenstmentAmount * math.Pow(1+expectetdReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
