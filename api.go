package main

import (
  "path/filepath"
  "fmt"
  // "io"
  "io/ioutil"
  // "os"
  "os/user"
  "log"
  "strings"
)

func getHomePath(file string) string {
  usr, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }
  return filepath.Join(usr.HomeDir, file)
}

func readConfig() []string {
  b, err := ioutil.ReadFile(getHomePath(".webbynode"))
  if err != nil {
    panic(err)
  }
  return strings.Split(string(b), "\n")
}

func GetCredentials() {
  config := WebbynodeConfig{}
  config.Load()
  fmt.Println(config)
}

func (cfg *WebbynodeConfig) Load() error {
  for _, line := range readConfig() {
    if strings.TrimSpace(line) != "" {
      parts := strings.SplitN(line, "=", 2)
      fmt.Println(parts)
      if parts[0] == "email" {
        cfg.Email = parts[1]
      }
    }
  }
  return nil
}

type WebbynodeConfig struct {
  Email       string
  Token       string
  System      string
  AwsKey      string
  AwsSecret   string
}

