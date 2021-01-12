/*
 * @Author: your name
 * @Date: 2021-01-06 18:30:38
 * @LastEditTime: 2021-01-12 18:04:36
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /fake-gin/fg.go
 */

package fg

import (
	"fmt"
	"net/http"
)

// HandleFunc is a ....
type HandleFunc func(c *Context)

// App is a ...
type App struct {
	router *Router
}

// New is a
func New() *App {
	return &App{router: NewRouter()}
}

func (app *App) Run(port string) (err error) {
	fmt.Println("Localhost run on 9090")
	return http.ListenAndServe(port, app)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	app.router.handle(c)
}

func (app *App) GET(path string, handleFunc HandleFunc) {
	app.router.addRoute("GET", path, handleFunc)
}

func (app *App) POST(path string, handleFunc HandleFunc) {
	app.router.addRoute("POST", path, handleFunc)
}
