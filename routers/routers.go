package routers

import (
    "github.com/julienschmidt/httprouter"
    "github.com/gowatson/controllers"
)

func GetControllers() (router *httprouter.Router) {
    router = httprouter.New()
    // controllers
    home := &controllers.HomeController{}
    router.GET("/", home.Home)
    router.GET("/home", home.Home)
    return
}