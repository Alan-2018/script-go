package isyntax

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

/*
	伪随机数 pseudo-random
*/
func TestISyntaxMathRandFuncs() {
	/*
		闰年 leap year
			被400整除 或者 被4整除 但不被100整除

			地球公转 与 公历 误差
				地球公转 略小于 365.25
				所以 4年一闰 闰一天 366
				但是 因为略小于 365.25
				所以 400年 要少闰3天
				所以 规定 100年不闰 400年闰一天
	*/
	var flag, year = false, 2000
	for !flag {
		year = rand.Intn(400) + 2001
		flag = year%400 == 0 || (year%4 == 0 && year%100 != 0)

		if flag {
			log.Printf("%v is leap year\n", year)
		} else {
			log.Printf("%v is not leap year\n", year)
		}

		time.Sleep(1e9)
	}

	/*
		随机生成日期 某年某月某日
			闰年 闰月
	*/

	/*
		切片
	*/
	arr := [...]string{"apple", "banana"}
	slc := arr[:]
	slc2 := append(slc, "tomato")
	slc3 := append(slc2, "potato")

	slc[1] = "stawberry"
	arr[1] = "litchi"

	slc3[1] = "blueberry"

	fmt.Println(arr)
	fmt.Println(slc)
	fmt.Println(slc2)
	fmt.Println(slc3)

	/*
		map 计数器 分组 集合
	*/

	temperatures := [...]float64{
		-28.0, 32.0, -22.0, 38.0, -27.0, -25.0, -24.0, -23.0, -22.0, -21.0, -20.0,
	}

	counter := make(map[float64]int, 10)
	groups := make(map[float64][]float64, 10)
	set := make(map[float64]bool, 10)

	for _, i := range temperatures {
		//
		counter[i]++

		// 若，不为2位数呢
		// 取余
		g := math.Trunc(i/10) * 10
		groups[g] = append(groups[g], i)

		//
		set[i] = true
	}

	fmt.Println(counter)
	fmt.Println(groups)
	fmt.Println(set)

	// 排序
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}

	sort.Float64s(unique)

	fmt.Println(unique)
}

func TestISyntaxMathFuncs() {
	var (
		// min float64 = math.SmallestNonzeroFloat64
		// max float64 = math.MaxFloat64
		min float64 = -1e10
		max float64 = 1e10
	)

	minStr := strconv.FormatFloat(min, 'f', 2, 64)
	maxStr := strconv.FormatFloat(max, 'f', 2, 64)

	log.Println(
		minStr,
		maxStr,
	)

}
