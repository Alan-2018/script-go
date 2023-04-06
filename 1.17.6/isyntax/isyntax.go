package isyntax

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/flower/script-go/iutils"
	"github.com/goinggo/mapstructure"
)

type (
	jmap   = map[string]interface{}
	jarray = []interface{}

	ptjmap = *map[string]interface{}
)

var (
	// 换行符
	nlStr = "\n"

	// https://www.cnblogs.com/liang1101/p/6741262.html
	t *testing.T = new(testing.T)
)

func init() {

}

/*
换行
	ln (line)
	nl (new line)

打印换行符 & 代码内换行

代码内换行
	"("
	","
	"+"  字符串加法
	"``" raw string
*/
func TestISyntaxFuncsPrints() {

	fmt.Print(nlStr)
	fmt.Println("")

	// []interface{} & ...interface{}
	fmt.Println(
		"x",
		"y",
		"z",
	)

	fmt.Println([]interface{}{"x", "y", "z"}...)

	log.Println(
		"x",
		"y",
		"z",
	)

	log.Printf("x: %v\n", nil)

	var (
		s string
		b []byte
	)

	s = `
	{}
	`
	fmt.Println(s)

	b = []byte(`
	{}
	`)
	fmt.Println(string(b))

	// ? how to try-catch
	log.Fatal()

	panic("panic")

	// glog

}

/*
条件语句
	conditional statements
*/
func TestISyntaxConditionalStatements() {

	var lambda func(inp interface{})

	// 匿名函数
	// type-switch 不支持 fallthrough
	lambda = func(inp interface{}) {
		switch i := inp.(type) {
		default:
			log.Printf("default type %T, value %v", i, i)
		case int:
			log.Printf("type %T, value %v", inp, inp)
		case int8, int16, int32, int64:
			log.Printf("type %T, value %v", inp, inp)
		case uint8, uint16, uint32, uint64:
			log.Printf("type %T, value %v", inp, inp)
		case float64:
			log.Printf("type %T, value %v", inp, inp)
		case float32:
			log.Printf("type %T, value %v", inp, inp)
		case bool:
			log.Printf("type %T, value %v", inp, inp)
		case string:
			log.Printf("type %T, value %v", inp, inp)
		case []byte:
			log.Printf("type %T, value %v", inp, inp)
		case interface{}:
			log.Printf("interface{} type %T, value %v", inp, inp)
		case nil:
			log.Printf("type %T, value %v", inp, inp)
		}
	}

	lambda(0)
	lambda(0.)
	lambda(rune(0))
	lambda(byte(0))
	lambda([]byte(nlStr))

	// 因为 interface{}；所以不进 default 分支
	lambda(uint(0))

	// type <nil>, value <nil>
	// 注释 nil 分支；则进 default 分支
	// ? nil 可以为 interface{} 值；但不是 interface{} 类型
	// type <nil> -> untyped nil
	lambda(nil)
	// lambda(&nil)

	// type <nil>, value <nil>
	// 空指针
	// 野指针
	// 未初始化
	var i interface{}
	lambda(i)
	lambda(&i)

	// interface{} type *string, value <nil>
	var ptStr *string
	lambda(ptStr)

	lambda = func(inp interface{}) {
		switch {
		case true:
			inp = true
			fallthrough
		case false:
			inp = false
			fallthrough
		default:
			inp = nil
		}

		log.Printf("interface{} type %T, value %v", inp, inp)
	}

	lambda(true)
	lambda(nlStr)

}

/*
循环语句
	loop statements
*/
func TestISyntaxLoopStatements() {

}

/*
interface{}
	空接口

	类型
	表示 任意类型
	? 但是 不包括 type <nil>
	但是 值 可以等于 value <nil>

	// jbody := make(jmap)
	// var inp interface{} = jbody
	// var oup interface{}

断言
	仅限于 interface{}

	i := interface{}.(type)
	i, err := interface{}.(int)
	i, err := interface{}.(string)

	? 形参 inp interface{} 不等于 nil
*/

/*
类型 & 零值（默认值）

nil
	不是类型，而是 引用类型 零值 或者 空值
	nil may be the value of pointer type AND reference type

	也是类型
*/
func TestISyntaxTypesAndDefaults() {

	var lambda func(inp interface{})

	lambda = func(inp interface{}) {
		log.Printf("%T, %v, %v\n", inp, inp, nil == inp)

		if i, ok := inp.([]byte); ok {
			log.Printf("%T, %v, %v\n", i, i, nil == i)
		}

		if i, ok := inp.(*string); ok {
			log.Printf("%T, %v, %v\n", i, i, nil == i)
		}
		// ... type-switch
	}

	/*
		types
	*/

	// char & string
	lambda('c') // int32, 99
	lambda("c") // string, c

	// string
	var str string
	str = "黑化肥发灰"
	str += "\n"
	str += "灰化肥发黑"
	str += "\n"
	str += "黑化肥发灰会挥发"
	str += "\n"
	str += "灰化肥挥发会发黑"
	fmt.Println(str)

	str = "黑化肥发灰" +
		"\n" +
		"灰化肥发黑" +
		"\n" +
		"黑化肥发灰会挥发" +
		"\n" +
		"灰化肥挥发会发黑"
	fmt.Println(str)

	// ?
	str = `
	黑化肥发灰
	灰化肥发黑
	黑化肥发灰会挥发
	灰化肥挥发会发黑
	`
	fmt.Println(str)

	// string array
	// invalid use of '...'
	// var strArr [...]string

	var strArr [3]string
	fmt.Println(strArr)

	strArr2 := [...]string{} // [0]string
	fmt.Println(strArr2)

	strArr3 := []string{"X", "Y", "Z"}
	fmt.Println(strArr3)

	// ? "\x001" NOT WORK
	fmt.Println(strings.Join(strArr3, "\x001"))

	fmt.Println(strArr3[:])
	fmt.Println(strArr3[len(strArr3)-1])
	fmt.Println(append(strArr3, ""))

	// int & int...

	var n int = 1
	lambda(1 / 3)
	lambda(1 / 3.)
	lambda(float64(1) / 3)
	lambda((float64(1) / 3) * 100)
	// cannot convert float64(1) / 2 (constant 0.5 of type float64) to int64
	// lambda((int64)(float64(1) / 2))

	// overflow & underflow
	// ? int 兼容 int32 & int64
	// ? int 值域 等于 int64
	log.Println("test overflow 2^64")
	lambda(math.Pow(2, 64))
	lambda(int64(math.Pow(2, 64)))
	lambda(int(math.Pow(2, 64)))

	log.Println("test overflow 2^65")
	lambda(math.Pow(2, 65))
	lambda(int64(math.Pow(2, 65)))
	lambda(int(math.Pow(2, 65)))

	// 科学计数法

	// 引用类型

	// 指针
	// type pointer

	// use of untyped nil in assignment
	// pt := nil
	//
	var pt *int
	pt, _ = nil, ""
	lambda(pt)

	n = math.MaxInt64
	pt = &n
	lambda(pt)
	lambda(&pt)
	lambda(*pt)

	// 空指针异常
	// panic: runtime error: invalid memory address or nil pointer dereference
	// https://blog.csdn.net/weixin_39616603/article/details/100574671
	// var pt2 *int
	// *pt2 = 1
	//
	var pt2 *int = new(int)
	lambda(pt2)
	lambda(&pt2)
	lambda(*pt2)
	*pt2 = 1
	lambda(pt2)
	lambda(&pt2)
	lambda(*pt2)

	arr := []interface{}{"X", "Y", "Z"}
	arr2 := make([]string, len(arr))
	for idx, _ := range arr {
		arr2[idx] = arr[idx].(string)
	}
	lambda(arr)
	lambda(arr2)

	var j jmap = make(jmap)
	lambda(j["default"])
	// nil (untyped nil value) is not a type
	// lambda(j["default"].(nil))
	// panic: interface conversion: interface {} is nil, not ...
	// lambda(j["default"].(interface{}))

	j["default"] = false
	lambda(j["default"])
	delete(j, "default")
	lambda(j["default"])

	// variable assignment maybe subject to point instead of interface
	var pt3 ptjmap = &j
	lambda((*pt3)["default"])

	/*
		defaults
	*/

	// 基本类型 & 引用类型
	// 基本类型 零值 为 0 （不同格式，内容为0）
	// 引用类型 零值 为 nil

	var bl bool
	lambda(bl)

	var i int
	lambda(i)

	var s string
	lambda(s)

	var bs []byte
	lambda(bs)
	log.Println(nil == bs)

	var ptStr *string
	lambda(ptStr)
	log.Println(nil == ptStr)

	var i2 interface{}
	lambda(i2)
	log.Println(nil == i2)

	lambda(nil)

	var err error
	lambda(err)

	j = make(jmap)
	lambda(j)
	log.Println(nil == j)

	var j2 jmap
	lambda(j2)
	log.Println(nil == j2)

	var buf bytes.Buffer
	lambda(buf)
	// cannot convert nil (untyped nil value) to struct{}
	// log.Println(nil == bf)

	var ptBuf *bytes.Buffer
	lambda(ptBuf)
	log.Println(nil == ptBuf)

	var ptBuf2 *bytes.Buffer = bytes.NewBuffer(nil)
	lambda(ptBuf2)
	log.Println(nil == ptBuf2)

}

func TestISyntaxStringsFuncs() {

	var s string = "\n---- body ----\n"

	fmt.Println(strings.Replace(s, "\n", "", -1))
	fmt.Println()

	fmt.Println(strings.TrimLeft(s, "-\n"))
	fmt.Println()

	fmt.Println(strings.TrimRight(s, "-\n"))
	fmt.Println()

	fmt.Println(strings.Trim(s, "-\n"))
	fmt.Println()

}

func TestISyntaxChars() {

	// 不可见字符集
	// invisible charset
	// https://blog.csdn.net/whynottrythis/article/details/104643145

	// "\r\n\t " & ...

}

func TestISyntaxIOFuncs() (err error) {

	var lambda func(dst io.Writer, src io.Reader) (err error) = func(dst io.Writer, src io.Reader) (err error) {
		_, err = io.Copy(dst, src)
		if nil != err {
			log.Println(err)
		}

		return
	}

	var buf *bytes.Buffer
	lambda(os.Stdout, buf)

	buf = new(bytes.Buffer)
	lambda(os.Stdout, buf)

	buf = bytes.NewBufferString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lambda(os.Stdout, buf)

	return
}

func TestISyntaxTimesFuncs() {

	// generic paradigm
	//    any &    interface{}
	// ...any & ...interface{}
	/*
		var lambda func(a ...interface{}) = func(a ...interface{}) {
			for idx, i := range a {
				fmt.Println(idx, i)
			}

			fmt.Println(a...)
		}
	*/

	iutils.Log(
		time.Local,
		// *time.Local,
		time.Now(),
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
		time.Now().Nanosecond(),
		time.Now().Weekday(),
		time.Date(
			time.Now().Year(),
			time.Now().Month(),
			time.Now().Day(),
			time.Now().Hour(),
			time.Now().Minute(),
			time.Now().Second(),
			time.Now().Nanosecond(),
			time.Local,
		),
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02"),
		time.Now().Format("2006-01-02 15:04:05 +0800 CST"),
		time.Now().Unix(),
		time.Now().UnixNano(),
		time.Now().UnixNano()/1e6,
	)

	fmt.Println()

	/*
		duration
	*/

	m, _ := time.ParseDuration("-1m")
	fmt.Println(time.Now().Add(m))

	h, _ := time.ParseDuration("1h")
	fmt.Println(time.Now().Add(h))

	d := time.Now().Add(h).Sub(time.Now())
	fmt.Println(
		d.String(),
		d.Hours(),
		d.Minutes(),
		d.Seconds(),
	)

	fmt.Println()

	// ? func math.Pow(x float64, y float64) float64
	time.Sleep(time.Second * time.Duration(math.Pow(1, 0)))

	/*
		默认使用 操作系统 时区

		layout
			"2006-01-02 15:04:05"
			? T Z
			"2006-01-02T15:04:05Z"

			https://blog.csdn.net/qq_26140191/article/details/98037067
			https://www.uoften.com/article/182775.html

		时间点 & 时间段
			Time & Duration
			M(MVC)

		时区
			Location
			时区转换 -> C(MVC)

		格式化
			Format & layout
			V(MVC)
	*/
	var (
		t       time.Time
		layout  string = "2006-01-02 15:04:05"
		layout2 string = "2006-01-02 15:04:05 Mon"
		layout3 string = "2006-01-02T15:04:05Z"
	)

	/*
		time.Time 结构体 主要是 标准 unix timestamp 为 标准
		string -> parse -> time.Time
		time.Time -> format -> string
	*/

	// 2-valued time.Parse("2006-01-02", "2022-05-01") (value of type (time.Time, error)) where single value is expected
	/*
		fmt.Println(
			time.Now().Format("2006-01-02"),
			time.Parse("2006-01-02", "2022-05-01"),
		)
	*/
	fmt.Println(time.Parse("2006-01-02", "2022-05-01"))
	fmt.Println(time.Parse("2006-01-02 15:04:05", "2022-05-01 15:04:05"))
	fmt.Println(time.Parse("2006-01-02T15:04:05Z", "2022-05-01T15:04:05Z"))

	fmt.Println(time.Parse("2006-01-02", "2022-05-01 15:04:05"))
	fmt.Println(time.Parse("2006-01-02 15:04:05 +0000 UTC", "2022-05-01 15:04:05 +0800 CST"))
	fmt.Println(time.Parse(time.RFC3339, "2022-05-01 15:04:05"))

	fmt.Println()

	/*
		"Asia/Chongqing" & "Asia/Shanghai"
	*/

	loc, _ := time.LoadLocation("Asia/Chongqing")

	var lambda func(t time.Time) = func(t time.Time) {
		iutils.Log(
			t,
			t.In(loc),
			t.Unix(),
			t.Local().Unix(),
			t.UTC().Unix(),
			t.Format(layout),
			t.Local().Format(layout),
			t.UTC().Format(layout),
			t.Format(layout2),
			t.Local().Format(layout2),
			t.UTC().Format(layout2),
			t.Format(layout3),
			t.Local().Format(layout3),
			t.UTC().Format(layout3),
		)
	}

	t, _ = time.Parse(layout, "2023-03-28 00:00:00")
	lambda(t)

	fmt.Println()

	t, _ = time.Parse(time.RFC3339, "2023-03-28T00:00:00Z")
	lambda(t)

	fmt.Println()

	t, _ = time.ParseInLocation(layout, "2023-03-28 00:00:00", loc)
	lambda(t)

	fmt.Println()

	/*
		loc
	*/

	// default to unix timestamp
	loc, _ = time.LoadLocation("")
	fmt.Println(time.Now().In(loc).Format(layout))

	// unix timestamp
	loc, _ = time.LoadLocation("UTC")
	fmt.Println(time.Now().In(loc).Format(layout2))

	loc, _ = time.LoadLocation("Local")
	fmt.Println(time.Now().In(loc).Format(layout2))

	fmt.Println()

	/*
		unix timestamp <-> time.Time
	*/

	fmt.Println(
		time.Unix(
			time.Now().UTC().Unix(),
			int64(time.Now().UTC().Nanosecond()),
		),
	)

}

func TestISyntaxErrors() {

	var err error

	statusCodeStr := fmt.Sprintf("%d", 502)

	err = fmt.Errorf("statusCode: %v; msg: %v", statusCodeStr, "UNKNOWNS")
	log.Println(err)
	log.Println(err.Error())

	// ? %w -> object
	// ? warp
	err = fmt.Errorf("error: %w", err)
	log.Println(err)

}

func TestISyntaxJsons() {
	var (
		j  = make(jmap)
		j2 = make(jmap)
	)

	// json str <-> map
	j["str"] = "s"
	log.Printf("%+v", j)

	b, _ := json.Marshal(j)
	log.Printf("%+v", string(b))

	err := json.Unmarshal(b, &j2)
	if err != nil {
		log.Printf("%w", err)
	}

	log.Printf("%+v", j2)

}

func TestISyntaxStructs() {

	// json: omitempty -> point cause of nil

	// import field: 导出项
	type Structs struct {
		Str string `db:"str" json:"str,omitempty"`
		Bl  bool   `db:"bl" json:"bl,omitempty"`
	}

	var structs Structs

	structs = Structs{
		Str: "s",
	}

	log.Printf("%+v", structs)

	// struct <-> json str
	var s string = `
	{
		"str": "s2",
		"bl": true
	}
	`

	json.Unmarshal([]byte(s), &structs)
	// t.Log(structs)

	b, _ := json.Marshal(&structs)
	log.Printf("%+v", string(b))

	// struct <-> map
	var j jmap = make(jmap)
	j["str"] = "jmap"

	if err := mapstructure.Decode(j, &structs); err != nil {
		log.Println(err)
	}

	// 结构体 属性 大写
	// panic: reflect.Value.Interface: cannot return value obtained from unexported field or method

	var Struct2Map func(obj interface{}) (j jmap) = func(obj interface{}) (j jmap) {
		t, v := reflect.TypeOf(obj), reflect.ValueOf(obj)

		j = make(jmap)
		for i := 0; i < t.NumField(); i++ {
			j[t.Field(i).Name] = v.Field(i).Interface()
			//
			j[v.Type().Field(i).Tag.Get("json")] = v.Field(i).Interface()
		}

		return j
	}

	log.Printf("%+v", Struct2Map(structs))

	// pointer

	var ptStructs *Structs = &structs
	log.Printf("%+v", ptStructs)
	log.Printf("%+v", *ptStructs)

}

// https://zhuanlan.zhihu.com/p/137060307
func TestISyntaxUnsafeFuncs() {

}
