package fibonacci

import (
	"math/big"
	"time"
)

//CalculateFibonacciLoop calculate Fibonacci with simple loop
//return
func CalculateFibonacciLoop(nElement int) (big.Int, int64) {
	start := time.Now()
	current := big.NewInt(0)
	next := big.NewInt(1)
	accumulator := big.NewInt(0)
	for count := 1; count < nElement; count++ {
		accumulator = next
		next = current.Add(current, next)
		current = accumulator
	}
	return *next, time.Since(start).Milliseconds()
}
