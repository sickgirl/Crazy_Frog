package main

import "fmt"

//1603. 设计停车系统
//https://leetcode.cn/problems/design-parking-system/
func main() {
	p := Constructor(1, 1, 0)
	fmt.Printf("%v", p)
	x := &p
	r := []bool{}

	ex := []int{1, 2, 3, 1}
	for _, v := range ex {
		res := x.AddCar(v)
		r = append(r, res)
	}
	fmt.Printf("%v", r)

}

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big > 0 {
			this.big--
			return true
		} else {
			return false
		}
	case 2:
		if this.medium > 0 {
			this.medium--
			return true
		} else {
			return false
		}
	case 3:
		if this.small > 0 {
			this.small--
			return true
		} else {
			return false
		}
	}
	return false
}
