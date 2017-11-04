package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	NCPU   = 7
	NITERS = 100000
)

func Avg(d time.Duration, v int) float64 {
	return float64(d.Nanoseconds()) / float64(v)
}

var ncpu, niters int

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.IntVar(&ncpu, "ncpu", NCPU, "GOMAXPROCS")
	flag.IntVar(&niters, "niters", NITERS, "ITERS")
	flag.Parse()
	wg := new(sync.WaitGroup)
	wg.Add(ncpu + 1)

	conn := make([]chan int, ncpu)
	for i := 0; i < ncpu; i++ {
		conn[i] = make(chan int, 100)
	}

	serverCode := func() {

		for i := 0; i < niters; i++ {
			for j := 0; j < ncpu; j++ {
				<-conn[j]
			}
		}

		wg.Done()
	}

	go serverCode()
	time.Sleep(100 * time.Millisecond)

	clientCode := func(i int) {

		for j := 0; j < niters; j++ {
			conn[i-1] <- (41 + i)
		}
		wg.Done()
	}

	run_startt := time.Now()
	for i := 1; i <= ncpu; i++ {
		go clientCode(i)
	}
	wg.Wait()
	run_endt := time.Now()
	fmt.Println(ncpu, "\t", Avg(run_endt.Sub(run_startt), niters))
}
