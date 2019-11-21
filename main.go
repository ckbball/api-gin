package main

import (
  "github.com/ckbball/api-gin/routers"
  "github.com/gin-gonic/gin"
)

func main() {

  //common.InitURL()

  r := gin.Default()

  v1 := r.Group("/api")

  {
    routers.Register(v1.Group(""))
  }

  r.Run()
}
