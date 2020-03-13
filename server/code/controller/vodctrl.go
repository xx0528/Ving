package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"server/service"
	"server/util"

	"github.com/88250/gulu"
	"github.com/gin-gonic/gin"
)

func GetAnimation(ctx *gin.Context) {
	page := util.GetPage(ctx)

	typeId, typeId1, pageSize := 4, 0, 20

	var aniData = make(map[string]interface{})
	aniData["itemList"] = GetVodsList(typeId, typeId1, page, pageSize)
	aniData["typeId"] = typeId
	aniData["typeId1"] = typeId1

	aniData["releaseTime"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["date"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["total"] = rand.Intn(5)
	aniData["publishTime"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["count"] = rand.Intn(5)
	aniData["nextPage"] = ""
	defer ctx.JSON(http.StatusOK, aniData)
}

func GetMoreAnimation(ctx *gin.Context) {
	result := gulu.Ret.NewResult()

	page := util.GetPage(ctx)
	pageSize := util.GetIntParam(ctx, "num")
	if pageSize > 50 {
		result.Msg = "请求参数错误"
		defer ctx.JSON(http.StatusOK, result)
		return
	}

	typeId, typeId1 := 4, 0
	var aniData = make(map[string]interface{})
	aniData["itemList"] = GetVodsList(typeId, typeId1, page, pageSize)
	aniData["typeId"] = typeId
	aniData["typeId1"] = typeId1

	aniData["releaseTime"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["date"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["total"] = rand.Intn(5)
	aniData["publishTime"] = time.Now().Unix() - int64(rand.Intn(1000))
	aniData["count"] = rand.Intn(5)
	aniData["nextPage"] = ""
	defer ctx.JSON(http.StatusOK, aniData)
}

func GetVodsList(typeId int, typeId1 int, page int, pageSize int) (aniList []interface{}) {
	vodList := service.Vod.GetVods(typeId, typeId1, page, pageSize)

	for i := 0; i < len(vodList); i++ {
		var authorData = make(map[string]interface{})
		vodInfo := vodList[i]
		authorData["id"] = 50000 + rand.Intn(10000)
		authorData["videoNum"] = rand.Intn(100)
		authorData["icon"] = "http://img.kaiyanapp.com/98beab66d3885a139b54f21e91817c4f.jpeg"
		authorData["name"] = "VING动漫精选"
		authorData["description"] = "提供免费精彩动漫资源"

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
		coverData["feed"] = vodInfo.VodPic     //"http://img.kaiyanapp.com/543451a3991b08eabb1a5f8a2f9f5d39.jpeg?imageMogr2/quality/60/format/jpg"
		coverData["blurred"] = vodInfo.VodPic  //"http://img.kaiyanapp.com/0379d60d05c7abdd12112bcda784d5cd.jpeg?imageMogr2/quality/60/format/jpg"
		coverData["detail"] = vodInfo.VodPic   //"http://img.kaiyanapp.com/543451a3991b08eabb1a5f8a2f9f5d39.jpeg?imageMogr2/quality/60/format/jpg"
		coverData["homepage"] = vodInfo.VodPic //"http://img.kaiyanapp.com/543451a3991b08eabb1a5f8a2f9f5d39.jpeg?imageView2/1/w/720/h/560/format/jpg/q/75|watermark/1/image/aHR0cDovL2ltZy5rYWl5YW5hcHAuY29tL2JsYWNrXzMwLnBuZw==/dissolve/100/gravity/Center/dx/0/dy/0|imageslim"

		var consumption = make(map[string]interface{})
		consumption["collectionCount"] = rand.Intn(50000)
		consumption["shareCount"] = rand.Intn(50000)
		consumption["replyCount"] = rand.Intn(50000)

		var data = make(map[string]interface{})
		data["id"] = vodInfo.VodID
		data["name"] = vodInfo.VodName
		data["title"] = vodInfo.VodSub
		data["aniNum"] = vodInfo.VodSerial
		data["curPlayNum"] = rand.Intn(10)
		data["aniTitle"] = vodInfo.VodSub
		data["dataType"] = fmt.Sprintf("Animation-%d-%d", typeId, typeId1)
		data["desc"] = vodInfo.VodBlurb
		data["author"] = authorData
		data["cover"] = coverData
		data["hits"] = vodInfo.VodHits
		data["category"] = "动漫"
		data["area"] = vodInfo.VodArea
		data["lang"] = vodInfo.VodLang
		data["year"] = vodInfo.VodYear
		data["likeCount"] = vodInfo.VodUp
		data["score"] = float32(3+rand.Intn(7)) + (float32(rand.Intn(10)) / 10)
		// data["playUrl"] = "http://ali.cdn.kaiyanapp.com/1582687191736_7a642380_1280x720.mp4?auth_key=1582881284-0-0-fe8a2d1121cb9ee797af8d67f3622a0a"
		// data["playUrl"] = "http://meigui.qqqq-kuyun.com/20190420/2629_3ee00295/index.m3u8"
		data["playUrl"] = GetPlayUrlInfo(vodInfo.VodPlayURL, vodInfo.VodPlayFrom, vodInfo.VodPlayServer, vodInfo.VodPlayNote)

		data["duration"] = 0
		data["createTime"] = time.Now().Unix() - int64(rand.Intn(1000))
		data["tags"] = tagList
		data["aType"] = fmt.Sprintf("videoType-%d-%d", typeId, typeId1)
		data["date"] = time.Now().Unix() - int64(rand.Intn(1000))
		data["idx"] = vodInfo.VodID
		data["label"] = "动画标签"
		data["promotion"] = "推广"
		data["played"] = false
		data["subtitles"] = vodInfo.VodSub
		data["lastViewTime"] = "最后观看时间"
		data["consumption"] = consumption

		var aniData = make(map[string]interface{})
		aniData["aType"] = typeId
		aniData["tag"] = "动画片"
		aniData["data"] = data
		aniList = append(aniList, aniData)
	}

	return
}

func GetPlayUrlInfo(vodPlayURL string, vodPlayFrom string, vodPlayServer string, vodPlayNote string) (playUrls string) {
	fromList := strings.Split(vodPlayFrom, "$$$")
	urlList := strings.Split(vodPlayURL, "$$$")
	for i := 0; i < len(fromList); i++ {
		if fromList[i] == "zuidam3u8" || fromList[i] == "yjm3u8" {
			logger.Info("url info --------- ", urlList[i])
			playUrls = urlList[i]
			urls := strings.Split(urlList[i], "#")
			for j := 0; j < len(urls); j++ {
				url := urls[j]
				logger.Info("url ------- ", url)
				info := strings.Split(url, "$")
				logger.Info("集数 -- ", info[0], ", 地址--", info[1])
			}
		}
	}
	return
}
