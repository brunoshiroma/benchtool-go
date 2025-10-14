package main

import (
	"C"
	"fmt"
	"math/big"
	"os"
	"sort"

	"github.com/brunoshiroma/benchtool-go/pkg/fibonacci"
)

//export benchtoolGoCall
func benchtoolGoCall(iteration int, repeats int, bench_type int) *C.char {
	var lastResult big.Int = big.Int{}
	lastResult.SetInt64(0)
	resultsTimes := make([]int64, 0, repeats)

	var benchFunction func(int) (big.Int, int64)
	if bench_type == 1 {
		benchFunction = fibonacci.CalculateFibonacciLoop
	} else {
		print(fmt.Sprintf("benctype not found %v ", bench_type))
	}

	for count := 0; count < repeats; count++ {
		result, execTime := benchFunction(iteration)

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

	return C.CString(fmt.Sprintf("%v %v", average, lastResult.Text(10)))
}

func main() {}
