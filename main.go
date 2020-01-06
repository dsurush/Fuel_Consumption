package main

import "fmt"

func minusPercent(x int) int {
	needPercent := 90
	needPercent = needPercent * x / 100
	return needPercent
}

func fuelConsumption(neededLiters int, haveLiters int)int{
	if haveLiters <= 0 || neededLiters <= 0 {
		return 0
	}
	canGo := haveLiters * 100 / neededLiters
	return minusPercent(canGo)
}

func main() {
	var needLiters int = 5
	var haveLiters int = 2
	restOfTheWay := fuelConsumption(needLiters, haveLiters)
	fmt.Println(restOfTheWay)
}
