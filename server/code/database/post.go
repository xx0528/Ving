package database

import (
	"context"
	"server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
)

// GetPosts returns all posts.
// start, end int, order, sort string
func (d *MovieDatabase) GetPosts(paging *model.Paging) []*model.Post {
	posts := []*model.Post{}
	condition := bson.D{}
	if paging.Condition != nil {
		condition = (paging.Condition).(bson.D)
	}
	cursor, err := d.DB.Collection("posts").
		Find(context.Background(), condition,
			&options.FindOptions{
				Skip:  paging.Skip,
				Sort:  bson.D{bson.E{Key: paging.SortKey, Value: paging.SortVal}},
				Limit: paging.Limit,
			})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		post := &model.Post{}
		if err := cursor.Decode(post); err != nil {
			return nil
		}
		posts = append(posts, post)
	}

	return posts
}

// CreatePost creates a post.
func (d *MovieDatabase) CreatePost(post *model.Post) *model.Post {
	// Specifies the order in which to return results.
	upsert := true
	result := d.DB.Collection("posts").
		FindOneAndReplace(context.Background(),
			bson.D{{Key: "_id", Value: post.ID}},
			post,
			&options.FindOneAndReplaceOptions{
				Upsert: &upsert,
			},
		)
	if result != nil {
		return post
	}
	return nil
}

// GetPostByID returns the post by the given id or nil.
func (d *MovieDatabase) GetPostByID(id primitive.ObjectID) *model.Post {
	var post *model.Post
	err := d.DB.Collection("posts").
		FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).
		Decode(&post)
	if err != nil {
		return nil
	}
	return post
}

// DeletePostByID deletes a post by its id.
func (d *MovieDatabase) DeletePostByID(id primitive.ObjectID) error {
	_, err := d.DB.Collection("posts").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	return err
}

// UpdatePost updates a post.
func (d *MovieDatabase) UpdatePost(post *model.Post) *model.Post {
	result := d.DB.Collection("posts").
		FindOneAndReplace(context.Background(),
			bson.D{{Key: "_id", Value: post.ID}},
			post,
			&options.FindOneAndReplaceOptions{},
		)
	if result != nil {
		return post
	}
	return nil
}

// CountPost returns the post count
func (d *MovieDatabase) CountPost(condition interface{}) string {
	cd := bson.D{}
	if condition != nil {
		cd = (condition).(bson.D)
	}
	total, err := d.DB.Collection("posts").CountDocuments(context.Background(), cd, &options.CountOptions{})
	if err != nil {
		return "0"
	}
	return strconv.Itoa(int(total))
}
