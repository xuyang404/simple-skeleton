package starter

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-skeleton/boot"
)

var database *gorm.DB

func Database() *gorm.DB {
	return database
}

type MysqlStater struct {
	boot.BaseStarter
}

func (s *MysqlStater) Setup(c boot.StaterContext) {
	conf := c.Conf()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Mysql.User,
		conf.Mysql.Pass,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.DbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	d, err := db.DB()
	if err != nil {
		panic(err)
	}

	d.SetMaxIdleConns(conf.Mysql.MaxIdleConn)
	d.SetMaxOpenConns(conf.Mysql.MaxConn)

	err = d.Ping()
	if err != nil {
		panic(err)
	}

	database = db

	logrus.Info("mysql connect success!!!")
}
