// Package util defines variety of utilities.
package model

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/88250/gulu"
	"github.com/jinzhu/gorm"
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

// Version of Movie.
const Version = "1.0.0"

// Conf of Movie.
var Conf *Configuration

// UserAgent represents HTTP client user agent.
var UserAgent = "Pipe/" + Version + "; +https://github.com/b3log/pipe"

// Models represents all models..
var Models = []interface{}{
	&MacCjNode{}, &MacGbook{}, &MacGroup{}, &MacOrder{}, &MacTmpwebsite{},
	&MacVod{}, &MacAdmin{}, &MacCash{}, &MacCollect{}, &MacComment{},
	&MacRole{}, &MacTmpvod{}, &MacTopic{}, &MacUlog{}, &MacUser{},
	&MacActor{}, &MacArt{}, &MacCard{}, &MacCjContent{}, &MacType{},
	&MacWebsite{}, &MacCjHistory{}, &MacLink{}, &MacMsg{}, &MacPlog{},
	&MacVisit{},
}

const tablePrefix = "ving_"

// ZeroPushTime represents zero push time.
var ZeroPushTime, _ = time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")

// Configuration (pipe.json).
type Configuration struct {
	Server                string // server scheme, host and port
	StaticServer          string // static resources server scheme, host and port
	StaticResourceVersion string // version of static resources
	LogLevel              string // logging level: trace/debug/info/warn/error/fatal
	ShowSQL               bool   // whether print sql in log
	SessionSecret         string // HTTP session secret
	SessionMaxAge         int    // HTTP session max age (in second)
	RuntimeMode           string // runtime mode (dev/prod)
	SQLite                string // SQLite database file path
	MySQL                 string // MySQL connection URL
	AxiosBaseURL          string // axio base URL
	MockServer            string // mock server
	Port                  string // listen port
}

// LoadConf loads the configurations. Command-line arguments will override configuration file.
func LoadConf() {
	showCfg := flag.Bool("show_cfg", true, "prints current cfg")
	confPath := flag.String("conf", "ving.json", "path of ving.json")
	confServer := flag.String("server", "", "this will override Conf.Server if specified")
	confStaticServer := flag.String("static_server", "", "this will override Conf.StaticServer if specified")
	confStaticResourceVer := flag.String("static_resource_ver", "", "this will override Conf.StaticResourceVersion if specified")
	confLogLevel := flag.String("log_level", "", "this will override Conf.LogLevel if specified")
	confSessionSecret := flag.String("session_secret", "", "this will override Conf.SessionSecret")
	confSessionMaxAge := flag.Int("session_max_age", 0, "this will override Conf.SessionMaxAge")
	confRuntimeMode := flag.String("runtime_mode", "", "this will override Conf.RuntimeMode if specified")
	confSQLite := flag.String("sqlite", "", "this will override Conf.SQLite if specified")
	confMySQL := flag.String("mysql", "", "this will override Conf.MySQL if specified")
	confPort := flag.String("port", "", "this will override Conf.Port if specified")

	flag.Parse()

	bytes, err := ioutil.ReadFile(*confPath)
	if nil != err {
		logger.Fatal("loads configuration file [" + *confPath + "] failed: " + err.Error())
	}

	Conf = &Configuration{}
	if err = json.Unmarshal(bytes, Conf); nil != err {
		logger.Fatal("parses [ving.json] failed: ", err)
	}

	gulu.Log.SetLevel(Conf.LogLevel)
	if "" != *confLogLevel {
		Conf.LogLevel = *confLogLevel
		gulu.Log.SetLevel(*confLogLevel)
	}

	if "" != *confSessionSecret {
		Conf.SessionSecret = *confSessionSecret
	}

	if 0 < *confSessionMaxAge {
		Conf.SessionMaxAge = *confSessionMaxAge
	}

	if "" == Conf.SessionSecret {
		Conf.SessionSecret = gulu.Rand.String(32)
	}

	home, err := gulu.OS.Home()
	if nil != err {
		logger.Fatal("can't find user home directory: " + err.Error())
	}
	logger.Debugf("${home} [%s]", home)

	if "" != *confRuntimeMode {
		Conf.RuntimeMode = *confRuntimeMode
	}

	if "" != *confServer {
		Conf.Server = *confServer
	}

	if "" != *confStaticServer {
		Conf.StaticServer = *confStaticServer
	}
	if "" == Conf.StaticServer {
		Conf.StaticServer = Conf.Server
	}

	time := strconv.FormatInt(time.Now().UnixNano(), 10)
	logger.Debugf("${time} [%s]", time)
	Conf.StaticResourceVersion = strings.Replace(Conf.StaticResourceVersion, "${time}", time, 1)
	if "" != *confStaticResourceVer {
		Conf.StaticResourceVersion = *confStaticResourceVer
	}

	Conf.SQLite = strings.Replace(Conf.SQLite, "${home}", home, 1)
	if "" != *confSQLite {
		Conf.SQLite = *confSQLite
	}
	if "" != *confMySQL {
		Conf.MySQL = *confMySQL
		Conf.SQLite = ""
	}

	if "" != *confPort {
		Conf.Port = *confPort
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	if *showCfg {
		fmt.Println("版本号                  :   ", Version)
		fmt.Println("日志级别                :   ", Conf.LogLevel)
		fmt.Println("SessionSecret          :   ", Conf.SessionSecret)
		fmt.Println("SessionMaxAge          :   ", Conf.SessionMaxAge)
		fmt.Println("RuntimeMode            :   ", Conf.RuntimeMode)
		fmt.Println("服务器地址               :   ", Conf.Server)
		fmt.Println("端口号                  :   ", Conf.Port)
		fmt.Println("StaticServer           :   ", Conf.StaticServer)
		fmt.Println("StaticResourceVersion  :   ", Conf.StaticResourceVersion)
		fmt.Println("数据库连接               :   ", Conf.MySQL)
	}

	logger.Debugf("configurations [%#v]", Conf)
}
