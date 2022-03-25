package main

import (
    "github.com/ant0ine/go-json-rest/rest"
)

type WebApp struct {
}

type JsonApp interface {
    RestApp() (rest.App, error)
}

type APIGet struct {
    Test          string          `json:"test"`
}

func (self *WebApp) Get(w rest.ResponseWriter, req *rest.Request) {
    out := APIGet{}
    out.Test = "test"

    w.WriteJson(out)
}

func RestApp() (rest.App, error) {
    app := &WebApp {
    }
    return rest.MakeRouter(
        rest.Get("/",           app.Get),
    )
}
