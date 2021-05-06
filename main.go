package main

import (
    "github.com/secmohammed/anonymous-message-board-golang/config"
    "github.com/secmohammed/anonymous-message-board-golang/controllers"
    "github.com/secmohammed/anonymous-message-board-golang/database"
    "github.com/secmohammed/anonymous-message-board-golang/routes"
    "github.com/secmohammed/anonymous-message-board-golang/services"
)

func main() {
    c := config.NewConfig()
    r := routes.NewRouter(c)
    conn := database.NewDatabaseConnection(c)
    ts := services.NewThreadService(conn)
    tc := controllers.NewThreadController(ts)
    r.RegisterThreadRoutes(tc)
    r.Serve()
}
