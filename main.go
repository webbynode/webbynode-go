package main

import (
  "flag"
  // "fmt"
  "log"
  "os"
)

func main() {
  flag.Parse()
  if err := ParseCommands(flag.Args()...); err != nil {
    log.Fatal(err)
    os.Exit(-1)
  }
}
