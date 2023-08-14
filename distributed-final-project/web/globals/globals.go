package globals

var Secret = []byte("secret")

const Userkey = "user"

var Userhash = make(map[string]string)

type User struct {
	Id        int32   `json:"id" binding:"required"`
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Email     string  `json:"email" binding:"required"`
	Password  string  `json:"password" binding:"required"`
	Following []int32 `json:"following"`
}

type LoginUser struct {
	Email    string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthUser struct {
	UserId string
	Token  string
}

type Auth struct {
	AuthDb string
}

type UserDbObject struct {
	UserDb string
}

type Token struct {
	TokenHash string `json:"tokenHash" binding:"required"`
}

type Post struct {
	PostId     int32
	UserId     int32
	Content    string
	DatePosted string
}

type Followers struct {
	UserId    int32
	Following []int
}

var Users []User
var Posts []Post

func init() {
	print("Initializing")
	Users = []User{
		{
			Id:        1,
			FirstName: "Irfan",
			LastName:  "Ahmed",
			Email:     "ia2167@nyu.edu",
			Password:  "pass123",
			Following: make([]int32, 0),
		},
		{
			Id:        2,
			FirstName: "Sagar",
			LastName:  "Pandav",
			Email:     "srp8070@nyu.edu",
			Password:  "sagar123",
			Following: make([]int32, 0),
		},
	}

	Posts = []Post{
		{
			PostId:     1,
			UserId:     1,
			Content:    "My First Post",
			DatePosted: "10-10-2022",
		},
		{
			PostId:     2,
			UserId:     2,
			Content:    "Its A lovely Day",
			DatePosted: "10-15-2021",
		},
	}
}
