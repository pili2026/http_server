package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const userKey = "session_id"

func SetSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(userKey))
	return sessions.Sessions("mySession", store)
}

func AuthSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionId := session.Get(userKey)

		if sessionId == nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "This page need to login",
			})
			return
		}
		ctx.Next()
	}
}

func SaveSession(ctx *gin.Context, userId int) {
	session := sessions.Default(ctx)
	session.Set(userKey, userId)
	session.Save()
}

func ClearSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
}

func GetSession(ctx *gin.Context) int {
	session := sessions.Default(ctx)
	sessionId := session.Get(userKey)

	if sessionId == nil {
		return 0
	}
	return sessionId.(int)
}

func CheckSession(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	sessionId := session.Get(userKey)
	return sessionId != nil
}
