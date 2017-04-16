package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	old := time.Now()
	var now time.Time

	i := 0
	for {
		i = (i + 1) % (1 << 30)

		fmt.Printf("\r%d", i)

		if i%1000000 == 0 {
			now = time.Now()
			fmt.Println("\t", now.Sub(old).Nanoseconds()/1000, "µs")
			old = now
		}
	}

	os.Exit(0)

	BenchMark(`IntegerAdd20`, IntegerAdd20)
	BenchMark(`FloatAdd20`, FloatAdd20)
	BenchMark(`FloatMultiply20`, FloatMultiply20)
}

func IntegerAdd20() {
	s := 0
	for i := 0; i < 1<<20; i++ {
		s += i
	}
}

func FloatAdd20() {
	s := 0.0
	for i := 0.0; i < 1<<20; i++ {
		s += i
	}
}

func FloatMultiply20() {
	s := 0.0
	for i := 0.0; i < 1<<20; i++ {
		s *= i
	}
}

func BenchMark(title string, fn func()) {

	numGoroutine := 10000

	cSum := make(chan int64, numGoroutine)

	for i := 0; i < numGoroutine; i++ {
		go func(c chan int64) {
			now := time.Now()
			fn()
			c <- time.Now().Sub(now).Nanoseconds()
		}(cSum)
	}

	total := int64(0)
	for i := 0; i < numGoroutine; i++ {
		time := <-cSum
		total += time
	}

	fmt.Printf("%s\t\tTotal: %5v s / %dop\tAve: %6v µs/op\n", title, total/1000000000, numGoroutine, total*1.0/int64(numGoroutine)/1000)

	time.Sleep(2 * time.Second)
}
