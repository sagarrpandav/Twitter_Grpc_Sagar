package gRPCImpl

import (
	"context"
	"crypto/md5"
	pb "distributed-final-project/web/gen/auth"
	db "distributed-final-project/web/globals"
	"encoding/hex"
)

type AuthApiServer struct {
	pb.UnimplementedAuthServer
}

type ErrorCustomKey struct {
	errorString string
	message string
}

func (e *ErrorCustomKey) Error() string {
	return e.errorString
}

func (*AuthApiServer) SignUp(c context.Context ,user *pb.User) (*pb.User, error) {
	err := ErrorCustomKey{}
	for _, registeredUser := range db.Users {
		if user.Email == registeredUser.Email {
			err.errorString = "Error"
			return nil, &err
		}
	}
	if len(err.errorString) == 0 {
		hash := md5.Sum([]byte(user.Password + user.Email))
		user.UserHash = hex.EncodeToString(hash[:])
		newUser := db.User{
			Id: len(db.Users),
			FirstName: user.GetFirstName(),
			LastName:  user.GetLastName(),
			Email:     user.GetEmail(),
			Password:  user.GetPassword(),
		}
		db.Userhash[user.GetUserHash()] = user.GetEmail()
		db.Users = append(db.Users, newUser)
	}
	return user, nil
}

func (*AuthApiServer) SignIn(ctx context.Context, userCredentials *pb.LoginUser) (*pb.User, error) {
	for _, registeredUser := range db.Users {
		if registeredUser.Email == userCredentials.Email &&  registeredUser.Password == userCredentials.Password {
			hash := md5.Sum([]byte(userCredentials.Password + userCredentials.Email))
			userHash := hex.EncodeToString(hash[:])
			newUser := pb.User{
				Id: uint32(len(db.Users)),
				FirstName: registeredUser.FirstName,
				LastName:  registeredUser.LastName,
				Email:     registeredUser.Email,
				Password:  registeredUser.Password,
				UserHash:  userHash,
			}
			db.Userhash[userHash] = newUser.GetEmail()
			return &newUser, nil
		}
	}
	err := ErrorCustomKey{
		errorString: "error",
	}
	return nil, &err
}

func (*AuthApiServer) SignOut(ctx context.Context, token *pb.Token) (*pb.ResponseMessage, error) {
	if _, ok := db.Userhash[token.GetUserHash()]; ok {
		delete(db.Userhash, token.GetUserHash())
	}
	res := pb.ResponseMessage{Message: token.UserHash}
	return &res, nil
}

func (*AuthApiServer) ValidateUserLoggedIn(c context.Context, token *pb.Token) (*pb.ResponseMessage, error) {
	if _, ok := db.Userhash[token.GetUserHash()]; ok {
		res := pb.ResponseMessage{Message: "User Authenticated"}
		return &res, nil
	}
	err := ErrorCustomKey{
		errorString: "Error",
		message:     "User not Authenticated",
	}
	return nil, &err
}