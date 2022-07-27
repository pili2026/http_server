package service

import (
	"booking_system/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/pquerna/ffjson/ffjson"
)

func CacheUserDecorator(h gin.HandlerFunc, id string, keyPattern string, args interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		keyId := ctx.Param(id)
		redisKey := fmt.Sprintf(keyPattern, keyId)
		conn := database.RedisDefaultPool.Get()
		defer conn.Close()

		data, err := redis.Bytes(conn.Do("GET", redisKey))

		if err != nil {
			h(ctx)
			dbResult, exists := ctx.Get("dbResult")

			if !exists {
				dbResult = args
			}

			redisData, _ := ffjson.Marshal(dbResult)
			conn.Do("SETEX", redisKey, 30, redisData)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "From DB",
				"data":    dbResult,
			})
			return
		}

		ffjson.Unmarshal(data, &args)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "From redis",
			"data":    args,
		})
	}
}

func CacheUsersDecorator(h gin.HandlerFunc, redisKey string, args interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		conn := database.RedisDefaultPool.Get()
		defer conn.Close()

		data, err := redis.Bytes(conn.Do("GET", redisKey))

		if err != nil {
			h(ctx)
			dbUsers, exists := ctx.Get("dbUsers")

			if !exists {
				dbUsers = args
			}

			redisData, _ := ffjson.Marshal(dbUsers)
			conn.Do("SETEX", redisKey, 30, redisData)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "From db",
				"data":    dbUsers,
			})
			return
		}

		ffjson.Unmarshal(data, &args)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "From redis",
			"data":    args,
		})

	}
}
