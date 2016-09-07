package main

import (
	"fmt"
	"math"
)
const DELTA = .0001

func newtM(x float64) float64 { 
	result, difference := 1.0, 1.0
	for difference > DELTA {
		oldResult := result
		result = oldResult - ( ( ( oldResult*oldResult ) - x) / (2*oldResult) );
		difference = math.Abs(oldResult - result)
	}
	return result
}

func newtMRec(x float64, result float64, difference float64) float64 { 
	if difference <= DELTA {
		return result
	} else{
		oldResult := result
		result = oldResult - ( ( ( oldResult*oldResult ) - x) / (2*oldResult) );
		difference = math.Abs(oldResult - result)
		return newtMRec(x,result,difference)
	}
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))
	x, z := 2.0,1.0
	fmt.Println( newtM(2.0) )
	fmt.Println( newtMRec( x,z,z ) )
}
