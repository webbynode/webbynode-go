package main

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

func FileExists(name string) bool {
  if _, err := os.Stat(name); err == nil {
    return true
  }
  return false
}

func Ask(prompt string) (string, error) {
  fmt.Print(prompt)
  reader := bufio.NewReader(os.Stdin)
  value, err := reader.ReadString('\n')
  return strings.TrimSpace(value), err
}

func AskYN(prompt string) (bool, error) {
  v, err := Ask(prompt + " (y/n)? ")
  if (err != nil) {
    return false, err
  }
  return strings.ToLower(v) == "y", err
}

func CopyFile(source, target string) error {
  // read whole the file
  b, err := ioutil.ReadFile(source)
  if err != nil {
    return err
  }

  // write whole the body
  err = ioutil.WriteFile(target, b, 0644)
  if err != nil {
    return err
  }

  return nil
}

func RenameFile(source, target string) error {
  return os.Rename(source, target)
}
