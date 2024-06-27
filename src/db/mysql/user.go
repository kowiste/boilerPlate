package mysql

import "boiler/src/model/user"

func (m MySQL) CreateUser(u *user.User) (err error) {
	result := m.db.Create(u)
	if result.Error != nil {
		return result.Error
	}
	return
}

func (m MySQL) GetUsers() (users *user.Users, err error) {
	result := m.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return
}
