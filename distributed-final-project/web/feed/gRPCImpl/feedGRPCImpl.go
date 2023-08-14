package gRPCImpl

import (
	"context"
	pb "distributed-final-project/web/gen/feed"
	db "distributed-final-project/web/globals"
)

type FeedApiServer struct {
	pb.UnimplementedFeedServer
}

type ErrorCustomKey struct {
	errorString string
	message     string
}

func (e *ErrorCustomKey) Error() string {
	return e.errorString
}

func (*FeedApiServer) AddPost(c context.Context, newPost *pb.Post) (*pb.FeedResponseMessage, error) {
	post := db.Post{
		PostId:     newPost.GetPostId(),
		UserId:     newPost.GetUserId(),
		Content:    newPost.GetContent(),
		DatePosted: newPost.GetDatePosted(),
	}
	db.Posts = append(db.Posts, post)
	return nil, nil
}

func (*FeedApiServer) GetPost(c context.Context, user *pb.UserId) (*pb.MultiplePost, error) {
	gRPCPosts := &pb.MultiplePost{
		Posts: make([]*pb.Post, 0),
	}
	for _, post := range db.Posts {
		if post.UserId == user.GetId() {
			p := pb.Post{
				PostId:     post.PostId,
				UserId:     post.UserId,
				Content:    post.Content,
				DatePosted: post.DatePosted,
			}
			gRPCPosts.Posts = append(gRPCPosts.Posts, &p)
		}
	}
	return gRPCPosts, nil
}

func (*FeedApiServer) GetUsers(context.Context, *pb.FeedEmptyRequest) (*pb.MultipleUser, error) {
	gRPCUsers := &pb.MultipleUser{
		Users: make([]*pb.FeedUser, 0),
	}
	for _, user := range db.Users {
		u := pb.FeedUser{
			Id:        uint32(user.Id),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}
		gRPCUsers.Users = append(gRPCUsers.Users, &u)
	}
	return gRPCUsers, nil
}
