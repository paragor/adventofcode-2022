package main

import (
	"fmt"
	"image"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type CircleArea struct {
	Radios int
	Center image.Point
}

func main() {
	pairs := [][]image.Point{}
	maxPoint := image.Point{X: math.MinInt, Y: math.MinInt}
	minPoint := image.Point{X: math.MaxInt, Y: math.MaxInt}
	re := regexp.MustCompilePOSIX("Sensor at x=(-?[0-9]+), y=(-?[0-9]+): closest beacon is at x=(-?[0-9]+), y=(-?[0-9]+)")
	for _, inputStr := range strings.Split(inputA, "\n") {
		match := re.FindStringSubmatch(inputStr)
		if len(match) != 4+1 {
			panic("wtf")
		}

		pair := []image.Point{
			{X: mustAtoI(match[1]), Y: mustAtoI(match[2])},
			{X: mustAtoI(match[3]), Y: mustAtoI(match[4])},
		}
		pairs = append(pairs, pair)
		maxPoint.X = max(max(maxPoint.X, pair[0].X), pair[1].X)
		maxPoint.Y = max(max(maxPoint.Y, pair[0].X), pair[1].Y)

		minPoint.X = min(min(minPoint.X, pair[0].X), pair[1].X)
		minPoint.Y = min(min(minPoint.Y, pair[0].X), pair[1].Y)
	}
	areas := []CircleArea{}
	maxRadios := 0
	for _, pair := range pairs {
		area := CircleArea{
			Radios: abs(pair[1].Y-pair[0].Y) + abs(pair[1].X-pair[0].X),
			Center: pair[0],
		}
		areas = append(areas, area)
		maxRadios = max(maxRadios, area.Radios+1)
	}
	addX := maxRadios
	addY := maxRadios
	if minPoint.X < 0 {
		addX -= minPoint.X
	}
	if minPoint.Y < 0 {
		addY -= minPoint.Y
	}
	maxPoint.X = maxPoint.X + addX
	minPoint.X = 0
	maxPoint.Y = maxPoint.Y + addY
	minPoint.Y = 0
	for _, pair := range pairs {
		pair[0].X += addX
		pair[0].Y += addY
		pair[1].X += addX
		pair[1].Y += addY
	}
	for i := range areas {
		areas[i].Center.X += addX
		areas[i].Center.Y += addY
	}

	searchY := 10 + addY
	searchY = 2000000 + addY

	maxSide := max(maxPoint.X+maxRadios, maxPoint.Y+maxRadios)
	maxSide = maxSide + (maxSide % 2)
	line := make([]uint8, maxSide)

	for _, area := range areas {
		for x := range line {
			diff := abs(area.Center.Y-searchY) + abs(area.Center.X-x)
			if diff <= area.Radios {
				line[x] = 1
			}
		}
	}
	for _, pair := range pairs {
		if pair[0].Y == searchY {
			line[pair[0].X] = 7
		}
		if pair[1].Y == searchY {
			line[pair[1].X] = 0
		}
	}
	result := 0
	for _, x := range line {
		result += int(x)
	}
	fmt.Println(result)
	fmt.Println("addX", addX, "addY", addY)

}

func mustAtoI(a string) int {
	res, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var input = strings.TrimSpace(`
Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`)
var inputDebug = strings.TrimSpace(`
Sensor at x=8, y=7: closest beacon is at x=2, y=10
`)
var inputA = strings.TrimSpace(`
Sensor at x=3772068, y=2853720: closest beacon is at x=4068389, y=2345925
Sensor at x=78607, y=2544104: closest beacon is at x=-152196, y=4183739
Sensor at x=3239531, y=3939220: closest beacon is at x=3568548, y=4206192
Sensor at x=339124, y=989831: closest beacon is at x=570292, y=1048239
Sensor at x=3957534, y=2132743: closest beacon is at x=3897332, y=2000000
Sensor at x=1882965, y=3426126: closest beacon is at x=2580484, y=3654136
Sensor at x=1159443, y=3861139: closest beacon is at x=2580484, y=3654136
Sensor at x=2433461, y=287013: closest beacon is at x=2088099, y=-190228
Sensor at x=3004122, y=3483833: closest beacon is at x=2580484, y=3654136
Sensor at x=3571821, y=799602: closest beacon is at x=3897332, y=2000000
Sensor at x=2376562, y=1539540: closest beacon is at x=2700909, y=2519581
Sensor at x=785113, y=1273008: closest beacon is at x=570292, y=1048239
Sensor at x=1990787, y=38164: closest beacon is at x=2088099, y=-190228
Sensor at x=3993778, y=3482849: closest beacon is at x=4247709, y=3561264
Sensor at x=3821391, y=3986080: closest beacon is at x=3568548, y=4206192
Sensor at x=2703294, y=3999015: closest beacon is at x=2580484, y=3654136
Sensor at x=1448314, y=2210094: closest beacon is at x=2700909, y=2519581
Sensor at x=3351224, y=2364892: closest beacon is at x=4068389, y=2345925
Sensor at x=196419, y=3491556: closest beacon is at x=-152196, y=4183739
Sensor at x=175004, y=138614: closest beacon is at x=570292, y=1048239
Sensor at x=1618460, y=806488: closest beacon is at x=570292, y=1048239
Sensor at x=3974730, y=1940193: closest beacon is at x=3897332, y=2000000
Sensor at x=2995314, y=2961775: closest beacon is at x=2700909, y=2519581
Sensor at x=105378, y=1513086: closest beacon is at x=570292, y=1048239
Sensor at x=3576958, y=3665667: closest beacon is at x=3568548, y=4206192
Sensor at x=2712265, y=2155055: closest beacon is at x=2700909, y=2519581
`)
