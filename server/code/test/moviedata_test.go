package test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"os"
	"server/model"
	"server/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
func TestInsertUsers(t *testing.T) {

	var (
		err           error
		collection    *mongo.Collection
		insertManyRes *mongo.InsertManyResult
		usersArray    = getUserInfoArray()
	)

	rand.Seed(time.Now().UTC().UnixNano())

	db, err := database.ConnectDB("mongodb://root:123456@127.0.0.1:27017/admin", "Server")

	if err != nil {
		t.Log("连接数据库失败")
		panic(err)
	}
	defer db.Close()
	t.Log("数据库连接成功")

	// collection = db.DB.Collection("Films")
	collection = db.DB.Collection("Users")
	
	//批量插入数据
	if insertManyRes, err = collection.InsertMany(getContext(), usersArray[0:]); err != nil {
		checkErr(err)
	}
	fmt.Printf("InsertMany插入的消息ID:%v\n", insertManyRes.InsertedIDs)
	var Dinfo = make(map[string]interface{})
	err = collection.FindOne(getContext(), bson.D{{"Name", "testName_8"}, {"Phone", "188212318"}}).Decode(&Dinfo)
	fmt.Println(Dinfo)
	fmt.Println("_id", Dinfo["_id"])
}

func TestInsertMovies(t *testing.T) {

	var (
		err           error
		collection    *mongo.Collection
		insertManyRes *mongo.InsertManyResult
		usersArray    = getFilmsInfoArray()
	)

	rand.Seed(time.Now().UTC().UnixNano())

	db, err := database.ConnectDB("mongodb://root:123456@127.0.0.1:27017/admin", "Server")

	if err != nil {
		t.Log("连接数据库失败")
		panic(err)
	}
	defer db.Close()
	t.Log("数据库连接成功")

	collection = db.DB.Collection("Films")

	//批量插入数据
	if insertManyRes, err = collection.InsertMany(getContext(), usersArray[0:]); err != nil {
		checkErr(err)
	}
	fmt.Printf("InsertMany插入的消息ID:%v\n", insertManyRes.InsertedIDs)
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

func getFilmsInfoArray() (data []interface{}) {
	var i int64
	for i = 0; i <= 100; i++ {
		data = append(data, model.Video{
			ID:       primitive.NewObjectID(),
			UserID:     primitive.NewObjectID(),
			UserName:    "Ving精选",
			Desc:    fmt.Sprintf("18821231%d", i+1),
			VideoURL:  "http://baobab.kaiyanapp.com/api/v1/playUrl?vid=183169&resourceType=video&editionType=default&source=aliyun&playUrlType=url_oss&f=iphone&u=&vc=0",
			Tags:  "美食|搞笑",
			Praises:  12335,
			Treads:	 12335,
			Comments:	 12335,
			Transmits:	 12335,
			UpdateTime:	 time.Now(),
			Duration: 70,
		})
	}
	return
}

func getUserInfoArray() (data []interface{}) {
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
