package utils

import (
    "fmt"
    "testing"
)


func TestMD5_32(t *testing.T) {
	str := "25704cf9-9cd4-4775-8c35-7e81e0cb1741,DZ2206HT09,18926016466"
	fmt.Println(MD5_32(str))
}

func TestMD5_16(t *testing.T) {
	str := "25704cf9-9cd4-4775-8c35-7e81e0cb1741,DZ2206HT09,18926016466"
	fmt.Println(MD5_16(str))
}
