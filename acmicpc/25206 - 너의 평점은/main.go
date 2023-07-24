package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	m := map[string]float32{
		"A+": 4.5,
		"A0": 4.0,
		"B+": 3.5,
		"B0": 3.0,
		"C+": 2.5,
		"C0": 2.0,
		"D+": 1.5,
		"D0": 1.0,
		"F":  0,
	}
	var name, grade string
	var credit, gpa, totalCredit float32
	for i := 0; i < 20; i++ {
		Fscanf(r, "%s %f %s\n", &name, &credit, &grade)
		if grade == "P" {
			continue
		}
		totalCredit += credit
		gpa += credit * m[grade]
	}
	Fprint(w, gpa/totalCredit)
}
