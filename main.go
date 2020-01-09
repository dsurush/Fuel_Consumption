package main

func minusPercent(x int) int {
	needPercent := 90
	needPercent = needPercent * x / 100
	return needPercent
}

func fuelConsumption(neededLiters int, haveLiters int)int{
	if haveLiters <= 0 || neededLiters <= 0 {
		return 0
	}
	const distanceCanGoByNeedLiters = 100
	canGo := haveLiters * distanceCanGoByNeedLiters / neededLiters
	return minusPercent(canGo)
}

func main() {
	needLitersByHundredKm := 5
	haveLiters := 2
	fuelConsumption(needLitersByHundredKm, haveLiters)

}
