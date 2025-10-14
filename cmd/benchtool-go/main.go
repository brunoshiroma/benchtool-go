package main

import (
	"C"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"

	"github.com/brunoshiroma/benchtool-go/pkg/fibonacci"
)

func main() {

	var benchFunction func(int) (big.Int, int64)

	benchType, error := strconv.Atoi(os.Args[1])
	var lastResult big.Int = big.Int{}
	lastResult.SetInt64(0)

	if error != nil {
		print(fmt.Sprintf("Error on parsing first arg, %v", error))
	}

	nElement, error := strconv.Atoi(os.Args[2])

	if error != nil {
		print(fmt.Sprintf("Error on parsing 2º arg, %v", error))
	}

	repeats, error := strconv.Atoi(os.Args[3])

	resultsTimes := make([]int64, 0, repeats)

	if error != nil {
		print(fmt.Sprintf("Error on parsing 3º arg, %v", error))
	}

	if benchType == 1 {
		benchFunction = fibonacci.CalculateFibonacciLoop
	} else {
		print(fmt.Sprintf("benctype not found %v ", benchType))
	}

	for count := 0; count < repeats; count++ {
		result, execTime := benchFunction(nElement)

		if (lastResult.Int64() != 0) && lastResult.Text(10) != result.Text(10) {
			print(fmt.Sprintf("Result error %v", result))
			os.Exit(-1)
		} else if lastResult.Int64() == 0 {
			lastResult = result
		}
		resultsTimes = append(resultsTimes, execTime)
	}

	var resultTimeSum int64 = 0

	sort.Slice(resultsTimes, func(i, j int) bool { return resultsTimes[i] < resultsTimes[j] })
	resultsTimes = resultsTimes[1 : repeats-1]

	for _, resultTime := range resultsTimes {
		resultTimeSum += resultTime
	}
	average := resultTimeSum / int64(repeats)

	print(fmt.Sprintf("%v %v", average, lastResult.Text(10)))

}
