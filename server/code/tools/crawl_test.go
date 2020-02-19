package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"server/msg"
	"testing"
	"time"
)

var visited = make(map[string]bool)

func TestGetData(t *testing.T) {
	fmt.Println("Hello, world")
	// url := "http://baobab.kaiyanapp.com/api/v4/tabs/selected"
	url := "http://baobab.kaiyanapp.com/api/v2/feed?&num=1&udid=d2807c895f0348a180148c9dfa6f2feeac0781b5&deviceModel=Android%20SDK%20built%20for%20x86"
	queue := make(chan string)
	go func() {
		queue <- url
	}()
	for uri := range queue {
		download(uri, queue)
	}
}

func download(url string, queue chan string) {
	visited[url] = true
	timeout := time.Duration(5 * time.Second)

	client := &http.Client{
		Timeout: timeout,
	}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(result))

	msgHome := &msg.Home{}
	if err = json.Unmarshal(result, msgHome); nil != err {
		fmt.Println("parses [msg home] failed: ", err)
	}

	// fmt.Println("next page url -- ", msgHome.NextPageUrl)
	// fmt.Println("IssueList 长度 -- ", len(msgHome.IssueList))
	for _, issue := range msgHome.IssueList {
		// fmt.Println("Itype - -", issue.Itype, " Count -- ", issue.Count)
		for _, item := range issue.ItemList {
			fmt.Println("type -- ", item.Itype, "title -- ", item.IData.Title)
			if item.Itype == "video" {
				fmt.Println("播放地址--", item.IData.PlayUrl)
				fmt.Println("作者--", item.IData.IAuthor.Name, "分类--", item.IData.Category)
			}
		}
	}

	if msgHome.NextPageUrl != "" {
		fmt.Println("parse url", msgHome.NextPageUrl)
		time.Sleep(3 * time.Second)
		go func() {
			queue <- msgHome.NextPageUrl
		}()
	} else {
		fmt.Println(string(result))
	}
	// links := collectlinks.All(resp.Body)
	// for _, link := range links {
	// 	fmt.Println("url---", url)
	// 	absolute := urlJoin(link, url)
	// 	if url != " " {
	// 		if !visited[absolute] {
	// 			fmt.Println("parse url", absolute)
	// 			// go func() {
	// 			// 	queue <- absolute
	// 			// }()
	// 		}
	// 	}
	// }
}

func urlJoin(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return " "
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return " "
	}
	return baseUrl.ResolveReference(uri).String()
}
