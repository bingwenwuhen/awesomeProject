package main

import (
	"fmt"
	"time"
)

func main() {
	unix := time.Now().Unix()
	fmt.Println(unix)
}
