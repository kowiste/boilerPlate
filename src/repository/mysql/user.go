package mysql

import (
	"boiler/src/model/user"
	"context"
)

func (m MySQL) CreateUser(c context.Context, user *user.User) (id string, err error) {
	result := m.db.Create(user)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = user.ID
	return
}

func (m MySQL) GetUsers(c context.Context) (users user.Users, err error) {
	result := m.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
func (m MySQL) GetUserByID(c context.Context, id string) (u *user.User, err error) {
	u=new(user.User)
	result := m.db.Find(&u)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
func (m MySQL) UpdateUser(c context.Context,  user *user.User) (err error) {
	result := m.db.Model(user).Where("id = ?", user.ID).Updates(user)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
func (m MySQL) DeleteUser(c context.Context, id string) (err error) {
	result := m.db.Delete(&user.User{}, "id = ?", id)
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
