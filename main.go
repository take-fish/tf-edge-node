package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// version 由构建时通过 -ldflags "-X main.version=..." 注入
var version = "dev"

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "tf-edge-node", "version": version})
	})

	// 伪装路径示例：边缘节点状态上报（运维巡检风格）
	api := r.Group("/api/v1")
	{
		api.GET("/edge/status", func(c *gin.Context) {
			hostname, _ := os.Hostname()
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "success",
				"data": gin.H{"node": hostname, "online": true, "version": version},
			})
		})
	}
	return r
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	gin.SetMode(gin.ReleaseMode)
	log.Printf("tf-edge-node %s listening on :%s", version, port)
	if err := setupRouter().Run(":" + port); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
