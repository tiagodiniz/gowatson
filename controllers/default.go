package controllers

import (
    "gopkg.in/unrolled/render.v1"
    "github.com/gowatson/tools"
)

type MainController struct {
    Render *render.Render
}

func (c *MainController) Get() *MainController {
    c.Render = tools.GetRender()
    return c
}
