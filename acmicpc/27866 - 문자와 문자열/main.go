package main

import . "fmt"

func main() {
	S, i := "", 0
	Scan(&S)
	Scan(&i)
	Print(string(S[i-1]))
}
