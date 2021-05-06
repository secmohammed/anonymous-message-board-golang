package routes

import (
    "errors"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/secmohammed/anonymous-message-board-golang/utils"
)

// x-api-key: key-secret
func ApiKey(param, key string) gin.HandlerFunc {
    return func(c *gin.Context) {
        attempt := c.Request.Header.Get(param)
        if attempt != key {
            c.AbortWithStatusJSON(utils.CreateApiError(http.StatusUnauthorized, errors.New("Invalid api key")))
            return
        }
        c.Next()
    }
}
