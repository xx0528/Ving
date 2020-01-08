package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The User holds
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`				//用户ID
	Name     string             `bson:"name" json:"name"`			//用户昵称
	Email    string             `bson:"email" json:"email"`			//email
	Address  UserAddress        `bson:"address" json:"address"`		//地址
	Phone    string             `bson:"phone" json:"phone"`			//手机号
	Website  string             `bson:"website" json:"website"`		//主页
	Created  time.Time          `bson:"created" json:"created"`		//建号时间
	Updated  time.Time          `bson:"updated" json:"updated"`		//更新时间
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
		ID:       primitive.NewObjectID(),
		Name:     u.Name,
		Email:    u.Email,
		Address:  u.Address,
		Phone:    u.Phone,
		Website:  u.Website,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
}
