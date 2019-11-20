package router

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func Register(router *gin.RouterGroup) {
  router.GET("/ping", Ping)
  router.GET("/posts", GetPosts)
}

func Ping(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"success": true})
}

func GetPosts(c *gin.Context) {
  tags := c.Query("tags")
  sortBy := c.Query("sortBy")
  direction := c.Query("direction")

  if sortBy == "" {
    sortBy = "id"
  }
  if direction == "" {
    direction = "asc"
  }

  if tags == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Tags parameter is required"})
    return
  }

}

func ValidateParams(sort, direction string) bool {
  var valid_sort = []string{"id", "reads", "likes", "popularity"}
  valid_direction := []string{"desc", "asc"}
  if !contains(valid_sort, sort) {
    return false
  }
  if !contains(valid_direction, direction) {
    return false
  }
  return true
}

func contains(s []string, e string) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}
