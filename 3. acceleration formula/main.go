package main

import "fmt"

func GetForceByAcceleration(mass float64, acc float64) float64 {
	force := mass * acc
	return force
}

func GetAccelByForce(mass float64, force float64) float64 {
	return float64(force / mass)
}

func main() {
	fmt.Println("ga jelas anjir gw ngoding apaan wkwk yg penting ngoding dah")
}
