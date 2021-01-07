/*
 * @Author: your name
 * @Date: 2021-01-06 18:30:38
 * @LastEditTime: 2021-01-07 17:15:47
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
type HandleFunc func(http.ResponseWriter, *http.Request)

// App is a ...
type App struct {
	router map[string]HandleFunc
}

// New is a
func New() *App {
	return &App{
		router: make(map[string]HandleFunc),
	}
}

func getKey(m string, p string) string {
	return m + "-" + p
}

func (app *App) addRoute(method string, path string, handleFunc HandleFunc) {
	key := getKey(method, path)
	app.router[key] = handleFunc
}

func (app *App) GET(path string, handleFunc HandleFunc) {
	app.addRoute("GET", path, handleFunc)
}

func (app *App) POST(path string, handleFunc HandleFunc) {
	app.addRoute("POST", path, handleFunc)
}

func (app *App) Run(port string) (err error) {
	fmt.Println("Localhost run on 9090")
	return http.ListenAndServe(port, app)
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	handler, ok := app.router[key]
	if !ok {
		fmt.Fprintf(w, "Status 404 %s\n", r.URL)
		return
	}
	handler(w, r)
}
