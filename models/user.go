package models

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"xiaorun", "123456", "1489456049", "1489456097"}
	UserList["xiaorun"] = &u
}

type User struct {
	Username string
	Password string
	// Profile  Profile
	Create_time int
	Last_time   int
}

// type Profile struct {
// 	Gender  string
// 	Age     int
// 	Address string
// 	Email   string
// }

func Signin(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func Signup(username, password string) string {
	u.Username = username
	u.Password = password
	UserList[u.Username] = &u
	return u.Username
}
