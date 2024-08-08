package main

import (
  "fmt"
  "strings"
)

func main() {

  const s = "Hello, world!"

  fmt.Println("Len, ", len(s))

  for i := 0; i < len(s); i++ {
    fmt.Printf("%c ", s[i])
  }

  fmt.Println()
  fmt.Println("Contains: ", strings.Contains(s, "Hello"))

  fmt.Println("Count: ", strings.Count(s, "l"))

  fmt.Println("HasPrefix: ", strings.HasPrefix(s, "Hello"))

  fmt.Println("HasSuffix: ", strings.HasSuffix(s, "world!"))

  fmt.Println("Index: ", strings.Index(s, "world"))

  fmt.Println("Join: ", strings.Join([]string{"Hello", "world"}, " "))

  fmt.Println("Repeat: ", strings.Repeat(s, 2))

  fmt.Println("Replace: ", strings.Replace(s, "world", "Go", 1))

  fmt.Println("Split: ", strings.Split(s, " "))
  // print type of Split
  fmt.Printf("%T \n", strings.Split(s, " "))

  fmt.Println("ToLower: ", strings.ToLower(s))

  fmt.Println("ToUpper: ", strings.ToUpper(s))

  fmt.Println("TrimSpace: ", strings.TrimSpace(" Hello, world! "))


  fmt.Println("TrimSpace: ", strings.TrimSpace(" Hello, world! "))

  fmt.Println("Trim: ", strings.Trim(" !!! Hello, world! !!! ", "! "))

}
