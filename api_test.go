package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "testing"
)

func TestLoadConfig(t *testing.T) {
  fileContents := `
email=felipe.coury@gmail.com
token=TOKEN12345
system=manager
aws_key=AWS_KEY
aws_secret=AWS_SECRET
`

  config := WebbynodeCfg{configFile: tempFixture(fileContents)}
  config.Load()

  assertEquals(t, "email", "felipe.coury@gmail.com", config.email)
  assertEquals(t, "token", "TOKEN12345", config.token)
  assertEquals(t, "system", "manager", config.system)
  assertEquals(t, "aws_key", "AWS_KEY", config.awsKey)
  assertEquals(t, "aws_secret", "AWS_SECRET", config.awsSecret)
}

func TestSaveConfig(t *testing.T) {
  tempFile, err := ioutil.TempFile("", "")
  if err != nil {
    panic(err)
  }

  fileName := tempFile.Name()

  config := WebbynodeCfg{
    configFile: fileName,
    email:      "email@something.com",
    token:      "token",
    system:     "manager",
    awsKey:     "awsKey",
    awsSecret:  "awsSecret",
  }
  config.Save()

  b, err := ioutil.ReadFile(fileName)
  if err != nil {
    panic(err)
  }
  contents := string(b)

  assertContains(t, contents, "email=email@something.com")
  assertContains(t, contents, "token=token")
  assertContains(t, contents, "system=manager")
  assertContains(t, contents, "aws_key=awsKey")
  assertContains(t, contents, "aws_secret=awsSecret\n")
}

func assertEquals(t *testing.T, item, expected, actual string) {
  if expected != actual {
    failure := fmt.Sprintf("Expected %s to be '%s' but got '%s'",
      item, expected, actual)
    t.Error(failure)
  }
}

func assertContains(t *testing.T, contents, snippet string) {
  if !strings.Contains(contents, snippet) {
    failure := fmt.Sprintf("Expected '%s' to include '%s' but it didn't.",
      contents, snippet)
    t.Error(failure)
  }
}

func tempFixture(contents string) string {
  tempFile, err := ioutil.TempFile("", "")
  if err != nil {
    panic(err)
  }
  err = ioutil.WriteFile(tempFile.Name(), []byte(contents), 0644)
  if err != nil {
    panic(err)
  }
  return tempFile.Name()
}
