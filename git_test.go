package main

import (
  // "fmt"
  // "io/ioutil"
  // "strings"
  "testing"
)

func TestParseSimpleUrl(t *testing.T) {
  fileContents := `
[remote "webbynode"]
    url = git@200.15.2.13:crm_bliss
    fetch = +refs/heads/*:refs/remotes/webbynode/*
`

  git := GitConfig{}
  git.ReadFromFile(tempFixture(fileContents))
  git.Parse()

  assertEquals(t, "ip", git.ip, "200.15.2.13")
  assertEquals(t, "port", git.port, "22")
}

func TestParseSshUrl(t *testing.T) {
  fileContents := `
[remote "webbynode"]
    url = ssh://git@200.15.4.13:389/var/apps/crm_bliss
    fetch = +refs/heads/*:refs/remotes/webbynode/*
`

  git := GitConfig{}
  git.ReadFromFile(tempFixture(fileContents))
  git.Parse()

  assertEquals(t, "ip", git.ip, "200.15.4.13")
  assertEquals(t, "port", git.port, "389")
}
