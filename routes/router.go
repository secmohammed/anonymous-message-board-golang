package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/secmohammed/anonymous-message-board-golang/config"
    "github.com/secmohammed/anonymous-message-board-golang/controllers"
)

type Router interface {
    gin.IRouter
    Serve() error
    RegisterThreadRoutes(c controllers.ThreadController)
    RegisterReplyRoutes(c controllers.ReplyController)
    RegisterAdminRoutes(c controllers.AdminController)
}
type router struct {
    *gin.Engine
    c *config.Config
}

func NewRouter(c *config.Config) Router {
    config := c.Get()
    r := gin.New()
    if config.GetString("ENVIRONMENT") == "production" {
        gin.SetMode(gin.ReleaseMode)
    }
    if config.GetBool("app.log") {
        r.Use(gin.Logger())
    }
    setupDefaults(r)
    return &router{Engine: r, c: c}
}
func (r *router) Serve() error {
    port := r.c.Get().GetString("app.port")
    return r.Run(":" + port)
}
