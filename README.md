# 🚀 Go Goroutine Concurrency vs Parallelism Benchmark

A comprehensive benchmark suite that demonstrates when Go goroutines run concurrently versus in parallel, and how different workload types benefit from parallelism.

## 📋 Table of Contents

- [Overview](#overview)
- [Key Concepts](#key-concepts)
- [Installation & Usage](#installation--usage)
- [Benchmark Types](#benchmark-types)
- [Understanding Results](#understanding-results)
- [System Requirements](#system-requirements)
- [Advanced Usage](#advanced-usage)

## 🎯 Overview

This benchmark helps you understand:
- **When goroutines actually run in parallel** vs just concurrently
- **Performance impact** of `GOMAXPROCS` settings
- **Different workload characteristics** (CPU-bound vs I/O-bound)
- **Scalability patterns** with varying numbers of goroutines
- **Real-world performance implications** for Go applications

## 🧠 Key Concepts

### Concurrency vs Parallelism
- **Concurrency**: Multiple tasks making progress by taking turns (time-slicing)
- **Parallelism**: Multiple tasks literally running simultaneously on different CPU cores

### GOMAXPROCS Impact
- `GOMAXPROCS=1`: All goroutines share one OS thread (concurrency only)
- `GOMAXPROCS=N`: Up to N goroutines can run simultaneously (true parallelism)

## 🚀 Installation & Usage

### Quick Start

```bash
# Clone or download the benchmark
git clone <repository-url>
cd goroutine-benchmark

# Run the benchmark
go run benchmark.go
```

### Alternative: Copy & Paste

1. Copy the benchmark code into a file named `benchmark.go`
2. Run: `go run benchmark.go`

### No Dependencies Required
- Uses only Go standard library
- Works with Go 1.16+
- Cross-platform (Windows, macOS, Linux)

## 📊 Benchmark Types

### 1. CPU-Intensive Tasks
**What it tests**: Prime number calculations
```go
// Heavy mathematical computations that benefit from parallelism
func cpuIntensiveTask() {
  // Calculate prime numbers up to 100,000
}
```

**Expected Results**:
- **Concurrency**: ~1.2s
- **Parallelism**: ~300ms
- **Speedup**: 3-4x on quad-core systems

### 2. I/O-Intensive Tasks
**What it tests**: Simulated network/file operations
```go
// I/O operations that benefit from concurrency but not parallelism
func ioIntensiveTask() {
  // Simulate network requests with sleep
  time.Sleep(5 * time.Millisecond)
}
```

**Expected Results**:
- **Concurrency**: ~200ms
- **Parallelism**: ~200ms
- **Speedup**: ~1x (no improvement)

### 3. Mixed Workload
**What it tests**: Combination of CPU and I/O tasks
- Demonstrates real-world application behavior
- Shows moderate parallelism benefits

### 4. Scalability Test
**What it tests**: Performance with different numbers of goroutines
- Tests 1, 2, 4, 8, 16 goroutines
- Shows optimal goroutine count for your system

## 📈 Understanding Results

### Sample Output
```
🚀 Enhanced Goroutine Concurrency vs Parallelism Benchmark
============================================================
CPU Cores: 8
GOMAXPROCS: 8

📊 CPU-Intensive Tasks (Prime Number Calculation)
------------------------------------------------------------
📈 CPU-Intensive Results (avg of 5 runs):
 Concurrent:  1.234s (±45.2ms)
 Parallel:    312ms (±12.1ms)
 Speedup:     3.95x
 Efficiency:  49.4%
 Theoretical Max: 8x
```

### Key Metrics Explained

| Metric | Description |
|--------|-------------|
| **Concurrent** | Time with `GOMAXPROCS=1` (concurrency only) |
| **Parallel** | Time with `GOMAXPROCS=all cores` (true parallelism) |
| **Speedup** | How much faster parallel execution is |
| **Efficiency** | Percentage of theoretical maximum speedup achieved |
| **±Standard Deviation** | Measurement consistency across runs |

### Performance Expectations

#### CPU-Bound Tasks ✅
- **High speedup** (2-8x depending on cores)
- **Linear scaling** with more cores
- **Examples**: Image processing, mathematical calculations, data compression

#### I/O-Bound Tasks ⚠️
- **Minimal speedup** (~1x)
- **Concurrency sufficient** for performance
- **Examples**: HTTP requests, database queries, file operations

#### Mixed Workloads 🔀
- **Moderate speedup** (1.5-3x)
- **Most real-world applications** fall here
- **Balance** between CPU and I/O operations

## 💻 System Requirements

### Minimum Requirements
- **Go**: 1.16 or later
- **RAM**: 512MB available
- **CPU**: Any modern processor

### Optimal Testing Environment
- **Multi-core CPU**: 4+ cores for meaningful parallelism testing
- **Stable system load**: Close other applications during benchmarking
- **Consistent power**: Disable CPU frequency scaling for accurate results

### Platform-Specific Optimizations

#### Linux
```bash
# Disable CPU frequency scaling for consistent results
sudo cpupower frequency-set --governor performance

# Set high process priority
sudo nice -n -20 go run benchmark.go
```

#### macOS
```bash
# Set high process priority
sudo nice -n -20 go run benchmark.go
```

#### Windows
```cmd
# Run as Administrator for best results
# Set high priority in Task Manager if needed
```

## 🔧 Advanced Usage

### Custom Benchmark Parameters

Modify these constants in the code for different test characteristics:

```go
// CPU workload intensity
const PrimeLimit = 100_000  // Increase for longer CPU tests

// I/O simulation timing
const IODelay = 5 * time.Millisecond  // Adjust I/O wait time

// Number of test iterations
const Iterations = 5  // More iterations = better accuracy
```

### Integration with Go Benchmark Framework

Create a `benchmark_test.go` file:

```go
package main

import (
  "testing"
  "runtime"
)

func BenchmarkCPUConcurrent(b *testing.B) {
  runtime.GOMAXPROCS(1)
  for i := 0; i < b.N; i++ {
      runCPUTasksImproved(1)
  }
}

func BenchmarkCPUParallel(b *testing.B) {
  runtime.GOMAXPROCS(runtime.NumCPU())
  for i := 0; i < b.N; i++ {
      runCPUTasksImproved(runtime.NumCPU())
  }
}
```

Run with:
```bash
go test -bench=. -benchmem -count=5
```

### Environment Variables

```bash
# Control Go runtime behavior
export GOMAXPROCS=4        # Limit to 4 cores
export GOGC=100           # Garbage collection frequency
export GODEBUG=gctrace=1  # Enable GC tracing

go run benchmark.go
```

## 📚 Educational Use Cases

### For Learning Go Concurrency
- **Visualize** the difference between concurrency and parallelism
- **Understand** when to use goroutines effectively
- **Learn** about GOMAXPROCS impact on performance

### For Performance Analysis
- **Identify** whether your workload is CPU-bound or I/O-bound
- **Optimize** goroutine usage in your applications
- **Benchmark** different concurrency patterns

### For System Architecture
- **Determine** optimal server configurations
- **Understand** scaling characteristics
- **Plan** resource allocation for Go services

## 🔍 Troubleshooting

### Common Issues

#### Low or No Speedup
```
Problem: Parallel execution shows minimal improvement
Causes:
- System has only 1 CPU core
- Other processes consuming CPU
- Workload is I/O-bound, not CPU-bound
- Insufficient work per goroutine
```

#### Inconsistent Results
```
Problem: High standard deviation in measurements
Solutions:
- Close other applications
- Run multiple iterations
- Use performance CPU governor (Linux)
- Increase workload size
```

#### Memory Issues
```
Problem: Out of memory errors
Solutions:
- Reduce number of goroutines
- Decrease workload size per goroutine
- Monitor with: go run -race benchmark.go
```

### Performance Tips

1. **Warm up the system** before benchmarking
2. **Run multiple iterations** for statistical significance
3. **Force garbage collection** between tests
4. **Use consistent system load** during testing
5. **Monitor CPU temperature** to avoid throttling

## 📊 Interpreting Your Results

### Good Parallelism (CPU-bound)
```
Speedup: 3.8x on 4-core system
Efficiency: 95%
Standard deviation: <10% of mean
```

### Expected Concurrency (I/O-bound)
```
Speedup: 1.1x
Efficiency: N/A (not applicable)
Consistent timing across runs
```

### System Bottlenecks
```
Speedup: <50% of expected
High standard deviation
Possible thermal throttling or resource contention
```

### Go Concurrency Resources
- [Go Concurrency Patterns](https://golang.org/doc/effective_go.html#concurrency)
- [The Go Memory Model](https://golang.org/ref/mem)
- [Go Runtime Scheduler](https://golang.org/src/runtime/proc.go)

### Performance Analysis
- [Go Performance Tuning](https://golang.org/doc/diagnostics.html)
- [pprof Profiling](https://golang.org/pkg/net/http/pprof/)
- [Benchmarking Best Practices](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)

## 📄 License

This benchmark is provided as educational material. Feel free to use, modify, and distribute for learning and development purposes.

---

**Happy benchmarking! 🚀**

*For questions or suggestions, please open an issue or contribute to the project.*