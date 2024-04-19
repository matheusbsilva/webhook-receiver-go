package main

import (
	"encoding/json"
	"fmt"
	"os"
  "time"
	"github.com/gin-gonic/gin"
)

type file struct {
  filepath string
  content map[string]any
}

func (j file) writeJsonFile() error {
  jsonString, _ := json.Marshal(j.content)
  return os.WriteFile(j.filepath, jsonString, os.ModePerm)
}

func main() {
  router := gin.Default()

  router.GET("/status", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "status": "ok",
    })
  })

  router.POST("/webhook", func(c *gin.Context) {
    var requestData map[string]any
    if err:= c.BindJSON(&requestData); err != nil {
      c.JSON(400, gin.H{"error": "Invalid request"})
      return
    }

    go func(requestData map[string]interface{}) {
      jsonData := file{
        filepath: fmt.Sprintf("data/%s.json", requestData["id"]),
        content: requestData,
      }

      err := jsonData.writeJsonFile()
      time.Sleep(5)

      if err != nil {
        fmt.Println(err)
      }
    }(requestData)

    c.JSON(200, gin.H{"message": "Webhook received!"})
  })

  router.Run()
}
