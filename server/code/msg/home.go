package msg

import "server/model"

type Home struct {
	IssueList       []Issue `json:"issueList"`
	NextPageUrl     string  `json:"nextPageUrl"`
	NextPublishTime int64   `json:"nextPublishTime"`
	NewestIssueType string  `json:"newestIssueType"`
}

type Issue struct {
	RleaseTime  int64  `json:"releaseTime"`
	Itype       string `json:"type"`
	Date        int64  `json:"date"`
	Total       int    `json:"total"`
	PublishTime int64  `json:"publishTime"`
	ItemList    []Item `json:"itemList"`
	Count       int    `json:"count"`
	NextPageUrl string `json:"nextPageUrl"`
}

type Item struct {
	Itype string `json:"type"`
	IData IData  `json:"data"`
	Tag   string `json:"Tag"`
}

type IData struct {
	DataType          string      `json:"DataType"`
	Text              string      `json:"text"`
	VideoTitle        string      `json:"videoTitle"`
	Id                int64       `json:"id"`
	Title             string      `json:"title"`
	Slogan            string      `json:"slogan"`
	Description       string      `json:"description"`
	ActionUrl         string      `json:"actionUrl"`
	IProvider         Provider    `json:"provider"`
	Category          string      `json:"category"`
	IParentReply      ParentReply `json:"parentReply"`
	IAuthor           Author      `json:"author"`
	ICover            Cover       `json:"cover"`
	LikeCount         int         `json:"likeCount"`
	PlayUrl           string      `json:"playUrl"`
	ThumbPlayUrl      string      `json:"thumbPlayUrl"`
	Duration          int64       `json:"duration"`
	Message           string      `json:"message"`
	CreateTime        int64       `json:"createTime"`
	IWebUrl           Weburl      `json:"webUrl"`
	Library           string      `json:"library"`
	IUser             model.User  `json:"user"`
	PlayInfos         []PlayInfo  `json:"playInfo"`
	IConsumption      Consumption `json:"consumption"`
	Campaign          interface{} `json:"campaign"`
	WaterMarks        interface{} `json:"waterMarks"`
	AdTrack           interface{} `json:"adTrack"`
	Tags              []Tag       `json:"tags"`
	Type              string      `json:"type"`
	TitlePgc          interface{} `json:"titlePgc"`
	DescriptionPgc    interface{} `json:"descriptionPgc"`
	Remark            string      `json:"remark"`
	Idx               int         `json:"idx"`
	ShareAdTrack      interface{} `json:"shareAdTrack"`
	FavoriteAdTrack   interface{} `json:"favoriteAdTrack"`
	WebAdTrack        interface{} `json:"webAdTrack"`
	Date              int64       `json:"date"`
	Promotion         interface{} `json:"promotion"`
	Label             interface{} `json:"label"`
	LabelList         interface{} `json:"labelList"`
	DescriptionEditor string      `json:"descriptionEditor"`
	Collected         bool        `json:"collected"`
	Played            bool        `json:"played"`
	Subtitles         interface{} `json:"subtitles"`
	LastViewTime      interface{} `json:"lastViewTime"`
	Playlists         interface{} `json:"playlists"`
	IHeader           Header      `json:"header"`
	ItemList          []Item      `json:"itemList"`
}

type Tag struct {
	Id        int         `json:"id"`
	Name      string      `json:"name"`
	ActionUrl string      `json:"actionUrl"`
	AdTrack   interface{} `json:"adTrack"`
}

type Provider struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
	Icon  string `json:"icon"`
}

type ParentReply struct {
	User    model.User `json:"user"`
	Message string     `json:"message"`
}

type Author struct {
	Icon        string `json:"icon"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Cover struct {
	Feed     string `json:"feed"`
	Detail   string `json:"detail"`
	Blurred  string `json:"blurred"`
	Sharing  string `json:"sharing"`
	Homepage string `json:"homepage"`
}

type Weburl struct {
	Raw      string `json:"raw"`
	ForWeibo string `json:"forWeibo"`
}

type PlayInfo struct {
	Name    string  `json:"name"`
	Url     string  `json:"url"`
	Ptype   string  `json:"ptype"`
	UrlList []int64 `json:"urlList"`
}

type Consumption struct {
	CollectionCount int `json:"collectionCount"`
	ShareCount      int `json:"shareCount"`
	ReplyCount      int `json:"replyCount"`
}

type Header struct {
	Id          int     `json:"id"`
	Icon        string  `json:"icon"`
	IconType    string  `json:"iconType"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Font        string  `json:"font"`
	Cover       string  `json:"cover"`
	ILabel      Label   `json:"label"`
	ActionUrl   string  `json:"actionUrl"`
	Subtitle    string  `json:"subtitle"`
	LabelList   []Label `json:"labelList"`
}

type Label struct {
	Text      string      `json:"text"`
	Card      string      `json:"card"`
	Detial    interface{} `json:"detial"`
	ActionUrl interface{} `json:"actionUrl"`
}
