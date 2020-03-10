package controller

import (
	"os"
	"server/api"
	"server/model"
	"server/util"

	// "errors"
	// "html/template"
	// "net/http"
	// "strconv"
	"strings"

	// "github.com/88250/gulu"
	"github.com/88250/gulu"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

// MapRoutes returns a gin engine and binds controllers with request URLs.
func MapRoutes() *gin.Engine {
	ret := gin.New()
	// ret.SetFuncMap(template.FuncMap{
	// "dict": func(values ...interface{}) (map[string]interface{}, error) {
	// 	if len(values)%2 != 0 {
	// 		return nil, errors.New("len(values) is " + strconv.Itoa(len(values)%2))
	// 	}
	// 	dict := make(map[string]interface{}, len(values)/2)
	// 	for i := 0; i < len(values); i += 2 {
	// 		key, ok := values[i].(string)
	// 		if !ok {
	// 			return nil, errors.New("")
	// 		}
	// 		dict[key] = values[i+1]
	// 	}
	// 	return dict, nil
	// },
	// "minus":    func(a, b int) int { return a - b },
	// "mod":      func(a, b int) int { return a % b },
	// "noescape": func(s string) template.HTML { return template.HTML(s) },
	// })

	if "dev" == model.Conf.RuntimeMode {
		ret.Use(gin.Logger())
	}
	ret.Use(gin.Recovery())
	store := cookie.NewStore([]byte(model.Conf.SessionSecret))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   model.Conf.SessionMaxAge,
		Secure:   strings.HasPrefix(model.Conf.Server, "https"),
		HttpOnly: true,
	})
	ret.Use(sessions.Sessions("ving", store))

	vodGroup := ret.Group("/vod")
	vodGroup.GET("/ani", GetAnimation)
	vodGroup.GET("/animore", GetMoreAnimation)

	commonHandler := api.CommonAPI{}
	userHandler := api.UserAPI{}
	postHandler := api.PostAPI{}

	postC := ret.Group(util.PathAPI)
	{
		//首页精选 /api/v2/home?
		postC.GET("/home", commonHandler.GetHome)
		//相关视频 api/v4/video/related?id=xxx
		postC.GET("/video", commonHandler.GetVideo)
		//获取分类 api/v4/categories
		postC.GET("/categories", commonHandler.GetCategories)
		//获取分类详情List api/v4/categories/videoList?id=xxx&udid=xxx
		postC.GET("/categories/video", commonHandler.GetVideoDetail)
		//获取排行榜的 api/v4/rankList
		postC.GET("/rankList", commonHandler.GetRank)
		//获取搜索信息 api/v1/search?&num=10&start=10&query=xxx
		postC.GET("/search", commonHandler.GetSearch)
		//热门搜索关键词 api/v3/queries/hot
		postC.GET("/queries/hot", commonHandler.GetHot)
		//关注 api/v4/tabs/follow
		postC.GET("/tabs/follow", commonHandler.GetFollow)
		//作者信息 api/v4/pgcs/detail/tab?id=571
		postC.GET("/pgcs/detail", commonHandler.GetAuthor)
		//动画列表 api/animation
		postC.GET("/animation", commonHandler.GetAnimation)
		//更多动画列表 api/animation/more
		postC.GET("/animation/more", commonHandler.GetMoreAnimation)
		//相关动画列表 api/animation/related
		postC.GET("/animation/related", commonHandler.GetRelatedAnimation)
	}

	postU := ret.Group("/users")
	{
		postU.GET("", userHandler.GetUsers)
		postU.DELETE(":id", userHandler.DeleteUserByID)
	}

	postG := ret.Group("/posts")
	{
		postG.GET("", postHandler.GetPosts)
		postG.POST("", postHandler.CreatePost)
		postG.GET(":id", postHandler.GetPostByID)
		postG.PUT(":id", postHandler.UpdatePostByID)
		postG.DELETE(":id", postHandler.DeletePostByID)
	}

	ret.NoRoute(func(c *gin.Context) {
		notFound(c)
	})

	return ret
}

func routePath(c *gin.Context) {
	path := c.Param("path")
	switch path {
	case util.PathAPI:
		return
	}

	logger.Tracef("can't handle path [" + path + "]")
	notFound(c)
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
