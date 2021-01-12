/*
 * @Author: your name
 * @Date: 2021-01-12 17:09:42
 * @LastEditTime: 2021-01-12 18:02:25
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /fake-gin/router.go
 */

package fg

import "fmt"

type Router struct {
	routers map[string]HandleFunc
}

func NewRouter() *Router {
	return &Router{routers: make(map[string]HandleFunc)}
}

func getKey(m string, p string) string {
	return m + "-" + p
}

func (r *Router) addRoute(method string, path string, handleFunc HandleFunc) {
	key := getKey(method, path)
	r.routers[key] = handleFunc
}

func (r *Router) handle(c *Context) {
	key := getKey(c.Method, c.Path)
	handleFunc, ok := r.routers[key]
	if !ok {
		fmt.Fprintf(c.ResWriter, "Status 404 %s\n", c.Req.URL)
		return
	}
	handleFunc(c)
}
