createAuth:
	protoc --proto_path=auth/authpb auth/authpb/*.proto --go_out=gen
	protoc --proto_path=auth/authpb auth/authpb/*.proto --go-grpc_out=gen

createPost:
	protoc --proto_path=post/postpb post/postpb/*.proto --go_out=gen
	protoc --proto_path=post/postpb post/postpb/*.proto --go-grpc_out=gen

createFollower:
	protoc --proto_path=follower/followerpb follower/followerpb/*.proto --go_out=gen
	protoc --proto_path=follower/followerpb follower/followerpb/*.proto --go-grpc_out=gen

clean:
	rm gen/*/*.go