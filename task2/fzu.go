package task2

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/xuri/excelize/v2"

	_ "github.com/go-sql-driver/mysql"
)

type Element struct {
	Author string
	Title  string
	Date   string
	Link   string

	Context                 string
	Is_Attachment           bool
	Attachment_Download_URL string
	Attachment_Download_NUM int
}

func saveToExcel(elements []Element, filename string) error {
	f := excelize.NewFile()
	sheetName := "Sheet1"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return fmt.Errorf("failed to create new sheet: %w", err)
	}

	headers := []string{
		"Author", "Title", "Date", "Link", "Context",
		"Is_Attachment", "Attachment_Download_URL", "Attachment_Download_NUM",
	}
	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	for i, el := range elements {
		row := i + 2
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), el.Author)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), el.Title)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), el.Date)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), el.Link)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", row), el.Context)
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", row), el.Is_Attachment)
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", row), el.Attachment_Download_URL)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", row), el.Attachment_Download_NUM)
	}

	f.SetActiveSheet(index)
	f.DeleteSheet("Sheet1")

	if err := f.SaveAs(filename); err != nil {
		return fmt.Errorf("failed to save excel file: %w", err)
	}

	return nil
}

func GetHttp(url string) (*goquery.Document, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败:%w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("执行请求失败:%w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("请求失败，状态码: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("加载HTML失败: %w", err)
	}
	return doc, nil
}

func GetHomePage(doc *goquery.Document, processItem func(author, title, date, link string)) {
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		date := strings.TrimSpace(s.Find("span.doclist_time").Text())

		aTag := s.Find("a[title]")
		title := aTag.AttrOr("title", "")
		link := aTag.AttrOr("href", "")

		fullText := s.Text()
		titleText := aTag.Text()
		tempText := strings.Replace(fullText, date, "", 1)
		author := strings.TrimSpace(strings.Replace(tempText, titleText, "", 1))

		if author != "" && title != "" && date != "" && link != "" {
			processItem(author, title, date, link)
		}
	})
}

func GetDetails(link string) (context string, is_attachment bool, attachment_download_url string, attachment_download_num int) {
	doc, err := GetHttp(link)
	if err != nil {
		context = "获取页面失败: " + err.Error()
		return
	}

	possibleSelectors := []string{
		"div.v_news_content",
		"div#vsb_content",
		"td.content",
	}
	for _, selector := range possibleSelectors {
		selection := doc.Find(selector)
		if selection.Length() > 0 {
			context = strings.TrimSpace(selection.Text())
			if context != "" {
				break
			}
		}
	}
	if context == "" {
		var sb strings.Builder
		doc.Find("p").Each(func(i int, s *goquery.Selection) {
			sb.WriteString(s.Text())
			sb.WriteString("\n")
		})
		context = strings.TrimSpace(sb.String())
	}

	attachmentTag := doc.Find("ul li a[href*='download.jsp']")
	if attachmentTag.Length() > 0 {
		is_attachment = true
		href, _ := attachmentTag.Attr("href")
		if !strings.HasPrefix(href, "http") {
			attachment_download_url = "https://jwch.fzu.edu.cn" + href
		} else {
			attachment_download_url = href
		}
	}

	scriptTag := doc.Find("span[id*='nattach'] script")
	if scriptTag.Length() > 0 {
		scriptContent := scriptTag.Text()
		re := regexp.MustCompile(`getClickTimes\((\d+),(\d+),".*?",".*?"\)`)
		matches := re.FindStringSubmatch(scriptContent)

		if len(matches) == 3 {
			wbnewsid := matches[1]
			owner := matches[2]

			ajaxURL := fmt.Sprintf("https://jwch.fzu.edu.cn/system/resource/code/news/click/clicktimes.jsp?wbnewsid=%s&owner=%s&type=wbnewsfile&randomid=nattach", wbnewsid, owner)

			resp, err := http.Get(ajaxURL)
			if err == nil {
				defer resp.Body.Close()
				body, _ := io.ReadAll(resp.Body)

				var data map[string]interface{}
				if json.Unmarshal(body, &data) == nil {
					if times, ok := data["wbshowtimes"].(float64); ok {
						attachment_download_num = int(times)
					}
				}
			}
		}
	}

	return
}

func working1(start, end int) {
	var url string
	var elements []Element

	for i := start; i <= end; i++ {
		if i != 204 {
			url = "https://jwch.fzu.edu.cn/jxtz/" + strconv.Itoa(i) + ".htm"
		} else {
			url = "https://jwch.fzu.edu.cn/jxtz.htm"
		}
		doc, err := GetHttp(url)
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}

		GetHomePage(doc, func(author, title, date, link string) {

			fullLink := "https://jwch.fzu.edu.cn/" + strings.TrimPrefix(link, "../")
			context, is_attachment, attachment_download_url, attachment_download_num := GetDetails(fullLink)

			element_piece := Element{
				Author:                  author,
				Title:                   title,
				Date:                    date,
				Link:                    fullLink,
				Context:                 context,
				Is_Attachment:           is_attachment,
				Attachment_Download_URL: attachment_download_url,
				Attachment_Download_NUM: attachment_download_num,
			}
			elements = append(elements, element_piece)
			fmt.Printf("  > 已处理并收入elements中: %s\n", title)

		})
	}
	err := saveToExcel(elements, "working1.xlsx")
	if err != nil {
		fmt.Println("ERR", err)
	}
}

func SpiderPage(i int, page chan int, elements *[]Element, mu *sync.Mutex, url string) {
	// 1. 用 defer 保证小助手无论如何都会报告！
	defer func() {
		page <- i
	}()

	if i != 204 {
		url = "https://jwch.fzu.edu.cn/jxtz/" + strconv.Itoa(i) + ".htm"
	} else {
		url = "https://jwch.fzu.edu.cn/jxtz.htm"
	}
	doc, err := GetHttp(url)
	if err != nil {
		fmt.Println("ERROR:", err)
		return // 因为有 defer，这里退出也会发送信号
	}

	GetHomePage(doc, func(author, title, date, link string) {
		fullLink := "https://jwch.fzu.edu.cn/" + strings.TrimPrefix(link, "../")
		context, is_attachment, attachment_download_url, attachment_download_num := GetDetails(fullLink)

		element_piece := Element{
			Author:                  author,
			Title:                   title,
			Date:                    date,
			Link:                    fullLink,
			Context:                 context,
			Is_Attachment:           is_attachment,
			Attachment_Download_URL: attachment_download_url,
			Attachment_Download_NUM: attachment_download_num,
		}

		// 3. 让小助手们学会排队！
		mu.Lock() // 先拿到“可以放东西”的许可
		// 2. 正确使用魔法信封，用 * 打开它！
		*elements = append(*elements, element_piece)
		mu.Unlock() // 放好后就让出许可
		fmt.Printf("  > (并发)已处理并收入elements中: %s\n", title)

	})
}

func working2(start, end int) {
	var url string
	var elements []Element
	// 准备一个“许可”（Mutex），保证一次只有一个小助手能放东西
	var mu sync.Mutex

	page := make(chan int)
	for i := start; i <= end; i++ {
		// 把“许可”也交给小助手
		go SpiderPage(i, page, &elements, &mu, url)
	}

	for i := start; i <= end; i++ {
		p := <-page
		fmt.Printf("并发任务：第 %d 页处理完成。\n", p)
	}

	err := saveToExcel(elements, "working2.xlsx")
	if err != nil {
		fmt.Println("ERR", err)
	}
}

func get_data() {
	var start int
	var end int

	fmt.Println("Please input start num(>=1)")
	fmt.Scanln(&start)
	fmt.Println("Please input end num(>=start && <=204)")
	fmt.Scanln(&end)

	startTime1 := time.Now()
	working1(start, end)
	duration1 := time.Since(startTime1)

	startTime2 := time.Now()
	working2(start, end)
	duration2 := time.Since(startTime2)

	fmt.Println("working1 use time:", duration1)
	fmt.Println("working1 use time:", duration2)

	speedup := float64(duration1) / float64(duration2)
	fmt.Println("working1/working2=", speedup)
}

func main(){
	get_data()
}
