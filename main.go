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
    rs := services.NewReplyService(conn)
    tc := controllers.NewThreadController(ts)
    rc := controllers.NewReplyController(rs)
    r.RegisterThreadRoutes(tc)
    r.RegisterReplyRoutes(rc)
    r.Serve()
}
