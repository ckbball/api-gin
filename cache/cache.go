package cache

import (
  "github.com/ckbball/api-gin/handlers"
  "strings"
)

var cache = make(map[string][]*handlers.Post)

func InsertPosts(posts []*handlers.Post, tag string) {
  for _, post := range posts {
    cache[tag] = append(cache[tag], post)
  }
}

func InsertPost(post *handlers.Post, tag string) {
  cache[tag] = append(cache[tag], post)
}

func CheckCache(tag string) bool {
  return cache[tag] != nil
}

func GetTag(tag string) []*handlers.Post {
  return cache[tag]
}

func FilterCache(tags string) (string, []*handlers.Post) {
  list := strings.Split(tags, ",")
  out := []*handlers.Post{}
  newTags := []string{}
  for _, val := range list {
    if cache[val] != nil {
      for _, value := range cache[val] {
        out = append(out, value)
      }
    } else {
      newTags = append(newTags, val)
    }
  }
  res := strings.Join(newTags, ",")
  return res, out
}

func contains(s []*handlers.Post, e *handlers.Post) bool {
  for _, a := range s {
    if a.Id == e.Id {
      return true
    }
  }
  return false
}
