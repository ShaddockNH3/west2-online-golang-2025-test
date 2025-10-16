package gobyexample

import (
	"fmt"
	"time"
)

func f1(from string) {
	for i := range 2 {
		fmt.Println(from, ":", i)
	}
}

func f2(from string) {
	for i := range 2 {
		fmt.Println(from, ":", i)
	}
}

func test6() {

	f1("direct")

	go f1("goroutine")
	go f2("hhh")
	go func(from string) {
		for i := range 2 {
			fmt.Println(from, ":", i)
		}
	}("jjj")

	time.Sleep(time.Second)
	fmt.Println("done")

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	// 1. 建造“时空隧道”
	message := make(chan string)

	// 2. 派遣“信使”去另一个宇宙执行任务
	go func() {
		fmt.Println("【信使】: 出发！我需要花2秒钟准备包裹...")
		time.Sleep(time.Second * 2) // 模拟信使准备包裹花费的时间

		fmt.Println("【信使】: 包裹'ping'准备好了，正要放入管道...")
		message <- "ping" // 把包裹放入管道，如果main还没准备好接收，这里会等待
		fmt.Println("【信使】: 包裹已被接收！我的任务完成了。")
	}()

	// 3. 我（主程序）在自己的世界里继续做别的事
	fmt.Println("【我】: 已经派出了信使，我可以先做点别的事情...")
	time.Sleep(time.Second * 1) // 模拟我做了1秒钟的其他工作
	fmt.Println("【我】: OK，现在我需要信使送来的那个包裹了，开始等待...")

	// 4. 在管道出口等待包裹
	msg := <-message // 程序会在这里卡住(阻塞)，直到信使把包裹放进来

	fmt.Println("【我】: 终于收到了包裹！内容是:", msg)
	fmt.Println("【我】: 所有事情都做完了。")

	fmt.Println()
	fmt.Println()
	fmt.Println()

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	messages <- "hhh"
	fmt.Println(<-messages)

}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ping(pings chan<- string, msg string) { // 只写
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // 从只读通道 pings 里取出一个值，然后把这个值，再发送到只写通道 pongs 里去。
	msg := <-pings
	pongs <- msg
}

func directions() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "pasted message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func select_try() {
	c1 := make(chan string)
	c2 := make(chan string)
	c3 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "1s"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "2s"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "1s"
	}()

	go func() {
		time.Sleep(100 * time.Second)
		c3 <- "100s"
	}()

	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)

		case msg3 := <-c3:
			fmt.Println("received", msg3)

		case <-time.After(2 * time.Second):
			fmt.Println("Out of Time")
		}

	}
}

func Non_Blocking() {
	messages := make(chan string)

	go func() {
		messages <- "Hello World!"
	}()

	time.Sleep(1 * time.Second)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

}

func close_try() {
	jobs := make(chan int, 5)
	done := make(chan bool, 1)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive job", j)
			} else {
				fmt.Println("all jobs done")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 4; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("receive more jobs", ok)
}

// func main() {
// 	// test6()
// 	// done:=make(chan bool,1)
// 	// go worker(done)

// 	// fmt.Println("111")

// 	// <-done

// 	// fmt.Println("222")
// 	// directions()
// 	// select_try()

// 	// Non_Blocking()

// 	close_try()

// 	fmt.Println()

// 	queue := make(chan int, 2)
// 	queue <- 1
// 	queue <- 2
// 	close(queue)
// 	for elem := range queue {
// 		fmt.Println(elem)
// 	}
// }
