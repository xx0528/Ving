package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"server/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}

type DatabaseSuite struct {
	suite.Suite
	db *MovieDatabase
}

func (s *DatabaseSuite) BeforeTest(suiteName, testName string) {
	s.T().Log("--BeforeTest--")
	db, _ := New("mongodb://root:123456@localhost:27017", "tenapi")
	s.db = db
}

func (s *DatabaseSuite) AfterTest(suiteName, testName string) {
	s.db.Close()
}

func (s *DatabaseSuite) TestPost() {
	s.db.DB.Collection("posts").Drop(nil)

	var err error
	for i := 1; i <= 25; i++ {
		// user1
		UserID, _ := primitive.ObjectIDFromHex("5c99bd941ba7b2304ad8c52a")
		article := (&model.Post{
			UserID: UserID,
			Title:  fmt.Sprintf("tile%d", i),
			Body:   "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
		}).New()
		s.db.CreatePost(article)
	}
	assert.Nil(s.T(), err)
}

func (s *DatabaseSuite) TestUpdatePost() {
	id, _ := primitive.ObjectIDFromHex("5c92e6199929adef73bceea1")
	userID, _ := primitive.ObjectIDFromHex("5c8f9a83da2c3fed4eee9dc1")

	post := &model.Post{
		ID:     id,
		UserID: userID,
		Title:  "title1",
		Body:   "title1bodytitle1body",
	}

	assert.Equal(s.T(), post, s.db.UpdatePost(post))
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
