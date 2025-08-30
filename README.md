# 🚀 Go Goroutine Concurrency vs Parallelism Benchmark

A comprehensive benchmark suite that demonstrates when Go goroutines run concurrently versus in parallel, and how different workload types benefit from parallelism.

## 📋 Table of Contents

- [Overview](#overview)
- [Key Concepts](#key-concepts)
- [Benchmark Types](#benchmark-types)
- [Understanding Results](#understanding-results)

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

### Environment Variables

```bash
# Control Go runtime behavior
export GOMAXPROCS=4        # Limit to 4 cores
export GOGC=100           # Garbage collection frequency
export GODEBUG=gctrace=1  # Enable GC tracing

go run benchmark.go
```

## 📄 License

This benchmark is provided as educational material. Feel free to use, modify, and distribute for learning and development purposes.

---

**Happy benchmarking! 🚀**

*For questions or suggestions, please open an issue or contribute to the project.*
