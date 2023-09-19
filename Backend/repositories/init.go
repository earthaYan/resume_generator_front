package repositories

import (
	"context"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	db_user    = "root"
	db_pwd     = "123456"
	db_name    = "resume"
	db_host    = "116.204.108.126"
	db_port    = "3306"
	db_charset = "utf8mb4"
)

var _db *gorm.DB

func InitMysql() {
	conn := strings.Join([]string{db_user, ":", db_pwd, "@tcp(", db_host, ":", db_port, ")/", db_name, "?charset=", db_charset, "&parseTime=true"}, "")
	var ormLogger = logger.Default
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,
		DefaultStringSize:         256,   //string类型的默认字段
		DisableDatetimePrecision:  true,  //禁用dateTime精度，mysql5.6不支持
		DontSupportRenameIndex:    true,  //重命名索引的时候，采用删除并重建的方式，5.7之前不支持
		DontSupportRenameColumn:   true,  //mysql8之前不支持
		SkipInitializeWithVersion: false, //不根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger, //打印日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表名不加s
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  //设置连接池
	sqlDB.SetMaxOpenConns(100)                 //最大打开数
	sqlDB.SetConnMaxLifetime(time.Second * 30) //超时时间
	_db = db
	migrate()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
