package user

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
}
type Users []User

func (u User) TableName() string {
	return "users"
}

