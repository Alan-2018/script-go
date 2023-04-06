package iutils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type (
	jmap   = map[string]interface{}
	jarray = []interface{}
)

func TestIUtils() {

}

func Log(inp ...interface{}) {
	// https://blog.csdn.net/wudebao5220150/article/details/128095696
	// https://mp.weixin.qq.com/s/MuZYTSC6Q4YSYgmyAojzIw

	var oup []interface{}

	for idx, _ := range inp {
		// t, v := reflect.TypeOf(nil), reflect.ValueOf(nil)
		t, v := reflect.TypeOf(inp[idx]), reflect.ValueOf(inp[idx])

		if nil == t {
			oup = append(oup, nil)
		} else if t.Kind() == reflect.Ptr && v.IsNil() {
			oup = append(oup, nil)
		} else if t.Kind() == reflect.Ptr {
			oup = append(oup, reflect.Indirect(v))
		} else {
			oup = append(oup, v)
		}
	}

	// log.Println(oup...)

	for idx, i := range oup {
		fmt.Println(idx, i)
	}

}

func TestIUtilsLog() {
	var (
		str    string  = "hello world"
		ptStr  *string = &str
		ptStr2 *string
	)

	Log(
		str,
		ptStr,
		ptStr2,
		nil,
	)
}

func Stringify(inp interface{}) (oup string) {

	switch inp.(type) {
	case nil:
		oup = ""

	case string:
		oup = inp.(string)

	case bool:
		oup = strconv.FormatBool(inp.(bool))

	case int:
		oup = strconv.Itoa(inp.(int))
	case uint:
		oup = strconv.Itoa(int(inp.(uint)))
	case int8:
		oup = strconv.Itoa(int(inp.(int8)))
	case uint8:
		oup = strconv.Itoa(int(inp.(uint8)))
	case int16:
		oup = strconv.Itoa(int(inp.(int16)))
	case uint16:
		oup = strconv.Itoa(int(inp.(uint16)))
	case int32:
		oup = strconv.Itoa(int(inp.(int32)))
	case uint32:
		oup = strconv.Itoa(int(inp.(uint32)))
	case int64:
		oup = strconv.FormatInt(inp.(int64), 10)
	case uint64:
		oup = strconv.FormatUint(inp.(uint64), 10)

	case float32:
		oup = strconv.FormatFloat(inp.(float64), 'f', -1, 32)
	case float64:
		oup = strconv.FormatFloat(inp.(float64), 'f', -1, 64)

	case []byte:
		oup = string(inp.([]byte))

	default:
		b, _ := json.Marshal(inp)
		oup = string(b)
	}

	return oup

}
