package main

import (
  "testing"
)

func TestCapitalize(t *testing.T) {
  assertEquals(t, `Capitalize("hello")`, "Hello", Capitalize("hello"))
}

func TestClassify(t *testing.T) {
  assertEquals(t, `Classify("something")`, "Something", Classify("something"))
  assertEquals(t, `Classify("add_key")`, "AddKey", Classify("add_key"))
  assertEquals(t, `Classify("add_other_key")`, "AddOtherKey", Classify("add_other_key"))
}
