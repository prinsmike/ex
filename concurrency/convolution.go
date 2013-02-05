// http://www.nada.kth.se/~snilsson/concurrency/
package main

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func init() {
	log.SetFlags(0) // no extra info in log messages
	//log.SetOutput(ioutil.Discard) // turns off logging

	numcpu := runtime.NumCPU()
	log.Println("CPU count:", numcpu)
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}

func main() {
	v := make(Vector, 32*1000)
	for i := range v {
		v[i] = 1
	}

	before := time.Now()
	w := Convolve(v, v)
	fmt.Println("time:", time.Now().Sub(before))

	fmt.Println(w)
}

type Vector []float64

func (v Vector) String() string {
	sb := new(bytes.Buffer)

	n, trunc := len(v), ""
	if n > 8 {
		n, trunc = 8, "..."
	}
	for i := 0; i < n; i++ {
		fmt.Fprintf(sb, "%.4g ", v[i])
	}
	fmt.Fprint(sb, trunc)

	return sb.String()
}

// Convolve computes w = u * v, where w[k] = Σ u[i]*v[j], i + j = k.
// Precondition: len(u) > 0, len(v) > 0.
func Convolve(u, v Vector) (w Vector) {
	n := len(u) + len(v) - 1
	w = make(Vector, n)
	log.Println("vector size:", n)

	// Divide w into work units that take ~100μs-1ms to compute.
	size := max(1, 1<<20/n)
	log.Println("work unit size:", size)

	wg := new(sync.WaitGroup)
	wg.Add(1 + (n-1)/size)
	for i := 0; i < n && i >= 0; i += size { // i < 0 after int overflow
		j := i + size
		if j > n || j < 0 { // j < 0 after int overflow
			j = n
		}
		// These goroutines share memory, but only for reading.
		go func(i, j int) {
			//before := time.Now()
			for k := i; k < j; k++ {
				w[k] = mul(u, v, k)
			}
			//log.Printf("time (%d-%d): %v\n", i, j, time.Now().Sub(before))
			wg.Done()
		}(i, j)
	}
	wg.Wait()
	return
}

// Returns Σ u[i]*v[j], i + j = k.
func mul(u, v Vector, k int) (res float64) {
	n := min(k+1, len(u))
	j := min(k, len(v)-1)
	for i := k - j; i < n; i, j = i+1, j-1 {
		res += u[i] * v[j]
	}
	return
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
