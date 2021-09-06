package main

func ping(pings chan string, msg string) {
	pings <- msg
}

func pong(pings chan string, pongs chan string) {
	// 可读性
	//msg := <-pings
	//pongs <- msg
	// 简洁性
	pongs <- <-pings
}

//func main() {
//	pings := make(chan string, 1)
//	pongs := make(chan string, 1)
//
//	ping(pings, "passed message")
//	pong(pings, pongs)
//
//	fmt.Println(<-pongs)
//}
