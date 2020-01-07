package controller

import (
	"server/model"
	"server/util"
	"server/api"
	"server/database"
	// "errors"
	"html/template"
	"net/http"
	// "strconv"
	"strings"

	"github.com/88250/gulu"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Logger
// var logger = gulu.Log.NewLogger(os.Stdout)

// MapRoutes returns a gin engine and binds controllers with request URLs.
func MapRoutes(db *database.MovieDatabase) *gin.Engine {
	ret := gin.New()
	ret.SetFuncMap(template.FuncMap{
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
		"minus":    func(a, b int) int { return a - b },
		"mod":      func(a, b int) int { return a % b },
		"noescape": func(s string) template.HTML { return template.HTML(s) },
	})

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
	// ret.Use(sessions.Sessions("movie", store))

	userHandler := api.UserAPI{DB: db}
	postHandler := api.PostAPI{DB: db}
	
	postApi := ret.Group(util.PathAPI)
	{
		postApi.POST("/logout", logoutAction)
		postApi.GET("/status", getStatusAction)
		postApi.GET("/index", getIndex)
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

func getStatusAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

}

func getIndex(c *gin.Context) {
	result := gulu.Ret.NewResult()
	result.Msg = "Hello web server"
	// c.String(http.StatusOK, result)
	defer c.JSON(http.StatusOK, result)
}
