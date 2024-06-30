package user

type User struct {
	ID       string `json:"id" validate:"required,uuid"`
	Name     string `json:"name" validate:"gt=0,lte=130"`
	LastName string `json:"lastName" validate:"gt=0,lte=130"`
}
type Users []User

func (u User) TableName() string {
	return "users"
}
