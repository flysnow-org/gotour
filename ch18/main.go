package main

//缓存已经计算的结果
var cache = map[int]int{}

func Fibonacci(n int) int {
	if v, ok := cache[n]; ok {
		return v
	}
	result := 0

	switch {
	case n < 0:
		result = 0
	case n == 0:
		result = 0
	case n == 1:
		result = 1
	default:
		result = Fibonacci(n-1) + Fibonacci(n-2)
	}
	cache[n] = result
	return result
}
