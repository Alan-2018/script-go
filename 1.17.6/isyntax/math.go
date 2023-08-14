package isyntax

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
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

type SimpleStruct struct {
	Name string `json:"name"`
}

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "not set"
	}

	return fmt.Sprintf("%d", n.value)
}

type SafeWriter struct {
	w   io.Writer
	err error
}

func (sw *SafeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}

	_, sw.err = fmt.Fprintln(sw.w, s)
}

const (
	rows    = 9
	columns = 9
)

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

type Grid [rows][columns]int8

type SudukuError []error

func (se SudukuError) Error() string {
	var elems []string
	for _, e := range se {
		elems = append(elems, e.Error())
	}

	return strings.Join(elems, ", ")

}

func (g Grid) Set(x, y int, digit int8) error {
	var err SudukuError
	if !g.inBounds(x, y) {
		// return ErrBounds
		err = append(err, ErrBounds)
	}

	if !g.validDigit(digit) {
		// return ErrDigit
		err = append(err, ErrDigit)
	}

	if len(err) > 0 {
		return err
	}

	g[x][y] = digit
	return nil
}

func (g Grid) inBounds(x, y int) bool {
	if x < 0 || x >= rows {
		return false
	} else if y < 0 || y >= columns {
		return false
	}

	return true
}

func (g Grid) validDigit(digit int8) bool {
	return false
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

	x := SimpleStruct{
		"x",
	}
	x2 := &SimpleStruct{
		"x",
	}

	fmt.Printf(
		"%T\t%+v\n",
		x,
		x,
	)
	fmt.Printf(
		"%T\t%+v\n",
		x2,
		x2,
	)

	/*
		map 按引用传递
		? slice chan
		? struct interface
	*/

	var lambda func(j jmap) jmap = func(j jmap) jmap {
		j["x"] = "x"
		return j
	}

	j := make(jmap)
	log.Printf("%+v\n", j)

	j2 := lambda(j)
	log.Printf("%+v\n", j2)
	log.Printf("%+v\n", j)

	/*
		数组 按引用传递
	*/

	var lambda2 func(arr [8][8]rune) = func(arr [8][8]rune) {
		arr[0][0] = 'r'
	}

	arr := [8][8]rune{}
	lambda2(arr)
	fmt.Println(arr)

	/*
		切片
	*/
	// ...

	var i interface{} = nil
	var i2 *int
	i = i2

	fmt.Printf("%#v\n", nil)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%#v\n", i2)

	var s fmt.Stringer
	fmt.Printf("%#v\n", s)

	n := newNumber(42)
	fmt.Println(n)

	e := number{}
	fmt.Println(e)

	// proverbs
	f, err := os.Create("proverbs.txt")
	if err != nil {
		fmt.Printf("%#v\n", err)
		// return err
		os.Exit(1)
	}

	defer f.Close()

	sw := SafeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don’t just check errors, handle them gracefully.")
	sw.writeln("Don’t panic.")

	// return sw.err

	// New error
	// SUDUKU
	var g Grid
	err = g.Set(10, 0, 5)
	if err != nil {
		switch err {
		case ErrBounds, ErrDigit:
			fmt.Printf("request params error.\n")
		default:
			fmt.Printf("%#v\n", err)
		}
	}

	// recover

	// defer func() {
	// 	if e := recover(); e != nil {
	// 		fmt.Println(e)
	// 	}
	// }()

	// panic("i forget my towel")

	// goroutine
	var sleepyGopher func(sec int64, id int, c chan int) = func(sec int64, id int, c chan int) {
		time.Sleep(time.Second * time.Duration(sec))
		fmt.Println("...snore...", id)
		c <- id
	}

	c := make(chan int)
	for i := 0; i < 3; i++ {
		go sleepyGopher(3, i, c)
	}

	// time.Sleep(time.Second * time.Duration(4))

	for i := 0; i < 3; i++ {
		id := <-c
		fmt.Printf("%v over\n", id)
	}

	// c2 := make(chan int)
	// for i := 0; i < 3; i++ {
	// 	go sleepyGopher(int64(rand.Intn(3)), i, c2)
	// }

	// timeout := time.After(3 * time.Second)

	// // i < 4
	// for i := 0; i < 4; i++ {
	// 	select {
	// 	case id := <-c2:
	// 		fmt.Printf("c2 %v over\n", id)
	// 	case <-timeout:
	// 		fmt.Printf("c2 timeout\n")
	// 		return
	// 	}
	// }

	// 死锁
	c3 := make(chan int)
	go func() { c3 <- 2 }()
	<-c3

	// stream
	// fmt.Println("stream")

	var sourceGopher func(downstream chan string) = func(downstream chan string) {
		ss := []string{"hello", "world", "bad gay", "bad ..."}
		for _, s := range ss {
			// fmt.Println("sourceGopher: ", s)
			downstream <- s
		}
		// downstream <- ""
		close(downstream)
	}

	var filterGopher func(upstream, downstream chan string) = func(upstream, downstream chan string) {
		for {
			v, ok := <-upstream
			// fmt.Println("filterGopher: ", v)
			// if v == "" {
			// 	downstream <- ""
			// 	return
			// }
			if !ok {
				close(downstream)
				break
			}

			if !strings.HasPrefix(v, "bad") {
				downstream <- v
			}
		}

		// for v := range upstream {
		// 	downstream <- v
		// }
		// close(downstream)
	}

	var printGopher func(upstream chan string) = func(upstream chan string) {
		for {
			v := <-upstream
			// fmt.Println("printGopher: ", v)
			if v == "" {
				return
			}

			fmt.Println(v)
		}
	}

	c4 := make(chan string)
	c5 := make(chan string)
	go sourceGopher(c4)
	go filterGopher(c4, c5)
	printGopher(c5)

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	// ...

}
