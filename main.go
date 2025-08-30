package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("🚀 Goroutine Concurrency vs Parallelism Benchmark")
	fmt.Println(strings.Repeat("=", 60))

	// Show system info
	fmt.Printf("CPU Cores: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("Go Version: %s\n", runtime.Version())
	fmt.Printf("OS/Arch: %s/%s\n\n", runtime.GOOS, runtime.GOARCH)

	// Warm up the system
	fmt.Println("🔥 Warming up...")
	warmUp()

	// Run multiple iterations for better accuracy
	fmt.Println("📊 Running benchmarks with multiple iterations...")

	// Test different workload types
	testCPUWorkImproved()
	testIOWorkImproved()
	testMixedWorkload()
	testScalability()
}

func warmUp() {
	// Run a quick warm-up to stabilize CPU frequency and caches
	var wg sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum := 0
			for j := 0; j < 1_000_000; j++ {
				sum += j * j
			}
		}()
	}
	wg.Wait()
	time.Sleep(100 * time.Millisecond)
}

func testCPUWorkImproved() {
	fmt.Println("\n📊 CPU-Intensive Tasks (Prime Number Calculation)")
	fmt.Println(strings.Repeat("-", 60))

	iterations := 5
	var concurrentTimes, parallelTimes []time.Duration

	// Run multiple iterations
	for i := 0; i < iterations; i++ {
		fmt.Printf("   Iteration %d/%d...\n", i+1, iterations)

		// Force garbage collection before each test
		runtime.GC()
		time.Sleep(10 * time.Millisecond)

		concurrentTime := runCPUTasksImproved(1)
		concurrentTimes = append(concurrentTimes, concurrentTime)

		runtime.GC()
		time.Sleep(10 * time.Millisecond)

		parallelTime := runCPUTasksImproved(runtime.NumCPU())
		parallelTimes = append(parallelTimes, parallelTime)
	}

	// Calculate averages and statistics
	avgConcurrent := average(concurrentTimes)
	avgParallel := average(parallelTimes)
	speedup := float64(avgConcurrent) / float64(avgParallel)
	efficiency := (speedup / float64(runtime.NumCPU())) * 100

	fmt.Printf("\n📈 CPU-Intensive Results (avg of %d runs):\n", iterations)
	fmt.Printf("   Concurrent:  %v (±%.1fms)\n", avgConcurrent, stdDev(concurrentTimes).Seconds()*1000)
	fmt.Printf("   Parallel:    %v (±%.1fms)\n", avgParallel, stdDev(parallelTimes).Seconds()*1000)
	fmt.Printf("   Speedup:     %.2fx\n", speedup)
	fmt.Printf("   Efficiency:  %.1f%%\n", efficiency)
	fmt.Printf("   Theoretical Max: %dx\n\n", runtime.NumCPU())
}

func testIOWorkImproved() {
	fmt.Println("💾 I/O-Intensive Tasks (Simulated Network Operations)")
	fmt.Println(strings.Repeat("-", 60))

	iterations := 5
	var concurrentTimes, parallelTimes []time.Duration

	for i := 0; i < iterations; i++ {
		fmt.Printf("   Iteration %d/%d...\n", i+1, iterations)

		runtime.GC()
		time.Sleep(10 * time.Millisecond)

		concurrentTime := runIOTasksImproved(1)
		concurrentTimes = append(concurrentTimes, concurrentTime)

		runtime.GC()
		time.Sleep(10 * time.Millisecond)

		parallelTime := runIOTasksImproved(runtime.NumCPU())
		parallelTimes = append(parallelTimes, parallelTime)
	}

	avgConcurrent := average(concurrentTimes)
	avgParallel := average(parallelTimes)
	speedup := float64(avgConcurrent) / float64(avgParallel)

	fmt.Printf("\n📈 I/O-Intensive Results (avg of %d runs):\n", iterations)
	fmt.Printf("   Concurrent:  %v (±%.1fms)\n", avgConcurrent, stdDev(concurrentTimes).Seconds()*1000)
	fmt.Printf("   Parallel:    %v (±%.1fms)\n", avgParallel, stdDev(parallelTimes).Seconds()*1000)
	fmt.Printf("   Speedup:     %.2fx\n", speedup)
	fmt.Printf("   Note: I/O tasks show minimal improvement with parallelism\n\n")
}

func testMixedWorkload() {
	fmt.Println("🔀 Mixed Workload (CPU + I/O)")
	fmt.Println(strings.Repeat("-", 60))

	concurrentTime := runMixedTasks(1)
	parallelTime := runMixedTasks(runtime.NumCPU())
	speedup := float64(concurrentTime) / float64(parallelTime)

	fmt.Printf("   Concurrent:  %v\n", concurrentTime)
	fmt.Printf("   Parallel:    %v\n", parallelTime)
	fmt.Printf("   Speedup:     %.2fx\n", speedup)
	fmt.Printf("   Note: Mixed workloads show moderate improvement\n\n")
}

func testScalability() {
	fmt.Println("📈 Scalability Test (Different Numbers of Goroutines)")
	fmt.Println(strings.Repeat("-", 60))

	goroutineCounts := []int{1, 2, 4, 8, 16}

	fmt.Printf("   Goroutines | Time     | Speedup\n")
	fmt.Printf("   -----------|----------|--------\n")

	baseTime := time.Duration(0)

	for _, count := range goroutineCounts {
		if count > runtime.NumCPU()*4 {
			continue // Skip if too many goroutines
		}

		duration := runScalabilityTest(count)
		if baseTime == 0 {
			baseTime = duration
		}

		speedup := float64(baseTime) / float64(duration)
		fmt.Printf("   %-10d | %-8v | %.2fx\n", count, duration, speedup)
	}
	fmt.Println()
}

func runCPUTasksImproved(maxProcs int) time.Duration {
	oldMaxProcs := runtime.GOMAXPROCS(maxProcs)
	defer runtime.GOMAXPROCS(oldMaxProcs)

	var wg sync.WaitGroup
	start := time.Now()

	// Use number of goroutines equal to CPU cores for better measurement
	numTasks := runtime.NumCPU()
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go cpuIntensiveTaskImproved(i, &wg)
	}

	wg.Wait()
	return time.Since(start)
}

func runIOTasksImproved(maxProcs int) time.Duration {
	oldMaxProcs := runtime.GOMAXPROCS(maxProcs)
	defer runtime.GOMAXPROCS(oldMaxProcs)

	var wg sync.WaitGroup
	start := time.Now()

	// Use more goroutines for I/O tasks to show concurrency benefit
	numTasks := runtime.NumCPU() * 2
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go ioIntensiveTaskImproved(i, &wg)
	}

	wg.Wait()
	return time.Since(start)
}

func runMixedTasks(maxProcs int) time.Duration {
	oldMaxProcs := runtime.GOMAXPROCS(maxProcs)
	defer runtime.GOMAXPROCS(oldMaxProcs)

	var wg sync.WaitGroup
	start := time.Now()

	numTasks := runtime.NumCPU()
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		if i%2 == 0 {
			go cpuIntensiveTaskImproved(i, &wg)
		} else {
			go ioIntensiveTaskImproved(i, &wg)
		}
	}

	wg.Wait()
	return time.Since(start)
}

func runScalabilityTest(numGoroutines int) time.Duration {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	start := time.Now()

	workPerGoroutine := 10_000_000 / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum := 0
			for j := 0; j < workPerGoroutine; j++ {
				sum += j * j
			}
		}()
	}

	wg.Wait()
	return time.Since(start)
}

func cpuIntensiveTaskImproved(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Calculate prime numbers - more realistic CPU work
	count := 0
	limit := 100_000

	for n := 2; n < limit; n++ {
		isPrime := true
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
	}

	// Don't print during benchmark for cleaner output
	_ = count
}

func ioIntensiveTaskImproved(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate realistic I/O pattern
	for i := 0; i < 20; i++ {
		// Simulate network request or file I/O
		time.Sleep(5 * time.Millisecond)

		// Small CPU work between I/O (like JSON parsing)
		sum := 0
		for j := 0; j < 50_000; j++ {
			sum += j
		}
	}
}

// Utility functions for statistics
func average(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		return 0
	}

	total := time.Duration(0)
	for _, d := range durations {
		total += d
	}
	return total / time.Duration(len(durations))
}

func stdDev(durations []time.Duration) time.Duration {
	if len(durations) <= 1 {
		return 0
	}

	avg := average(durations)
	variance := float64(0)

	for _, d := range durations {
		diff := float64(d - avg)
		variance += diff * diff
	}

	variance /= float64(len(durations) - 1)
	return time.Duration(variance)
}
