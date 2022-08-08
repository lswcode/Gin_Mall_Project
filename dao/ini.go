package dao

import (
	"time"

	"fmt"
	"gin_mall/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var (
	_db *gorm.DB
)

// 数据库初始化连接函数
func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info) // 如果是调试模式，设置logModel为info模式
	} else {
		ormLogger = logger.Default // 生产模式则使用Default
	}
	// 开始数据库连接
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger, // 配置logger模式
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不要自动加s
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  //设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) //打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	// 主从配置 数据库读写分离
	_ = _db.Use(dbresolver.
		Register(dbresolver.Config{
			// `db2` 作为 sources，`db3`、`db4` 作为 replicas
			Sources:  []gorm.Dialector{mysql.Open(connRead)},                         // 写操作
			Replicas: []gorm.Dialector{mysql.Open(connWrite), mysql.Open(connWrite)}, // 读操作
			Policy:   dbresolver.RandomPolicy{},                                      // sources/replicas 负载均衡策略
		}))

	// 开启自动迁移模式，即自动创建表，实际工作环境中，不推荐在后台api服务器中创建表
	err = _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{},
			&model.Product{},
			&model.Carousel{},
			&model.Category{},
			&model.Favorite{},
			&model.ProductImg{},
			&model.Order{},
			&model.Cart{},
			&model.Admin{},
			&model.Address{},
			&model.Notice{})
	if err != nil {
		fmt.Println("register table fail")
	}
	fmt.Println("register table success")

}
