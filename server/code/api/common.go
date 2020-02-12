package api

import (
	// "errors"
	"server/msg"
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
