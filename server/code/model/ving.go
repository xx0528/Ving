package model

// MacCjNode [...]
type MacCjNode struct {
	Nodeid           int16  `gorm:"primary_key;AUTO_INCREMENT;column:nodeid;type:smallint(6) unsigned;not null" json:"nodeid"` //
	Name             string `gorm:"column:name;type:varchar(20);not null" json:"name"`                                         //
	Lastdate         int    `gorm:"column:lastdate;type:int(10) unsigned;not null" json:"lastdate"`                            //
	Sourcecharset    string `gorm:"column:sourcecharset;type:varchar(8);not null" json:"sourcecharset"`                        //
	Sourcetype       int8   `gorm:"column:sourcetype;type:tinyint(1) unsigned;not null" json:"sourcetype"`                     //
	URLpage          string `gorm:"column:urlpage;type:text;not null" json:"urlpage"`                                          //
	PagesizeStart    int8   `gorm:"column:pagesize_start;type:tinyint(3) unsigned;not null" json:"pagesize_start"`             //
	PagesizeEnd      string `gorm:"column:pagesize_end;type:mediumint(8) unsigned;not null" json:"pagesize_end"`               //
	PageBase         string `gorm:"column:page_base;type:char(255);not null" json:"page_base"`                                 //
	ParNum           int8   `gorm:"column:par_num;type:tinyint(3) unsigned;not null" json:"par_num"`                           //
	URLContain       string `gorm:"column:url_contain;type:char(100);not null" json:"url_contain"`                             //
	URLExcept        string `gorm:"column:url_except;type:char(100);not null" json:"url_except"`                               //
	URLStart         string `gorm:"column:url_start;type:char(100);not null" json:"url_start"`                                 //
	URLEnd           string `gorm:"column:url_end;type:char(100);not null" json:"url_end"`                                     //
	TitleRule        string `gorm:"column:title_rule;type:char(100);not null" json:"title_rule"`                               //
	TitleHTMLRule    string `gorm:"column:title_html_rule;type:text;not null" json:"title_html_rule"`                          //
	TypeRule         string `gorm:"column:type_rule;type:char(100);not null" json:"type_rule"`                                 //
	TypeHTMLRule     string `gorm:"column:type_html_rule;type:text;not null" json:"type_html_rule"`                            //
	ContentRule      string `gorm:"column:content_rule;type:char(100);not null" json:"content_rule"`                           //
	ContentHTMLRule  string `gorm:"column:content_html_rule;type:text;not null" json:"content_html_rule"`                      //
	ContentPageStart string `gorm:"column:content_page_start;type:char(100);not null" json:"content_page_start"`               //
	ContentPageEnd   string `gorm:"column:content_page_end;type:char(100);not null" json:"content_page_end"`                   //
	ContentPageRule  int8   `gorm:"column:content_page_rule;type:tinyint(1) unsigned;not null" json:"content_page_rule"`       //
	ContentPage      int8   `gorm:"column:content_page;type:tinyint(1) unsigned;not null" json:"content_page"`                 //
	ContentNextpage  string `gorm:"column:content_nextpage;type:char(100);not null" json:"content_nextpage"`                   //
	DownAttachment   int8   `gorm:"column:down_attachment;type:tinyint(1) unsigned;not null" json:"down_attachment"`           //
	Watermark        int8   `gorm:"column:watermark;type:tinyint(1) unsigned;not null" json:"watermark"`                       //
	CollOrder        int8   `gorm:"column:coll_order;type:tinyint(3) unsigned;not null" json:"coll_order"`                     //
	CustomizeConfig  string `gorm:"column:customize_config;type:text;not null" json:"customize_config"`                        //
	ProgramConfig    string `gorm:"column:program_config;type:text;not null" json:"program_config"`                            //
	Mid              int8   `gorm:"column:mid;type:tinyint(1) unsigned;not null" json:"mid"`                                   //
}

// TableName get sql table name.获取数据库表名
func (m *MacCjNode) TableName() string {
	return "mac_cj_node"
}

// MacGbook [...]
type MacGbook struct {
	GbookID        int    `gorm:"primary_key;AUTO_INCREMENT;column:gbook_id;type:int(10) unsigned;not null" json:"gbook_id"`             //编号
	GbookRid       int    `gorm:"index:gbook_rid;column:gbook_rid;type:int(10) unsigned;not null" json:"gbook_rid"`                      //
	UserID         int    `gorm:"index:user_id;column:user_id;type:int(10) unsigned;not null" json:"user_id"`                            //
	GbookStatus    int8   `gorm:"column:gbook_status;type:tinyint(1) unsigned;not null" json:"gbook_status"`                             //状态0未审核1已审核
	GbookName      string `gorm:"column:gbook_name;type:varchar(60);not null" json:"gbook_name"`                                         //昵称
	GbookIP        int    `gorm:"column:gbook_ip;type:int(10) unsigned;not null" json:"gbook_ip"`                                        //ip地址
	GbookTime      int    `gorm:"index:gbook_time;column:gbook_time;type:int(10) unsigned;not null" json:"gbook_time"`                   //时间
	GbookReplyTime int    `gorm:"index:gbook_reply_time;column:gbook_reply_time;type:int(10) unsigned;not null" json:"gbook_reply_time"` //回复时间
	GbookContent   string `gorm:"column:gbook_content;type:varchar(255);not null" json:"gbook_content"`                                  //留言内容
	GbookReply     string `gorm:"index:gbook_reply;column:gbook_reply;type:varchar(255);not null" json:"gbook_reply"`                    //回复内容
}

// TableName get sql table name.获取数据库表名
func (m *MacGbook) TableName() string {
	return "mac_gbook"
}

// MacGroup [...]
type MacGroup struct {
	GroupID          int16  `gorm:"primary_key;AUTO_INCREMENT;column:group_id;type:smallint(6);not null" json:"group_id"`         //
	GroupName        string `gorm:"column:group_name;type:varchar(30);not null" json:"group_name"`                                //
	GroupStatus      int8   `gorm:"index:group_status;column:group_status;type:tinyint(1) unsigned;not null" json:"group_status"` //
	GroupType        string `gorm:"column:group_type;type:text;not null" json:"group_type"`                                       //
	GroupPopedom     string `gorm:"column:group_popedom;type:text;not null" json:"group_popedom"`                                 //
	GroupPointsDay   int16  `gorm:"column:group_points_day;type:smallint(6) unsigned;not null" json:"group_points_day"`           //
	GroupPointsWeek  int16  `gorm:"column:group_points_week;type:smallint(6);not null" json:"group_points_week"`                  //
	GroupPointsMonth int16  `gorm:"column:group_points_month;type:smallint(6) unsigned;not null" json:"group_points_month"`       //
	GroupPointsYear  int16  `gorm:"column:group_points_year;type:smallint(6) unsigned;not null" json:"group_points_year"`         //
	GroupPointsFree  int8   `gorm:"column:group_points_free;type:tinyint(1) unsigned;not null" json:"group_points_free"`          //
}

// TableName get sql table name.获取数据库表名
func (m *MacGroup) TableName() string {
	return "mac_group"
}

// MacOrder [...]
type MacOrder struct {
	OrderID      int     `gorm:"primary_key;AUTO_INCREMENT;column:order_id;type:int(10) unsigned;not null" json:"order_id"` //
	UserID       int     `gorm:"index:user_id;column:user_id;type:int(10) unsigned;not null" json:"user_id"`                //
	OrderStatus  int8    `gorm:"column:order_status;type:tinyint(1) unsigned;not null" json:"order_status"`                 //
	OrderCode    string  `gorm:"index:order_code;column:order_code;type:varchar(30);not null" json:"order_code"`            //
	OrderPrice   float64 `gorm:"column:order_price;type:decimal(12,2) unsigned;not null" json:"order_price"`                //
	OrderTime    int     `gorm:"index:order_time;column:order_time;type:int(10) unsigned;not null" json:"order_time"`       //
	OrderPoints  string  `gorm:"column:order_points;type:mediumint(8) unsigned;not null" json:"order_points"`               //
	OrderPayType string  `gorm:"column:order_pay_type;type:varchar(10);not null" json:"order_pay_type"`                     //
	OrderPayTime int     `gorm:"column:order_pay_time;type:int(10) unsigned;not null" json:"order_pay_time"`                //
	OrderRemarks string  `gorm:"column:order_remarks;type:varchar(100);not null" json:"order_remarks"`                      //
}

// TableName get sql table name.获取数据库表名
func (m *MacOrder) TableName() string {
	return "mac_order"
}

// MacTmpwebsite [...]
type MacTmpwebsite struct {
	ID1   int    `gorm:"column:id1;type:int(10) unsigned" json:"id1"`         //
	Name1 string `gorm:"column:name1;type:varchar(60);not null" json:"name1"` //
}

// TableName get sql table name.获取数据库表名
func (m *MacTmpwebsite) TableName() string {
	return "mac_tmpwebsite"
}

// MacVod [...]
type MacVod struct {
	VodID          int     `gorm:"primary_key;AUTO_INCREMENT;column:vod_id;type:int(10) unsigned;not null" json:"vod_id"`                  //视频id
	TypeID         int16   `gorm:"index:type_id;column:type_id;type:smallint(6);not null" json:"type_id"`                                  //分类id
	TypeID1        int16   `gorm:"index:type_id_1;column:type_id_1;type:smallint(6) unsigned;not null" json:"type_id_1"`                   //一级分类id
	GroupID        int16   `gorm:"index:group_id;column:group_id;type:smallint(6) unsigned;not null" json:"group_id"`                      //用户组id
	VodName        string  `gorm:"index:vod_name;column:vod_name;type:varchar(255);not null" json:"vod_name"`                              //视频名
	VodSub         string  `gorm:"column:vod_sub;type:varchar(255);not null" json:"vod_sub"`                                               //副标题
	VodEn          string  `gorm:"index:vod_en;column:vod_en;type:varchar(255);not null" json:"vod_en"`                                    //别名
	VodStatus      int8    `gorm:"column:vod_status;type:tinyint(1) unsigned;not null" json:"vod_status"`                                  //状态0未审1已审
	VodLetter      string  `gorm:"index:vod_letter;column:vod_letter;type:char(1);not null" json:"vod_letter"`                             //首字母
	VodColor       string  `gorm:"column:vod_color;type:varchar(6);not null" json:"vod_color"`                                             //颜色
	VodTag         string  `gorm:"index:vod_tag;column:vod_tag;type:varchar(100);not null" json:"vod_tag"`                                 //tags
	VodClass       string  `gorm:"index:vod_class;column:vod_class;type:varchar(255);not null" json:"vod_class"`                           //扩展分类
	VodPic         string  `gorm:"column:vod_pic;type:varchar(255);not null" json:"vod_pic"`                                               //图片
	VodPicThumb    string  `gorm:"column:vod_pic_thumb;type:varchar(255);not null" json:"vod_pic_thumb"`                                   //缩略图
	VodPicSlide    string  `gorm:"column:vod_pic_slide;type:varchar(255);not null" json:"vod_pic_slide"`                                   //幻灯图
	VodActor       string  `gorm:"index:vod_actor;column:vod_actor;type:varchar(255);not null" json:"vod_actor"`                           //主演
	VodDirector    string  `gorm:"index:vod_director;column:vod_director;type:varchar(255);not null" json:"vod_director"`                  //导演
	VodWriter      string  `gorm:"column:vod_writer;type:varchar(100);not null" json:"vod_writer"`                                         //编剧
	VodBehind      string  `gorm:"column:vod_behind;type:varchar(100);not null" json:"vod_behind"`                                         //幕后
	VodBlurb       string  `gorm:"column:vod_blurb;type:varchar(255);not null" json:"vod_blurb"`                                           //简介
	VodRemarks     string  `gorm:"column:vod_remarks;type:varchar(100);not null" json:"vod_remarks"`                                       //备注
	VodPubdate     string  `gorm:"column:vod_pubdate;type:varchar(100);not null" json:"vod_pubdate"`                                       //上映日期
	VodTotal       string  `gorm:"index:vod_total;column:vod_total;type:mediumint(8) unsigned;not null" json:"vod_total"`                  //总集数
	VodSerial      string  `gorm:"column:vod_serial;type:varchar(20);not null" json:"vod_serial"`                                          //连载数
	VodTv          string  `gorm:"column:vod_tv;type:varchar(30);not null" json:"vod_tv"`                                                  //上映电视台
	VodWeekday     string  `gorm:"column:vod_weekday;type:varchar(30);not null" json:"vod_weekday"`                                        //节目周期
	VodArea        string  `gorm:"index:vod_area;column:vod_area;type:varchar(20);not null" json:"vod_area"`                               //地区
	VodLang        string  `gorm:"index:vod_lang;column:vod_lang;type:varchar(10);not null" json:"vod_lang"`                               //语言
	VodYear        string  `gorm:"index:vod_year;column:vod_year;type:varchar(10);not null" json:"vod_year"`                               //年代
	VodVersion     string  `gorm:"index:vod_version;column:vod_version;type:varchar(30);not null" json:"vod_version"`                      //版本-dvd,hd,720p
	VodState       string  `gorm:"index:vod_state;column:vod_state;type:varchar(30);not null" json:"vod_state"`                            //资源类别-正片,预告片,花絮
	VodAuthor      string  `gorm:"column:vod_author;type:varchar(60);not null" json:"vod_author"`                                          //编辑人员
	VodJumpurl     string  `gorm:"column:vod_jumpurl;type:varchar(150);not null" json:"vod_jumpurl"`                                       //跳转url
	VodTpl         string  `gorm:"column:vod_tpl;type:varchar(30);not null" json:"vod_tpl"`                                                //独立模板
	VodTplPlay     string  `gorm:"column:vod_tpl_play;type:varchar(30);not null" json:"vod_tpl_play"`                                      //独立播放页模板
	VodTplDown     string  `gorm:"column:vod_tpl_down;type:varchar(30);not null" json:"vod_tpl_down"`                                      //独立下载页模板
	VodIsend       int8    `gorm:"index:vod_isend;column:vod_isend;type:tinyint(1) unsigned;not null" json:"vod_isend"`                    //是否完结
	VodLock        int8    `gorm:"index:vod_lock;column:vod_lock;type:tinyint(1) unsigned;not null" json:"vod_lock"`                       //锁定1
	VodLevel       int8    `gorm:"index:vod_level;column:vod_level;type:tinyint(1) unsigned;not null" json:"vod_level"`                    //推荐级别
	VodCopyright   int8    `gorm:"column:vod_copyright;type:tinyint(1) unsigned;not null" json:"vod_copyright"`                            //访问整个视频所需积分
	VodPoints      int16   `gorm:"column:vod_points;type:smallint(6) unsigned;not null" json:"vod_points"`                                 //是否开启版权提示
	VodPointsPlay  int16   `gorm:"index:vod_points_play;column:vod_points_play;type:smallint(6) unsigned;not null" json:"vod_points_play"` //每集点播付费
	VodPointsDown  int16   `gorm:"index:vod_points_down;column:vod_points_down;type:smallint(6) unsigned;not null" json:"vod_points_down"` //每集下载付费
	VodHits        string  `gorm:"index:vod_hits;column:vod_hits;type:mediumint(8) unsigned;not null" json:"vod_hits"`                     //总点击量
	VodHitsDay     string  `gorm:"index:vod_hits_day;column:vod_hits_day;type:mediumint(8) unsigned;not null" json:"vod_hits_day"`         //日点击量
	VodHitsWeek    string  `gorm:"index:vod_hits_week;column:vod_hits_week;type:mediumint(8) unsigned;not null" json:"vod_hits_week"`      //周点击量
	VodHitsMonth   string  `gorm:"index:vod_hits_month;column:vod_hits_month;type:mediumint(8) unsigned;not null" json:"vod_hits_month"`   //月点击量
	VodDuration    string  `gorm:"column:vod_duration;type:varchar(10);not null" json:"vod_duration"`                                      //时长
	VodUp          string  `gorm:"index:vod_up;column:vod_up;type:mediumint(8) unsigned;not null" json:"vod_up"`                           //顶数
	VodDown        string  `gorm:"index:vod_down;column:vod_down;type:mediumint(8) unsigned;not null" json:"vod_down"`                     //踩数
	VodScore       float64 `gorm:"index:vod_score;column:vod_score;type:decimal(3,1) unsigned;not null" json:"vod_score"`                  //平均分
	VodScoreAll    string  `gorm:"index:vod_score_all;column:vod_score_all;type:mediumint(8) unsigned;not null" json:"vod_score_all"`      //总评分
	VodScoreNum    string  `gorm:"index:vod_score_num;column:vod_score_num;type:mediumint(8) unsigned;not null" json:"vod_score_num"`      //评分次数
	VodTime        int     `gorm:"index:vod_time;column:vod_time;type:int(10) unsigned;not null" json:"vod_time"`                          //更新时间
	VodTimeAdd     int     `gorm:"index:vod_time_add;column:vod_time_add;type:int(10) unsigned;not null" json:"vod_time_add"`              //添加时间
	VodTimeHits    int     `gorm:"column:vod_time_hits;type:int(10) unsigned;not null" json:"vod_time_hits"`                               //点击时间
	VodTimeMake    int     `gorm:"index:vod_time_make;column:vod_time_make;type:int(10) unsigned;not null" json:"vod_time_make"`           //生成时间
	VodTrysee      int16   `gorm:"column:vod_trysee;type:smallint(6) unsigned;not null" json:"vod_trysee"`                                 //试看时长分
	VodDoubanID    int     `gorm:"column:vod_douban_id;type:int(10) unsigned;not null" json:"vod_douban_id"`                               //豆瓣id
	VodDoubanScore float64 `gorm:"column:vod_douban_score;type:decimal(3,1) unsigned;not null" json:"vod_douban_score"`                    //豆瓣评分
	VodReurl       string  `gorm:"column:vod_reurl;type:varchar(255);not null" json:"vod_reurl"`                                           //来源地址
	VodRelVod      string  `gorm:"column:vod_rel_vod;type:varchar(255);not null" json:"vod_rel_vod"`                                       //关联视频ids
	VodRelArt      string  `gorm:"column:vod_rel_art;type:varchar(255);not null" json:"vod_rel_art"`                                       //关联文章ids
	VodPwd         string  `gorm:"column:vod_pwd;type:varchar(10);not null" json:"vod_pwd"`                                                //访问内容页密码
	VodPwdURL      string  `gorm:"column:vod_pwd_url;type:varchar(255);not null" json:"vod_pwd_url"`                                       //获取密码链接
	VodPwdPlay     string  `gorm:"column:vod_pwd_play;type:varchar(10);not null" json:"vod_pwd_play"`                                      //访问播放页密码
	VodPwdPlayURL  string  `gorm:"column:vod_pwd_play_url;type:varchar(255);not null" json:"vod_pwd_play_url"`                             //获取密码链接
	VodPwdDown     string  `gorm:"column:vod_pwd_down;type:varchar(10);not null" json:"vod_pwd_down"`                                      //访问下载页密码
	VodPwdDownURL  string  `gorm:"column:vod_pwd_down_url;type:varchar(255);not null" json:"vod_pwd_down_url"`                             //获取密码链接
	VodContent     string  `gorm:"column:vod_content;type:text;not null" json:"vod_content"`                                               //详细介绍
	VodPlayFrom    string  `gorm:"column:vod_play_from;type:varchar(255);not null" json:"vod_play_from"`                                   //播放组
	VodPlayServer  string  `gorm:"column:vod_play_server;type:varchar(255);not null" json:"vod_play_server"`                               //播放服务器组
	VodPlayNote    string  `gorm:"column:vod_play_note;type:varchar(255);not null" json:"vod_play_note"`                                   //播放备注
	VodPlayURL     string  `gorm:"column:vod_play_url;type:mediumtext;not null" json:"vod_play_url"`                                       //播放地址
	VodDownFrom    string  `gorm:"column:vod_down_from;type:varchar(255);not null" json:"vod_down_from"`                                   //下载租
	VodDownServer  string  `gorm:"column:vod_down_server;type:varchar(255);not null" json:"vod_down_server"`                               //下载服务器组
	VodDownNote    string  `gorm:"column:vod_down_note;type:varchar(255);not null" json:"vod_down_note"`                                   //下载备注
	VodDownURL     string  `gorm:"column:vod_down_url;type:mediumtext;not null" json:"vod_down_url"`                                       //下载地址
	VodPlot        int8    `gorm:"index:vod_plot;column:vod_plot;type:tinyint(1) unsigned;not null" json:"vod_plot"`                       //是否包含分集剧情
	VodPlotName    string  `gorm:"column:vod_plot_name;type:mediumtext;not null" json:"vod_plot_name"`                                     //分集剧情名称
	VodPlotDetail  string  `gorm:"column:vod_plot_detail;type:mediumtext;not null" json:"vod_plot_detail"`                                 //分集剧情详情
}

// TableName get sql table name.获取数据库表名
func (m *MacVod) TableName() string {
	return "mac_vod"
}

// MacAdmin [...]
type MacAdmin struct {
	AdminID            int16  `gorm:"primary_key;AUTO_INCREMENT;column:admin_id;type:smallint(6) unsigned;not null" json:"admin_id"` //
	AdminName          string `gorm:"index:admin_name;column:admin_name;type:varchar(30);not null" json:"admin_name"`                //
	AdminPwd           string `gorm:"column:admin_pwd;type:char(32);not null" json:"admin_pwd"`                                      //
	AdminRandom        string `gorm:"column:admin_random;type:char(32);not null" json:"admin_random"`                                //
	AdminStatus        int8   `gorm:"column:admin_status;type:tinyint(1) unsigned;not null" json:"admin_status"`                     //
	AdminAuth          string `gorm:"column:admin_auth;type:text;not null" json:"admin_auth"`                                        //
	AdminLoginTime     int    `gorm:"column:admin_login_time;type:int(10) unsigned;not null" json:"admin_login_time"`                //
	AdminLoginIP       int    `gorm:"column:admin_login_ip;type:int(10) unsigned;not null" json:"admin_login_ip"`                    //
	AdminLoginNum      int    `gorm:"column:admin_login_num;type:int(10) unsigned;not null" json:"admin_login_num"`                  //
	AdminLastLoginTime int    `gorm:"column:admin_last_login_time;type:int(10) unsigned;not null" json:"admin_last_login_time"`      //
	AdminLastLoginIP   int    `gorm:"column:admin_last_login_ip;type:int(10) unsigned;not null" json:"admin_last_login_ip"`          //
}

// TableName get sql table name.获取数据库表名
func (m *MacAdmin) TableName() string {
	return "mac_admin"
}

// MacCash [...]
type MacCash struct {
	CashID        int     `gorm:"primary_key;AUTO_INCREMENT;column:cash_id;type:int(10) unsigned;not null" json:"cash_id"`   //
	UserID        int     `gorm:"index:user_id;column:user_id;type:int(10) unsigned;not null" json:"user_id"`                //
	CashStatus    int8    `gorm:"index:cash_status;column:cash_status;type:tinyint(1) unsigned;not null" json:"cash_status"` //状态
	CashPoints    int16   `gorm:"column:cash_points;type:smallint(6) unsigned;not null" json:"cash_points"`                  //积分
	CashMoney     float64 `gorm:"column:cash_money;type:decimal(12,2) unsigned;not null" json:"cash_money"`                  //金额
	CashBankName  string  `gorm:"column:cash_bank_name;type:varchar(60);not null" json:"cash_bank_name"`                     //银行名称
	CashBankNo    string  `gorm:"column:cash_bank_no;type:varchar(30);not null" json:"cash_bank_no"`                         //银行账号
	CashPayeeName string  `gorm:"column:cash_payee_name;type:varchar(30);not null" json:"cash_payee_name"`                   //收款人姓名
	CashTime      int     `gorm:"column:cash_time;type:int(10) unsigned;not null" json:"cash_time"`                          //
	CashTimeAudit int     `gorm:"column:cash_time_audit;type:int(10) unsigned;not null" json:"cash_time_audit"`              //
}

// TableName get sql table name.获取数据库表名
func (m *MacCash) TableName() string {
	return "mac_cash"
}

// MacCollect [...]
type MacCollect struct {
	CollectID         int    `gorm:"primary_key;AUTO_INCREMENT;column:collect_id;type:int(10) unsigned;not null" json:"collect_id"` //
	CollectName       string `gorm:"column:collect_name;type:varchar(30);not null" json:"collect_name"`                             //
	CollectURL        string `gorm:"column:collect_url;type:varchar(255);not null" json:"collect_url"`                              //
	CollectType       int8   `gorm:"column:collect_type;type:tinyint(1) unsigned;not null" json:"collect_type"`                     //
	CollectMid        int8   `gorm:"column:collect_mid;type:tinyint(1) unsigned;not null" json:"collect_mid"`                       //
	CollectAppid      string `gorm:"column:collect_appid;type:varchar(30);not null" json:"collect_appid"`                           //
	CollectAppkey     string `gorm:"column:collect_appkey;type:varchar(30);not null" json:"collect_appkey"`                         //
	CollectParam      string `gorm:"column:collect_param;type:varchar(100);not null" json:"collect_param"`                          //
	CollectFilter     int8   `gorm:"column:collect_filter;type:tinyint(1) unsigned;not null" json:"collect_filter"`                 //
	CollectFilterFrom string `gorm:"column:collect_filter_from;type:varchar(255);not null" json:"collect_filter_from"`              //
	CollectOpt        int8   `gorm:"column:collect_opt;type:tinyint(1) unsigned;not null" json:"collect_opt"`                       //
}

// TableName get sql table name.获取数据库表名
func (m *MacCollect) TableName() string {
	return "mac_collect"
}

// MacComment [...]
type MacComment struct {
	CommentID      int    `gorm:"primary_key;AUTO_INCREMENT;column:comment_id;type:int(10) unsigned;not null" json:"comment_id"`     //编号
	CommentMid     int8   `gorm:"index:comment_mid;column:comment_mid;type:tinyint(1) unsigned;not null" json:"comment_mid"`         //模块id，1视频2文字3专题
	CommentRid     int    `gorm:"index:comment_rid;column:comment_rid;type:int(10) unsigned;not null" json:"comment_rid"`            //
	CommentPid     int    `gorm:"index:comment_pid;column:comment_pid;type:int(10) unsigned;not null" json:"comment_pid"`            //
	UserID         int    `gorm:"index:user_id;column:user_id;type:int(10) unsigned;not null" json:"user_id"`                        //
	CommentStatus  int8   `gorm:"column:comment_status;type:tinyint(1) unsigned;not null" json:"comment_status"`                     //状态0未审核1已审核
	CommentName    string `gorm:"column:comment_name;type:varchar(60);not null" json:"comment_name"`                                 //昵称
	CommentIP      int    `gorm:"column:comment_ip;type:int(10) unsigned;not null" json:"comment_ip"`                                //ip地址
	CommentTime    int    `gorm:"index:comment_time;column:comment_time;type:int(10) unsigned;not null" json:"comment_time"`         //时间
	CommentContent string `gorm:"column:comment_content;type:varchar(255);not null" json:"comment_content"`                          //留言内容
	CommentUp      string `gorm:"column:comment_up;type:mediumint(8) unsigned;not null" json:"comment_up"`                           //顶数
	CommentDown    string `gorm:"column:comment_down;type:mediumint(8) unsigned;not null" json:"comment_down"`                       //踩数
	CommentReply   string `gorm:"index:comment_reply;column:comment_reply;type:mediumint(8) unsigned;not null" json:"comment_reply"` //回复
	CommentReport  string `gorm:"column:comment_report;type:mediumint(8) unsigned;not null" json:"comment_report"`                   //举报
}

// TableName get sql table name.获取数据库表名
func (m *MacComment) TableName() string {
	return "mac_comment"
}

// MacRole [...]
type MacRole struct {
	RoleID        int     `gorm:"primary_key;AUTO_INCREMENT;column:role_id;type:int(10) unsigned;not null" json:"role_id"`              //角色id
	RoleRid       int     `gorm:"index:role_rid;column:role_rid;type:int(10) unsigned;not null" json:"role_rid"`                        //关联视频id
	RoleName      string  `gorm:"index:role_name;column:role_name;type:varchar(255);not null" json:"role_name"`                         //角色名
	RoleEn        string  `gorm:"index:role_en;column:role_en;type:varchar(255);not null" json:"role_en"`                               //拼音
	RoleStatus    int8    `gorm:"column:role_status;type:tinyint(1) unsigned;not null" json:"role_status"`                              //状态
	RoleLock      int8    `gorm:"column:role_lock;type:tinyint(1) unsigned;not null" json:"role_lock"`                                  //锁定
	RoleLetter    string  `gorm:"index:role_letter;column:role_letter;type:char(1);not null" json:"role_letter"`                        //首字母
	RoleColor     string  `gorm:"column:role_color;type:varchar(6);not null" json:"role_color"`                                         //高亮颜色
	RoleActor     string  `gorm:"index:role_actor;column:role_actor;type:varchar(255);not null" json:"role_actor"`                      //演员名称
	RoleRemarks   string  `gorm:"column:role_remarks;type:varchar(100);not null" json:"role_remarks"`                                   //备注
	RolePic       string  `gorm:"column:role_pic;type:varchar(255);not null" json:"role_pic"`                                           //图片
	RoleSort      int16   `gorm:"column:role_sort;type:smallint(6) unsigned;not null" json:"role_sort"`                                 //排序
	RoleLevel     int8    `gorm:"index:role_level;column:role_level;type:tinyint(1) unsigned;not null" json:"role_level"`               //推荐值
	RoleTime      int     `gorm:"index:role_time;column:role_time;type:int(10) unsigned;not null" json:"role_time"`                     //更新时间
	RoleTimeAdd   int     `gorm:"index:role_time_add;column:role_time_add;type:int(10) unsigned;not null" json:"role_time_add"`         //添加时间
	RoleTimeHits  int     `gorm:"column:role_time_hits;type:int(10) unsigned;not null" json:"role_time_hits"`                           //点击时间
	RoleTimeMake  int     `gorm:"column:role_time_make;type:int(10) unsigned;not null" json:"role_time_make"`                           //生成时间
	RoleHits      string  `gorm:"column:role_hits;type:mediumint(8) unsigned;not null" json:"role_hits"`                                //总点击数
	RoleHitsDay   string  `gorm:"column:role_hits_day;type:mediumint(8) unsigned;not null" json:"role_hits_day"`                        //日点击数
	RoleHitsWeek  string  `gorm:"column:role_hits_week;type:mediumint(8) unsigned;not null" json:"role_hits_week"`                      //周点击数
	RoleHitsMonth string  `gorm:"column:role_hits_month;type:mediumint(8) unsigned;not null" json:"role_hits_month"`                    //月点击数
	RoleScore     float64 `gorm:"index:role_score;column:role_score;type:decimal(3,1) unsigned;not null" json:"role_score"`             //平均分
	RoleScoreAll  string  `gorm:"index:role_score_all;column:role_score_all;type:mediumint(8) unsigned;not null" json:"role_score_all"` //总评分
	RoleScoreNum  string  `gorm:"index:role_score_num;column:role_score_num;type:mediumint(8) unsigned;not null" json:"role_score_num"` //评分次数
	RoleUp        string  `gorm:"index:role_up;column:role_up;type:mediumint(8) unsigned;not null" json:"role_up"`                      //顶数
	RoleDown      string  `gorm:"index:role_down;column:role_down;type:mediumint(8) unsigned;not null" json:"role_down"`                //踩数
	RoleTpl       string  `gorm:"column:role_tpl;type:varchar(30);not null" json:"role_tpl"`                                            //自定义模板
	RoleJumpurl   string  `gorm:"column:role_jumpurl;type:varchar(150);not null" json:"role_jumpurl"`                                   //跳转url
	RoleContent   string  `gorm:"column:role_content;type:text;not null" json:"role_content"`                                           //详情
}

// TableName get sql table name.获取数据库表名
func (m *MacRole) TableName() string {
	return "mac_role"
}

// MacTmpvod [...]
type MacTmpvod struct {
	ID1   int    `gorm:"column:id1;type:int(10) unsigned" json:"id1"`          //
	Name1 string `gorm:"column:name1;type:varchar(255);not null" json:"name1"` //
}

// TableName get sql table name.获取数据库表名
func (m *MacTmpvod) TableName() string {
	return "mac_tmpvod"
}

// MacTopic [...]
type MacTopic struct {
	TopicID        int16   `gorm:"primary_key;AUTO_INCREMENT;column:topic_id;type:smallint(6) unsigned;not null" json:"topic_id"`              //专题id
	TopicName      string  `gorm:"index:topic_name;column:topic_name;type:varchar(255);not null" json:"topic_name"`                            //名称
	TopicEn        string  `gorm:"index:topic_en;column:topic_en;type:varchar(255);not null" json:"topic_en"`                                  //别名
	TopicSub       string  `gorm:"column:topic_sub;type:varchar(255);not null" json:"topic_sub"`                                               //副标
	TopicStatus    int8    `gorm:"column:topic_status;type:tinyint(1) unsigned;not null" json:"topic_status"`                                  //状态
	TopicSort      int16   `gorm:"index:topic_sort;column:topic_sort;type:smallint(6) unsigned;not null" json:"topic_sort"`                    //排序号
	TopicLetter    string  `gorm:"column:topic_letter;type:char(1);not null" json:"topic_letter"`                                              //首字母
	TopicColor     string  `gorm:"column:topic_color;type:varchar(6);not null" json:"topic_color"`                                             //高亮颜色
	TopicTpl       string  `gorm:"column:topic_tpl;type:varchar(30);not null" json:"topic_tpl"`                                                //模板文件
	TopicType      string  `gorm:"column:topic_type;type:varchar(255);not null" json:"topic_type"`                                             //扩展分类
	TopicPic       string  `gorm:"column:topic_pic;type:varchar(255);not null" json:"topic_pic"`                                               //图片
	TopicPicThumb  string  `gorm:"column:topic_pic_thumb;type:varchar(255);not null" json:"topic_pic_thumb"`                                   //缩略图
	TopicPicSlide  string  `gorm:"column:topic_pic_slide;type:varchar(255);not null" json:"topic_pic_slide"`                                   //幻灯图
	TopicKey       string  `gorm:"column:topic_key;type:varchar(255);not null" json:"topic_key"`                                               //seo关键字
	TopicDes       string  `gorm:"column:topic_des;type:varchar(255);not null" json:"topic_des"`                                               //seo描述
	TopicTitle     string  `gorm:"column:topic_title;type:varchar(255);not null" json:"topic_title"`                                           //seo标题
	TopicBlurb     string  `gorm:"column:topic_blurb;type:varchar(255);not null" json:"topic_blurb"`                                           //简介
	TopicRemarks   string  `gorm:"column:topic_remarks;type:varchar(100);not null" json:"topic_remarks"`                                       //备注
	TopicLevel     int8    `gorm:"index:topic_level;column:topic_level;type:tinyint(1) unsigned;not null" json:"topic_level"`                  //推荐值
	TopicUp        string  `gorm:"index:topic_up;column:topic_up;type:mediumint(8) unsigned;not null" json:"topic_up"`                         //顶数
	TopicDown      string  `gorm:"index:topic_down;column:topic_down;type:mediumint(8) unsigned;not null" json:"topic_down"`                   //踩数
	TopicScore     float64 `gorm:"index:topic_score;column:topic_score;type:decimal(3,1) unsigned;not null" json:"topic_score"`                //平均分
	TopicScoreAll  string  `gorm:"index:topic_score_all;column:topic_score_all;type:mediumint(8) unsigned;not null" json:"topic_score_all"`    //总评分
	TopicScoreNum  string  `gorm:"index:topic_score_num;column:topic_score_num;type:mediumint(8) unsigned;not null" json:"topic_score_num"`    //总评次
	TopicHits      string  `gorm:"index:topic_hits;column:topic_hits;type:mediumint(8) unsigned;not null" json:"topic_hits"`                   //总点击
	TopicHitsDay   string  `gorm:"index:topic_hits_day;column:topic_hits_day;type:mediumint(8) unsigned;not null" json:"topic_hits_day"`       //日点击
	TopicHitsWeek  string  `gorm:"index:topic_hits_week;column:topic_hits_week;type:mediumint(8) unsigned;not null" json:"topic_hits_week"`    //周点击
	TopicHitsMonth string  `gorm:"index:topic_hits_month;column:topic_hits_month;type:mediumint(8) unsigned;not null" json:"topic_hits_month"` //月点击
	TopicTime      int     `gorm:"index:topic_time;column:topic_time;type:int(10) unsigned;not null" json:"topic_time"`                        //更新时间
	TopicTimeAdd   int     `gorm:"index:topic_time_add;column:topic_time_add;type:int(10) unsigned;not null" json:"topic_time_add"`            //添加时间
	TopicTimeHits  int     `gorm:"index:topic_time_hits;column:topic_time_hits;type:int(10) unsigned;not null" json:"topic_time_hits"`         //
	TopicTimeMake  int     `gorm:"column:topic_time_make;type:int(10) unsigned;not null" json:"topic_time_make"`                               //
	TopicTag       string  `gorm:"column:topic_tag;type:varchar(255);not null" json:"topic_tag"`                                               //专题标签
	TopicRelVod    string  `gorm:"column:topic_rel_vod;type:text" json:"topic_rel_vod"`                                                        //专题包含视频数量
	TopicRelArt    string  `gorm:"column:topic_rel_art;type:text" json:"topic_rel_art"`                                                        //专题包含文章数量
	TopicContent   string  `gorm:"column:topic_content;type:text" json:"topic_content"`                                                        //详细介绍
	TopicExtend    string  `gorm:"column:topic_extend;type:text" json:"topic_extend"`                                                          //扩展配置json
}

// TableName get sql table name.获取数据库表名
func (m *MacTopic) TableName() string {
	return "mac_topic"
}

// MacUlog [...]
type MacUlog struct {
	UlogID     int   `gorm:"primary_key;AUTO_INCREMENT;column:ulog_id;type:int(10) unsigned;not null" json:"ulog_id"` //
	UserID     int   `gorm:"index:user_id;column:user_id;type:int(10) unsigned;not null" json:"user_id"`              //
	UlogMid    int8  `gorm:"index:ulog_mid;column:ulog_mid;type:tinyint(1) unsigned;not null" json:"ulog_mid"`        //
	UlogType   int8  `gorm:"index:ulog_type;column:ulog_type;type:tinyint(1) unsigned;not null" json:"ulog_type"`     //
	UlogRid    int   `gorm:"index:ulog_rid;column:ulog_rid;type:int(10) unsigned;not null" json:"ulog_rid"`           //
	UlogSid    int8  `gorm:"column:ulog_sid;type:tinyint(3) unsigned;not null" json:"ulog_sid"`                       //
	UlogNid    int16 `gorm:"column:ulog_nid;type:smallint(6) unsigned;not null" json:"ulog_nid"`                      //
	UlogPoints int16 `gorm:"column:ulog_points;type:smallint(6) unsigned;not null" json:"ulog_points"`                //
	UlogTime   int   `gorm:"column:ulog_time;type:int(10) unsigned;not null" json:"ulog_time"`                        //
}

// TableName get sql table name.获取数据库表名
func (m *MacUlog) TableName() string {
	return "mac_ulog"
}

// MacUser [...]
type MacUser struct {
	UserID            int    `gorm:"primary_key;AUTO_INCREMENT;column:user_id;type:int(10) unsigned;not null" json:"user_id"`      //用户编号
	GroupID           int16  `gorm:"index:group_id;column:group_id;type:smallint(6) unsigned;not null" json:"group_id"`            //用户组编号
	UserName          string `gorm:"index:user_name;column:user_name;type:varchar(30);not null" json:"user_name"`                  //登录名
	UserPwd           string `gorm:"column:user_pwd;type:varchar(32);not null" json:"user_pwd"`                                    //用户密码
	UserNickName      string `gorm:"column:user_nick_name;type:varchar(30);not null" json:"user_nick_name"`                        //昵称
	UserQq            string `gorm:"column:user_qq;type:varchar(16);not null" json:"user_qq"`                                      //QQ
	UserEmail         string `gorm:"column:user_email;type:varchar(30);not null" json:"user_email"`                                //邮箱
	UserPhone         string `gorm:"column:user_phone;type:varchar(16);not null" json:"user_phone"`                                //联系电话
	UserStatus        int8   `gorm:"column:user_status;type:tinyint(1) unsigned;not null" json:"user_status"`                      //用户状态
	UserPortrait      string `gorm:"column:user_portrait;type:varchar(100);not null" json:"user_portrait"`                         //头像
	UserPortraitThumb string `gorm:"column:user_portrait_thumb;type:varchar(100);not null" json:"user_portrait_thumb"`             //头像缩略图
	UserOpenidQq      string `gorm:"column:user_openid_qq;type:varchar(40);not null" json:"user_openid_qq"`                        //绑定qq
	UserOpenidWeixin  string `gorm:"column:user_openid_weixin;type:varchar(40);not null" json:"user_openid_weixin"`                //绑定微信
	UserQuestion      string `gorm:"column:user_question;type:varchar(255);not null" json:"user_question"`                         //密保问题
	UserAnswer        string `gorm:"column:user_answer;type:varchar(255);not null" json:"user_answer"`                             //密保答案
	UserPoints        int    `gorm:"column:user_points;type:int(10) unsigned;not null" json:"user_points"`                         //积分
	UserPointsFroze   int    `gorm:"column:user_points_froze;type:int(10) unsigned;not null" json:"user_points_froze"`             //冻结积分
	UserRegTime       int    `gorm:"index:user_reg_time;column:user_reg_time;type:int(10) unsigned;not null" json:"user_reg_time"` //注册时间
	UserRegIP         int    `gorm:"column:user_reg_ip;type:int(10) unsigned;not null" json:"user_reg_ip"`                         //注册ip
	UserLoginTime     int    `gorm:"column:user_login_time;type:int(10) unsigned;not null" json:"user_login_time"`                 //登录时间
	UserLoginIP       int    `gorm:"column:user_login_ip;type:int(10) unsigned;not null" json:"user_login_ip"`                     //登录ip
	UserLastLoginTime int    `gorm:"column:user_last_login_time;type:int(10) unsigned;not null" json:"user_last_login_time"`       //上次登录时间
	UserLastLoginIP   int    `gorm:"column:user_last_login_ip;type:int(10) unsigned;not null" json:"user_last_login_ip"`           //上次登录ip
	UserLoginNum      int16  `gorm:"column:user_login_num;type:smallint(6) unsigned;not null" json:"user_login_num"`               //登录次数
	UserExtend        int16  `gorm:"column:user_extend;type:smallint(6) unsigned;not null" json:"user_extend"`                     //额外数据
	UserRandom        string `gorm:"column:user_random;type:varchar(32);not null" json:"user_random"`                              //随机码
	UserEndTime       int    `gorm:"column:user_end_time;type:int(10) unsigned;not null" json:"user_end_time"`                     //vip截止期限
	UserPid           int    `gorm:"column:user_pid;type:int(10) unsigned;not null" json:"user_pid"`                               //一级分销
	UserPid2          int    `gorm:"column:user_pid_2;type:int(10) unsigned;not null" json:"user_pid_2"`                           //二级分销
	UserPid3          int    `gorm:"column:user_pid_3;type:int(10) unsigned;not null" json:"user_pid_3"`                           //三级分销
}

// TableName get sql table name.获取数据库表名
func (m *MacUser) TableName() string {
	return "mac_user"
}

// MacActor [...]
type MacActor struct {
	ActorID        int     `gorm:"primary_key;AUTO_INCREMENT;column:actor_id;type:int(10) unsigned;not null" json:"actor_id"`               //演员id
	TypeID         int16   `gorm:"index:type_id;column:type_id;type:smallint(6) unsigned;not null" json:"type_id"`                          //分类id
	TypeID1        int16   `gorm:"index:type_id_1;column:type_id_1;type:smallint(6) unsigned;not null" json:"type_id_1"`                    //一级分类id
	ActorName      string  `gorm:"index:actor_name;column:actor_name;type:varchar(255);not null" json:"actor_name"`                         //姓名
	ActorEn        string  `gorm:"index:actor_en;column:actor_en;type:varchar(255);not null" json:"actor_en"`                               //拼音
	ActorAlias     string  `gorm:"column:actor_alias;type:varchar(255);not null" json:"actor_alias"`                                        //别名
	ActorStatus    int8    `gorm:"column:actor_status;type:tinyint(1) unsigned;not null" json:"actor_status"`                               //状态
	ActorLock      int8    `gorm:"column:actor_lock;type:tinyint(1) unsigned;not null" json:"actor_lock"`                                   //锁定
	ActorLetter    string  `gorm:"index:actor_letter;column:actor_letter;type:char(1);not null" json:"actor_letter"`                        //首字母
	ActorSex       string  `gorm:"index:actor_sex;column:actor_sex;type:char(1);not null" json:"actor_sex"`                                 //性别
	ActorColor     string  `gorm:"column:actor_color;type:varchar(6);not null" json:"actor_color"`                                          //高亮颜色
	ActorPic       string  `gorm:"column:actor_pic;type:varchar(255);not null" json:"actor_pic"`                                            //图片
	ActorBlurb     string  `gorm:"column:actor_blurb;type:varchar(255);not null" json:"actor_blurb"`                                        //简介
	ActorRemarks   string  `gorm:"column:actor_remarks;type:varchar(100);not null" json:"actor_remarks"`                                    //备注
	ActorArea      string  `gorm:"index:actor_area;column:actor_area;type:varchar(20);not null" json:"actor_area"`                          //地区
	ActorHeight    string  `gorm:"column:actor_height;type:varchar(10);not null" json:"actor_height"`                                       //身高
	ActorWeight    string  `gorm:"column:actor_weight;type:varchar(10);not null" json:"actor_weight"`                                       //体重
	ActorBirthday  string  `gorm:"column:actor_birthday;type:varchar(10);not null" json:"actor_birthday"`                                   //生日
	ActorBirtharea string  `gorm:"column:actor_birtharea;type:varchar(20);not null" json:"actor_birtharea"`                                 //出生地
	ActorBlood     string  `gorm:"column:actor_blood;type:varchar(10);not null" json:"actor_blood"`                                         //血型
	ActorStarsign  string  `gorm:"column:actor_starsign;type:varchar(10);not null" json:"actor_starsign"`                                   //星座
	ActorSchool    string  `gorm:"column:actor_school;type:varchar(20);not null" json:"actor_school"`                                       //毕业院校
	ActorWorks     string  `gorm:"column:actor_works;type:varchar(255);not null" json:"actor_works"`                                        //主要作品多个逗号相连
	ActorTag       string  `gorm:"index:actor_tag;column:actor_tag;type:varchar(255);not null" json:"actor_tag"`                            //tags
	ActorClass     string  `gorm:"index:actor_class;column:actor_class;type:varchar(255);not null" json:"actor_class"`                      //扩展分类
	ActorLevel     int8    `gorm:"index:actor_level;column:actor_level;type:tinyint(1) unsigned;not null" json:"actor_level"`               //推荐值
	ActorTime      int     `gorm:"index:actor_time;column:actor_time;type:int(10) unsigned;not null" json:"actor_time"`                     //更新时间
	ActorTimeAdd   int     `gorm:"index:actor_time_add;column:actor_time_add;type:int(10) unsigned;not null" json:"actor_time_add"`         //添加时间
	ActorTimeHits  int     `gorm:"column:actor_time_hits;type:int(10) unsigned;not null" json:"actor_time_hits"`                            //点击时间
	ActorTimeMake  int     `gorm:"column:actor_time_make;type:int(10) unsigned;not null" json:"actor_time_make"`                            //生成时间
	ActorHits      string  `gorm:"column:actor_hits;type:mediumint(8) unsigned;not null" json:"actor_hits"`                                 //总点击数
	ActorHitsDay   string  `gorm:"column:actor_hits_day;type:mediumint(8) unsigned;not null" json:"actor_hits_day"`                         //日点击数
	ActorHitsWeek  string  `gorm:"column:actor_hits_week;type:mediumint(8) unsigned;not null" json:"actor_hits_week"`                       //周点击数
	ActorHitsMonth string  `gorm:"column:actor_hits_month;type:mediumint(8) unsigned;not null" json:"actor_hits_month"`                     //月点击数
	ActorScore     float64 `gorm:"index:actor_score;column:actor_score;type:decimal(3,1) unsigned;not null" json:"actor_score"`             //平均分
	ActorScoreAll  string  `gorm:"index:actor_score_all;column:actor_score_all;type:mediumint(8) unsigned;not null" json:"actor_score_all"` //总评分
	ActorScoreNum  string  `gorm:"index:actor_score_num;column:actor_score_num;type:mediumint(8) unsigned;not null" json:"actor_score_num"` //评分次数
	ActorUp        string  `gorm:"index:actor_up;column:actor_up;type:mediumint(8) unsigned;not null" json:"actor_up"`                      //顶数
	ActorDown      string  `gorm:"index:actor_down;column:actor_down;type:mediumint(8) unsigned;not null" json:"actor_down"`                //踩数
	ActorTpl       string  `gorm:"column:actor_tpl;type:varchar(30);not null" json:"actor_tpl"`                                             //自定义模板
	ActorJumpurl   string  `gorm:"column:actor_jumpurl;type:varchar(150);not null" json:"actor_jumpurl"`                                    //跳转url
	ActorContent   string  `gorm:"column:actor_content;type:text;not null" json:"actor_content"`                                            //详情
}

// TableName get sql table name.获取数据库表名
func (m *MacActor) TableName() string {
	return "mac_actor"
}

// MacArt [...]
type MacArt struct {
	ArtID           int     `gorm:"primary_key;AUTO_INCREMENT;column:art_id;type:int(10) unsigned;not null" json:"art_id"`                //文章id
	TypeID          int16   `gorm:"index:type_id;column:type_id;type:smallint(6) unsigned;not null" json:"type_id"`                       //分类id
	TypeID1         int16   `gorm:"index:type_id_1;column:type_id_1;type:smallint(6) unsigned;not null" json:"type_id_1"`                 //一级分类id
	GroupID         int16   `gorm:"column:group_id;type:smallint(6) unsigned;not null" json:"group_id"`                                   //用户组id
	ArtName         string  `gorm:"index:art_name;column:art_name;type:varchar(255);not null" json:"art_name"`                            //标题
	ArtSub          string  `gorm:"column:art_sub;type:varchar(255);not null" json:"art_sub"`                                             //副标题
	ArtEn           string  `gorm:"index:art_en;column:art_en;type:varchar(255);not null" json:"art_en"`                                  //别名
	ArtStatus       int8    `gorm:"column:art_status;type:tinyint(1) unsigned;not null" json:"art_status"`                                //状态0未审1已审
	ArtLetter       string  `gorm:"index:art_letter;column:art_letter;type:char(1);not null" json:"art_letter"`                           //首字母
	ArtColor        string  `gorm:"column:art_color;type:varchar(6);not null" json:"art_color"`                                           //颜色
	ArtFrom         string  `gorm:"column:art_from;type:varchar(30);not null" json:"art_from"`                                            //来源
	ArtAuthor       string  `gorm:"column:art_author;type:varchar(30);not null" json:"art_author"`                                        //作者
	ArtTag          string  `gorm:"index:art_tag;column:art_tag;type:varchar(100);not null" json:"art_tag"`                               //tags
	ArtClass        string  `gorm:"column:art_class;type:varchar(255);not null" json:"art_class"`                                         //扩展分类
	ArtPic          string  `gorm:"column:art_pic;type:varchar(255);not null" json:"art_pic"`                                             //主图
	ArtPicThumb     string  `gorm:"column:art_pic_thumb;type:varchar(255);not null" json:"art_pic_thumb"`                                 //缩略图
	ArtPicSlide     string  `gorm:"column:art_pic_slide;type:varchar(255);not null" json:"art_pic_slide"`                                 //幻灯图
	ArtBlurb        string  `gorm:"column:art_blurb;type:varchar(255);not null" json:"art_blurb"`                                         //简介
	ArtRemarks      string  `gorm:"column:art_remarks;type:varchar(100);not null" json:"art_remarks"`                                     //备注
	ArtJumpurl      string  `gorm:"column:art_jumpurl;type:varchar(150);not null" json:"art_jumpurl"`                                     //跳转url
	ArtTpl          string  `gorm:"column:art_tpl;type:varchar(30);not null" json:"art_tpl"`                                              //独立模板
	ArtLevel        int8    `gorm:"index:art_level;column:art_level;type:tinyint(1) unsigned;not null" json:"art_level"`                  //推荐等级
	ArtLock         int8    `gorm:"index:art_lock;column:art_lock;type:tinyint(1) unsigned;not null" json:"art_lock"`                     //锁定
	ArtPoints       int16   `gorm:"column:art_points;type:smallint(6) unsigned;not null" json:"art_points"`                               //访问整个文章所需点数
	ArtPointsDetail int16   `gorm:"column:art_points_detail;type:smallint(6) unsigned;not null" json:"art_points_detail"`                 //访问每一页所需点数
	ArtUp           string  `gorm:"index:art_up;column:art_up;type:mediumint(8) unsigned;not null" json:"art_up"`                         //顶数
	ArtDown         string  `gorm:"index:art_down;column:art_down;type:mediumint(8) unsigned;not null" json:"art_down"`                   //踩数
	ArtHits         string  `gorm:"index:art_hits;column:art_hits;type:mediumint(8) unsigned;not null" json:"art_hits"`                   //总点击量
	ArtHitsDay      string  `gorm:"index:art_hits_day;column:art_hits_day;type:mediumint(8) unsigned;not null" json:"art_hits_day"`       //日点击量
	ArtHitsWeek     string  `gorm:"index:art_hits_week;column:art_hits_week;type:mediumint(8) unsigned;not null" json:"art_hits_week"`    //周点击量
	ArtHitsMonth    string  `gorm:"index:art_hits_month;column:art_hits_month;type:mediumint(8) unsigned;not null" json:"art_hits_month"` //月点击量
	ArtTime         int     `gorm:"index:art_time;column:art_time;type:int(10) unsigned;not null" json:"art_time"`                        //更新时间
	ArtTimeAdd      int     `gorm:"index:art_time_add;column:art_time_add;type:int(10) unsigned;not null" json:"art_time_add"`            //添加时间
	ArtTimeHits     int     `gorm:"column:art_time_hits;type:int(10) unsigned;not null" json:"art_time_hits"`                             //点击时间
	ArtTimeMake     int     `gorm:"index:art_time_make;column:art_time_make;type:int(10) unsigned;not null" json:"art_time_make"`         //生成时间
	ArtScore        float64 `gorm:"index:art_score;column:art_score;type:decimal(3,1) unsigned;not null" json:"art_score"`                //平均分
	ArtScoreAll     string  `gorm:"index:art_score_all;column:art_score_all;type:mediumint(8) unsigned;not null" json:"art_score_all"`    //总评分
	ArtScoreNum     string  `gorm:"index:art_score_num;column:art_score_num;type:mediumint(8) unsigned;not null" json:"art_score_num"`    //评分次数
	ArtRelArt       string  `gorm:"column:art_rel_art;type:varchar(255);not null" json:"art_rel_art"`                                     //关联文章
	ArtRelVod       string  `gorm:"column:art_rel_vod;type:varchar(255);not null" json:"art_rel_vod"`                                     //关联视频
	ArtPwd          string  `gorm:"column:art_pwd;type:varchar(10);not null" json:"art_pwd"`                                              //访问密码
	ArtPwdURL       string  `gorm:"column:art_pwd_url;type:varchar(255);not null" json:"art_pwd_url"`                                     //密码获取链接
	ArtTitle        string  `gorm:"column:art_title;type:mediumtext;not null" json:"art_title"`                                           //页标题
	ArtNote         string  `gorm:"column:art_note;type:mediumtext;not null" json:"art_note"`                                             //页备注
	ArtContent      string  `gorm:"column:art_content;type:mediumtext;not null" json:"art_content"`                                       //页详细介绍
}

// TableName get sql table name.获取数据库表名
func (m *MacArt) TableName() string {
	return "mac_art"
}

// MacCard [...]
type MacCard struct {
	CardID         int    `gorm:"primary_key;AUTO_INCREMENT;column:card_id;type:int(10) unsigned;not null" json:"card_id"`      //
	CardNo         string `gorm:"index:card_no;column:card_no;type:varchar(16);not null" json:"card_no"`                        //
	CardPwd        string `gorm:"index:card_pwd;column:card_pwd;type:varchar(8);not null" json:"card_pwd"`                      //
	CardMoney      int16  `gorm:"column:card_money;type:smallint(6) unsigned;not null" json:"card_money"`                       //
	CardPoints     int16  `gorm:"column:card_points;type:smallint(6) unsigned;not null" json:"card_points"`                     //
	CardUseStatus  int8   `gorm:"column:card_use_status;type:tinyint(1) unsigned;not null" json:"card_use_status"`              //
	CardSaleStatus int8   `gorm:"column:card_sale_status;type:tinyint(1) unsigned;not null" json:"card_sale_status"`            //
	UserID         int    `gorm:"index:user_id;column:user_id;type:int(10) unsigned;not null" json:"user_id"`                   //
	CardAddTime    int    `gorm:"index:card_add_time;column:card_add_time;type:int(10) unsigned;not null" json:"card_add_time"` //
	CardUseTime    int    `gorm:"index:card_use_time;column:card_use_time;type:int(10) unsigned;not null" json:"card_use_time"` //
}

// TableName get sql table name.获取数据库表名
func (m *MacCard) TableName() string {
	return "mac_card"
}

// MacCjContent [...]
type MacCjContent struct {
	ID     int    `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int(10) unsigned;not null" json:"-"` //
	Nodeid int    `gorm:"index:nodeid;column:nodeid;type:int(10) unsigned;not null" json:"nodeid"`      //
	Status int8   `gorm:"index:status;column:status;type:tinyint(1) unsigned;not null" json:"status"`   //
	URL    string `gorm:"column:url;type:char(255);not null" json:"url"`                                //
	Title  string `gorm:"column:title;type:char(100);not null" json:"title"`                            //
	Data   string `gorm:"column:data;type:mediumtext;not null" json:"data"`                             //
}

// TableName get sql table name.获取数据库表名
func (m *MacCjContent) TableName() string {
	return "mac_cj_content"
}

// MacType [...]
type MacType struct {
	TypeID        int16  `gorm:"primary_key;AUTO_INCREMENT;column:type_id;type:smallint(6) unsigned;not null" json:"type_id"` //
	TypeName      string `gorm:"index:type_name;column:type_name;type:varchar(60);not null" json:"type_name"`                 //
	TypeEn        string `gorm:"index:type_en;column:type_en;type:varchar(60);not null" json:"type_en"`                       //
	TypeSort      int16  `gorm:"index:type_sort;column:type_sort;type:smallint(6) unsigned;not null" json:"type_sort"`        //
	TypeMid       int16  `gorm:"index:type_mid;column:type_mid;type:smallint(6) unsigned;not null" json:"type_mid"`           //
	TypePid       int16  `gorm:"index:type_pid;column:type_pid;type:smallint(6) unsigned;not null" json:"type_pid"`           //
	TypeStatus    int8   `gorm:"column:type_status;type:tinyint(1) unsigned;not null" json:"type_status"`                     //
	TypeTpl       string `gorm:"column:type_tpl;type:varchar(30);not null" json:"type_tpl"`                                   //
	TypeTplList   string `gorm:"column:type_tpl_list;type:varchar(30);not null" json:"type_tpl_list"`                         //
	TypeTplDetail string `gorm:"column:type_tpl_detail;type:varchar(30);not null" json:"type_tpl_detail"`                     //
	TypeTplPlay   string `gorm:"column:type_tpl_play;type:varchar(30);not null" json:"type_tpl_play"`                         //
	TypeTplDown   string `gorm:"column:type_tpl_down;type:varchar(30);not null" json:"type_tpl_down"`                         //
	TypeKey       string `gorm:"column:type_key;type:varchar(255);not null" json:"type_key"`                                  //
	TypeDes       string `gorm:"column:type_des;type:varchar(255);not null" json:"type_des"`                                  //
	TypeTitle     string `gorm:"column:type_title;type:varchar(255);not null" json:"type_title"`                              //
	TypeUnion     string `gorm:"column:type_union;type:varchar(255);not null" json:"type_union"`                              //
	TypeExtend    string `gorm:"column:type_extend;type:text" json:"type_extend"`                                             //
	TypeLogo      string `gorm:"column:type_logo;type:varchar(255);not null" json:"type_logo"`                                //
	TypePic       string `gorm:"column:type_pic;type:varchar(255);not null" json:"type_pic"`                                  //
	TypeJumpurl   string `gorm:"column:type_jumpurl;type:varchar(150);not null" json:"type_jumpurl"`                          //
}

// TableName get sql table name.获取数据库表名
func (m *MacType) TableName() string {
	return "mac_type"
}

// MacWebsite [...]
type MacWebsite struct {
	WebsiteID           int     `gorm:"primary_key;AUTO_INCREMENT;column:website_id;type:int(10) unsigned;not null" json:"website_id"`                             //网址id
	TypeID              int16   `gorm:"index:type_id;column:type_id;type:smallint(5) unsigned;not null" json:"type_id"`                                            //分类id
	TypeID1             int16   `gorm:"index:type_id_1;column:type_id_1;type:smallint(5) unsigned;not null" json:"type_id_1"`                                      //一级分类id
	WebsiteName         string  `gorm:"index:website_name;column:website_name;type:varchar(60);not null" json:"website_name"`                                      //网址名
	WebsiteSub          string  `gorm:"column:website_sub;type:varchar(255);not null" json:"website_sub"`                                                          //副标
	WebsiteEn           string  `gorm:"index:website_en;column:website_en;type:varchar(255);not null" json:"website_en"`                                           //拼音
	WebsiteStatus       int8    `gorm:"column:website_status;type:tinyint(1) unsigned;not null" json:"website_status"`                                             //状态
	WebsiteLetter       string  `gorm:"index:website_letter;column:website_letter;type:char(1);not null" json:"website_letter"`                                    //首字母
	WebsiteColor        string  `gorm:"column:website_color;type:varchar(6);not null" json:"website_color"`                                                        //高亮颜色
	WebsiteLock         int8    `gorm:"index:website_lock;column:website_lock;type:tinyint(1) unsigned;not null" json:"website_lock"`                              //锁定
	WebsiteSort         int     `gorm:"index:website_sort;column:website_sort;type:int(10);not null" json:"website_sort"`                                          //
	WebsiteJumpurl      string  `gorm:"column:website_jumpurl;type:varchar(255);not null" json:"website_jumpurl"`                                                  //跳转url
	WebsitePic          string  `gorm:"column:website_pic;type:varchar(255);not null" json:"website_pic"`                                                          //排序
	WebsiteLogo         string  `gorm:"column:website_logo;type:varchar(255);not null" json:"website_logo"`                                                        //logo
	WebsiteArea         string  `gorm:"column:website_area;type:varchar(20);not null" json:"website_area"`                                                         //
	WebsiteLang         string  `gorm:"column:website_lang;type:varchar(10);not null" json:"website_lang"`                                                         //
	WebsiteLevel        int8    `gorm:"index:website_level;column:website_level;type:tinyint(1) unsigned;not null" json:"website_level"`                           //推荐值
	WebsiteTime         int     `gorm:"index:website_time;column:website_time;type:int(10) unsigned;not null" json:"website_time"`                                 //更新时间
	WebsiteTimeAdd      int     `gorm:"index:website_time_add;column:website_time_add;type:int(10) unsigned;not null" json:"website_time_add"`                     //添加时间
	WebsiteTimeHits     int     `gorm:"column:website_time_hits;type:int(10) unsigned;not null" json:"website_time_hits"`                                          //点击时间
	WebsiteTimeMake     int     `gorm:"index:website_time_make;column:website_time_make;type:int(10) unsigned;not null" json:"website_time_make"`                  //生成时间
	WebsiteHits         string  `gorm:"index:website_hits;column:website_hits;type:mediumint(8) unsigned;not null" json:"website_hits"`                            //总点击量
	WebsiteHitsDay      string  `gorm:"index:website_hits_day;column:website_hits_day;type:mediumint(8) unsigned;not null" json:"website_hits_day"`                //日点击量
	WebsiteHitsWeek     string  `gorm:"index:website_hits_week;column:website_hits_week;type:mediumint(8) unsigned;not null" json:"website_hits_week"`             //周点击量
	WebsiteHitsMonth    string  `gorm:"index:website_hits_month;column:website_hits_month;type:mediumint(8) unsigned;not null" json:"website_hits_month"`          //月点击量
	WebsiteScore        float64 `gorm:"index:website_score;column:website_score;type:decimal(3,1) unsigned;not null" json:"website_score"`                         //平均分
	WebsiteScoreAll     string  `gorm:"index:website_score_all;column:website_score_all;type:mediumint(8) unsigned;not null" json:"website_score_all"`             //总评分
	WebsiteScoreNum     string  `gorm:"index:website_score_num;column:website_score_num;type:mediumint(8) unsigned;not null" json:"website_score_num"`             //评分次数
	WebsiteUp           string  `gorm:"index:website_up;column:website_up;type:mediumint(8) unsigned;not null" json:"website_up"`                                  //顶数
	WebsiteDown         string  `gorm:"index:website_down;column:website_down;type:mediumint(8) unsigned;not null" json:"website_down"`                            //踩数
	WebsiteReferer      string  `gorm:"index:website_referer;column:website_referer;type:mediumint(8) unsigned;not null" json:"website_referer"`                   //总来路
	WebsiteRefererDay   string  `gorm:"index:website_referer_day;column:website_referer_day;type:mediumint(8) unsigned;not null" json:"website_referer_day"`       //日来路
	WebsiteRefererWeek  string  `gorm:"index:website_referer_week;column:website_referer_week;type:mediumint(8) unsigned;not null" json:"website_referer_week"`    //周来路
	WebsiteRefererMonth string  `gorm:"index:website_referer_month;column:website_referer_month;type:mediumint(8) unsigned;not null" json:"website_referer_month"` //月来路
	WebsiteTag          string  `gorm:"index:website_tag;column:website_tag;type:varchar(100);not null" json:"website_tag"`                                        //tags
	WebsiteClass        string  `gorm:"index:website_class;column:website_class;type:varchar(255);not null" json:"website_class"`                                  //扩展分类
	WebsiteRemarks      string  `gorm:"column:website_remarks;type:varchar(100);not null" json:"website_remarks"`                                                  //备注
	WebsiteTpl          string  `gorm:"column:website_tpl;type:varchar(30);not null" json:"website_tpl"`                                                           //自定义模板
	WebsiteBlurb        string  `gorm:"column:website_blurb;type:varchar(255);not null" json:"website_blurb"`                                                      //简介
	WebsiteContent      string  `gorm:"column:website_content;type:mediumtext;not null" json:"website_content"`                                                    //详情
}

// TableName get sql table name.获取数据库表名
func (m *MacWebsite) TableName() string {
	return "mac_website"
}

// MacCjHistory [...]
type MacCjHistory struct {
	Md5 string `gorm:"primary_key;column:md5;type:char(32);not null" json:"md5"` //
}

// TableName get sql table name.获取数据库表名
func (m *MacCjHistory) TableName() string {
	return "mac_cj_history"
}

// MacLink [...]
type MacLink struct {
	LinkID      int16  `gorm:"primary_key;AUTO_INCREMENT;column:link_id;type:smallint(6) unsigned;not null" json:"link_id"`  //编号
	LinkType    int8   `gorm:"index:link_type;column:link_type;type:tinyint(1) unsigned;not null" json:"link_type"`          //类型0文字1图片
	LinkName    string `gorm:"column:link_name;type:varchar(60);not null" json:"link_name"`                                  //名称
	LinkSort    int16  `gorm:"index:link_sort;column:link_sort;type:smallint(6);not null" json:"link_sort"`                  //排序
	LinkAddTime int    `gorm:"index:link_add_time;column:link_add_time;type:int(10) unsigned;not null" json:"link_add_time"` //添加时间
	LinkTime    int    `gorm:"index:link_time;column:link_time;type:int(10) unsigned;not null" json:"link_time"`             //更新时间
	LinkURL     string `gorm:"column:link_url;type:varchar(255);not null" json:"link_url"`                                   //链接
	LinkLogo    string `gorm:"column:link_logo;type:varchar(255);not null" json:"link_logo"`                                 //logo
}

// TableName get sql table name.获取数据库表名
func (m *MacLink) TableName() string {
	return "mac_link"
}

// MacMsg [...]
type MacMsg struct {
	MsgID      int    `gorm:"primary_key;AUTO_INCREMENT;column:msg_id;type:int(10) unsigned;not null" json:"msg_id"` //
	UserID     int    `gorm:"index:user_id;column:user_id;type:int(10) unsigned;not null" json:"user_id"`            //
	MsgType    int8   `gorm:"column:msg_type;type:tinyint(1) unsigned;not null" json:"msg_type"`                     //
	MsgStatus  int8   `gorm:"column:msg_status;type:tinyint(1) unsigned;not null" json:"msg_status"`                 //
	MsgTo      string `gorm:"column:msg_to;type:varchar(30);not null" json:"msg_to"`                                 //
	MsgCode    string `gorm:"index:msg_code;column:msg_code;type:varchar(10);not null" json:"msg_code"`              //
	MsgContent string `gorm:"column:msg_content;type:varchar(255);not null" json:"msg_content"`                      //
	MsgTime    int    `gorm:"index:msg_time;column:msg_time;type:int(10) unsigned;not null" json:"msg_time"`         //
}

// TableName get sql table name.获取数据库表名
func (m *MacMsg) TableName() string {
	return "mac_msg"
}

// MacPlog [...]
type MacPlog struct {
	PlogID      int    `gorm:"primary_key;AUTO_INCREMENT;column:plog_id;type:int(10) unsigned;not null" json:"plog_id"` //
	UserID      int    `gorm:"index:user_id;column:user_id;type:int(10) unsigned;not null" json:"user_id"`              //
	UserID1     int    `gorm:"column:user_id_1;type:int(10);not null" json:"user_id_1"`                                 //
	PlogType    int8   `gorm:"index:plog_type;column:plog_type;type:tinyint(1) unsigned;not null" json:"plog_type"`     //
	PlogPoints  int16  `gorm:"column:plog_points;type:smallint(6) unsigned;not null" json:"plog_points"`                //
	PlogTime    int    `gorm:"column:plog_time;type:int(10) unsigned;not null" json:"plog_time"`                        //
	PlogRemarks string `gorm:"column:plog_remarks;type:varchar(100);not null" json:"plog_remarks"`                      //
}

// TableName get sql table name.获取数据库表名
func (m *MacPlog) TableName() string {
	return "mac_plog"
}

// MacVisit [...]
type MacVisit struct {
	VisitID   int    `gorm:"primary_key;AUTO_INCREMENT;column:visit_id;type:int(10) unsigned;not null" json:"visit_id"` //
	UserID    int    `gorm:"index:user_id;column:user_id;type:int(10) unsigned" json:"user_id"`                         //
	VisitIP   int    `gorm:"column:visit_ip;type:int(10) unsigned;not null" json:"visit_ip"`                            //
	VisitLy   string `gorm:"column:visit_ly;type:varchar(100);not null" json:"visit_ly"`                                //
	VisitTime int    `gorm:"index:visit_time;column:visit_time;type:int(10) unsigned;not null" json:"visit_time"`       //
}

// TableName get sql table name.获取数据库表名
func (m *MacVisit) TableName() string {
	return "mac_visit"
}
