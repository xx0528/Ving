package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"server/database"
	"server/model"
	"server/msg"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var visited = make(map[string]bool)

func getContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}

func checkErr(err error) {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("没有查到数据")
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}

func TestGetData(t *testing.T) {
	var (
		err     error
		msgHome *msg.Home
	)

	fmt.Println("Hello, world")
	// url := "http://baobab.kaiyanapp.com/api/v4/tabs/selected"
	url := "http://baobab.kaiyanapp.com/api/v2/feed?&num=1&udid=d2807c895f0348a180148c9dfa6f2feeac0781b5&deviceModel=Android%20SDK%20built%20for%20x86"
	queue := make(chan string)

	rand.Seed(time.Now().UTC().UnixNano())

	db, err := database.ConnectDB("mongodb://root:123456@127.0.0.1:27017/admin", "Server")

	if err != nil {
		t.Log("连接数据库失败")
		panic(err)
	}
	defer db.Close()
	t.Log("数据库连接成功")

	videosCollection := db.DB.Collection("Videos")
	usersCollection := db.DB.Collection("Users")

	var videosIds = map[int64]bool{}
	var usersIds = map[int64]bool{}

	go func() {
		queue <- url
	}()
	for uri := range queue {
		msgHome = download(uri, queue)
		fmt.Println("next page url -- ", msgHome.NextPageUrl)
		fmt.Println("IssueList 长度 -- ", len(msgHome.IssueList[0].ItemList))

		var videosData []interface{}
		var usersData []interface{}
		for _, issue := range msgHome.IssueList {
			// fmt.Println("Itype - -", issue.Itype, " Count -- ", issue.Count)
			for _, item := range issue.ItemList {
				fmt.Println("type -- ", item.Itype, "title -- ", item.IData.Title)
				if item.Itype == "video" {
					fmt.Println("播放地址--", item.IData.PlayUrl)
					fmt.Println("作者--", item.IData.IAuthor.Name, "分类--", item.IData.Category)
					iData := item.IData
					uData := item.IData.IAuthor

					if _, ok := videosIds[iData.Id]; ok == false {
						videosData = append(videosData, model.Video{
							ID:         iData.Id,
							UserID:     uData.ID,
							UserName:   uData.Name,
							Title:      iData.Title,
							Desc:       iData.Description,
							PlayURL:    iData.PlayUrl,
							Category:   iData.Category,
							UpdateTime: iData.Date,
							Duration:   iData.Duration,
						})
						videosIds[iData.Id] = true
					} else {
						fmt.Println("已经有视频---", iData.Id)
					}

					if _, ok := usersIds[uData.ID]; ok == false {
						usersData = append(usersData, model.User{
							ID:   uData.ID,
							Name: uData.Name,
							Icon: uData.Icon,
							Desc: uData.Description,
						})
						usersIds[uData.ID] = true
					} else {
						fmt.Println("已经有作者---", uData.ID)
					}

				}
			}
		}
		// fmt.Println("user 长度---", len(usersData))
		var insertManyRes *mongo.InsertManyResult
		if len(usersData) > 0 {
			if insertManyRes, err = usersCollection.InsertMany(getContext(), usersData[0:]); err != nil {
				checkErr(err)
			}
			fmt.Printf("usersCollection InsertMany插入的消息ID:%v\n", insertManyRes.InsertedIDs)
		}

		if len(videosData) > 0 {
			if insertManyRes, err = videosCollection.InsertMany(getContext(), videosData[0:]); err != nil {
				checkErr(err)
			}
			fmt.Printf("videosCollection InsertMany插入的消息ID:%v\n", insertManyRes.InsertedIDs)
		}

		// var insertManyRes *mongo.InsertManyResult
		// if insertManyRes, err = videosCollection.InsertMany(getContext(), videosData[0:]); err != nil {
		// 	checkErr(err)
		// }
		// if insertManyRes, err := usersCollection.InsertMany(getContext(), usersData[0:]); err != nil {
		// 	checkErr(err)
		// }
	}
}

func download(url string, queue chan string) (data *msg.Home) {
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

	data = &msg.Home{}
	if err = json.Unmarshal(result, data); nil != err {
		fmt.Println("parses [msg home] failed: ", err)
	}

	// fmt.Println("next page url -- ", data.NextPageUrl)
	// fmt.Println("IssueList 长度 -- ", len(data.IssueList))
	// for _, issue := range data.IssueList {
	// 	// fmt.Println("Itype - -", issue.Itype, " Count -- ", issue.Count)
	// 	for _, item := range issue.ItemList {
	// 		fmt.Println("type -- ", item.Itype, "title -- ", item.IData.Title)
	// 		if item.Itype == "video" {
	// 			fmt.Println("播放地址--", item.IData.PlayUrl)
	// 			fmt.Println("作者--", item.IData.IAuthor.Name, "分类--", item.IData.Category)
	// 		}
	// 	}
	// }

	if data.NextPageUrl != "" {
		fmt.Println("parse url", data.NextPageUrl)
		time.Sleep(3 * time.Second)
		go func() {
			queue <- data.NextPageUrl
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
	return
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
