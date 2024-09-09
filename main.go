package main

import (
	"fmt"
	"time"

	"github.com/mergestat/timediff"
)
func main() {
	fmt.Println("hello world")
	str1 := timediff.TimeDiff(time.Now().Add(-10 * time.Second))
	fmt.Print(str1)
}
