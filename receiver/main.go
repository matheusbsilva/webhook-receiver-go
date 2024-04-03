package main

import (
	"encoding/json"
	"fmt"
	"os"
  "time"
	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/status", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "status": "ok",
    })
  })

  router.POST("/webhook", func(c *gin.Context) {
    var requestData map[string]interface{}
    if err:= c.BindJSON(&requestData); err != nil {
      c.JSON(400, gin.H{"error": "Invalid request"})
      return
    }

    go func(requestData map[string]interface{}) {
      jsonString, _ := json.Marshal(requestData)
      id, _ := requestData["id"]
      filepath := fmt.Sprintf("data/%s.json", id)
      fmt.Println(jsonString)

      err := os.WriteFile(filepath, jsonString, os.ModePerm)
      time.Sleep(5)

      if err != nil {
        fmt.Println(err)
      }
    }(requestData)

    c.JSON(200, gin.H{"message": "Webhook received!"})
  })

  router.Run()
}
