package main

import (
    "fmt"

    "github.com/secmohammed/anonymous-message-board-golang/config"
    "github.com/secmohammed/anonymous-message-board-golang/database"
    "github.com/secmohammed/anonymous-message-board-golang/routes"
)

func main() {
    c := config.NewConfig()
    r := routes.NewRouter(c)
    conn := database.NewDatabaseConnection(c)
    fmt.Println(conn)
    r.Serve()
}
