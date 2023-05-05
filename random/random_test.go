package utils

import (
    "testing"
    "fmt"
)


func TestRandom(t *testing.T) {
  fmt.Println(RandString(8, []rune{}))
}
