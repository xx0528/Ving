package api

import (
	// "errors"
	"github.com/gin-gonic/gin"
	// "server/model"
	"github.com/88250/gulu"
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

	collection = a.DB.Collection("Films")

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