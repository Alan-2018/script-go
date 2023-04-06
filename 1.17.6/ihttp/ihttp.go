package ihttp

import (
	"io/ioutil"
	"log"
	"net/http"
	urlib "net/url"
	"strings"

	"github.com/flower/script-go/iutils"
)

type (
	jmap   = map[string]interface{}
	jarray = []interface{}
)

var lambda = func(inp interface{}) {
	log.Printf("%T, %v, %v\n", inp, inp, nil == inp)
}

/*
	http status code
	HTTP状态码 可以用于 函数返回值 或者 错误码 等等
*/
func TestIHttpStatusCode() {
	lambda(http.StatusOK)
}

/*
* ref
	* ! https://blog.csdn.net/qq_29176323/article/details/109745009
	* ! https://blog.csdn.net/qq_29176323/category_10058293.html

	* ! https://www.cnblogs.com/Goden/p/4658287.html
	* ! https://www.cnblogs.com/Goden/p/4639672.html
*/

func TestIHttpFuncs() {
	var (
		urlStr string = "https://www.baidu.com/s?wd=flower.cong@basebit.ai"

		err error
	)

	/*
		转义
	*/

	urlStrEscape := urlib.QueryEscape(urlStr)
	urlStrUnescape, _ := urlib.QueryUnescape(urlStrEscape)

	iutils.Log(
		urlStrEscape,
		urlStrUnescape,
	)

	/*
		HTTP request
			path & query (GET & ?)
			header
			body
	*/

	/*
		GET
	*/

	url, _ := urlib.Parse(urlStr)

	params := make(urlib.Values)
	params.Set("wd", "word")
	params.Set("ie", "utf-8")
	url.RawQuery = params.Encode()

	resp, err := http.Get(url.String())

	iutils.Log(
		resp,
		err,
	)

	req, _ := http.NewRequest(http.MethodGet, url.String(), nil)

	req.Header.Add("Content-type", "application/json;charset=utf-8")
	req.Header.Add("header", "?")

	req.AddCookie(
		&http.Cookie{
			Name:  "key",
			Value: "value",
		},
	)

	resp, err = (&http.Client{}).Do(req)

	iutils.Log(
		resp,
		err,
	)

	/*
		body -> io.Reader -> Read(p []byte) (n int, err error)

		jmap -> []byte -> *bytes.Buffer -> Read(p []byte) (n int, err error)
		urlib.Values{}.Encode() -> "bar=baz&foo=quux" -> ? how to body -> ? application/x-www-form-urlencoded
	*/

	urlStr = "http://jsonplaceholder.typicode.com/posts"
	body := urlib.Values{}
	iutils.Log(body)

	body.Set("userId", "")
	body.Set("title", "test")
	body.Set("body", "test")
	iutils.Log(body)

	req, _ = http.NewRequest(http.MethodPost, urlStr, strings.NewReader(body.Encode()))
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	body2, err := ioutil.ReadAll(req.Body)
	iutils.Log(string(body2))

	resp, err = (&http.Client{}).Do(req)

	iutils.Log(
		resp,
		err,
	)

}
