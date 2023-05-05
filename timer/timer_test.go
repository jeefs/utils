package utils

import (
    "testing"
    "time"
)

func TestTimer(t *testing.T) {
    t1 := NewTimer("计时器01")
    t1.Start()
    time.Sleep(2 * time.Second)
    t1.End() 
}
