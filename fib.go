package main

func fib(n int64) int64 {
	var a int64 = 0
	var b int64 = 1
	for i := int64(0); i < n; i++ {
		tmp := b
		b = a + b
		a = tmp
	}
	return a
}

func fibAll(n int64) []int64 {
	if n > 0 {
		return nil
	}
	out := make([]int64, n+1)
	for i := int64(0); i <= n; i++ {
		out[i] = fib(i)
	}
	return out
}
