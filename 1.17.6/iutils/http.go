package iutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"time"
)

/*
	HTTP utils
*/

// retry request to be more robust just for exception & http status code unauth & ...
func HttpRequestLoops(client *http.Client, req *http.Request) (*http.Response, error) {
	var (
		body []byte = nil
		err  error  = nil
	)

	if req.Method == "POST" || req.Method == "PUT" {
		body, err = ioutil.ReadAll(req.Body)
		if err != nil {
			err = fmt.Errorf("E HttpRequestLoops err: %w", err)
			log.Println(err)
			return nil, err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}

	i, cnt := 0, 3
	for {
		resp, err := client.Do(req)
		if err == nil && resp.StatusCode != http.StatusUnauthorized &&
			resp.StatusCode != http.StatusInternalServerError {
			return resp, nil
		}

		if resp != nil {
			err = fmt.Errorf("E HttpRequestLoops err: %w && req: %+v && resp body: %s", err, req, resp.Body)
		} else {
			err = fmt.Errorf("E HttpRequestLoops err: %w && req: %+v && resp body: %s", err, req, nil)
		}
		log.Println(err)

		if i > cnt {
			return resp, err
		}

		time.Sleep(time.Second * time.Duration(math.Pow(2, float64(i))))

		i++

		if req.Method == "POST" || req.Method == "PUT" {
			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			// fmt.Printf(string(body))
		}
	}
}

type HTTPRequest struct {
	URL    string
	Token  string
	Method string
	Body   []byte
	JBody  jmap
}

type HTTPResponse struct {
	Error      error
	StatusCode int
	Body       []byte
	JBody      jmap
}

// HTTP method default `POST`
// build `GET` request
func HttpRequest(params *HTTPRequest) *HTTPResponse {
	var (
		body []byte
		buf  *bytes.Buffer = new(bytes.Buffer)
	)

	if "" == params.Method {
		params.Method = http.MethodPost
	}

	if http.MethodPost == params.Method || http.MethodPut == params.Method {
		if nil == params.Body {
			body, _ = json.Marshal(params.JBody)
		} else {
			body = params.Body
		}

		buf = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(params.Method, params.URL, buf)
	if err != nil {
		err = fmt.Errorf("E HttpRequest http.NewRequest err: %w", err)
		log.Println(err)
		return &HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	req.Header.Set("Content-Type", "application/json")
	if params.Token != "" {
		// req.Header.Set("Authorization", "Bearer "+params.Token)
		req.Header.Set("Authorization", params.Token)
	}

	resp, err := HttpRequestLoops(&http.Client{}, req)
	if err != nil && resp == nil {
		return &HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	statusCode := resp.StatusCode

	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("E HttpRequest ioutil.ReadAll err: %w", err)
		log.Println(err)
		return &HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      err,
		}
	}

	jbody := make(jmap)
	err = json.Unmarshal(rbody, &jbody)
	if err != nil {
		err = fmt.Errorf("E HttpRequest json.Unmarshal err: %w && resp body: %s", err, rbody)
		log.Println(err)
		return &HTTPResponse{
			StatusCode: statusCode,
			Error:      err,
			Body:       rbody,
		}
	}

	if statusCode/100 != 2 {
		err = fmt.Errorf("E HttpRequest resp.StatusCode err: %v && resp body: %s", resp.StatusCode, rbody)
		log.Println(err)
		return &HTTPResponse{
			StatusCode: statusCode,
			Error:      err,
			Body:       rbody,
			JBody:      jbody,
		}
	}

	return &HTTPResponse{
		StatusCode: statusCode,
		Error:      nil,
		Body:       rbody,
		JBody:      jbody,
	}
}
