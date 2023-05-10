package utils 

import (
    "testing"
    "fmt"
)

func TestA(t *testing.T) {
        s1 := []int{2, 3, 3, 5}
	fmt.Println(IsContinuous(s1))
}


func TestB(t *testing.T) {
	chunkSize := 2
        s1 := []int{1, 2, 3}
        res := SliceChunk(s1,chunkSize)
        fmt.Printf("chunk is:%v\n", res)

}
