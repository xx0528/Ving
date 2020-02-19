package model

import (
	"time"
)

// The User holds
type User struct {
	ID       int64       `bson:"_id" json:"id"`            //用户ID
	Name     string      `bson:"name" json:"name"`         //用户昵称
	Icon     string      `bson:"icon" json:"icon"`         //头像
	Email    string      `bson:"email" json:"email"`       //email
	Address  UserAddress `bson:"address" json:"address"`   //地址
	Phone    string      `bson:"phone" json:"phone"`       //手机号
	Website  string      `bson:"website" json:"website"`   //主页
	Desc     string      `bson:"desc" json:"descripthon"`  //简介
	Created  time.Time   `bson:"created" json:"created"`   //建号时间
	Updated  time.Time   `bson:"updated" json:"updated"`   //更新时间
	VideoNum int         `bson:"videoNum" json:"videoNum"` //视频个数
}

// The UserAddress holds
type UserAddress struct {
	Street  string         `bson:"street" json:"street"`
	Suite   string         `bson:"suite" json:"suite"`
	City    string         `bson:"city" json:"city"`
	Zipcode string         `bson:"zipcode" json:"zipcode"`
	Geo     UserAddressGeo `bson:"geo" json:"geo"`
}

// The UserAddressGeo holds
type UserAddressGeo struct {
	Lat string `bson:"lat" json:"lat"`
	Lng string `bson:"lng" json:"lng"`
}

// New is
func (u *User) New() *User {
	return &User{
		ID:       u.ID,
		Name:     u.Name,
		Icon:     u.Icon,
		Email:    u.Email,
		Address:  u.Address,
		Phone:    u.Phone,
		Website:  u.Website,
		Created:  time.Now(),
		Updated:  time.Now(),
		VideoNum: u.VideoNum,
	}
}
