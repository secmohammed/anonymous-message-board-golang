package routes

import "github.com/secmohammed/anonymous-message-board-golang/controllers"

func (r *router) RegisterThreadRoutes(c controllers.ThreadController) {
    rg := r.Group("/api/threads")
    rg.GET("/", c.ListThreads)
    rg.POST("/", c.CreateThread)
    rg.PUT("/:id", c.Report)
    rg.DELETE("/:id", c.Delete)
    rg.GET("/:id", c.GetThread)
}
