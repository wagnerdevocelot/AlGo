package main

func fibonacci(numero int) int {
	if numero <= 1 {
		return 1
	}
	return fibonacci(numero-1) + fibonacci(numero-2)
}
