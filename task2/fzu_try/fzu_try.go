package fzu_try

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func HttpGet(url string) (result string, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// SpiderPage 函数负责抓取单个页面并保存
func SpiderPage(i int, page chan int) {
	// 为了不让并发和串行创建的文件混在一起，我们给并发抓取的文件名加个前缀
	filename := "并发-第" + strconv.Itoa(i) + "页.html"
	url := "https://jwch.fzu.edu.cn/jxtz/" + strconv.Itoa(i) + ".htm"

	result, err := HttpGet(url)
	if err != nil {
		// 在并发中，一个页面的失败不应该让整个程序停止，所以我们只打印错误
		fmt.Printf("\n抓取第 %d 页失败: %v\n", i, err)
		page <- i // 即使失败了，也要通知主goroutine，表示这个任务结束了
		return
	}

	// os.WriteFile 是一个更简洁的写入文件的方式
	err = os.WriteFile(filename, []byte(result), 0644)
	if err != nil {
		fmt.Printf("\n保存第 %d 页失败: %v\n", i, err)
	}

	page <- i // 与主goroutine完成同步
}

// working2 是我们的并发选手
func working2(start, end int) {
	fmt.Printf("\n--- 并发选手入场，开始抓取第 %d 页到 %d 页 ---\n", start, end)
	
    // 计时开始！
	startTime := time.Now()

	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}

	// 等待所有goroutine完成
	for i := start; i <= end; i++ {
		// 接收一个完成信号，打印一下进度
		p := <-page
		fmt.Printf("并发任务：第 %d 页处理完成。\n", p)
	}
    
    // 计时结束！
	duration := time.Since(startTime)
	fmt.Printf("--- 并发选手比赛结束，总共用时: %v ---\n", duration)
}

// working1 是我们的串行选手
func working1(start, end int) {
	fmt.Printf("\n--- 串行选手入场，开始抓取第 %d 页到 %d 页 ---\n", start, end)

    // 计时开始！
	startTime := time.Now()

	for i := start; i <= end; i++ {
		url := "https://jwch.fzu.edu.cn/jxtz/" + strconv.Itoa(i) + ".htm"
		result, err := HttpGet(url)
		if err != nil {
			fmt.Printf("\n抓取第 %d 页失败: %v\n", i, err)
			continue // 跳过这个失败的页面，继续下一个
		}

		filename := "串行-第" + strconv.Itoa(i) + "页.html"
		err = os.WriteFile(filename, []byte(result), 0644)
		if err != nil {
			fmt.Printf("\n保存第 %d 页失败: %v\n", i, err)
		}
		fmt.Printf("串行任务：第 %d 页处理完成。\n", i)
	}
    
    // 计时结束
	duration := time.Since(startTime)
	fmt.Printf("--- 串行选手比赛结束，总共用时: %v ---\n", duration)
}

func main() {
	var start, end int
	fmt.Println("请输入爬取的起始页 (>=1):")
	fmt.Scan(&start)
	fmt.Println("请输入爬取的结束页 (>=start):")
	fmt.Scan(&end)

	working1(start, end)

	working2(start, end)
}
