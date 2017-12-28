package dao

// GetUserByID is to get user by a specified id
func GetUserByID(id int) User {
	return User{"Jack", 20, "jack@gmail.com"}
}

// DelUserByID is to delete user by a specified id
func DelUserByID(id int) bool {
	return true
}

// User is a pojo struct
type User struct {
	Name  string
	Age   int
	Email string
}
