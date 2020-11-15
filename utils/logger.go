package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Logger struct{}

func (l Logger) LogRequest(req *http.Request) {
	s, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(s))
}

func (l Logger) LogResponse(res *http.Response) {
	s, _ := httputil.DumpResponse(res, true)
	fmt.Println(string(s))
}
