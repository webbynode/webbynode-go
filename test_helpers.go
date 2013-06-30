package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "testing"
)

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
