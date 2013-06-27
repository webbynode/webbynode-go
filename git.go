package main

import (
  // "bytes"
  // "io"
  // "flag"
  "fmt"
  // "os"
  // "os/exec"
  // "reflect"
  "regexp"
  // "strings"
  "github.com/vaughan0/go-ini"
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
  if regexp.MustCompile(`^ssh://(\w+)@(.+)/(.+)$`).MatchString(url) == true {
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

func (git *GitConfig) SshConsole() {
  // args := []string { "-p", git.port, git.user + "@" + git.ip }
  // process, err := os.StartProcess("/usr/bin/ssh", args, &os.ProcAttr{})
  // if err != nil {
  //   panic(err)
  // }
  // fmt.Println(process)
  // process.Wait()
  // fmt.Println(process.String())

  // cmd := exec.Command("ssh", "-p", git.port, git.user + "@" + git.ip)
  // err := cmd.Start()
  // if err != nil {
  //   panic(err)
  // }
  // err = cmd.Wait()
  // if err != nil {
  //   fmt.Fprintln(os.Stderr, err)
  //   return
  // }

  // var b bytes.Buffer
  // io.Copy(&b, cmd.Stdout)
  // fmt.Println(b.String())
  fmt.Println("This command is not yet implemented.")
  fmt.Println("Meanwhile use this to log into your Webby:")
  fmt.Printf("  ssh -p %s %s@%s\n", git.port, git.user, git.ip)
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
