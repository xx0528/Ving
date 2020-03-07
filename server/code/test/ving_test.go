package test

// MacCjNode [...]
type MacCjNode struct {
	Nodeid           int16  `gorm:"primary_key;column:nodeid;type:smallint(6) unsigned;not null" json:"nodeid"`
	Name             string `gorm:"column:name;type:varchar(20);not null" json:"name"`
	Lastdate         int    `gorm:"column:lastdate;type:int(10) unsigned;not null" json:"lastdate"`
	Sourcecharset    string `gorm:"column:sourcecharset;type:varchar(8);not null" json:"sourcecharset"`
	Sourcetype       int8   `gorm:"column:sourcetype;type:tinyint(1) unsigned;not null" json:"sourcetype"`
	URLpage          string `gorm:"column:urlpage;type:text;not null" json:"urlpage"`
	PagesizeStart    int8   `gorm:"column:pagesize_start;type:tinyint(3) unsigned;not null" json:"pagesize_start"`
	PagesizeEnd      string `gorm:"column:pagesize_end;type:mediumint(8) unsigned;not null" json:"pagesize_end"`
	PageBase         string `gorm:"column:page_base;type:char(255);not null" json:"page_base"`
	ParNum           int8   `gorm:"column:par_num;type:tinyint(3) unsigned;not null" json:"par_num"`
	URLContain       string `gorm:"column:url_contain;type:char(100);not null" json:"url_contain"`
	URLExcept        string `gorm:"column:url_except;type:char(100);not null" json:"url_except"`
	URLStart         string `gorm:"column:url_start;type:char(100);not null" json:"url_start"`
	URLEnd           string `gorm:"column:url_end;type:char(100);not null" json:"url_end"`
	TitleRule        string `gorm:"column:title_rule;type:char(100);not null" json:"title_rule"`
	TitleHTMLRule    string `gorm:"column:title_html_rule;type:text;not null" json:"title_html_rule"`
	TypeRule         string `gorm:"column:type_rule;type:char(100);not null" json:"type_rule"`
	TypeHTMLRule     string `gorm:"column:type_html_rule;type:text;not null" json:"type_html_rule"`
	ContentRule      string `gorm:"column:content_rule;type:char(100);not null" json:"content_rule"`
	ContentHTMLRule  string `gorm:"column:content_html_rule;type:text;not null" json:"content_html_rule"`
	ContentPageStart string `gorm:"column:content_page_start;type:char(100);not null" json:"content_page_start"`
	ContentPageEnd   string `gorm:"column:content_page_end;type:char(100);not null" json:"content_page_end"`
	ContentPageRule  int8   `gorm:"column:content_page_rule;type:tinyint(1) unsigned;not null" json:"content_page_rule"`
	ContentPage      int8   `gorm:"column:content_page;type:tinyint(1) unsigned;not null" json:"content_page"`
	ContentNextpage  string `gorm:"column:content_nextpage;type:char(100);not null" json:"content_nextpage"`
	DownAttachment   int8   `gorm:"column:down_attachment;type:tinyint(1) unsigned;not null" json:"down_attachment"`
	Watermark        int8   `gorm:"column:watermark;type:tinyint(1) unsigned;not null" json:"watermark"`
	CollOrder        int8   `gorm:"column:coll_order;type:tinyint(3) unsigned;not null" json:"coll_order"`
	CustomizeConfig  string `gorm:"column:customize_config;type:text;not null" json:"customize_config"`
	ProgramConfig    string `gorm:"column:program_config;type:text;not null" json:"program_config"`
	Mid              int8   `gorm:"column:mid;type:tinyint(1) unsigned;not null" json:"mid"`
}

// TableName get sql table name.获取数据库表名
func (m *MacCjNode) TableName() string {
	return "mac_cj_node"
}

// MacGbook [...]
type MacGbook struct {
	GbookID        int    `gorm:"primary_key;column:gbook_id;type:int(10) unsigned;not null" json:"gbook_id"`
	GbookRid       int    `gorm:"index;column:gbook_rid;type:int(10) unsigned;not null" json:"gbook_rid"`
	UserID         int    `gorm:"index;column:user_id;type:int(10) unsigned;not null" json:"user_id"`
	GbookStatus    int8   `gorm:"column:gbook_status;type:tinyint(1) unsigned;not null" json:"gbook_status"`
	GbookName      string `gorm:"column:gbook_name;type:varchar(60);not null" json:"gbook_name"`
	GbookIP        int    `gorm:"column:gbook_ip;type:int(10) unsigned;not null" json:"gbook_ip"`
	GbookTime      int    `gorm:"index;column:gbook_time;type:int(10) unsigned;not null" json:"gbook_time"`
	GbookReplyTime int    `gorm:"index;column:gbook_reply_time;type:int(10) unsigned;not null" json:"gbook_reply_time"`
	GbookContent   string `gorm:"column:gbook_content;type:varchar(255);not null" json:"gbook_content"`
	GbookReply     string `gorm:"index;column:gbook_reply;type:varchar(255);not null" json:"gbook_reply"`
}

// TableName get sql table name.获取数据库表名
func (m *MacGbook) TableName() string {
	return "mac_gbook"
}

// MacGroup [...]
type MacGroup struct {
	GroupID          int16  `gorm:"primary_key;column:group_id;type:smallint(6);not null" json:"group_id"`
	GroupName        string `gorm:"column:group_name;type:varchar(30);not null" json:"group_name"`
	GroupStatus      int8   `gorm:"index;column:group_status;type:tinyint(1) unsigned;not null" json:"group_status"`
	GroupType        string `gorm:"column:group_type;type:text;not null" json:"group_type"`
	GroupPopedom     string `gorm:"column:group_popedom;type:text;not null" json:"group_popedom"`
	GroupPointsDay   int16  `gorm:"column:group_points_day;type:smallint(6) unsigned;not null" json:"group_points_day"`
	GroupPointsWeek  int16  `gorm:"column:group_points_week;type:smallint(6);not null" json:"group_points_week"`
	GroupPointsMonth int16  `gorm:"column:group_points_month;type:smallint(6) unsigned;not null" json:"group_points_month"`
	GroupPointsYear  int16  `gorm:"column:group_points_year;type:smallint(6) unsigned;not null" json:"group_points_year"`
	GroupPointsFree  int8   `gorm:"column:group_points_free;type:tinyint(1) unsigned;not null" json:"group_points_free"`
}

// TableName get sql table name.获取数据库表名
func (m *MacGroup) TableName() string {
	return "mac_group"
}

// MacOrder [...]
type MacOrder struct {
	OrderID      int     `gorm:"primary_key;column:order_id;type:int(10) unsigned;not null" json:"order_id"`
	UserID       int     `gorm:"index;column:user_id;type:int(10) unsigned;not null" json:"user_id"`
	OrderStatus  int8    `gorm:"column:order_status;type:tinyint(1) unsigned;not null" json:"order_status"`
	OrderCode    string  `gorm:"index;column:order_code;type:varchar(30);not null" json:"order_code"`
	OrderPrice   float64 `gorm:"column:order_price;type:decimal(12,2) unsigned;not null" json:"order_price"`
	OrderTime    int     `gorm:"index;column:order_time;type:int(10) unsigned;not null" json:"order_time"`
	OrderPoints  string  `gorm:"column:order_points;type:mediumint(8) unsigned;not null" json:"order_points"`
	OrderPayType string  `gorm:"column:order_pay_type;type:varchar(10);not null" json:"order_pay_type"`
	OrderPayTime int     `gorm:"column:order_pay_time;type:int(10) unsigned;not null" json:"order_pay_time"`
	OrderRemarks string  `gorm:"column:order_remarks;type:varchar(100);not null" json:"order_remarks"`
}

// TableName get sql table name.获取数据库表名
func (m *MacOrder) TableName() string {
	return "mac_order"
}

// MacTmpwebsite [...]
type MacTmpwebsite struct {
	ID1   int    `gorm:"column:id1;type:int(10) unsigned" json:"id1"`
	Name1 string `gorm:"column:name1;type:varchar(60);not null" json:"name1"`
}

// TableName get sql table name.获取数据库表名
func (m *MacTmpwebsite) TableName() string {
	return "mac_tmpwebsite"
}

// MacVod [...]
type MacVod struct {
	VodID          int     `gorm:"primary_key;column:vod_id;type:int(10) unsigned;not null" json:"vod_id"`
	TypeID         int16   `gorm:"index;column:type_id;type:smallint(6);not null" json:"type_id"`
	TypeID1        int16   `gorm:"index;column:type_id_1;type:smallint(6) unsigned;not null" json:"type_id_1"`
	GroupID        int16   `gorm:"index;column:group_id;type:smallint(6) unsigned;not null" json:"group_id"`
	VodName        string  `gorm:"index;column:vod_name;type:varchar(255);not null" json:"vod_name"`
	VodSub         string  `gorm:"column:vod_sub;type:varchar(255);not null" json:"vod_sub"`
	VodEn          string  `gorm:"index;column:vod_en;type:varchar(255);not null" json:"vod_en"`
	VodStatus      int8    `gorm:"column:vod_status;type:tinyint(1) unsigned;not null" json:"vod_status"`
	VodLetter      string  `gorm:"index;column:vod_letter;type:char(1);not null" json:"vod_letter"`
	VodColor       string  `gorm:"column:vod_color;type:varchar(6);not null" json:"vod_color"`
	VodTag         string  `gorm:"index;column:vod_tag;type:varchar(100);not null" json:"vod_tag"`
	VodClass       string  `gorm:"index;column:vod_class;type:varchar(255);not null" json:"vod_class"`
	VodPic         string  `gorm:"column:vod_pic;type:varchar(255);not null" json:"vod_pic"`
	VodPicThumb    string  `gorm:"column:vod_pic_thumb;type:varchar(255);not null" json:"vod_pic_thumb"`
	VodPicSlide    string  `gorm:"column:vod_pic_slide;type:varchar(255);not null" json:"vod_pic_slide"`
	VodActor       string  `gorm:"index;column:vod_actor;type:varchar(255);not null" json:"vod_actor"`
	VodDirector    string  `gorm:"index;column:vod_director;type:varchar(255);not null" json:"vod_director"`
	VodWriter      string  `gorm:"column:vod_writer;type:varchar(100);not null" json:"vod_writer"`
	VodBehind      string  `gorm:"column:vod_behind;type:varchar(100);not null" json:"vod_behind"`
	VodBlurb       string  `gorm:"column:vod_blurb;type:varchar(255);not null" json:"vod_blurb"`
	VodRemarks     string  `gorm:"column:vod_remarks;type:varchar(100);not null" json:"vod_remarks"`
	VodPubdate     string  `gorm:"column:vod_pubdate;type:varchar(100);not null" json:"vod_pubdate"`
	VodTotal       string  `gorm:"index;column:vod_total;type:mediumint(8) unsigned;not null" json:"vod_total"`
	VodSerial      string  `gorm:"column:vod_serial;type:varchar(20);not null" json:"vod_serial"`
	VodTv          string  `gorm:"column:vod_tv;type:varchar(30);not null" json:"vod_tv"`
	VodWeekday     string  `gorm:"column:vod_weekday;type:varchar(30);not null" json:"vod_weekday"`
	VodArea        string  `gorm:"index;column:vod_area;type:varchar(20);not null" json:"vod_area"`
	VodLang        string  `gorm:"index;column:vod_lang;type:varchar(10);not null" json:"vod_lang"`
	VodYear        string  `gorm:"index;column:vod_year;type:varchar(10);not null" json:"vod_year"`
	VodVersion     string  `gorm:"index;column:vod_version;type:varchar(30);not null" json:"vod_version"`
	VodState       string  `gorm:"index;column:vod_state;type:varchar(30);not null" json:"vod_state"`
	VodAuthor      string  `gorm:"column:vod_author;type:varchar(60);not null" json:"vod_author"`
	VodJumpurl     string  `gorm:"column:vod_jumpurl;type:varchar(150);not null" json:"vod_jumpurl"`
	VodTpl         string  `gorm:"column:vod_tpl;type:varchar(30);not null" json:"vod_tpl"`
	VodTplPlay     string  `gorm:"column:vod_tpl_play;type:varchar(30);not null" json:"vod_tpl_play"`
	VodTplDown     string  `gorm:"column:vod_tpl_down;type:varchar(30);not null" json:"vod_tpl_down"`
	VodIsend       int8    `gorm:"index;column:vod_isend;type:tinyint(1) unsigned;not null" json:"vod_isend"`
	VodLock        int8    `gorm:"index;column:vod_lock;type:tinyint(1) unsigned;not null" json:"vod_lock"`
	VodLevel       int8    `gorm:"index;column:vod_level;type:tinyint(1) unsigned;not null" json:"vod_level"`
	VodCopyright   int8    `gorm:"column:vod_copyright;type:tinyint(1) unsigned;not null" json:"vod_copyright"`
	VodPoints      int16   `gorm:"column:vod_points;type:smallint(6) unsigned;not null" json:"vod_points"`
	VodPointsPlay  int16   `gorm:"index;column:vod_points_play;type:smallint(6) unsigned;not null" json:"vod_points_play"`
	VodPointsDown  int16   `gorm:"index;column:vod_points_down;type:smallint(6) unsigned;not null" json:"vod_points_down"`
	VodHits        string  `gorm:"index;column:vod_hits;type:mediumint(8) unsigned;not null" json:"vod_hits"`
	VodHitsDay     string  `gorm:"index;column:vod_hits_day;type:mediumint(8) unsigned;not null" json:"vod_hits_day"`
	VodHitsWeek    string  `gorm:"index;column:vod_hits_week;type:mediumint(8) unsigned;not null" json:"vod_hits_week"`
	VodHitsMonth   string  `gorm:"index;column:vod_hits_month;type:mediumint(8) unsigned;not null" json:"vod_hits_month"`
	VodDuration    string  `gorm:"column:vod_duration;type:varchar(10);not null" json:"vod_duration"`
	VodUp          string  `gorm:"index;column:vod_up;type:mediumint(8) unsigned;not null" json:"vod_up"`
	VodDown        string  `gorm:"index;column:vod_down;type:mediumint(8) unsigned;not null" json:"vod_down"`
	VodScore       float64 `gorm:"index;column:vod_score;type:decimal(3,1) unsigned;not null" json:"vod_score"`
	VodScoreAll    string  `gorm:"index;column:vod_score_all;type:mediumint(8) unsigned;not null" json:"vod_score_all"`
	VodScoreNum    string  `gorm:"index;column:vod_score_num;type:mediumint(8) unsigned;not null" json:"vod_score_num"`
	VodTime        int     `gorm:"index;column:vod_time;type:int(10) unsigned;not null" json:"vod_time"`
	VodTimeAdd     int     `gorm:"index;column:vod_time_add;type:int(10) unsigned;not null" json:"vod_time_add"`
	VodTimeHits    int     `gorm:"column:vod_time_hits;type:int(10) unsigned;not null" json:"vod_time_hits"`
	VodTimeMake    int     `gorm:"index;column:vod_time_make;type:int(10) unsigned;not null" json:"vod_time_make"`
	VodTrysee      int16   `gorm:"column:vod_trysee;type:smallint(6) unsigned;not null" json:"vod_trysee"`
	VodDoubanID    int     `gorm:"column:vod_douban_id;type:int(10) unsigned;not null" json:"vod_douban_id"`
	VodDoubanScore float64 `gorm:"column:vod_douban_score;type:decimal(3,1) unsigned;not null" json:"vod_douban_score"`
	VodReurl       string  `gorm:"column:vod_reurl;type:varchar(255);not null" json:"vod_reurl"`
	VodRelVod      string  `gorm:"column:vod_rel_vod;type:varchar(255);not null" json:"vod_rel_vod"`
	VodRelArt      string  `gorm:"column:vod_rel_art;type:varchar(255);not null" json:"vod_rel_art"`
	VodPwd         string  `gorm:"column:vod_pwd;type:varchar(10);not null" json:"vod_pwd"`
	VodPwdURL      string  `gorm:"column:vod_pwd_url;type:varchar(255);not null" json:"vod_pwd_url"`
	VodPwdPlay     string  `gorm:"column:vod_pwd_play;type:varchar(10);not null" json:"vod_pwd_play"`
	VodPwdPlayURL  string  `gorm:"column:vod_pwd_play_url;type:varchar(255);not null" json:"vod_pwd_play_url"`
	VodPwdDown     string  `gorm:"column:vod_pwd_down;type:varchar(10);not null" json:"vod_pwd_down"`
	VodPwdDownURL  string  `gorm:"column:vod_pwd_down_url;type:varchar(255);not null" json:"vod_pwd_down_url"`
	VodContent     string  `gorm:"column:vod_content;type:text;not null" json:"vod_content"`
	VodPlayFrom    string  `gorm:"column:vod_play_from;type:varchar(255);not null" json:"vod_play_from"`
	VodPlayServer  string  `gorm:"column:vod_play_server;type:varchar(255);not null" json:"vod_play_server"`
	VodPlayNote    string  `gorm:"column:vod_play_note;type:varchar(255);not null" json:"vod_play_note"`
	VodPlayURL     string  `gorm:"column:vod_play_url;type:mediumtext;not null" json:"vod_play_url"`
	VodDownFrom    string  `gorm:"column:vod_down_from;type:varchar(255);not null" json:"vod_down_from"`
	VodDownServer  string  `gorm:"column:vod_down_server;type:varchar(255);not null" json:"vod_down_server"`
	VodDownNote    string  `gorm:"column:vod_down_note;type:varchar(255);not null" json:"vod_down_note"`
	VodDownURL     string  `gorm:"column:vod_down_url;type:mediumtext;not null" json:"vod_down_url"`
	VodPlot        int8    `gorm:"index;column:vod_plot;type:tinyint(1) unsigned;not null" json:"vod_plot"`
	VodPlotName    string  `gorm:"column:vod_plot_name;type:mediumtext;not null" json:"vod_plot_name"`
	VodPlotDetail  string  `gorm:"column:vod_plot_detail;type:mediumtext;not null" json:"vod_plot_detail"`
}

// TableName get sql table name.获取数据库表名
func (m *MacVod) TableName() string {
	return "mac_vod"
}

// MacAdmin [...]
type MacAdmin struct {
	AdminID            int16  `gorm:"primary_key;column:admin_id;type:smallint(6) unsigned;not null" json:"admin_id"`
	AdminName          string `gorm:"index;column:admin_name;type:varchar(30);not null" json:"admin_name"`
	AdminPwd           string `gorm:"column:admin_pwd;type:char(32);not null" json:"admin_pwd"`
	AdminRandom        string `gorm:"column:admin_random;type:char(32);not null" json:"admin_random"`
	AdminStatus        int8   `gorm:"column:admin_status;type:tinyint(1) unsigned;not null" json:"admin_status"`
	AdminAuth          string `gorm:"column:admin_auth;type:text;not null" json:"admin_auth"`
	AdminLoginTime     int    `gorm:"column:admin_login_time;type:int(10) unsigned;not null" json:"admin_login_time"`
	AdminLoginIP       int    `gorm:"column:admin_login_ip;type:int(10) unsigned;not null" json:"admin_login_ip"`
	AdminLoginNum      int    `gorm:"column:admin_login_num;type:int(10) unsigned;not null" json:"admin_login_num"`
	AdminLastLoginTime int    `gorm:"column:admin_last_login_time;type:int(10) unsigned;not null" json:"admin_last_login_time"`
	AdminLastLoginIP   int    `gorm:"column:admin_last_login_ip;type:int(10) unsigned;not null" json:"admin_last_login_ip"`
}

// TableName get sql table name.获取数据库表名
func (m *MacAdmin) TableName() string {
	return "mac_admin"
}

// MacCash [...]
type MacCash struct {
	CashID        int     `gorm:"primary_key;column:cash_id;type:int(10) unsigned;not null" json:"cash_id"`
	UserID        int     `gorm:"index;column:user_id;type:int(10) unsigned;not null" json:"user_id"`
	CashStatus    int8    `gorm:"index;column:cash_status;type:tinyint(1) unsigned;not null" json:"cash_status"`
	CashPoints    int16   `gorm:"column:cash_points;type:smallint(6) unsigned;not null" json:"cash_points"`
	CashMoney     float64 `gorm:"column:cash_money;type:decimal(12,2) unsigned;not null" json:"cash_money"`
	CashBankName  string  `gorm:"column:cash_bank_name;type:varchar(60);not null" json:"cash_bank_name"`
	CashBankNo    string  `gorm:"column:cash_bank_no;type:varchar(30);not null" json:"cash_bank_no"`
	CashPayeeName string  `gorm:"column:cash_payee_name;type:varchar(30);not null" json:"cash_payee_name"`
	CashTime      int     `gorm:"column:cash_time;type:int(10) unsigned;not null" json:"cash_time"`
	CashTimeAudit int     `gorm:"column:cash_time_audit;type:int(10) unsigned;not null" json:"cash_time_audit"`
}

// TableName get sql table name.获取数据库表名
func (m *MacCash) TableName() string {
	return "mac_cash"
}

// MacCollect [...]
type MacCollect struct {
	CollectID         int    `gorm:"primary_key;column:collect_id;type:int(10) unsigned;not null" json:"collect_id"`
	CollectName       string `gorm:"column:collect_name;type:varchar(30);not null" json:"collect_name"`
	CollectURL        string `gorm:"column:collect_url;type:varchar(255);not null" json:"collect_url"`
	CollectType       int8   `gorm:"column:collect_type;type:tinyint(1) unsigned;not null" json:"collect_type"`
	CollectMid        int8   `gorm:"column:collect_mid;type:tinyint(1) unsigned;not null" json:"collect_mid"`
	CollectAppid      string `gorm:"column:collect_appid;type:varchar(30);not null" json:"collect_appid"`
	CollectAppkey     string `gorm:"column:collect_appkey;type:varchar(30);not null" json:"collect_appkey"`
	CollectParam      string `gorm:"column:collect_param;type:varchar(100);not null" json:"collect_param"`
	CollectFilter     int8   `gorm:"column:collect_filter;type:tinyint(1) unsigned;not null" json:"collect_filter"`
	CollectFilterFrom string `gorm:"column:collect_filter_from;type:varchar(255);not null" json:"collect_filter_from"`
	CollectOpt        int8   `gorm:"column:collect_opt;type:tinyint(1) unsigned;not null" json:"collect_opt"`
}

// TableName get sql table name.获取数据库表名
func (m *MacCollect) TableName() string {
	return "mac_collect"
}

// MacComment [...]
type MacComment struct {
	CommentID      int    `gorm:"primary_key;column:comment_id;type:int(10) unsigned;not null" json:"comment_id"`
	CommentMid     int8   `gorm:"index;column:comment_mid;type:tinyint(1) unsigned;not null" json:"comment_mid"`
	CommentRid     int    `gorm:"index;column:comment_rid;type:int(10) unsigned;not null" json:"comment_rid"`
	CommentPid     int    `gorm:"index;column:comment_pid;type:int(10) unsigned;not null" json:"comment_pid"`
	UserID         int    `gorm:"index;column:user_id;type:int(10) unsigned;not null" json:"user_id"`
	CommentStatus  int8   `gorm:"column:comment_status;type:tinyint(1) unsigned;not null" json:"comment_status"`
	CommentName    string `gorm:"column:comment_name;type:varchar(60);not null" json:"comment_name"`
	CommentIP      int    `gorm:"column:comment_ip;type:int(10) unsigned;not null" json:"comment_ip"`
	CommentTime    int    `gorm:"index;column:comment_time;type:int(10) unsigned;not null" json:"comment_time"`
	CommentContent string `gorm:"column:comment_content;type:varchar(255);not null" json:"comment_content"`
	CommentUp      string `gorm:"column:comment_up;type:mediumint(8) unsigned;not null" json:"comment_up"`
	CommentDown    string `gorm:"column:comment_down;type:mediumint(8) unsigned;not null" json:"comment_down"`
	CommentReply   string `gorm:"index;column:comment_reply;type:mediumint(8) unsigned;not null" json:"comment_reply"`
	CommentReport  string `gorm:"column:comment_report;type:mediumint(8) unsigned;not null" json:"comment_report"`
}

// TableName get sql table name.获取数据库表名
func (m *MacComment) TableName() string {
	return "mac_comment"
}

// MacRole [...]
type MacRole struct {
	RoleID        int     `gorm:"primary_key;column:role_id;type:int(10) unsigned;not null" json:"role_id"`
	RoleRid       int     `gorm:"index;column:role_rid;type:int(10) unsigned;not null" json:"role_rid"`
	RoleName      string  `gorm:"index;column:role_name;type:varchar(255);not null" json:"role_name"`
	RoleEn        string  `gorm:"index;column:role_en;type:varchar(255);not null" json:"role_en"`
	RoleStatus    int8    `gorm:"column:role_status;type:tinyint(1) unsigned;not null" json:"role_status"`
	RoleLock      int8    `gorm:"column:role_lock;type:tinyint(1) unsigned;not null" json:"role_lock"`
	RoleLetter    string  `gorm:"index;column:role_letter;type:char(1);not null" json:"role_letter"`
	RoleColor     string  `gorm:"column:role_color;type:varchar(6);not null" json:"role_color"`
	RoleActor     string  `gorm:"index;column:role_actor;type:varchar(255);not null" json:"role_actor"`
	RoleRemarks   string  `gorm:"column:role_remarks;type:varchar(100);not null" json:"role_remarks"`
	RolePic       string  `gorm:"column:role_pic;type:varchar(255);not null" json:"role_pic"`
	RoleSort      int16   `gorm:"column:role_sort;type:smallint(6) unsigned;not null" json:"role_sort"`
	RoleLevel     int8    `gorm:"index;column:role_level;type:tinyint(1) unsigned;not null" json:"role_level"`
	RoleTime      int     `gorm:"index;column:role_time;type:int(10) unsigned;not null" json:"role_time"`
	RoleTimeAdd   int     `gorm:"index;column:role_time_add;type:int(10) unsigned;not null" json:"role_time_add"`
	RoleTimeHits  int     `gorm:"column:role_time_hits;type:int(10) unsigned;not null" json:"role_time_hits"`
	RoleTimeMake  int     `gorm:"column:role_time_make;type:int(10) unsigned;not null" json:"role_time_make"`
	RoleHits      string  `gorm:"column:role_hits;type:mediumint(8) unsigned;not null" json:"role_hits"`
	RoleHitsDay   string  `gorm:"column:role_hits_day;type:mediumint(8) unsigned;not null" json:"role_hits_day"`
	RoleHitsWeek  string  `gorm:"column:role_hits_week;type:mediumint(8) unsigned;not null" json:"role_hits_week"`
	RoleHitsMonth string  `gorm:"column:role_hits_month;type:mediumint(8) unsigned;not null" json:"role_hits_month"`
	RoleScore     float64 `gorm:"index;column:role_score;type:decimal(3,1) unsigned;not null" json:"role_score"`
	RoleScoreAll  string  `gorm:"index;column:role_score_all;type:mediumint(8) unsigned;not null" json:"role_score_all"`
	RoleScoreNum  string  `gorm:"index;column:role_score_num;type:mediumint(8) unsigned;not null" json:"role_score_num"`
	RoleUp        string  `gorm:"index;column:role_up;type:mediumint(8) unsigned;not null" json:"role_up"`
	RoleDown      string  `gorm:"index;column:role_down;type:mediumint(8) unsigned;not null" json:"role_down"`
	RoleTpl       string  `gorm:"column:role_tpl;type:varchar(30);not null" json:"role_tpl"`
	RoleJumpurl   string  `gorm:"column:role_jumpurl;type:varchar(150);not null" json:"role_jumpurl"`
	RoleContent   string  `gorm:"column:role_content;type:text;not null" json:"role_content"`
}

// TableName get sql table name.获取数据库表名
func (m *MacRole) TableName() string {
	return "mac_role"
}

// MacTmpvod [...]
type MacTmpvod struct {
	ID1   int    `gorm:"column:id1;type:int(10) unsigned" json:"id1"`
	Name1 string `gorm:"column:name1;type:varchar(255);not null" json:"name1"`
}

// TableName get sql table name.获取数据库表名
func (m *MacTmpvod) TableName() string {
	return "mac_tmpvod"
}

// MacTopic [...]
type MacTopic struct {
	TopicID        int16   `gorm:"primary_key;column:topic_id;type:smallint(6) unsigned;not null" json:"topic_id"`
	TopicName      string  `gorm:"index;column:topic_name;type:varchar(255);not null" json:"topic_name"`
	TopicEn        string  `gorm:"index;column:topic_en;type:varchar(255);not null" json:"topic_en"`
	TopicSub       string  `gorm:"column:topic_sub;type:varchar(255);not null" json:"topic_sub"`
	TopicStatus    int8    `gorm:"column:topic_status;type:tinyint(1) unsigned;not null" json:"topic_status"`
	TopicSort      int16   `gorm:"index;column:topic_sort;type:smallint(6) unsigned;not null" json:"topic_sort"`
	TopicLetter    string  `gorm:"column:topic_letter;type:char(1);not null" json:"topic_letter"`
	TopicColor     string  `gorm:"column:topic_color;type:varchar(6);not null" json:"topic_color"`
	TopicTpl       string  `gorm:"column:topic_tpl;type:varchar(30);not null" json:"topic_tpl"`
	TopicType      string  `gorm:"column:topic_type;type:varchar(255);not null" json:"topic_type"`
	TopicPic       string  `gorm:"column:topic_pic;type:varchar(255);not null" json:"topic_pic"`
	TopicPicThumb  string  `gorm:"column:topic_pic_thumb;type:varchar(255);not null" json:"topic_pic_thumb"`
	TopicPicSlide  string  `gorm:"column:topic_pic_slide;type:varchar(255);not null" json:"topic_pic_slide"`
	TopicKey       string  `gorm:"column:topic_key;type:varchar(255);not null" json:"topic_key"`
	TopicDes       string  `gorm:"column:topic_des;type:varchar(255);not null" json:"topic_des"`
	TopicTitle     string  `gorm:"column:topic_title;type:varchar(255);not null" json:"topic_title"`
	TopicBlurb     string  `gorm:"column:topic_blurb;type:varchar(255);not null" json:"topic_blurb"`
	TopicRemarks   string  `gorm:"column:topic_remarks;type:varchar(100);not null" json:"topic_remarks"`
	TopicLevel     int8    `gorm:"index;column:topic_level;type:tinyint(1) unsigned;not null" json:"topic_level"`
	TopicUp        string  `gorm:"index;column:topic_up;type:mediumint(8) unsigned;not null" json:"topic_up"`
	TopicDown      string  `gorm:"index;column:topic_down;type:mediumint(8) unsigned;not null" json:"topic_down"`
	TopicScore     float64 `gorm:"index;column:topic_score;type:decimal(3,1) unsigned;not null" json:"topic_score"`
	TopicScoreAll  string  `gorm:"index;column:topic_score_all;type:mediumint(8) unsigned;not null" json:"topic_score_all"`
	TopicScoreNum  string  `gorm:"index;column:topic_score_num;type:mediumint(8) unsigned;not null" json:"topic_score_num"`
	TopicHits      string  `gorm:"index;column:topic_hits;type:mediumint(8) unsigned;not null" json:"topic_hits"`
	TopicHitsDay   string  `gorm:"index;column:topic_hits_day;type:mediumint(8) unsigned;not null" json:"topic_hits_day"`
	TopicHitsWeek  string  `gorm:"index;column:topic_hits_week;type:mediumint(8) unsigned;not null" json:"topic_hits_week"`
	TopicHitsMonth string  `gorm:"index;column:topic_hits_month;type:mediumint(8) unsigned;not null" json:"topic_hits_month"`
	TopicTime      int     `gorm:"index;column:topic_time;type:int(10) unsigned;not null" json:"topic_time"`
	TopicTimeAdd   int     `gorm:"index;column:topic_time_add;type:int(10) unsigned;not null" json:"topic_time_add"`
	TopicTimeHits  int     `gorm:"index;column:topic_time_hits;type:int(10) unsigned;not null" json:"topic_time_hits"`
	TopicTimeMake  int     `gorm:"column:topic_time_make;type:int(10) unsigned;not null" json:"topic_time_make"`
	TopicTag       string  `gorm:"column:topic_tag;type:varchar(255);not null" json:"topic_tag"`
	TopicRelVod    string  `gorm:"column:topic_rel_vod;type:text" json:"topic_rel_vod"`
	TopicRelArt    string  `gorm:"column:topic_rel_art;type:text" json:"topic_rel_art"`
	TopicContent   string  `gorm:"column:topic_content;type:text" json:"topic_content"`
	TopicExtend    string  `gorm:"column:topic_extend;type:text" json:"topic_extend"`
}

// TableName get sql table name.获取数据库表名
func (m *MacTopic) TableName() string {
	return "mac_topic"
}

// MacUlog [...]
type MacUlog struct {
	UlogID     int   `gorm:"primary_key;column:ulog_id;type:int(10) unsigned;not null" json:"ulog_id"`
	UserID     int   `gorm:"index;column:user_id;type:int(10) unsigned;not null" json:"user_id"`
	UlogMid    int8  `gorm:"index;column:ulog_mid;type:tinyint(1) unsigned;not null" json:"ulog_mid"`
	UlogType   int8  `gorm:"index;column:ulog_type;type:tinyint(1) unsigned;not null" json:"ulog_type"`
	UlogRid    int   `gorm:"index;column:ulog_rid;type:int(10) unsigned;not null" json:"ulog_rid"`
	UlogSid    int8  `gorm:"column:ulog_sid;type:tinyint(3) unsigned;not null" json:"ulog_sid"`
	UlogNid    int16 `gorm:"column:ulog_nid;type:smallint(6) unsigned;not null" json:"ulog_nid"`
	UlogPoints int16 `gorm:"column:ulog_points;type:smallint(6) unsigned;not null" json:"ulog_points"`
	UlogTime   int   `gorm:"column:ulog_time;type:int(10) unsigned;not null" json:"ulog_time"`
}

// TableName get sql table name.获取数据库表名
func (m *MacUlog) TableName() string {
	return "mac_ulog"
}

// MacUser [...]
type MacUser struct {
	UserID            int    `gorm:"primary_key;column:user_id;type:int(10) unsigned;not null" json:"user_id"`
	GroupID           int16  `gorm:"index;column:group_id;type:smallint(6) unsigned;not null" json:"group_id"`
	UserName          string `gorm:"index;column:user_name;type:varchar(30);not null" json:"user_name"`
	UserPwd           string `gorm:"column:user_pwd;type:varchar(32);not null" json:"user_pwd"`
	UserNickName      string `gorm:"column:user_nick_name;type:varchar(30);not null" json:"user_nick_name"`
	UserQq            string `gorm:"column:user_qq;type:varchar(16);not null" json:"user_qq"`
	UserEmail         string `gorm:"column:user_email;type:varchar(30);not null" json:"user_email"`
	UserPhone         string `gorm:"column:user_phone;type:varchar(16);not null" json:"user_phone"`
	UserStatus        int8   `gorm:"column:user_status;type:tinyint(1) unsigned;not null" json:"user_status"`
	UserPortrait      string `gorm:"column:user_portrait;type:varchar(100);not null" json:"user_portrait"`
	UserPortraitThumb string `gorm:"column:user_portrait_thumb;type:varchar(100);not null" json:"user_portrait_thumb"`
	UserOpenidQq      string `gorm:"column:user_openid_qq;type:varchar(40);not null" json:"user_openid_qq"`
	UserOpenidWeixin  string `gorm:"column:user_openid_weixin;type:varchar(40);not null" json:"user_openid_weixin"`
	UserQuestion      string `gorm:"column:user_question;type:varchar(255);not null" json:"user_question"`
	UserAnswer        string `gorm:"column:user_answer;type:varchar(255);not null" json:"user_answer"`
	UserPoints        int    `gorm:"column:user_points;type:int(10) unsigned;not null" json:"user_points"`
	UserPointsFroze   int    `gorm:"column:user_points_froze;type:int(10) unsigned;not null" json:"user_points_froze"`
	UserRegTime       int    `gorm:"index;column:user_reg_time;type:int(10) unsigned;not null" json:"user_reg_time"`
	UserRegIP         int    `gorm:"column:user_reg_ip;type:int(10) unsigned;not null" json:"user_reg_ip"`
	UserLoginTime     int    `gorm:"column:user_login_time;type:int(10) unsigned;not null" json:"user_login_time"`
	UserLoginIP       int    `gorm:"column:user_login_ip;type:int(10) unsigned;not null" json:"user_login_ip"`
	UserLastLoginTime int    `gorm:"column:user_last_login_time;type:int(10) unsigned;not null" json:"user_last_login_time"`
	UserLastLoginIP   int    `gorm:"column:user_last_login_ip;type:int(10) unsigned;not null" json:"user_last_login_ip"`
	UserLoginNum      int16  `gorm:"column:user_login_num;type:smallint(6) unsigned;not null" json:"user_login_num"`
	UserExtend        int16  `gorm:"column:user_extend;type:smallint(6) unsigned;not null" json:"user_extend"`
	UserRandom        string `gorm:"column:user_random;type:varchar(32);not null" json:"user_random"`
	UserEndTime       int    `gorm:"column:user_end_time;type:int(10) unsigned;not null" json:"user_end_time"`
	UserPid           int    `gorm:"column:user_pid;type:int(10) unsigned;not null" json:"user_pid"`
	UserPid2          int    `gorm:"column:user_pid_2;type:int(10) unsigned;not null" json:"user_pid_2"`
	UserPid3          int    `gorm:"column:user_pid_3;type:int(10) unsigned;not null" json:"user_pid_3"`
}

// TableName get sql table name.获取数据库表名
func (m *MacUser) TableName() string {
	return "mac_user"
}

// MacActor [...]
type MacActor struct {
	ActorID        int     `gorm:"primary_key;column:actor_id;type:int(10) unsigned;not null" json:"actor_id"`
	TypeID         int16   `gorm:"index;column:type_id;type:smallint(6) unsigned;not null" json:"type_id"`
	TypeID1        int16   `gorm:"index;column:type_id_1;type:smallint(6) unsigned;not null" json:"type_id_1"`
	ActorName      string  `gorm:"index;column:actor_name;type:varchar(255);not null" json:"actor_name"`
	ActorEn        string  `gorm:"index;column:actor_en;type:varchar(255);not null" json:"actor_en"`
	ActorAlias     string  `gorm:"column:actor_alias;type:varchar(255);not null" json:"actor_alias"`
	ActorStatus    int8    `gorm:"column:actor_status;type:tinyint(1) unsigned;not null" json:"actor_status"`
	ActorLock      int8    `gorm:"column:actor_lock;type:tinyint(1) unsigned;not null" json:"actor_lock"`
	ActorLetter    string  `gorm:"index;column:actor_letter;type:char(1);not null" json:"actor_letter"`
	ActorSex       string  `gorm:"index;column:actor_sex;type:char(1);not null" json:"actor_sex"`
	ActorColor     string  `gorm:"column:actor_color;type:varchar(6);not null" json:"actor_color"`
	ActorPic       string  `gorm:"column:actor_pic;type:varchar(255);not null" json:"actor_pic"`
	ActorBlurb     string  `gorm:"column:actor_blurb;type:varchar(255);not null" json:"actor_blurb"`
	ActorRemarks   string  `gorm:"column:actor_remarks;type:varchar(100);not null" json:"actor_remarks"`
	ActorArea      string  `gorm:"index;column:actor_area;type:varchar(20);not null" json:"actor_area"`
	ActorHeight    string  `gorm:"column:actor_height;type:varchar(10);not null" json:"actor_height"`
	ActorWeight    string  `gorm:"column:actor_weight;type:varchar(10);not null" json:"actor_weight"`
	ActorBirthday  string  `gorm:"column:actor_birthday;type:varchar(10);not null" json:"actor_birthday"`
	ActorBirtharea string  `gorm:"column:actor_birtharea;type:varchar(20);not null" json:"actor_birtharea"`
	ActorBlood     string  `gorm:"column:actor_blood;type:varchar(10);not null" json:"actor_blood"`
	ActorStarsign  string  `gorm:"column:actor_starsign;type:varchar(10);not null" json:"actor_starsign"`
	ActorSchool    string  `gorm:"column:actor_school;type:varchar(20);not null" json:"actor_school"`
	ActorWorks     string  `gorm:"column:actor_works;type:varchar(255);not null" json:"actor_works"`
	ActorTag       string  `gorm:"index;column:actor_tag;type:varchar(255);not null" json:"actor_tag"`
	ActorClass     string  `gorm:"index;column:actor_class;type:varchar(255);not null" json:"actor_class"`
	ActorLevel     int8    `gorm:"index;column:actor_level;type:tinyint(1) unsigned;not null" json:"actor_level"`
	ActorTime      int     `gorm:"index;column:actor_time;type:int(10) unsigned;not null" json:"actor_time"`
	ActorTimeAdd   int     `gorm:"index;column:actor_time_add;type:int(10) unsigned;not null" json:"actor_time_add"`
	ActorTimeHits  int     `gorm:"column:actor_time_hits;type:int(10) unsigned;not null" json:"actor_time_hits"`
	ActorTimeMake  int     `gorm:"column:actor_time_make;type:int(10) unsigned;not null" json:"actor_time_make"`
	ActorHits      string  `gorm:"column:actor_hits;type:mediumint(8) unsigned;not null" json:"actor_hits"`
	ActorHitsDay   string  `gorm:"column:actor_hits_day;type:mediumint(8) unsigned;not null" json:"actor_hits_day"`
	ActorHitsWeek  string  `gorm:"column:actor_hits_week;type:mediumint(8) unsigned;not null" json:"actor_hits_week"`
	ActorHitsMonth string  `gorm:"column:actor_hits_month;type:mediumint(8) unsigned;not null" json:"actor_hits_month"`
	ActorScore     float64 `gorm:"index;column:actor_score;type:decimal(3,1) unsigned;not null" json:"actor_score"`
	ActorScoreAll  string  `gorm:"index;column:actor_score_all;type:mediumint(8) unsigned;not null" json:"actor_score_all"`
	ActorScoreNum  string  `gorm:"index;column:actor_score_num;type:mediumint(8) unsigned;not null" json:"actor_score_num"`
	ActorUp        string  `gorm:"index;column:actor_up;type:mediumint(8) unsigned;not null" json:"actor_up"`
	ActorDown      string  `gorm:"index;column:actor_down;type:mediumint(8) unsigned;not null" json:"actor_down"`
	ActorTpl       string  `gorm:"column:actor_tpl;type:varchar(30);not null" json:"actor_tpl"`
	ActorJumpurl   string  `gorm:"column:actor_jumpurl;type:varchar(150);not null" json:"actor_jumpurl"`
	ActorContent   string  `gorm:"column:actor_content;type:text;not null" json:"actor_content"`
}

// TableName get sql table name.获取数据库表名
func (m *MacActor) TableName() string {
	return "mac_actor"
}

// MacArt [...]
type MacArt struct {
	ArtID           int     `gorm:"primary_key;column:art_id;type:int(10) unsigned;not null" json:"art_id"`
	TypeID          int16   `gorm:"index;column:type_id;type:smallint(6) unsigned;not null" json:"type_id"`
	TypeID1         int16   `gorm:"index;column:type_id_1;type:smallint(6) unsigned;not null" json:"type_id_1"`
	GroupID         int16   `gorm:"column:group_id;type:smallint(6) unsigned;not null" json:"group_id"`
	ArtName         string  `gorm:"index;column:art_name;type:varchar(255);not null" json:"art_name"`
	ArtSub          string  `gorm:"column:art_sub;type:varchar(255);not null" json:"art_sub"`
	ArtEn           string  `gorm:"index;column:art_en;type:varchar(255);not null" json:"art_en"`
	ArtStatus       int8    `gorm:"column:art_status;type:tinyint(1) unsigned;not null" json:"art_status"`
	ArtLetter       string  `gorm:"index;column:art_letter;type:char(1);not null" json:"art_letter"`
	ArtColor        string  `gorm:"column:art_color;type:varchar(6);not null" json:"art_color"`
	ArtFrom         string  `gorm:"column:art_from;type:varchar(30);not null" json:"art_from"`
	ArtAuthor       string  `gorm:"column:art_author;type:varchar(30);not null" json:"art_author"`
	ArtTag          string  `gorm:"index;column:art_tag;type:varchar(100);not null" json:"art_tag"`
	ArtClass        string  `gorm:"column:art_class;type:varchar(255);not null" json:"art_class"`
	ArtPic          string  `gorm:"column:art_pic;type:varchar(255);not null" json:"art_pic"`
	ArtPicThumb     string  `gorm:"column:art_pic_thumb;type:varchar(255);not null" json:"art_pic_thumb"`
	ArtPicSlide     string  `gorm:"column:art_pic_slide;type:varchar(255);not null" json:"art_pic_slide"`
	ArtBlurb        string  `gorm:"column:art_blurb;type:varchar(255);not null" json:"art_blurb"`
	ArtRemarks      string  `gorm:"column:art_remarks;type:varchar(100);not null" json:"art_remarks"`
	ArtJumpurl      string  `gorm:"column:art_jumpurl;type:varchar(150);not null" json:"art_jumpurl"`
	ArtTpl          string  `gorm:"column:art_tpl;type:varchar(30);not null" json:"art_tpl"`
	ArtLevel        int8    `gorm:"index;column:art_level;type:tinyint(1) unsigned;not null" json:"art_level"`
	ArtLock         int8    `gorm:"index;column:art_lock;type:tinyint(1) unsigned;not null" json:"art_lock"`
	ArtPoints       int16   `gorm:"column:art_points;type:smallint(6) unsigned;not null" json:"art_points"`
	ArtPointsDetail int16   `gorm:"column:art_points_detail;type:smallint(6) unsigned;not null" json:"art_points_detail"`
	ArtUp           string  `gorm:"index;column:art_up;type:mediumint(8) unsigned;not null" json:"art_up"`
	ArtDown         string  `gorm:"index;column:art_down;type:mediumint(8) unsigned;not null" json:"art_down"`
	ArtHits         string  `gorm:"index;column:art_hits;type:mediumint(8) unsigned;not null" json:"art_hits"`
	ArtHitsDay      string  `gorm:"index;column:art_hits_day;type:mediumint(8) unsigned;not null" json:"art_hits_day"`
	ArtHitsWeek     string  `gorm:"index;column:art_hits_week;type:mediumint(8) unsigned;not null" json:"art_hits_week"`
	ArtHitsMonth    string  `gorm:"index;column:art_hits_month;type:mediumint(8) unsigned;not null" json:"art_hits_month"`
	ArtTime         int     `gorm:"index;column:art_time;type:int(10) unsigned;not null" json:"art_time"`
	ArtTimeAdd      int     `gorm:"index;column:art_time_add;type:int(10) unsigned;not null" json:"art_time_add"`
	ArtTimeHits     int     `gorm:"column:art_time_hits;type:int(10) unsigned;not null" json:"art_time_hits"`
	ArtTimeMake     int     `gorm:"index;column:art_time_make;type:int(10) unsigned;not null" json:"art_time_make"`
	ArtScore        float64 `gorm:"index;column:art_score;type:decimal(3,1) unsigned;not null" json:"art_score"`
	ArtScoreAll     string  `gorm:"index;column:art_score_all;type:mediumint(8) unsigned;not null" json:"art_score_all"`
	ArtScoreNum     string  `gorm:"index;column:art_score_num;type:mediumint(8) unsigned;not null" json:"art_score_num"`
	ArtRelArt       string  `gorm:"column:art_rel_art;type:varchar(255);not null" json:"art_rel_art"`
	ArtRelVod       string  `gorm:"column:art_rel_vod;type:varchar(255);not null" json:"art_rel_vod"`
	ArtPwd          string  `gorm:"column:art_pwd;type:varchar(10);not null" json:"art_pwd"`
	ArtPwdURL       string  `gorm:"column:art_pwd_url;type:varchar(255);not null" json:"art_pwd_url"`
	ArtTitle        string  `gorm:"column:art_title;type:mediumtext;not null" json:"art_title"`
	ArtNote         string  `gorm:"column:art_note;type:mediumtext;not null" json:"art_note"`
	ArtContent      string  `gorm:"column:art_content;type:mediumtext;not null" json:"art_content"`
}

// TableName get sql table name.获取数据库表名
func (m *MacArt) TableName() string {
	return "mac_art"
}

// MacCard [...]
type MacCard struct {
	CardID         int    `gorm:"primary_key;column:card_id;type:int(10) unsigned;not null" json:"card_id"`
	CardNo         string `gorm:"index;column:card_no;type:varchar(16);not null" json:"card_no"`
	CardPwd        string `gorm:"index;column:card_pwd;type:varchar(8);not null" json:"card_pwd"`
	CardMoney      int16  `gorm:"column:card_money;type:smallint(6) unsigned;not null" json:"card_money"`
	CardPoints     int16  `gorm:"column:card_points;type:smallint(6) unsigned;not null" json:"card_points"`
	CardUseStatus  int8   `gorm:"column:card_use_status;type:tinyint(1) unsigned;not null" json:"card_use_status"`
	CardSaleStatus int8   `gorm:"column:card_sale_status;type:tinyint(1) unsigned;not null" json:"card_sale_status"`
	UserID         int    `gorm:"index;column:user_id;type:int(10) unsigned;not null" json:"user_id"`
	CardAddTime    int    `gorm:"index;column:card_add_time;type:int(10) unsigned;not null" json:"card_add_time"`
	CardUseTime    int    `gorm:"index;column:card_use_time;type:int(10) unsigned;not null" json:"card_use_time"`
}

// TableName get sql table name.获取数据库表名
func (m *MacCard) TableName() string {
	return "mac_card"
}

// MacCjContent [...]
type MacCjContent struct {
	ID     int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Nodeid int    `gorm:"index;column:nodeid;type:int(10) unsigned;not null" json:"nodeid"`
	Status int8   `gorm:"index;column:status;type:tinyint(1) unsigned;not null" json:"status"`
	URL    string `gorm:"column:url;type:char(255);not null" json:"url"`
	Title  string `gorm:"column:title;type:char(100);not null" json:"title"`
	Data   string `gorm:"column:data;type:mediumtext;not null" json:"data"`
}

// TableName get sql table name.获取数据库表名
func (m *MacCjContent) TableName() string {
	return "mac_cj_content"
}

// MacType [...]
type MacType struct {
	TypeID        int16  `gorm:"primary_key;column:type_id;type:smallint(6) unsigned;not null" json:"type_id"`
	TypeName      string `gorm:"index;column:type_name;type:varchar(60);not null" json:"type_name"`
	TypeEn        string `gorm:"index;column:type_en;type:varchar(60);not null" json:"type_en"`
	TypeSort      int16  `gorm:"index;column:type_sort;type:smallint(6) unsigned;not null" json:"type_sort"`
	TypeMid       int16  `gorm:"index;column:type_mid;type:smallint(6) unsigned;not null" json:"type_mid"`
	TypePid       int16  `gorm:"index;column:type_pid;type:smallint(6) unsigned;not null" json:"type_pid"`
	TypeStatus    int8   `gorm:"column:type_status;type:tinyint(1) unsigned;not null" json:"type_status"`
	TypeTpl       string `gorm:"column:type_tpl;type:varchar(30);not null" json:"type_tpl"`
	TypeTplList   string `gorm:"column:type_tpl_list;type:varchar(30);not null" json:"type_tpl_list"`
	TypeTplDetail string `gorm:"column:type_tpl_detail;type:varchar(30);not null" json:"type_tpl_detail"`
	TypeTplPlay   string `gorm:"column:type_tpl_play;type:varchar(30);not null" json:"type_tpl_play"`
	TypeTplDown   string `gorm:"column:type_tpl_down;type:varchar(30);not null" json:"type_tpl_down"`
	TypeKey       string `gorm:"column:type_key;type:varchar(255);not null" json:"type_key"`
	TypeDes       string `gorm:"column:type_des;type:varchar(255);not null" json:"type_des"`
	TypeTitle     string `gorm:"column:type_title;type:varchar(255);not null" json:"type_title"`
	TypeUnion     string `gorm:"column:type_union;type:varchar(255);not null" json:"type_union"`
	TypeExtend    string `gorm:"column:type_extend;type:text" json:"type_extend"`
	TypeLogo      string `gorm:"column:type_logo;type:varchar(255);not null" json:"type_logo"`
	TypePic       string `gorm:"column:type_pic;type:varchar(255);not null" json:"type_pic"`
	TypeJumpurl   string `gorm:"column:type_jumpurl;type:varchar(150);not null" json:"type_jumpurl"`
}

// TableName get sql table name.获取数据库表名
func (m *MacType) TableName() string {
	return "mac_type"
}

// MacWebsite [...]
type MacWebsite struct {
	WebsiteID           int     `gorm:"primary_key;column:website_id;type:int(10) unsigned;not null" json:"website_id"`
	TypeID              int16   `gorm:"index;column:type_id;type:smallint(5) unsigned;not null" json:"type_id"`
	TypeID1             int16   `gorm:"index;column:type_id_1;type:smallint(5) unsigned;not null" json:"type_id_1"`
	WebsiteName         string  `gorm:"index;column:website_name;type:varchar(60);not null" json:"website_name"`
	WebsiteSub          string  `gorm:"column:website_sub;type:varchar(255);not null" json:"website_sub"`
	WebsiteEn           string  `gorm:"index;column:website_en;type:varchar(255);not null" json:"website_en"`
	WebsiteStatus       int8    `gorm:"column:website_status;type:tinyint(1) unsigned;not null" json:"website_status"`
	WebsiteLetter       string  `gorm:"index;column:website_letter;type:char(1);not null" json:"website_letter"`
	WebsiteColor        string  `gorm:"column:website_color;type:varchar(6);not null" json:"website_color"`
	WebsiteLock         int8    `gorm:"index;column:website_lock;type:tinyint(1) unsigned;not null" json:"website_lock"`
	WebsiteSort         int     `gorm:"index;column:website_sort;type:int(10);not null" json:"website_sort"`
	WebsiteJumpurl      string  `gorm:"column:website_jumpurl;type:varchar(255);not null" json:"website_jumpurl"`
	WebsitePic          string  `gorm:"column:website_pic;type:varchar(255);not null" json:"website_pic"`
	WebsiteLogo         string  `gorm:"column:website_logo;type:varchar(255);not null" json:"website_logo"`
	WebsiteArea         string  `gorm:"column:website_area;type:varchar(20);not null" json:"website_area"`
	WebsiteLang         string  `gorm:"column:website_lang;type:varchar(10);not null" json:"website_lang"`
	WebsiteLevel        int8    `gorm:"index;column:website_level;type:tinyint(1) unsigned;not null" json:"website_level"`
	WebsiteTime         int     `gorm:"index;column:website_time;type:int(10) unsigned;not null" json:"website_time"`
	WebsiteTimeAdd      int     `gorm:"index;column:website_time_add;type:int(10) unsigned;not null" json:"website_time_add"`
	WebsiteTimeHits     int     `gorm:"column:website_time_hits;type:int(10) unsigned;not null" json:"website_time_hits"`
	WebsiteTimeMake     int     `gorm:"index;column:website_time_make;type:int(10) unsigned;not null" json:"website_time_make"`
	WebsiteHits         string  `gorm:"index;column:website_hits;type:mediumint(8) unsigned;not null" json:"website_hits"`
	WebsiteHitsDay      string  `gorm:"index;column:website_hits_day;type:mediumint(8) unsigned;not null" json:"website_hits_day"`
	WebsiteHitsWeek     string  `gorm:"index;column:website_hits_week;type:mediumint(8) unsigned;not null" json:"website_hits_week"`
	WebsiteHitsMonth    string  `gorm:"index;column:website_hits_month;type:mediumint(8) unsigned;not null" json:"website_hits_month"`
	WebsiteScore        float64 `gorm:"index;column:website_score;type:decimal(3,1) unsigned;not null" json:"website_score"`
	WebsiteScoreAll     string  `gorm:"index;column:website_score_all;type:mediumint(8) unsigned;not null" json:"website_score_all"`
	WebsiteScoreNum     string  `gorm:"index;column:website_score_num;type:mediumint(8) unsigned;not null" json:"website_score_num"`
	WebsiteUp           string  `gorm:"index;column:website_up;type:mediumint(8) unsigned;not null" json:"website_up"`
	WebsiteDown         string  `gorm:"index;column:website_down;type:mediumint(8) unsigned;not null" json:"website_down"`
	WebsiteReferer      string  `gorm:"index;column:website_referer;type:mediumint(8) unsigned;not null" json:"website_referer"`
	WebsiteRefererDay   string  `gorm:"index;column:website_referer_day;type:mediumint(8) unsigned;not null" json:"website_referer_day"`
	WebsiteRefererWeek  string  `gorm:"index;column:website_referer_week;type:mediumint(8) unsigned;not null" json:"website_referer_week"`
	WebsiteRefererMonth string  `gorm:"index;column:website_referer_month;type:mediumint(8) unsigned;not null" json:"website_referer_month"`
	WebsiteTag          string  `gorm:"index;column:website_tag;type:varchar(100);not null" json:"website_tag"`
	WebsiteClass        string  `gorm:"index;column:website_class;type:varchar(255);not null" json:"website_class"`
	WebsiteRemarks      string  `gorm:"column:website_remarks;type:varchar(100);not null" json:"website_remarks"`
	WebsiteTpl          string  `gorm:"column:website_tpl;type:varchar(30);not null" json:"website_tpl"`
	WebsiteBlurb        string  `gorm:"column:website_blurb;type:varchar(255);not null" json:"website_blurb"`
	WebsiteContent      string  `gorm:"column:website_content;type:mediumtext;not null" json:"website_content"`
}

// TableName get sql table name.获取数据库表名
func (m *MacWebsite) TableName() string {
	return "mac_website"
}

// MacCjHistory [...]
type MacCjHistory struct {
	Md5 string `gorm:"primary_key;column:md5;type:char(32);not null" json:"md5"`
}

// TableName get sql table name.获取数据库表名
func (m *MacCjHistory) TableName() string {
	return "mac_cj_history"
}

// MacLink [...]
type MacLink struct {
	LinkID      int16  `gorm:"primary_key;column:link_id;type:smallint(6) unsigned;not null" json:"link_id"`
	LinkType    int8   `gorm:"index;column:link_type;type:tinyint(1) unsigned;not null" json:"link_type"`
	LinkName    string `gorm:"column:link_name;type:varchar(60);not null" json:"link_name"`
	LinkSort    int16  `gorm:"index;column:link_sort;type:smallint(6);not null" json:"link_sort"`
	LinkAddTime int    `gorm:"index;column:link_add_time;type:int(10) unsigned;not null" json:"link_add_time"`
	LinkTime    int    `gorm:"index;column:link_time;type:int(10) unsigned;not null" json:"link_time"`
	LinkURL     string `gorm:"column:link_url;type:varchar(255);not null" json:"link_url"`
	LinkLogo    string `gorm:"column:link_logo;type:varchar(255);not null" json:"link_logo"`
}

// TableName get sql table name.获取数据库表名
func (m *MacLink) TableName() string {
	return "mac_link"
}

// MacMsg [...]
type MacMsg struct {
	MsgID      int    `gorm:"primary_key;column:msg_id;type:int(10) unsigned;not null" json:"msg_id"`
	UserID     int    `gorm:"index;column:user_id;type:int(10) unsigned;not null" json:"user_id"`
	MsgType    int8   `gorm:"column:msg_type;type:tinyint(1) unsigned;not null" json:"msg_type"`
	MsgStatus  int8   `gorm:"column:msg_status;type:tinyint(1) unsigned;not null" json:"msg_status"`
	MsgTo      string `gorm:"column:msg_to;type:varchar(30);not null" json:"msg_to"`
	MsgCode    string `gorm:"index;column:msg_code;type:varchar(10);not null" json:"msg_code"`
	MsgContent string `gorm:"column:msg_content;type:varchar(255);not null" json:"msg_content"`
	MsgTime    int    `gorm:"index;column:msg_time;type:int(10) unsigned;not null" json:"msg_time"`
}

// TableName get sql table name.获取数据库表名
func (m *MacMsg) TableName() string {
	return "mac_msg"
}

// MacPlog [...]
type MacPlog struct {
	PlogID      int    `gorm:"primary_key;column:plog_id;type:int(10) unsigned;not null" json:"plog_id"`
	UserID      int    `gorm:"index;column:user_id;type:int(10) unsigned;not null" json:"user_id"`
	UserID1     int    `gorm:"column:user_id_1;type:int(10);not null" json:"user_id_1"`
	PlogType    int8   `gorm:"index;column:plog_type;type:tinyint(1) unsigned;not null" json:"plog_type"`
	PlogPoints  int16  `gorm:"column:plog_points;type:smallint(6) unsigned;not null" json:"plog_points"`
	PlogTime    int    `gorm:"column:plog_time;type:int(10) unsigned;not null" json:"plog_time"`
	PlogRemarks string `gorm:"column:plog_remarks;type:varchar(100);not null" json:"plog_remarks"`
}

// TableName get sql table name.获取数据库表名
func (m *MacPlog) TableName() string {
	return "mac_plog"
}

// MacVisit [...]
type MacVisit struct {
	VisitID   int    `gorm:"primary_key;column:visit_id;type:int(10) unsigned;not null" json:"visit_id"`
	UserID    int    `gorm:"index;column:user_id;type:int(10) unsigned" json:"user_id"`
	VisitIP   int    `gorm:"column:visit_ip;type:int(10) unsigned;not null" json:"visit_ip"`
	VisitLy   string `gorm:"column:visit_ly;type:varchar(100);not null" json:"visit_ly"`
	VisitTime int    `gorm:"index;column:visit_time;type:int(10) unsigned;not null" json:"visit_time"`
}

// TableName get sql table name.获取数据库表名
func (m *MacVisit) TableName() string {
	return "mac_visit"
}
