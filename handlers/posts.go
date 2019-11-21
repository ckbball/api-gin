package handlers

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "sort"
  "strings"
)

type Post struct {
  Id         int      `json:"id"`
  Author     string   `json:"author"`
  AuthorId   int      `json:"authorId"`
  Likes      int      `json:"likes"`
  Popularity float32  `json:"popularity"`
  Reads      int      `json:"reads"`
  Tags       []string `json:"tags"`
}

type Posts struct {
  Posts []*Post `json:"posts"`
}

func GetPosts(tags string, url string) ([]*Post, error) {
  t := strings.Split(tags, ",")

  var out_posts []*Post
  keys := make(map[int]bool)

  for _, in := range t {
    full := urlBuilder(url, in)
    response, err := http.Get(full)
    if err != nil {
      fmt.Println("HTTP request failed with error %s\n", err)
      return nil, err
    }
    var posts Posts
    if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
      log.Println(err)
    }
    for _, val := range posts.Posts {
      if _, value := keys[val.Id]; !value {
        keys[val.Id] = true
        out_posts = append(out_posts, val)
      }
    }
  }
  return out_posts, nil
}

func urlBuilder(base, tag string) string {
  var sb strings.Builder
  sb.WriteString(base)
  sb.WriteString("?tag=")
  sb.WriteString(tag)
  out := sb.String()
  return out
}

// sor can be 'id' 'reads' 'likes' 'popularity'
// direction can be 'asc' 'desc'
func SortPosts(posts []*Post, sor, direction string) []*Post {

  if sor == "reads" && direction == "asc" {
    sort.Slice(posts, func(i, j int) bool { return posts[i].Reads < posts[j].Reads })
  } else if sor == "reads" && direction == "desc" {
    sort.Slice(posts, func(i, j int) bool { return posts[i].Reads > posts[j].Reads })
  } else if sor == "likes" && direction == "asc" {
    sort.Slice(posts, func(i, j int) bool { return posts[i].Likes < posts[j].Likes })
  } else if sor == "likes" && direction == "desc" {
    sort.Slice(posts, func(i, j int) bool { return posts[i].Likes > posts[j].Likes })
  } else if sor == "popularity" && direction == "asc" {
    sort.Slice(posts, func(i, j int) bool { return posts[i].Popularity < posts[j].Popularity })
  } else if sor == "popularity" && direction == "desc" {
    sort.Slice(posts, func(i, j int) bool { return posts[i].Popularity > posts[j].Popularity })
  } else if sor == "id" && direction == "asc" {
    sort.Slice(posts, func(i, j int) bool { return posts[i].Id < posts[j].Id })
  } else if sor == "id" && direction == "desc" {
    sort.Slice(posts, func(i, j int) bool { return posts[i].Id > posts[j].Id })
  }

  fmt.Printf("Sorted Posts: \n")
  for _, post := range posts {
    fmt.Printf("%+v \n", post.Likes)
  }
  return posts
}

// Deprecated here incase its needed
func Deduplicate(posts []*Post) []*Post {
  keys := make(map[int]bool)
  list := []*Post{}
  for _, entry := range posts {
    if _, value := keys[entry.Id]; !value {
      keys[entry.Id] = true
      list = append(list, entry)
    }
  }
  return list
}
