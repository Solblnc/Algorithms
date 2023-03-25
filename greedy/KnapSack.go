package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 3
	capacity := float64(50)

	goods := [][]float64{{60, 20}, {100, 50}, {120, 30}}
	knapSack(n, capacity, goods)

}

//type Item struct {
//	value  float64
//	volume float64
//}

func knapSack(n int, capacity float64, goods [][]float64) (int, error) {

	sortGoods(goods)

	var actualCap float64
	var actualValue float64
	var actualGood int

	for actualCap <= capacity && actualGood <= n-1 {
		if actualCap+goods[actualGood][1] <= capacity {
			actualValue = actualValue + goods[actualGood][0]
			actualCap = actualCap + goods[actualGood][1]
		} else {
			//actualValue = ((capacity - actualCap) / goods[actualGood].volume) * goods[actualGood].value
			actualCap = capacity
		}

		actualGood++
	}

	answer, _ := fmt.Printf("%.3f", actualValue)

	return answer, nil

}

func sortGoods(goods [][]float64) {
	sort.Slice(goods, func(i, j int) bool {
		return goods[i][0]/goods[i][1] > goods[j][0]/goods[j][1]
	})
}
