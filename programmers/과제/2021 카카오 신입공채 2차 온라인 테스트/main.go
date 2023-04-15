package main

import (
	"challenge1/api"
	"fmt"
)

func main() {
	fmt.Println("response Body:", api.Start(1))
	fmt.Println("response Body:", api.Locations())
	fmt.Println("response Body:", api.Trucks())
	fmt.Println("response Body:", api.Simulate())
	fmt.Println("response Body:", api.Score())
}
