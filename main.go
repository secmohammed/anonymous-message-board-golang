package main

import (
    "github.com/secmohammed/anonymous-message-board-golang/config"
    "github.com/secmohammed/anonymous-message-board-golang/routes"
)

func main() {
    c := config.NewConfig()
    r := routes.NewRouter(c)
    r.Serve()
}
