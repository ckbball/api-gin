package routers

import (
  "fmt"
  "github.com/ckbball/api-gin/handlers"
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
  "log"
  "net/http"
  "os"
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

  fmt.Println("Query Params: tags: , sort: , direction: ", tags, sortBy, direction)

  if !ValidateSort(sortBy) {
    c.JSON(http.StatusBadRequest, gin.H{"error": "sortBy parameter is invalid"})
    return
  }

  if !ValidateDirection(direction) {
    c.JSON(http.StatusBadRequest, gin.H{"error": "direction parameter is invalid"})
    return
  }

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

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  url := os.Getenv("URL_STRING")

  //get blog posts
  // check cache
  // call api func that would go through each tag and make a call and add to cache
  posts, err := handlers.GetPosts(tags, url)

  result := handlers.SortPosts(posts, sortBy, direction)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err})
    return
  }

  //set response
  c.JSON(http.StatusOK, gin.H{"posts": final})
}

func ValidateSort(sort string) bool {

  var valid_sort = []string{"id", "reads", "likes", "popularity", ""}
  if !contains(valid_sort, sort) {
    return false
  }
  return true
}

func ValidateDirection(direction string) bool {

  valid_direction := []string{"desc", "asc", ""}
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
