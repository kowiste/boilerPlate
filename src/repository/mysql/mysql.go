package mysql

import (
	conf "boiler/src/config"
	"boiler/src/model/asset"
	"boiler/src/model/user"
	"fmt"

	"github.com/kowiste/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	db *gorm.DB
}

func New() *MySQL {
	return &MySQL{}
}

func (m *MySQL) Init() (err error) {
	cnf, err := config.Get[conf.BoilerConfig]()
	if err != nil {
		fmt.Println("Error getting config:", err)
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cnf.DatabaseUser,
		cnf.DatabasePassword,
		cnf.DatabaseURL,
		cnf.DatabasePort,
		cnf.DatabaseName,
	)
	m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto Migrate User model
	err = m.db.AutoMigrate(&user.User{}, &asset.Asset{})
	if err != nil {
		return err
	}

	return nil
}


