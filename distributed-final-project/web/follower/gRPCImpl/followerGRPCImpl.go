package gRPCImpl

import (
	"context"
	pb "distributed-final-project/web/gen/follower"
	db "distributed-final-project/web/globals"
)

type FollowerServer struct {
	pb.UnimplementedFollowerServer
}

type ErrorCustomKey struct {
	errorString string
	message     string
}

func (e *ErrorCustomKey) Error() string {
	return e.errorString
}

func (*FollowerServer) GetFollowers(context context.Context, followerRequest *pb.GetFollowersRequest) (*pb.FollowResponseMessage, error) {
	for _, registeredUser := range db.Users {
		if registeredUser.Id == followerRequest.SelfId {
			response := pb.FollowResponseMessage{Message: string(registeredUser.Following)}
			return &response, nil
		}
	}
	return nil, nil
}

func (*FollowerServer) FollowUser(context context.Context, followRequest *pb.FollowRequest) (*pb.FollowResponseMessage, error) {
	tmpUser := db.User{}
	var slice = make([]db.User, len(db.Users))
	for _, registeredUser := range db.Users {
		if registeredUser.Id == followRequest.SelfId {
			registeredUser.Following = append(registeredUser.Following, followRequest.OtherUserId)
			slice = append(slice, registeredUser)
			response := pb.FollowResponseMessage{Message: string("success")}
			return &response, nil
		} else {
			slice = append(slice, registeredUser)
		}
	}
	db.Users = append(db.Users, tmpUser)
	return nil, nil
}

func (*FollowerServer) UnFollowUser(context context.Context, followRequest *pb.FollowRequest) (*pb.FollowResponseMessage, error) {
	for _, registeredUser := range db.Users {
		if registeredUser.Id == followRequest.SelfId {
			registeredUser.Following = remove(registeredUser.Following, followRequest.GetOtherUserId())
			response := pb.FollowResponseMessage{Message: string("success")}
			return &response, nil
		}
	}
	return nil, nil
}

func remove(s []int32, r int32) []int32 {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
