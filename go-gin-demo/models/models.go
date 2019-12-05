package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mirkowu/go-gin-demo/pkg/logging"
	"github.com/mirkowu/go-gin-demo/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID        int   `sql:"auto_increment;primary_key;unique" json:"-"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

//请在包的引号前加一个 "_" ，以表示自动调用相关包内的init方法(因为在main中使用过，故也会自动调用包内的init方法
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = setting.DB_TYPE
	dbName = setting.DB_NAME
	user = setting.DB_USER
	password = setting.DB_PASSWORD
	host = setting.DB_HOST
	tablePrefix = setting.DB_TABLE_PREFIX

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		logging.Error(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	//以实现结构体名为非复数形式：默认不设置的时候就是false
	db.SingularTable(true)

	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	//db.Callback().Create() /*.Before("gorm:create")*/ .Register("gorm:before_create", BeforeCreate)
	//	db.Callback().Create() /*.Before("gorm:create")*/ .Register("gorm:after_update", BeforeUpdate)
	//db.Callback().Create()/*.Before("gorm:create")*/.Register("update_created_at", BeforeCreate)
	//db.Callback().Create()/*.Before("gorm:update")*/.Register("update_created_at", BeforeUpdate)
	//db.Callback().Create().Before("gorm:update").Register("before_updated", BeforeUpdate)
}

//func BeforeCreate(scope *gorm.Scope) {
//	if scope.HasColumn("created_at") {
//		scope.SetColumn("created_at", time.Now().Unix())
//	}
//	return
//}

//func BeforeUpdate(scope *gorm.Scope) {
//	if scope.HasColumn("updated_at") {
//		scope.SetColumn("updated_at", time.Now().Unix())
//	}
//	return
//}

func CloseDB() {
	defer db.Close()
}
