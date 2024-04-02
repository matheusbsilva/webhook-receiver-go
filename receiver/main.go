package main

import "github.com/gin-gonic/gin"

func main() {
  router := gin.Default()

  router.GET("/status", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "status": "ok",
    })
  })


  router.Run()
}