package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	dataCur := "银川 \n300 \n400 \n500 \n600 \n800 \n1000 \n1300 \n1500 \n1700 \n2000 \n2300 \n2500 \n"
	for _, data := range strings.Split(dataCur, " \n\n") {
		// 把数据转换为整数数组
		sales := []int{}
		for i, s := range strings.Split(data, " \n") {
			if s != "" {
				if i == 0 {
					print("CityVipMap[\"")
					print(s)
					print("\"] = \n")
					continue
				}

				n, err := strconv.Atoi(s)
				if err != nil {
					println("当前序号 是:", i)

					panic(err)
				}
				sales = append(sales, n)
			}
		}

		// 创建包含每月销售额的数组
		monthlySales := [12]int{}
		for i := 0; i < 12; i++ {
			monthlySales[i] = sales[i]
		}

		// 创建键为月份，值为销售额的映射
		salesMap := map[string]int{}
		for i, s := range monthlySales {
			monthStr := fmt.Sprintf("%02d", i+1)
			salesMap[monthStr] = s
		}

		// 打印salesMap的值
		fmt.Printf("%#v", salesMap)

		println(" ")
	}

}
