package initialize

import (
	"fmt"
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitMysql 初始化Mysql数据库
func InitMysql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		global.GVA_CONFIG.Mysql.Username,
		global.GVA_CONFIG.Mysql.Password,
		global.GVA_CONFIG.Mysql.Path,
		global.GVA_CONFIG.Mysql.Dbname,
		global.GVA_CONFIG.Mysql.Config)

	log.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(global.GVA_CONFIG.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.GVA_CONFIG.Mysql.MaxOpenConns)

	if err := InitMysqlTables(db); err != nil {
		return nil, err
	}
	return db, nil
}

// InitMysqlTables 注册数据库表专用
func InitMysqlTables(db *gorm.DB) error {
	err := db.AutoMigrate(
		&business.BusCustomer{},
		&business.BusCustomerAsset{},
		&business.BusCustomerAssetBill{},
		&business.BusCustomerDepositOrder{},
		&business.BusCustomerDepositPayRecord{},
		&business.BusCustomerDepositPayment{},
	)
	if err != nil {
		global.GVA_LOG.Error("InitTables", zap.Any("err", err))
		return err
	}
	global.GVA_LOG.Info("register table success")
	return nil
}
