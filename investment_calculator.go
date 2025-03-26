package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

var invenstmentAmount float64
var expectetdReturnRate = 5.5
var years float64

func main() {

	// fmt.Print("invenstment Amount = ")
	outputText("invenstment Amount = ")
	fmt.Scan(&invenstmentAmount)

	// fmt.Print("Year = ")
	outputText("Year = ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(invenstmentAmount, years, expectetdReturnRate, inflationRate)

	// futureValue := invenstmentAmount * math.Pow(1+expectetdReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value : ", futureValue)
	// fmt.Println("Future Real Value : ", futureRealValue)

	formatFV := fmt.Sprintf("Future Value : %.1f \n", futureValue)
	formatRFV := fmt.Sprintf("Future Real Value : %.1f ", futureRealValue)

	fmt.Print(formatFV, formatRFV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(invenstmentAmount, years, expectetdReturnRate, inflationRate float64) (float64, float64) {
	fv := invenstmentAmount * math.Pow(1+expectetdReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)
	return fv, frv

}
