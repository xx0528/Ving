package api

import (
	// "errors"
	"fmt"
	"math/rand"
	"server/msg"
	"strconv"
	"time"

	"github.com/88250/gulu"
	"github.com/gin-gonic/gin"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	// "strconv"
)

// The CommonDatabase interface for encapsulating database access.
type CommonDatabase interface {
}

// The CommonAPI provides handlers for managing users.
type CommonAPI struct {
	DB CommonDatabase
}

func (a *CommonAPI) GetHome(ctx *gin.Context) {
	result := gulu.Ret.NewResult()
	result.Msg = "getHome"

	// collection = a.DB.Collection("Films")
	msgHome := &msg.Home{
		NextPageUrl:     "127.0.0.1:5678/api/home/date=1581296400000&num=1",
		NextPublishTime: time.Now().Unix(),
		NewestIssueType: "night",
	}
	for i := 0; i < 3; i++ {
		issue := &msg.Issue{
			Count:       i,
			Date:        time.Now().Unix(),
			PublishTime: time.Now().Unix(),
		}
		for i := 0; i < 2; i++ {
			msgItem := &msg.Item{
				Itype: "videoCollectionWithBrief",
				Tag:   "null",
			}
			msgItem.IData = msg.IData{
				ActionUrl:   "",
				DataType:    "Banner",
				Description: "",
				Id:          0,
				Title:       "",
			}
			issue.ItemList = append(issue.ItemList, *msgItem)
		}
		msgHome.IssueList = append(msgHome.IssueList, *issue)
	}
	result.Data = msgHome
	defer ctx.JSON(http.StatusOK, result)
}
func (a *CommonAPI) GetVideo(ctx *gin.Context) {
	result := gulu.Ret.NewResult()
	result.Msg = "getVideo"
	defer ctx.JSON(http.StatusOK, result)
}

func (a *CommonAPI) GetCategories(ctx *gin.Context) {
	result := gulu.Ret.NewResult()
	result.Msg = "getCategories"
	defer ctx.JSON(http.StatusOK, result)
}

func (a *CommonAPI) GetVideoDetail(ctx *gin.Context) {
	result := gulu.Ret.NewResult()
	result.Msg = "getVideoDetail"
	defer ctx.JSON(http.StatusOK, result)
}

func (a *CommonAPI) GetRank(ctx *gin.Context) {
	result := gulu.Ret.NewResult()
	result.Msg = "getRank"
	defer ctx.JSON(http.StatusOK, result)
}

func (a *CommonAPI) GetSearch(ctx *gin.Context) {
	result := gulu.Ret.NewResult()
	result.Msg = "getSearch"
	defer ctx.JSON(http.StatusOK, result)
}

func (a *CommonAPI) GetHot(ctx *gin.Context) {
	result := gulu.Ret.NewResult()
	result.Msg = "getHot"
	defer ctx.JSON(http.StatusOK, result)
}

func (a *CommonAPI) GetFollow(ctx *gin.Context) {
	result := gulu.Ret.NewResult()
	result.Msg = "getFollow"
	defer ctx.JSON(http.StatusOK, result)
}

func (a *CommonAPI) GetAuthor(ctx *gin.Context) {
	result := gulu.Ret.NewResult()
	result.Msg = "getAuthor"
	defer ctx.JSON(http.StatusOK, result)
}

func (a *CommonAPI) GetAnimation(ctx *gin.Context) {
	// result := gulu.Ret.NewResult()
	// result.Msg = "获取动画片列表"

	var aniData = make(map[string]interface{})
	rand.Seed(time.Now().UnixNano())
	num := 3 + rand.Intn(6)
	aniData["itemList"] = a.GetAniList("Video", num, "", 0, 0)

	aniData["aType"] = "TypeAnimation"
	aniData["releaseTime"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["date"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["total"] = rand.Intn(5)
	aniData["publishTime"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["count"] = rand.Intn(5)
	aniData["nextPage"] = "http://baobab.kaiyanapp.com/api/v2/feed?date=1582506000000&num=1"
	defer ctx.JSON(http.StatusOK, aniData)
}

func (a *CommonAPI) GetMoreAnimation(ctx *gin.Context) {

	pageNum := ctx.Query("num")
	pageIdx, _ := strconv.Atoi(pageNum)
	var aniData = make(map[string]interface{})
	rand.Seed(time.Now().UnixNano())
	num := 3 + rand.Intn(6)
	aniData["itemList"] = a.GetAniList("Video", num, "", 0, pageIdx)
	aniData["aType"] = "TypeAnimation"
	aniData["releaseTime"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["date"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["total"] = rand.Intn(5)
	aniData["publishTime"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["count"] = rand.Intn(5)
	aniData["nextPage"] = "http://baobab.kaiyanapp.com/api/v2/feed?date=1582506000000&num=1"
	defer ctx.JSON(http.StatusOK, aniData)
}

func (a *CommonAPI) GetRelatedAnimation(ctx *gin.Context) {
	aniId := ctx.Query("id")
	fmt.Println("aniId -- -", aniId)

	iAniId, _ := strconv.Atoi(aniId)

	var aniData = make(map[string]interface{})
	rand.Seed(time.Now().UnixNano())

	var dataList []interface{}

	for i := 0; i < 3; i++ {
		dataList = append(dataList, a.GetTextCard("textCard"))
		num := 3 + rand.Intn(6)
		smailList := a.GetAniList("aniSmallCard", num, "", 0, iAniId)
		for j := 0; j < len(smailList); j++ {
			dataList = append(dataList, smailList[j])
		}
	}

	aniData["itemList"] = dataList
	aniData["aType"] = "TypeAnimation"
	aniData["releaseTime"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["date"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["total"] = rand.Intn(5)
	aniData["publishTime"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["count"] = rand.Intn(5)
	aniData["nextPage"] = "http://baobab.kaiyanapp.com/api/v2/feed?date=1582506000000&num=1"
	defer ctx.JSON(http.StatusOK, aniData)
}

func (a *CommonAPI) GetAniList(aType string, num int, category string, aniID int64, pageIdx int) (aniList []interface{}) {
	for i := 0; i < num; i++ {
		var authorData = make(map[string]interface{})
		authorData["id"] = 50000 + rand.Intn(10000)
		authorData["videoNum"] = rand.Intn(100)
		authorData["icon"] = "http://img.kaiyanapp.com/98beab66d3885a139b54f21e91817c4f.jpeg"
		authorData["name"] = "VING广告精选"
		authorData["description"] = "为广告人的精彩创意点赞"

		var tagList []interface{}
		tagNum := rand.Intn(5)
		for i := 0; i < tagNum; i++ {
			var tagData = make(map[string]interface{})
			tagData["id"] = 50000 + rand.Intn(10000)
			tagData["name"] = fmt.Sprintf("脑洞广告%d", i)
			tagData["actionUrl"] = "eyepetizer://tag/766/?title=%E8%84%91%E6%B4%9E%E5%B9%BF%E5%91%8A"
			tagData["bgPicture"] = "http://img.kaiyanapp.com/7c46ad04ff913b87915615c78d226a40.jpeg?imageMogr2/quality/60/format/jpg"
			tagData["headerImage"] = "http://img.kaiyanapp.com/0d6ab7ed49de67eab89ada4f65353e8c.jpeg?imageMogr2/quality/60/format/jpg"
			tagList = append(tagList, tagData)
		}

		var coverData = make(map[string]interface{})
		coverData["feed"] = "http://img.kaiyanapp.com/543451a3991b08eabb1a5f8a2f9f5d39.jpeg?imageMogr2/quality/60/format/jpg"
		coverData["blurred"] = "http://img.kaiyanapp.com/0379d60d05c7abdd12112bcda784d5cd.jpeg?imageMogr2/quality/60/format/jpg"
		coverData["detail"] = "http://img.kaiyanapp.com/543451a3991b08eabb1a5f8a2f9f5d39.jpeg?imageMogr2/quality/60/format/jpg"
		coverData["homepage"] = "http://img.kaiyanapp.com/543451a3991b08eabb1a5f8a2f9f5d39.jpeg?imageView2/1/w/720/h/560/format/jpg/q/75|watermark/1/image/aHR0cDovL2ltZy5rYWl5YW5hcHAuY29tL2JsYWNrXzMwLnBuZw==/dissolve/100/gravity/Center/dx/0/dy/0|imageslim"

		var consumption = make(map[string]interface{})
		consumption["collectionCount"] = rand.Intn(50000)
		consumption["shareCount"] = rand.Intn(50000)
		consumption["replyCount"] = rand.Intn(50000)

		var data = make(map[string]interface{})
		data["id"] = 50000 + rand.Intn(10000)
		data["name"] = fmt.Sprintf("动画名字--%d", i)
		data["title"] = fmt.Sprintf("动画标题--%d", i)
		data["aniNum"] = 10 + rand.Intn(80)
		data["curPlayNum"] = rand.Intn(10)
		data["aniTitle"] = fmt.Sprintf("动画标题标题--%d", i)
		data["dataType"] = fmt.Sprintf("Animation-%s", aType)
		data["desc"] = fmt.Sprintf("动画描述-000%d", i)
		data["author"] = authorData
		data["cover"] = coverData
		data["category"] = "动画分类"
		data["likeCount"] = rand.Intn(100000)
		data["score"] = float32(3+rand.Intn(7)) + (float32(rand.Intn(10)) / 10)
		data["playUrl"] = "http://ali.cdn.kaiyanapp.com/1582517488402_72e64987.mp4?auth_key=1582823404-0-0-73bf6c2341ca8fb63abd52377763062f"
		data["duration"] = 50 + rand.Intn(200)
		data["createTime"] = time.Now().Unix() - int64(rand.Intn(1000))
		data["tags"] = tagList
		data["aType"] = fmt.Sprintf("videoType-%s", aType)
		data["date"] = time.Now().Unix() - int64(rand.Intn(1000))
		data["idx"] = rand.Intn(100)
		data["label"] = "动画标签"
		data["promotion"] = "推广"
		data["played"] = false
		data["subtitles"] = "动画子标题"
		data["lastViewTime"] = "最后观看时间"
		data["consumption"] = consumption

		var aniData = make(map[string]interface{})
		aniData["aType"] = aType
		aniData["tag"] = "动画片"
		aniData["data"] = data
		aniList = append(aniList, aniData)
	}

	return
}

func (a *CommonAPI) GetTextCard(aType string) (aniData map[string]interface{}) {

	var data = make(map[string]interface{})
	data["id"] = 0
	data["dataType"] = "TextCard"
	data["text"] = "最新热血番"
	data["aType"] = "header4"

	aniData = make(map[string]interface{})
	aniData["aType"] = aType
	aniData["data"] = data

	return aniData
}
