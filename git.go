package main

import (
  // "flag"
  // "fmt"
  // "reflect"
  "regexp"
  // "strings"
  "github.com/vaughan0/go-ini"
)

type GitConfig struct {
  file ini.File
  ip   string
  port string
  home string
}

func (git *GitConfig) Parse() {
  if git.file == nil {
    git.Read()
  }

  file := git.file
  url, _ := file.Get("remote \"webbynode\"", "url")
  if regexp.MustCompile(`^ssh://(\w+)@(.+)/(.+)$`).MatchString(url) == true {
    re1, err := regexp.Compile(`^ssh://(\w+)@(.+)/(.+)$`)
    if err != nil {
      panic(nil)
    }

    res := re1.FindStringSubmatch(url)
    part := res[2]
    re2, _ := regexp.Compile(`(.*):(\d*)\/(.*)$`)
    if re2.MatchString(part) {
      res := re2.FindStringSubmatch(part)
      git.ip = res[1]
      git.port = res[2]
      git.home = res[3]
    }
  } else {
    re1, err := regexp.Compile(`^(\w+)@(.+):(.+)$`)
    if err != nil {
      panic(nil)
    }
    res := re1.FindStringSubmatch(url)
    git.ip = res[2]
    git.port = "22"
  }
}

func (git *GitConfig) ReadFromFile(fileName string) {
  file, err := ini.LoadFile(fileName)
  if err != nil {
    panic(err)
  }
  git.file = file
}

func (git *GitConfig) Read() {
  git.ReadFromFile(".git/config")
}
