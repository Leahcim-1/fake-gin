/*
 * @Author: your name
 * @Date: 2021-01-07 15:48:28
 * @LastEditTime: 2021-01-07 17:15:08
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /fake-gin/context.go
 */

package fg

import "net/http"

// Context is a struct that
type Context struct {
	Req        *http.Request
	ResWriter  http.ResponseWriter
	Method     string
	Host       string
	Path       string
	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Req:       r,
		ResWriter: w,
		Method:    r.Method,
		Path:      r.URL.Path,
		Host:      r.Host,
	}
}
