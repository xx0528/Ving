package test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"os"
	"server/i18n"
	"server/model"
	"server/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestInsertMovie(t *testing.T) {

	var (
		err           error
		collection    *mongo.Collection
		insertOneRes  *mongo.InsertOneResult
		insertManyRes *mongo.InsertManyResult
		usersArray    = GetUserInfoArray()
	)

	rand.Seed(time.Now().UTC().UnixNano())

	i18n.Load()

	db, err := database.ConnectDB("mongodb://root:123456@127.0.0.1:27017/admin", "Movie")

	if err != nil {
		t.Log("连接数据库失败")
		panic(err)
	}
	defer db.Close()
	t.Log("数据库连接成功")

	collection = db.DB.Collection("Movie")
	//插入一条数据
	if insertOneRes, err = collection.InsertOne(getContext(), usersArray[0]); err != nil {
		checkErr(err)
	}

	fmt.Printf("InsertOne插入的消息ID:%v\n", insertOneRes.InsertedID)

	//批量插入数据
	if insertManyRes, err = collection.InsertMany(getContext(), usersArray[1:]); err != nil {
		checkErr(err)
	}
	fmt.Printf("InsertMany插入的消息ID:%v\n", insertManyRes.InsertedIDs)
	var Dinfo = make(map[string]interface{})
	err = collection.FindOne(getContext(), bson.D{{"Name", "testName_8"}, {"Phone", 188212318}}).Decode(&Dinfo)
	fmt.Println(Dinfo)
	fmt.Println("_id", Dinfo["_id"])
}

func checkErr(err error) {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("没有查到数据")
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(0)
		}

	}
}

func getContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}

func GetUserInfoArray() (data []interface{}) {
	var i int64
	for i = 0; i <= 100; i++ {
		data = append(data, model.User{
			ID:       primitive.NewObjectID(),
			Name:     fmt.Sprintf("testName_%d", i+1),
			Email:    fmt.Sprintf("1235_%d@gmail.com", i+1),
			Phone:    fmt.Sprintf("18821231%d", i+1),
			Website:  fmt.Sprintf("www.baidu.com"),
			Created:  time.Now(),
			Updated:  time.Now(),
		})
	}
	return
}
