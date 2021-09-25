package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机时间
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	t1 := time.Duration(r1.Intn(10)) * time.Minute
	fmt.Println("now :", time.Now())
	fmt.Println("rand time: ", t1)
	fmt.Println("now add rand time", time.Now().Add(t1))
}
