package main

import (
    "github.com/jessevdk/go-flags"
    "net/http"
    "log"
    "os"
    "github.com/ant0ine/go-json-rest/rest"
)

type Options struct {
    HttpDevel       bool        `long:"http-devel" description:"Use development mode"`
    HttpListen      string      `long:"http-listen" value-name:"HOST:PORT" default:":8282"`
}

func main() {
    var options Options

    if _, err := flags.Parse(&options); err != nil {
        os.Exit(1)
    }

    // http API
    api := rest.NewApi()

    // if options.HttpDevel {
        api.Use(rest.DefaultDevStack...)
    // }

    if app, err := RestApp(); err != nil {
        log.Fatalf("RestApp: %v\n", err)
    } else {
        api.SetApp(app)
    }

    http.Handle("/", api.MakeHandler())

    if err := http.ListenAndServe(options.HttpListen, nil); err != nil {
        log.Fatalf("http.ListenAndServe %v: %v\n", options.HttpListen, err)
    }
}
