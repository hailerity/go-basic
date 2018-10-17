package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	apic := make([][]uint8, dy)
	for i, _ := range apic {
		apic[i] = make([]uint8, dx)
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			if x == y {
				apic[x][y] = 0
			} else {
				apic[x][y] = 254
			}
		}
	}

	return epic
}

func main() {
	pic.Show(Pic)
}
