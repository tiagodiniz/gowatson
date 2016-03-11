package main

import (
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/satori/go.uuid"
	"net/http"
	"github.com/gowatson/routers"
)

func main() {
	// middleware stack
	middleware := negroni.Classic()
	middleware.Use(negroni.NewStatic(http.Dir("templates")))
	middleware.Use(sessions.Sessions("gowatson", cookiestore.New([]byte(uuid.NewV4().String()))))
	//middleware.UseHandler(routers.GetHandlers())
	middleware.UseHandler(routers.GetControllers())
	middleware.Run(":8080")

}
