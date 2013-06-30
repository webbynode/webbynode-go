package main

import (
  "fmt"
  "github.com/vaughan0/go-ini"
  "regexp"
  "syscall"
)

type GitConfig struct {
  file ini.File
  user string
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
  if isSshFormatted(url) {
    re1, err := regexp.Compile(`^ssh://(\w+)@(.+)/(.+)$`)
    if err != nil {
      panic(nil)
    }

    res := re1.FindStringSubmatch(url)
    git.user = res[1]
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
    git.user = res[1]
    git.ip = res[2]
    git.port = "22"
  }
}

func (git *GitConfig) AddSshKey(passphrase string) {
  sshKey := GetHomePath(".ssh/id_rsa.pub")
  if !FileExists(sshKey) {
    createSshKey(sshKey, passphrase)
    fmt.Println("Doesn't exist")
  }
}

func createSshKey(keyFile, passphrase string) {
  // command := `/bin/bash -c "test -f %s"
}

var SshFormat = regexp.MustCompile(`^ssh://(\w+)@(.+)/(.+)$`)

func isSshFormatted(url string) bool {
  return SshFormat.MatchString(url)
}

func (git *GitConfig) SshConsole() {
  syscall.Exec("/usr/bin/ssh",
    []string{"ssh", "-p", git.port, git.user + "@" + git.ip},
    []string{})
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
