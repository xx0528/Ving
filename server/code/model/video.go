package model

// The User holds
type Video struct {
	ID         int64  `bson:"_id" json:"id"`                //视频id
	UserID     int64  `bson:"userid" json:"userid"`         //用户id
	UserName   string `bson:"username" json:"username"`     //用户昵称
	Title      string `bson:"title" json:"title"`           //视频标题
	Desc       string `bson:"desc" json:"desc"`             //视频描述
	PlayURL    string `bson:"playurl" json:"playurl"`       //视频地址
	Category   string `bson:"category" json:"category"`     //视频分类
	Tags       []Tag  `bson:"tags" json:"tags"`             //视频分类标签
	Praises    int64  `bson:"praises" json:"praises"`       //👍点赞数量
	Treads     int64  `bson:"treads" json:"treads"`         //👎踩数量
	Comments   int64  `bson:"comments" json:"comments"`     //📄评论数量
	Transmits  int64  `bson:"transmits" json:"transmits"`   //🥠转发数量
	UpdateTime int64  `bson:"updatetime" json:"updatetime"` //上传时间
	Duration   int32  `bson:"duration" json:"duration"`     //视频时长
}

type Tag struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	ActionUrl   string      `json:"actionUrl"`
	AdTrack     interface{} `json:"adTrack"`
	BGPicture   string      `json:"bgPicture"`
	Desc        string      `json:"desc"`
	HeaderImage string      `json:"headerImage"`
}

// New is
func (u *Video) New() *Video {
	return &Video{
		ID:         u.ID,
		UserID:     u.UserID,
		UserName:   u.UserName,
		Desc:       u.Desc,
		PlayURL:    u.PlayURL,
		Category:   u.Category,
		Tags:       u.Tags,
		Praises:    u.Praises,
		Treads:     u.Treads,
		Comments:   u.Comments,
		Transmits:  u.Transmits,
		UpdateTime: u.UpdateTime,
		Duration:   u.Duration,
	}
}
