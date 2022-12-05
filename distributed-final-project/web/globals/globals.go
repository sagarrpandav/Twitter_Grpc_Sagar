package globals

var Secret = []byte("secret")

const Userkey = "user"

var Userhash = make(map[string]string)

type User struct {
	Id int `json:"id" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUser struct {
	Email string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	TokenHash string `json:"tokenHash" binding:"required"`
}

type Post struct {
	PostId int
	UserId int
	content string
	DatePosted string
}

type Followers struct {
	UserId int
	Following []int
}

var Users []User
var Posts []Post

func init()  {
	print("Initializing")
	Users = []User {
		{
			Id: 1,
			FirstName: "Irfan",
			LastName:  "Ahmed",
			Email:     "ia2167@nyu.edu",
			Password:  "pass123",
		},
		{
			Id: 2,
			FirstName: "Sagar",
			LastName:  "Pandav",
			Email:     "srp8070@nyu.edu",
			Password:  "sagar123",
		},
	}
}