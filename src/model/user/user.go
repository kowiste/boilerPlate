package user

type User struct {
	ID       string `json:"id" validate:"required,uuid"`
	Name     string `json:"name" validate:"gt=0,lte=130"`
	LastName string `json:"lastName" validate:"gt=0,lte=130"`
	Age      uint   `json:"age" validate:"gt=0"`
}
type Users []User

type FindUsersInput struct {
	Text string `json:"text"`
	Age  int    `json:"age"`
}

func (u User) TableName() string {
	return "users"
}
