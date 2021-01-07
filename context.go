/*
 * @Author: your name
 * @Date: 2021-01-07 15:48:28
 * @LastEditTime: 2021-01-07 19:02:52
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /fake-gin/context.go
 */

package fg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Context is a struct that
type Context struct {
	Req        *http.Request
	ResWriter  http.ResponseWriter
	Method     string
	Host       string
	Path       string
	StatusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Req:       r,
		ResWriter: w,
		Method:    r.Method,
		Path:      r.URL.Path,
		Host:      r.Host,
	}
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.ResWriter.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.ResWriter.Header().Set(key, value)

}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) resHandle(code int, err error) {
	c.Status(code)
	if err != nil {
		http.Error(c.ResWriter, err.Error(), 500)
		return
	}
}

func (c *Context) FormatText(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	text := fmt.Sprintf(format, values...)
	_, err := c.ResWriter.Write([]byte(text))
	c.resHandle(code, err)

}

func (c *Context) Data(code int, data []byte) {
	c.SetHeader("Content-Type", "text/plain")
	_, err := c.ResWriter.Write(data)
	c.resHandle(code, err)
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	encoder := json.NewEncoder(c.ResWriter)
	err := encoder.Encode(obj)
	c.resHandle(code, err)
}

func (c *Context) HTML(code int, html []byte) {
	c.SetHeader("Content-Type", "text/html")
	_, err := c.ResWriter.Write(html)
	c.resHandle(code, err)

}
