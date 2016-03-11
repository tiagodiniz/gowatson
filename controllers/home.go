package controllers

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

type HomeController struct  {
    MainController
}

func (this *HomeController) Home(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

    // recupera o user na session
   // session := sessions.GetSession(request)

    // renderiza nosso html
    this.Get().Render.HTML(response, http.StatusOK, "home", nil)

}