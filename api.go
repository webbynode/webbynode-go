package main

import (
  "fmt"
  // "io"
  "io/ioutil"
  // "os"
  "os/user"
  "log"
  "strings"
)

func GetCredentials() {
  usr, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }
  b, err := ioutil.ReadFile(usr.HomeDir + "/.webbynode")
  if err != nil {
    panic(err)
  }
  lines := strings.Split(string(b), "\n")
  config := &WebbynodeConfig{}
  for _, line := range lines {
    if strings.TrimSpace(line) != "" {
      parts := strings.SplitN(line, "=", 2)
      fmt.Println(parts[0])
      if parts[0] == "email" {
        config.Email = parts[1]
      }
    }
  }
  fmt.Println(config)
}

type WebbynodeConfig struct {
  Email       string
  Token       string
  System      string
  AwsKey      string
  AwsSecret   string
}

