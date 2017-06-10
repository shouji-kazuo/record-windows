package config

import (
  "testing"
  "fmt"
)

func TestConfig_ImportFromJsonPath(t *testing.T) {
  tConfig, tError := New().ImportFromJsonPath("./test.json")
  if tError != nil {
    t.Fatalf("Fail. error = %v", tError)
  } else {
    fmt.Println("tConfig = ", tConfig)
  }
}
